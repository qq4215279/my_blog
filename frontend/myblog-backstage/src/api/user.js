import request from '../router/request';

export const loginByUsername = (account, password) => request({
    url: '/api/root/login',
    method: 'post',
    params: {
      userName:account,
      password
    }
});

export const logout = () => request({
  url: '/api/user/loginout',
  method: 'get',
});


export const getUserInfo = () => request({
  url: '/api/user/userInfo',
  method: 'get',
});
