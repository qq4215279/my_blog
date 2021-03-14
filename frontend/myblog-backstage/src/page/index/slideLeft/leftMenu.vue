<template>
  <div class="left-menu" style="height: 100%">
    <el-scrollbar wrap-class="scrollbar-wrapper" style="height: 100%">
      <el-menu
        mode="vertical"
        background-color="#20222a"
        text-color="#bfcbd9"
        active-text-color="#409EFF"
        :default-active="activeItem.path"
      >
        <menu-item
          v-for="item in menu"
          :key="item.path"
          :item="item"
        />
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script>
import {mapGetters} from "vuex";
import menuItem from "./menuItem"
export default {
    data(){
        return {
            activeItem:{}
        }
    },
    components:{
        menuItem
    },
    computed:{
        ...mapGetters(["menu",'tag'])
    },
    watch:{
      tag(val) {
        this.setActive(val);
      },
    },
    methods:{
        setActive(tag){
            this.activeItem = tag
        }
    },
    created(){
        console.log("左侧菜单",this.menu)
        console.log("当前菜单",this.tag)
        this.setActive(this.tag)
    }
}
</script>

<style lang="scss">
.left-menu {
  .el-scrollbar {
    .el-scrollbar {
      overflow-y: scroll !important;
      // &__view{
      //     .menubar{
      //         background: $leftMenuBgColor;
      //     }

      // }
    }
  }
  .el-menu {
    div {
      background: $leftMenuBgColor !important;
      li {
        background: $leftMenuBgColor !important;
      }
    }

    border-right: none;
    text-align: left;
  }
}
</style>