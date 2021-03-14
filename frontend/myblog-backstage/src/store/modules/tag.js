import { setStore, getStore } from '@/util/store'
import website from '@/config/website'

const tagWel = website.fistPage;

//查找数组指定路径
function findTags(path, list) {
    const result = list.find((ele) => {
        return ele.path == path
    })
    if (result) return true
    return false
}

const navs = {
    state: {
        tagList: getStore({ name: 'tagList' }) || [tagWel],
        tag: getStore({ name: 'tag' }) || tagWel,
        tagWel: tagWel
    },
    mutations: {
        ADD_TAG: (state, action) => {
            state.tag = action;
            setStore({ name: 'tag', content: state.tag,type: 'session'})
            //如果存在当前标签则返回
            if (findTags(state.tag.path, state.tagList)) return
            //不存在则添加当前标签
            state.tagList.push(action)
            setStore({ name: 'tagList', content: state.tagList, type: 'session' })
        },
        DEL_TAG: (state, action) => {
            state.tagList = state.tagList.filter(item => {
                return !(item.path == action.path);
            })
            setStore({ name: 'tagList', content: state.tagList, type: 'session' })
        },
        DEL_ALL_TAG: (state) => {
            state.tagList = [state.tagWel];
            setStore({ name: 'tagList', content: state.tagList, type: 'session' })
        },
        DEL_TAG_OTHER: (state) => {
            state.tagList = state.tagList.filter(item => {
                if (item.path === state.tag.path) {
                    return true;
                } else if (!website.isFirstPage && item.path === website.fistPage.path) {
                    return true;
                }
            })
            setStore({ name: 'tagList', content: state.tagList, type: 'session' })
        },
        SET_TAG_LIST(state, tagList) {
            state.tagList = tagList;
            setStore({ name: 'tagList', content: state.tagList, type: 'session' })
        },
        SET_TAG: (state, action) => {
            console.log("设置的tag值",action)
            state.tag = action;
        }
    }
}
export default navs