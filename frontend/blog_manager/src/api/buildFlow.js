import request from '@/utils/request'

const BUILD_FLOW = 'buildFlow'

export default {

  // 获取构建流程
  getList(projectId, projectBranchId) {
    return request({
      url: `/${BUILD_FLOW}/getList`,
      method: 'post',
      data: {
        projectId: projectId,
        projectBranchId: projectBranchId
      }
    })
  },
  // 获取构建流程
  copy(data) {
    return request({
      url: `/${BUILD_FLOW}/copy`,
      method: 'post',
      data: data
    })
  },

  create(projectBranchId, flow, name, intro, concurrent, updateSource) {
    return request({
      url: `/${BUILD_FLOW}/create`,
      method: 'post',
      data: {
        projectBranchId: projectBranchId,
        flow: flow,
        name: name,
        intro: intro,
        concurrent: concurrent,
        updateSource: updateSource
      }
    })
  },

  update(id, projectBranchId, flow, name, intro, concurrent, updateSource) {
    return request({
      url: `/${BUILD_FLOW}/update`,
      method: 'post',
      data: {
        id: id,
        projectBranchId: projectBranchId,
        flow: flow,
        name: name,
        intro: intro,
        concurrent: concurrent,
        updateSource: updateSource
      }
    })
  },

  delete(id, updateSource) {
    return request({
      url: `/${BUILD_FLOW}/delete`,
      method: 'post',
      data: {
        id: id,
        updateSource: updateSource
      }
    })
  }

}
