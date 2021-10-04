/*
 * Copyright 2018-2020,上海哈里奥网络科技有限公司
 * All Right Reserved.
 */
import request from '@/utils/request'

const BRANCH = 'branch'

export default {

  // 打分支
  incrementVersion(projectBranchId, version) { // 分支id
    return request({
      url: `/${BRANCH}/incrementVersion`,
      method: 'post',
      data: {
        projectBranchId: projectBranchId,
        version: version
      }
    })
  },

  // 添加版本
  createBranch(projectBranchId, version) { // 分支id
    return request({
      url: `/${BRANCH}/createBranch`,
      method: 'post',
      data: {
        projectBranchId: projectBranchId,
        version: version
      }
    })
  },

  delete(projectBranchId) {
    return request({
      url: `/${BRANCH}/delete`,
      method: 'post',
      data: {
        projectBranchId: projectBranchId
      }
    })
  },

  // 更新子版本号
  updateSubVersion(projectBranchId, subVersion) {
    return request({
      url: `/${BRANCH}/update`,
      method: 'post',
      data: {
        projectBranchId: projectBranchId,
        subVersion: subVersion
      }
    })
  },

  star(projectBranchId) {
    return request({
      url: `/${BRANCH}/star`,
      method: 'post',
      data: {
        projectBranchId: projectBranchId
      }
    })
  },

  unStar(projectBranchId) {
    return request({
      url: `/${BRANCH}/unStar`,
      method: 'post',
      data: {
        projectBranchId: projectBranchId
      }
    })
  }

}
