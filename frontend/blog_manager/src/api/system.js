import request from '@/utils/request'

export function getCacheInfo() {
  return request({
    url: 'getCacheInfo.action',
    method: 'get'
  })
}

export function refreshDBCache() {
  return request({
    url: 'refreshDBCache.action',
    method: 'get'
  })
}

export function getEnvInfo() {
  return request({
    url: 'system/getEnvInfo',
    method: 'post'
  })
}

export function dicts() {
  return request({
    url: 'dicts.action',
    method: 'get'
  })
}

export function addDict(data) {
  return request({
    url: 'addDict.action',
    method: 'post',
    data: data
  })
}

export function updateDict(data) {
  return request({
    url: 'updateDict.action',
    method: 'post',
    data: data
  })
}

export function disableDict(data) {
  return request({
    url: 'disableDict.action',
    method: 'post',
    data: data
  })
}

export function jobs(params) {
  return request({
    url: 'jobs.action',
    method: 'post',
    data: params
  })
}

export function runJob(params) {
  return request({
    url: 'runJob.action',
    method: 'post',
    data: params
  })
}

export function reRunJob(params) {
  return request({
    url: 'reRunJob.action',
    method: 'post',
    data: params
  })
}

/**
 * 创建错误上报配置
 * @param params
 */
export function addReportError(params) {
  return request({
    url: 'addReportError.action',
    method: 'post',
    data: params
  })
}

/**
 * 删除错误上报配置
 * @param params
 */
export function removeReportError(params) {
  return request({
    url: 'removeReportError.action',
    method: 'post',
    data: params
  })
}

/**
 * 查询错误上报配置
 * @param params
 */
export function getReportError(params) {
  return request({
    url: 'getReportError.action',
    method: 'post',
    data: params
  })
}
