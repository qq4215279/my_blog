/**
 * 全站http配置
 *
 * axios参数说明
 * isToken是否需要token
 */
import axios from 'axios'
// import store from '@/store/';
// import router from '@/router/index'
import {getToken} from '@/util/auth'
import {Message} from 'element-ui'
import website from '@/config/website';
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style

const statusWhiteList = website.statusWhiteList || [];

let refreshLock = false;
let refreshToken = function(){return axios()};
let requestList = []

axios.defaults.timeout = 10000;
//返回其他状态吗
axios.defaults.validateStatus = function (status) {
  return status >= 200 && status <= 500; // 默认的
};
//跨域请求，允许保存cookie
axios.defaults.withCredentials = true;

//HTTPrequest拦截
axios.interceptors.request.use(config => {
  NProgress.start() // start progress bar
  const meta = (config.meta || {});
  const isToken = meta.isToken === false;
//   config.headers['Authorization'] = `Basic ${Base64.encode(`${website.clientId}:${website.clientSecret}`)}`;
  if (getToken() && !isToken) {
    config.headers['Blade-Auth'] = 'bearer ' + getToken() // 让每个请求携带token--['Authorization']为自定义key 请根据实际情况自行修改
  }
  return config
}, error => {
  return Promise.reject(error)
});
//HTTPresponse拦截
axios.interceptors.response.use(res => {
  NProgress.done();
  const status = res.data.code || 200
  const message = res.data.msg || '未知错误';
  const config = res.config;
  //如果在白名单里则自行catch逻辑处理
  if (statusWhiteList.includes(status)) return Promise.reject(res);
  //如果是401则无感刷新token
  if (status === 401) {
      //刷新中
      if(!refreshLock){
        refreshLock = true
        refreshToken().then((res)=>{
          const { token } = res.data
          requestList.forEach((fn)=>{
            fn(token)
          })
        })
        .catch(res => {
          //刷新失败
          console.error('refreshtoken error =>', res)
          window.location.href = '/'
        }).finally(() => {
          refreshLock = false
        })
      }
      return new Promise((resolve)=>{
        requestList.push((token)=>{
          config.baseURL = ''
          config.headers['Blade-Auth'] = token
          resolve(axios(config))
        })
      })
    }
  // 如果请求为非200否者默认统一处理
  if (status !== 200) {
    Message({
      message: message,
      type: 'error'
    })
    return Promise.reject(new Error(message))
  }
  return res;
}, error => {
  NProgress.done();
  return Promise.reject(new Error(error));
})

export default axios;
