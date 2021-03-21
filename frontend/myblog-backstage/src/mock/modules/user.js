const menu = [
    {
        label: '菜单管理',
        path: '/menuManage',
        name:'menuManage',
        icon: 'icon',
        children: []
    },
    {
        label: '文章管理',
        path: '/article',
        name:'article',
        icon: 'icon',
        children: []
    },
    {
        label: '用户管理',
        path: '/userManage',
        name:'userManage',
        icon: 'icon',
        children: []
    },
    {
        label: '测试管理',
        path: '/testManage',
        name:'testManage',
        icon: 'icon',
        children: [{
            label: '测试子管理',
            path: '/testManage/testManage1',
            name:'testManage1',
            icon: 'icon',
            children: []
        }]
    },
]
const topMenu = [
    {
        label: '移动端',
        path: '/yidong',
        icon: 'icon',
        children: []
    },
    {
        label: 'pc端',
        path: '/pc',
        icon: 'icon',
        children: []
    },
]
const user = {
    accessToken:""
};
export function login (option) {
    let params = JSON.parse(option.body);
    user.accessToken = `accessToken-${params.account}`;
    // console.log("登录信息返回==========>")
    const resp = {
        code:200,
        data:{
            accessToken:`accessToken-${params.account}`,
            userInfo:{
                account:params.account,
                userName:params.account,
                menu,
                topMenu,
            }
        }
    }
    return resp 
}
export function getUserInfo () {
    // let params = JSON.parse(option.body);
    const resp = {
        code:200,
        data:{
            account: user.accessToken.split('-')[1],
            userName:user.accessToken.split('-')[1],
            menu,
            topMenu,
            id:"1664005277"
        }
    }
    console.log("用户信息返回==========>",resp)
    return resp 
}