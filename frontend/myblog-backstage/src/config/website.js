/**
 * 全局配置文件
 */
export default{
    //网站关键字
    key:"zf",
    accessTokenKey:"zf-access-token",
    //是否刷新token
    ifRefreshToken:false,
    //配置菜单的属性
    menu: {
        iconDefault: 'iconfont icon-caidan',
        props: {
            label: 'name',
            path: 'path',
            icon: 'icon',
            children: 'children'
        }
    },
    //开发环境依赖
    dev:{
        //是否开启mock
        mock:false,
        //mock基础路径
        mockBaseUrl:"api"
    },
    //首页配置
    fistPage: {
        label: "首页",
        path: "/zhifou",
        name:"首页",
        params: {},
        query: {},
        meta: {}
    },
}