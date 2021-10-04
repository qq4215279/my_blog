import { login, loginByToken, logout } from '@/api/user'
import { setToken, removeToken } from '@/utils/auth'
import { setCookie, delCookie } from '@/utils/cookie'

const state = {
  userName: '',
  nickName: '',
  networkType: '',
  loginToken: '',
  roles: [],
  privileges: [],
  functions: []
}

const mutations = {
  SET_USERNAME: (state, userName) => {
    state.userName = userName
  },
  SET_NICKNAME: (state, nickName) => {
    state.nickName = nickName
  },
  SET_NETWORKTYPE: (state, networkType) => {
    state.networkType = networkType
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_PRIVILEGES: (state, privileges) => {
    state.privileges = privileges
  },
  SET_FUNCTIONS: (state, functions) => {
    state.functions = functions
  },
  SET_LOGINTOKEN: (state, loginToken) => {
    state.loginToken = loginToken
  }
}

const actions = {
  // update user
  updateUserInfo({ commit, dispatch }, data) {
    // console.log('data===>', data)
    data = data.data
    return new Promise((resolve) => {
      const dto = data.userDto
      // console.log('dto==>', dto)
      setToken(dto.loginToken)
      commit('SET_LOGINTOKEN', dto.loginToken)
      commit('SET_USERNAME', dto.userName)
      commit('SET_NICKNAME', dto.nickName)
      commit('SET_NETWORKTYPE', dto.networkType)
      if (dto.roles && dto.roles.length > 0) {
        commit('SET_ROLES', dto.roles)
      } else {
        dto.roles = ['comm']
        commit('SET_ROLES', ['comm'])
      }
      commit('SET_PRIVILEGES', data.privileges)
      commit('SET_FUNCTIONS', data.functions)

      var user = {
        userName: dto.userName,
        nickName: dto.nickName,
        networkType: dto.networkType,
        loginToken: dto.loginToken
      }
      sessionStorage.setItem('user', JSON.stringify(user))

      dispatch('permission/generateRoutes', data.privileges, { root: true })

      resolve()
    })
  },

  // user login
  login({ dispatch }, data) {
    return new Promise((resolve, reject) => {
      login({ userName: data.userName, password: data.password }).then(response => {
        if (response.state === 0) {
          resolve(response)
          return
        }
        const data = response
        dispatch('updateUserInfo', data)

        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // loginByToken
  loginByToken({ commit, dispatch }, form) {
    return new Promise((resolve, reject) => {
      loginByToken({ userName: form.userName, token: form.token }).then(response => {
        if (response.state === 0) {
          // 清除Token
          commit('SET_LOGINTOKEN', '')
          commit('SET_ROLES', '')
          commit('SET_PRIVILEGES', '')
          commit('SET_FUNCTIONS', '')
          removeToken()
          delCookie('session_id')

          resolve(response)
          return
        }
        const data = response

        console.log('sessionId=====>>>>', data.data.sessionId)
        setCookie('session_id', data.data.sessionId)
        dispatch('updateUserInfo', data)

        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({ commit }) {
    return new Promise((resolve, reject) => {
      logout().then(() => {
        commit('SET_LOGINTOKEN', '')
        commit('SET_ROLES', '')
        commit('SET_PRIVILEGES', '')
        commit('SET_FUNCTIONS', '')
        removeToken()
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
