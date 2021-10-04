import router from './router'
import store from './store'
// import { Message } from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import { getToken, getUser } from '@/utils/auth' // get token from cookie
import getPageTitle from '@/utils/get-page-title'

NProgress.configure({ showSpinner: false }) // NProgress Configuration

function hasPermission(roles, permissionRoles) {
  if (roles.indexOf('admin') >= 0) return true // admin permission passed directly
  if (!permissionRoles) return true
  return roles.some(role => permissionRoles.indexOf(role) >= 0)
}

const whiteList = ['/login'] // no redirect whitelist

router.beforeEach(async(to, from, next) => {
  // console.log('from', from)
  // console.log('to', to)
  // start progress bar
  NProgress.start()

  // set page title
  document.title = getPageTitle(to.meta.title)

  // determine whether the user has logged in
  const hasToken = getToken()
  // console.log('hasToken', hasToken)
  if (hasToken) {
    if (to.path === '/login') {
      // if is logged in, redirect to the home page
      if (config.loginType === 'sso') {
        self.location.href = config.ssoDomain + 'login?redirect=' + window.location.protocol + '//' + window.location.host + '/'
      } else if (config.loginType === 'ldap') {
        next(`/login?redirect=${to.path}`)
      }
      NProgress.done()
    } else {
      // console.log('needReload', store.getters.needReloadRouter, store.getters.roles.length)
      if (store.getters.needReloadRouter || store.getters.roles.length === 0) {
        if (store.getters.roles.length === 0) {
          // 重新登录
          const user = getUser()
          let reLoginSucc = false
          await store.dispatch('user/loginByToken', { userName: user.userName, token: user.loginToken }).then(data => {
            if (data.state === 0) {
              // 重登陆失败，重定向到主页
              if (config.loginType === 'sso') {
                self.location.href = config.ssoDomain + 'login?redirect=' + window.location.protocol + '//' + window.location.host + '/'
              } else if (config.loginType === 'ldap') {
                next(`/login?redirect=${to.path}`)
              }
              return
            }
            reLoginSucc = true
          })

          if (reLoginSucc) {
            // 重建router
            const accessRoutes = await store.dispatch('permission/generateRoutes', store.getters.privileges)

            // dynamically add accessible routes
            router.addRoutes(accessRoutes)
            store.dispatch('permission/routerRefreshDone') // 标记自己完成刷新
            next({ ...to, replace: true })
          }
          return
        }
        // 重建router
        const accessRoutes = await store.dispatch('permission/generateRoutes', store.getters.privileges)

        // dynamically add accessible routes
        router.addRoutes(accessRoutes)
        store.dispatch('permission/routerRefreshDone') // 标记自己完成刷新
        next({ ...to, replace: true })
      } else {
        // console.log('hasPermission', hasPermission(store.getters.roles, to.meta.roles))
        // 没有动态改变权限的需求可直接next() 删除下方权限判断 ↓
        if (hasPermission(store.getters.roles, to.meta.roles)) {
          next()
        } else {
          next({ path: '/401', replace: true, query: { noGoBack: true }})
        }
      }
    }
  } else {
    /* has no token*/
    if (whiteList.indexOf(to.path) !== -1) {
      // in the free login whitelist, go directly
      next()
    } else {
      // other pages that do not have permission to access are redirected to the login page.
      if (config.loginType === 'sso') {
        self.location.href = config.ssoDomain + 'login?redirect=' + window.location.protocol + '//' + window.location.host + '/'
      } else if (config.loginType === 'ldap') {
        next(`/login?redirect=${to.path}`)
      }
      NProgress.done()
    }
  }
})

router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})
