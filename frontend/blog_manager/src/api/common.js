import request from '@/utils/request'

// 查询部门列表
export function formatMoney() {
  return request({
    url: 'formatMoney.action',
    method: 'post'
  })
}
