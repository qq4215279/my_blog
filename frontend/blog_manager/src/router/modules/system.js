/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const systemRouter = {
  path: '/system',
  component: Layout,
  redirect: '/system/refreshdbcache',
  name: '系统管理',
  hidden: true,
  meta: {
    title: '系统管理',
    icon: 'setting'
  },
  children: [
    {
      path: 'refreshdbcache',
      hidden: true,
      component: () => import('@/views/system/refreshdb'),
      name: '刷新DB缓存',
      meta: { title: '刷新DB缓存' }
    },
    {
      path: 'systeminfo',
      hidden: true,
      component: () => import('@/views/system/systeminfo'),
      name: '系统信息',
      meta: { title: '系统信息' }
    }
  ]
}
export default systemRouter
