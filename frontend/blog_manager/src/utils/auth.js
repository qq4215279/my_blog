import { setCookie, getCookie, delCookie } from './cookie'

const TokenKey = 'Login-Token-Key'

export function getToken() {
  return getCookie(TokenKey)
}

export function setToken(token) {
  return setCookie(TokenKey, token)
}

export function removeToken() {
  return delCookie(TokenKey)
}

export function getUser() {
  var user = getCookie('op_account_info')
  if (user != null) {
    user = JSON.parse(user)
    return user
  }
  return {}
}

