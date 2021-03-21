import Vue from 'vue'
import Store from '../store/';
import myRouter from './myRouter'
import VueRouter from 'vue-router'
import Layout from '@/page/index/'

Vue.use(VueRouter)

//获取原型对象上的push函数
const originalPush = VueRouter.prototype.push
//修改原型对象中的push方法
VueRouter.prototype.push = function push(location) {
   return originalPush.call(this, location).catch(err => err)
}

const routes = [
  {
    path: '/',
    name: '主页',
    component: Layout,
    redirect: '/zhifou',
    children: [{
      path: '/zhifou',
      name: '首页',
      component: () => import(/* webpackChunkName: "about" */ '../views/Home.vue')
    }]
  },
  {
    path: '/login',
    name: '登录页',
    component: () =>
        import( /* webpackChunkName: "page" */ '@/page/login/index'),
    meta: {
        keepAlive: true,
        isTab: false,
        isAuth: false
    }
  },
  {
    path: '/404',
    component: () =>
        import( /* webpackChunkName: "page" */ '@/page/error/404'),
    name: '404',
    meta: {
        keepAlive: true,
        isTab: false,
        isAuth: false
    }
  },
  {
    path: '*',
    redirect: '/404'
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})
myRouter.init(router,Store)
console.log("菜单信息:",Store.state.user.menu)
router.$MyRouter.formatRoutes(Store.state.user.menu, true);
export default router
