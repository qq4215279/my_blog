<template>
  <div class="zf-layout-index">
    <!-- 头部导航 -->
      <div class="zf-layout-top">
        <top/>
      </div>
      <!-- 左侧导航 -->
      <div class="zf-layout-left">
        <slide-left/>
      </div>
      <div class="zf-layout-main">
        <div class="zf-layout-main-tags">
          <tags />
        </div>
        <div class="zf-layout-main-content" id="main-content">
          <!-- 主要视图区 -->
          <keep-alive>
            <router-view class="main-content" v-if="$route.meta.keepAlive"/>
          </keep-alive>
          <router-view class="main-content" v-if="!$route.meta.keepAlive"/>
        </div>
      </div>
  </div>
</template>
<script>
import top from "./top.vue"
import slideLeft from "./slideLeft/slideLeft.vue"
import tags from "./tags.vue"
export default {
  components:{
    top,
    slideLeft,
    tags
  },
}
</script>
<style lang="scss">
.zf-layout-index{
  width:100%;
  height:100%;
  .zf-layout-top{
    padding-left: 240px;
    width: 100%;
    height:60px;
    background-color: #fff;
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
  }
  .zf-layout-left{
    position: fixed;
    left: 0;
    top: 0;
    width: 240px;
    height: 100%;
    z-index: 1025;
    transition: all .3s;
  }
  .zf-layout-main{
    position: absolute;
    left: 240px;
    padding: 0;
    padding-bottom: 20px;
    width: calc(100% - 240px);
    height: calc(100% - 64px);
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    overflow: hidden;
    -webkit-transition: all 0.5s;
    transition: all 0.5s;
    background: #f0f2f5;
    z-index: 1026;
    display: flex;
    flex-direction: column;
    &-tags{
      height:40px;
      padding: 0 10px;
      margin-bottom: 10px;
      box-sizing: border-box;
      overflow: hidden;
      border-top: 1px solid $borderColor;
      background: $bgColor;
      color: $textColor;
      box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
    }
    &-content{
      flex:1;
    }
  }
}
</style>