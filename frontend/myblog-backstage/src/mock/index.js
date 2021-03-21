const Mock = require('mockjs');
import {login,getUserInfo} from './modules/user'

// console.log("开启mock监听")

Mock.setup({ timeout:"200-2000" })

Mock.mock('/api/user/login', 'post', login)
Mock.mock('/api/user/userInfo', 'get', getUserInfo)