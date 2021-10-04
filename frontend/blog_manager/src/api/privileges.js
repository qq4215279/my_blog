import request from '@/utils/request'

// 查询个人信息
export function getSelfInfo() {
  return request({
    url: 'getSelfInfo.action',
    method: 'post'
  })
}

