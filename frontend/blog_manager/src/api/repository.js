import request from '@/utils/request'

const REPOSITORY = 'repository'

export default {

  // 获取插件信息
  getPlugins() {
    return request({
      url: `/plugin/getList`,
      method: 'post'
    })
  },

  // 获取分支下所有仓库
  getList(projectBranchId) {
    return request({
      url: `/${REPOSITORY}/getList`,
      method: 'post',
      data: { projectBranchId: projectBranchId }
    })
  },

  create(projectBranchId, repositories, content) {
    return request({
      url: `/${REPOSITORY}/create`,
      method: 'post',
      data: {
        projectBranchId: projectBranchId,
        name: repositories.name,
        intro: repositories.intro,
        type: repositories.type,
        content: content
      }
    })
  },

  update(id, name, intro, content) {
    return request({
      url: `/${REPOSITORY}/update`,
      method: 'post',
      data: {
        id: id,
        name: name,
        intro: intro,
        content: content
      }
    })
  },

  delete(repositoryId) {
    return request({
      url: `/${REPOSITORY}/delete`,
      method: 'post',
      data: {
        id: repositoryId
      }
    })
  }

}
