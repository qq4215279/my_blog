/**
 * 全站路由配置
 *
 * meta参数说明
 * keepAlive是否缓冲页面
 * isTab是否加入到tag导航
 * isAuth是否需要授权
 */

let myRouter = function(){
    this.$router = null;
    this.$store = null;
}

myRouter.init = function (router, store) {
    //获取并将路由和仓库挂载到自定义路由上面
    this.$router = router;
    this.$store = store;
    this.$router.$MyRouter = {
        $website: this.$store.getters.website,
        routerRef: this,
        meta:{},
        formatRoutes(aMenu = [], first){
            if(!aMenu || aMenu.length === 0) return
            const actionRouter = [];
            const propsConfig = this.$website.menu.props;
            const propsDefault = {
                label: propsConfig.label || 'name',
                path: propsConfig.path || 'path',
                icon: propsConfig.icon || 'icon',
                children: propsConfig.children || 'children',
                meta: propsConfig.meta || 'meta',
            };
            for (let i = 0; i < aMenu.length; i++) {
                const oMenu = aMenu[i];
                let path = oMenu[propsDefault.path],
                    name = oMenu[propsDefault.label],
                    icon = oMenu[propsDefault.icon],
                    children = oMenu[propsDefault.children],
                    meta = oMenu[propsDefault.meta] || {},

                    component = "views" + path,
                    hasChild = children.length !== 0;

                let oRouter = {
                    path:path,
                    component(resolve){
                        if(first){
                            require(['../page/index/index'], resolve)
                            return
                        }
                        else if(!first && hasChild){
                            require(['../page/index/layout'], resolve)
                            return
                        }
                        else{
                            require([`../${component}.vue`], resolve)
                            return
                        }
                    },
                    name,
                    icon,
                    meta,
                    redirect: (() => {
                        if (!hasChild && first) return `${path}/index`
                        else return '';
                    })(),
                    children: !hasChild ? (() => {
                        if (first) {
                            return [{
                                component(resolve) { require([`../${component}.vue`], resolve) },
                                icon: icon,
                                name: `${name}Child`,
                                meta: meta,
                                path: 'index'
                            }]
                        }
                        return [];
                    })() : (() => {
                        return this.formatRoutes(children, false)
                    })()
                };
                actionRouter.push(oRouter)
            }
            if(first){
                this.routerRef.$router.addRoutes(actionRouter)
            }else{
                return actionRouter
            }
        }
    }
}

export default myRouter