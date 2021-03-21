<template>
  <div class="layout-left-slide">
        <div class="zf-logo">
            <span>
                <span class="zf-logo-title">
                    <div style="height:100%">
                        <img src="@/assets/logo.png" width="64" style="vertical-align: middle;">
                        <span >知否后台管理</span>
                    </div>
                    
                </span>
            </span>
        </div>
        <div class="left-menu-wrapper">
            <left-menu></left-menu>
        </div>
  </div>
</template>

<script>
import leftMenu from "./leftMenu"
export default {
    components:{
        leftMenu
    },
    // created(){
    //     console.log("左侧组件被创建")
    // },
    // mounted(){
    //     console.log("左侧组件被挂载")
    // },
    // beforeRouteEnter(){
    //     console.log("左侧进入路由")
    // },
    // updated(){
    //     console.log("左侧组件更新")
    // }
}
</script>

<style lang="scss">
.layout-left-slide{
    display: flex;
    flex-direction: column;
    height:100%;
    background: $leftMenuBgColor;
    -webkit-transition: width .2s;
    transition: width .2s;
    box-sizing: border-box;
    box-shadow: 2px 0 6px rgba(0, 21, 41, 0.35);
    .zf-logo{
        background: $leftMenuBgColor;
        width: 240px;
        height: 64px;
        line-height: 64px;
        font-size: 20px;
        overflow: hidden;
        box-sizing: border-box;
        box-shadow: 0 1px 2px 0 $leftBorderColor;
        // box-shadow: 0 1px 2px 0 $borderColor;
        color: rgba(255, 255, 255, 0.8);
        z-index: 1024;
        &-title{
            display: block;
            height:100%;
            text-align: center;
            font-weight: 300;
            font-size: 25px;
            line-height: 64px;
        }
    }
    .left-menu-wrapper{
        flex:1
    }
}
</style>