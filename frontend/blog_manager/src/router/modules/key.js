import Layout from '@/layout'

const keyRouter = {
  path: '/key',
  component: Layout,
  redirect: '/key/index',
  name: '秘钥',
  hidden: true,
  meta: {
    title: '秘钥',
    icon: 'lock'
  },
  children: [
    {
      path: 'index',
      hidden: true,
      component: () => import('@/views/key/index'),
      name: 'index',
      meta: { title: '秘钥' }
    }
  ]
}
export default keyRouter
