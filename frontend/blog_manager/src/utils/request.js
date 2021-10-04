import axios from 'axios'
import { Message } from 'element-ui'
// import store from '@/store'
import { getCookie } from '@/utils/cookie'
import router from '@/router'

// create an axios instance
const service = axios.create({
  baseURL: config.rootUrl, // url = base url + request url
  // baseURL: '/root', // url = base url + request url
  withCredentials: true // send cookies when cross-domain requests
})

// request interceptor
service.interceptors.request.use(
  config => {
    let newData = ''
    for (const k in config.data) {
      newData += encodeURIComponent(k) + '=' + encodeURIComponent(config.data[k]) + '&'
    }
    newData += 'csrftoken=' + getCookie('csrftoken')
    config.data = newData
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
   */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data

    // if the custom code is not 20000, it is judged as an error.
    if (res.state === 0) {
      Message({
        message: res.msg || '未知错误',
        type: 'error',
        dangerouslyUseHTMLString: true
        // duration: 5 * 1000
      })

      return res
    } else if (res.state === 2) {
      Message({
        message: res.msg || '未知错误',
        type: 'error',
        duration: 2 * 1000
      })
    } else if (res.state === 4) {
      Message({
        message: res.msg || '未知错误',
        type: 'error',
        duration: 2 * 1000
      })
      setTimeout(function() {
        router.push({ path: '/login' })
      }, 2000)

      return res
    } else {
      return res
    }
  },
  error => {
    console.log('err' + error) // for debug
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
