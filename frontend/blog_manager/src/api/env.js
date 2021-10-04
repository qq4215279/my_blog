/*
 * Copyright 2018-2020,上海哈里奥网络科技有限公司
 * All Right Reserved.
 */

import request from '@/utils/request'

const ENV = 'env'

export default {

  getEnvType() {
    return request({
      url: `${ENV}/getEnvType`,
      method: 'post'
    })
  },

  getEnvSelect() {
    return request({
      url: `${ENV}/getEnvSelect`,
      method: 'post'
    })
  },

  getList() {
    return request({
      url: `${ENV}/getList`,
      method: 'post'
    })
  },

  create(data) {
    return request({
      url: `${ENV}/create`,
      method: 'post',
      data: {
        type: data.type,
        key: data.key,
        value: data.value,
        intro: data.intro
      }
    })
  },

  update(data) {
    return request({
      url: `${ENV}/update`,
      method: 'post',
      data: {
        type: data.type,
        key: data.key,
        value: data.value,
        intro: data.intro
      }
    })
  },

  delete(key) {
    return request({
      url: `${ENV}/delete`,
      method: 'post',
      data: {
        key: key
      }
    })
  }

}
