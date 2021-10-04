/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const userRouter = {
  path: '/user',
  component: Layout,
  redirect: '/user/usermanager',
  name: '用户管理',
  hidden: true,
  meta: {
    title: '用户管理',
    icon: 'user'
  },
  children: [
    {
      path: 'usermanager',
      hidden: true,
      component: () => import('@/views/user/usermanager'),
      name: '用户管理',
      meta: { title: '用户管理' }
    },
    {
      path: 'selfmanager',
      hidden: true,
      component: () => import('@/views/user/selfmanager'),
      name: '个人中心',
      meta: { title: '个人中心' }
    }
  ]
}
export default userRouter
