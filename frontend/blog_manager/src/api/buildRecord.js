/*
 * Copyright 2018-2020,上海哈里奥网络科技有限公司
 * All Right Reserved.
 */

import request from '@/utils/request'

const BUILD_RECORD = 'buildRecord'

export default {

  // 获取最近构建记录
  getLastRecordList(projectId) {
    return request({
      url: `/${BUILD_RECORD}/getLastRecordList`,
      method: 'post',
      data: { projectId: projectId }
    })
  },

  // 获取构建记录
  getList(data) {
    return request({
      url: `/${BUILD_RECORD}/getList`,
      method: 'post',
      data: {
        projectId: data.projectId,
        projectBranchId: data.projectBranchId,
        buildFlowId: data.buildFlowId,
        user: data.user,
        dateArr: data.dateArr
      }
    })
  },

  // 创建构建记录
  startBuildJob(buildFlowId, projectBranchId, runTimeArgs) {
    return request({
      url: `/${BUILD_RECORD}/startBuildJob`,
      method: 'post',
      data: {
        buildFlowId: buildFlowId,
        projectBranchId: projectBranchId,
        runTimeArgs: runTimeArgs
      }
    })
  },

  // 停止创建构建记录
  stopBuildJob(recordId) {
    return request({
      url: `/${BUILD_RECORD}/stopBuildJob`,
      method: 'post',
      data: {
        id: recordId
      }
    })
  },

  // 读取日志文件
  readLog(buildRecordId) {
    return request({
      url: `/${BUILD_RECORD}/readLog`,
      method: 'get',
      data: {
        id: buildRecordId
      }
    })
  }

}
