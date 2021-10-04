import Layout from '@/layout'

const buildFlowRouter = {
  path: '/buildFlow',
  component: Layout,
  redirect: '/buildFlow/index',
  name: '构建流程',
  hidden: true,
  meta: {
    title: '构建流程',
    icon: 'setting'
  },
  children: [
    {
      path: 'index',
      hidden: true,
      component: () => import('@/views/buildFlow/index'),
      name: 'index',
      meta: { title: '构建流程' }
    }
  ]
}
export default buildFlowRouter
