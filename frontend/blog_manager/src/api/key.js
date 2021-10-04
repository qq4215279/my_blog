/*
 * Copyright 2018-2020,上海哈里奥网络科技有限公司
 * All Right Reserved.
 */

import request from '@/utils/request'

const KEY = 'key'

export default {

  getList() {
    return request({
      url: `${KEY}/getList`,
      method: 'post'
    })
  },

  create(data) {
    return request({
      url: `${KEY}/create`,
      method: 'post',
      data: {
        name: data.name,
        user: data.user,
        password: data.password,
        intro: data.intro
      }
    })
  },

  createSSH(data) {
    return request({
      url: `${KEY}/createSSH`,
      method: 'post',
      data: {
        name: data.name,
        email: data.email,
        intro: data.intro
      }
    })
  },

  delete(id) {
    return request({
      url: `${KEY}/delete`,
      method: 'post',
      data: {
        id: id
      }
    })
  },
  deleteSSHKey(id) {
    return request({
      url: `${KEY}/deleteSSHKey`,
      method: 'post',
      data: {
        id: id
      }
    })
  },
  downloadSSHKey(id) {
    return request({
      url: `/${KEY}/downloadSSHKey`,
      method: 'get',
      data: {
        id: id
      }
    })
  }

}
