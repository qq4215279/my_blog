export function format(date, format) {
  // eslint-disable-next-line no-extend-native
  // (new Date()).Format('yyyy-MM-dd hh:mm:ss.S') ==> 2006-07-02 08:09:04.423
  // (new Date()).Format('yyyy-M-d h:m:s.S')      ==> 2006-7-2 8:9:4.18
  Date.prototype.Format = function(fmt) {
    const o = {
      // 月份
      'M+': this.getMonth() + 1,
      // 日
      'd+': this.getDate(),
      // 小时
      'h+': this.getHours(),
      // 分
      'm+': this.getMinutes(),
      // 秒
      's+': this.getSeconds(),
      // 季度
      'q+': Math.floor((this.getMonth() + 3) / 3),
      // 毫秒
      'S': this.getMilliseconds()
    }
    if (/(y+)/.test(fmt)) {
      fmt = fmt.replace(RegExp.$1, (this.getFullYear() + '').substr(4 - RegExp.$1.length))
    }
    for (var k in o) {
      if (new RegExp('(' + k + ')').test(fmt)) {
        fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? (o[k]) : (('00' + o[k]).substr(('' + o[k]).length)))
      }
    }
    return fmt
  }
  return new Date(date).Format(format)
}
