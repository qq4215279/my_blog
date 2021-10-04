import request from '@/utils/request'

const PROJECT = 'project'

export default {

  // 查询项目
  getList() {
    return request({
      url: `/${PROJECT}/getList`,
      method: 'post'
    })
  },

  // 查询项目
  getProjectByID(id) {
    return request({
      url: `/${PROJECT}/getProject`,
      method: 'post',
      data: { id: id }
    })
  },

  // 创建项目
  create(pojo) {
    return request({
      url: `/${PROJECT}/create`,
      method: 'post',
      data: pojo
    })
  },

  // 修改项目
  update(pojo) {
    return request({
      url: `/${PROJECT}/update`,
      method: 'post',
      data: pojo
    })
  },

  // 删除项目
  delete(id) {
    return request({
      url: `/${PROJECT}/delete`,
      method: 'post',
      data: { id: id }
    })
  }
}

