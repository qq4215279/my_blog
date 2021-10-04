import Layout from '@/layout'

const buildRecordRouter = {
  path: '/buildRecord',
  component: Layout,
  redirect: '/buildRecord/index',
  name: '打包记录',
  hidden: true,
  meta: {
    title: '打包记录查询',
    icon: 'search'
  },
  children: [
    {
      path: 'index',
      hidden: true,
      component: () => import('@/views/buildRecord/index'),
      name: '打包记录查询首页',
      meta: { title: '打包记录查询' }
    }
  ]
}
export default buildRecordRouter
