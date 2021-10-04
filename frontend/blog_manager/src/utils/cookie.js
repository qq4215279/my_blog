export function setCookie(name, value) {
  const Days = 10
  const exp = new Date()
  exp.setTime(exp.getTime() + Days * 24 * 60 * 60 * 1000)
  document.cookie = name + '=' + escape(value) + ';expires=' + exp.toGMTString() + ';domain=' + config.baseDomain
}

export function getCookie(name) {
  let arr = null
  const reg = new RegExp('(^| )' + name + '=([^;]*)(;|$)')
  arr = document.cookie.match(reg)
  if (arr != null) { return unescape(arr[2]) } else { return null }
}

export function delCookie(name) {
  const exp = new Date()
  exp.setTime(exp.getTime() - 1)
  const cval = getCookie(name)
  if (cval != null) {
    document.cookie = name + '=' + cval + ';expires=' + exp.toGMTString() + ';domain=' + config.baseDomain
  }
}
