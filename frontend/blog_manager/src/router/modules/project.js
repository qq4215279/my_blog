import Layout from '@/layout'

const projectRouter = {
  path: '/project',
  component: Layout,
  redirect: '/project/index',
  name: '我的项目',
  hidden: true,
  meta: {
    title: '我的项目',
    icon: 'project'
  },
  props: true,
  children: [
    {
      props: true,
      path: 'index',
      hidden: true,
      component: () => import('@/views/project/index'),
      name: 'index',
      meta: { title: '我的项目' }
    }
  ]
}
export default projectRouter
