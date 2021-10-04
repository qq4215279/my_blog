import Layout from '@/layout'

const environmentRouter = {
  path: '/env',
  component: Layout,
  redirect: '/env/index',
  name: '环境变量',
  hidden: true,
  meta: {
    title: '环境变量',
    icon: 'environment'
  },
  children: [
    {
      path: 'index',
      hidden: true,
      component: () => import('@/views/environment/index'),
      name: 'index',
      meta: { title: '环境变量' }
    }
  ]
}
export default environmentRouter
