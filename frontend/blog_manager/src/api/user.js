import request from '@/utils/request'

const USER = 'user'

// 创建用户
export function createUser(data) {
  return request({
    url: `/${USER}/create`,
    method: 'post',
    data: {
      userName: data.userName,
      nickName: data.nickName,
      userMail: data.userMail
    }
  })
}

export function login(data) {
  return request({
    url: `/${USER}/login`,
    method: 'post',
    data
  })
}

export function loginByToken(data) {
  return request({
    url: `/${USER}/loginByToken`,
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: `/${USER}/info`,
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: `/${USER}/Logout`,
    method: 'post'
  })
}

export function getRole() {
  return request({
    url: 'getRole.action',
    method: 'post'
  })
}

// 查询下属信息
export function getSubordinateList() {
  return request({
    url: `/${USER}/getSubAccountList`,
    method: 'post'
  })
}

// 禁用用户
export function forbiddenUser(disableUserName) {
  return request({
    url: `/${USER}/disableUser`,
    method: 'post',
    data: {
      disableUserName: disableUserName
    }
  })
}

// 启用用户
export function allowedUser(enableUserName) {
  return request({
    url: `/${USER}/enableUser`,
    method: 'post',
    data: {
      enableUserName: enableUserName
    }
  })
}

// 获取当前用户的权限
export function getPrivilegeList() {
  return request({
    url: `/${USER}/getPrivilegeList`,
    method: 'post'
  })
}

// 修改用户权限
export function changePrivileges(data) {
  return request({
    url: `/${USER}/changePrivileges`,
    method: 'post',
    data: data
  })
}

export function changeProjectPrivileges(userName, projectIds) {
  return request({
    url: `/${USER}/changeProjectPrivileges`,
    method: 'post',
    data: {
      userName: userName,
      projectIds: projectIds
    }
  })
}

// 获取高级权限列表
export function getHighPrivileges() {
  return request({
    url: `/${USER}/getHighPrivileges`,
    method: 'post'
  })
}

// 修改高级权限
export function changeHighPrivileges(userName, privileges) {
  return request({
    url: `/${USER}/changeHighPrivileges`,
    method: 'post',
    data: {
      userName: userName,
      privileges: privileges
    }
  })
}
