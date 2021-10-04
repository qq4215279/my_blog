import { asyncRoutes, constantRoutes } from '@/router'
// import { stat } from 'fs'

/**
 * Use meta.role to determine if the current user has permission
 * @param roles
 * @param route
 */
function hasPermission(roles, route) {
  if (route.meta && route.meta.roles) {
    return roles.some(role => route.meta.roles.includes(role))
  } else {
    return true
  }
}

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
export function filterAsyncRoutes(routes, roles) {
  const res = []

  routes.forEach(route => {
    const tmp = { ...route }
    if (hasPermission(roles, tmp)) {
      if (tmp.children) {
        tmp.children = filterAsyncRoutes(tmp.children, roles)
      }
      res.push(tmp)
    }
  })

  return res
}
export function filterAsyncPrivileges(routes, privileges) {
  const res = []
  routes.forEach(route => {
    const projectChildren = []
    const tmp = { ...route }
    for (const menu in privileges) {
      if (menu !== route.path) {
        continue
      }
      route.hidden = false
      if (!tmp.children) {
        continue
      }
      privileges[menu].forEach((subMenu, index) => {
        tmp.children.forEach(child => {
          if (index === 0) {
            tmp.redirect = `/project/${subMenu.path}`
          }
          if (menu === '/project') {
            child.hidden = false
            child.name = subMenu.name
            child.path = subMenu.path
            child.meta.projectId = subMenu.projectId
            child.meta.title = subMenu.title
            const autoProject = { ...child }
            autoProject.meta = { ...child.meta }
            projectChildren.push(autoProject)
          }
          if (subMenu.path === child.path) {
            child.hidden = false
            child.meta.title = subMenu.title
          }
        })
      })
      if (menu === '/project' && projectChildren.length > 0) {
        tmp.children = projectChildren
      }
    }
    res.push(tmp)
  })
  return res
}

const state = {
  routes: [],
  addRoutes: [],
  needReload: false
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = constantRoutes.concat(routes)
  },
  SET_NEEDRELOAD: (state, val) => {
    state.needReload = val
  }
}

const actions = {
  // generateRoutes({ commit }, roles) {
  //   return new Promise(resolve => {
  //     let accessedRoutes
  //     if (roles.includes('admin')) {
  //       accessedRoutes = asyncRoutes || []
  //     } else {
  //       accessedRoutes = filterAsyncRoutes(asyncRoutes, roles)
  //     }
  //     commit('SET_ROUTES', accessedRoutes)
  //     commit('SET_NEEDRELOAD', true)
  //     resolve(accessedRoutes)
  //   })
  // },
  generateRoutes({ commit }, privileges) {
    return new Promise(resolve => {
      const accessedRoutes = filterAsyncPrivileges(asyncRoutes, privileges)
      commit('SET_ROUTES', accessedRoutes)
      commit('SET_NEEDRELOAD', true)
      resolve(accessedRoutes)
    })
  },
  routerRefreshDone({ commit }) {
    return new Promise(resolve => {
      commit('SET_NEEDRELOAD', false)
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
