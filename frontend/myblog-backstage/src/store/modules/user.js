import { loginByUsername, logout, getUserInfo } from '@/api/user'
import {setToken, removeToken, getToken} from '@/util/auth'
import {setStore, getStore} from '@/util/store'
import router from '../../router'

function addPath(amenu){
    //处理菜单数组
    return amenu
}
const user = {
    state:{
        token:getToken(),
        userInfo:getStore({name: 'userInfo'}),
        menu:getStore({name: 'menu'}),
        topMenu:getStore({name: 'topMenu'})
    },
    actions:{
        //根据用户名登录
        LoginByUsername({commit}, userInfo) {
            console.log(456)
            return new Promise((resolve, reject) => {
                loginByUsername(userInfo.username, userInfo.password).then(res => {
                    console.log(789)
                    console.log(res)
                    const data = res.data.data;
                    commit('SET_TOKEN', data.userDto.loginToken);
                    commit('SET_USER_INFO', data.userDto);
                    //----根据需求是否需要-------
                    commit('SET_MENU', data.privileges);
                    commit('SET_TOP_MENU', data.userDto.topMenu || []);
                    //--------------------------
                    // commit('DEL_ALL_TAG');
                    console.log(res)
                    resolve(res);
                }).catch(error => {
                    reject(error);
                })
            })
        },
        // 登出
        LogOut({commit}) {
            return new Promise((resolve, reject) => {
                logout().then(() => {
                    commit('SET_TOKEN', '');
                    commit('SET_MENU', []);
                    commit('SET_TOP_MENU', []);
                    commit('SET_USER_INFO',{})
                    // commit('DEL_ALL_TAG');
                    removeToken()
                    resolve()
                }).catch(error => {
                    reject(error)
                })
            })
        },
        //根据token值更新用户信息
        GetUserInfo({ commit }) {
            return new Promise((resolve, reject) => {
                getUserInfo().then((res) => {
                    const data = res.data.data;
                    console.log("用户信息",res)
                    commit('SET_USER_INFO', data);
                    commit('SET_MENU', data.menu);
                    commit('SET_TOP_MENU', data.topMenu);
                    
                    resolve(data);
                }).catch(err => {
                    reject(err);
                })
            })
        },
    },
    mutations:{
        //设置token值
        SET_TOKEN: (state, token) => {
            setToken(token)
            state.token = token;
        },
        SET_USER_INFO: (state, userInfo) => {
            state.userInfo = userInfo;
            setStore({name: 'userInfo', content: state.userInfo, type: 'session'})
        },
        SET_MENU: (state, menu) => {
            state.menu = addPath(menu)
            console.log("注册菜单")
            setStore({name: 'menu', content: state.menu, type: 'session'})
            router.$MyRouter.formatRoutes(state.menu,true)
        },
        SET_TOP_MENU: (state, data) => {
            state.topMenu = data
            setStore({name: 'topMenu', content: state.topMenu, type: 'session'})
        },
    }
}

export default user