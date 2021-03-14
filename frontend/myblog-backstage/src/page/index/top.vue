<template>
  <div class="layout-top">
      <div class="layout-top-left">
          <div>
              <i class="el-icon-s-fold"></i>
          </div>
      </div>
      <div class="layout-top-title">
          <div class="top-menu">
              <div 
                class="top-menu-item"
                v-for="item in topMenu"
                :key="item.id"
                @click="getMenu(item)"
              >
                {{item.label}}
              </div>
          </div>
      </div>
      <div class="layout-top-right">
            <div class="layout-top-right-theme">
                <theme-cat :theme="themeValue" />
                <el-switch
                    class="layout-top-right-theme-switch"
                    v-model="themeValue"
                >
                </el-switch>
            </div>
        <!-- 用户信息 -->
        <div class="layout-top-right-user">
            <p>
                <img src="/img/images/logo.png" />
            </p>
            <el-dropdown class="layout-top-user-drop">
                <span class="el-dropdown-link">
                    {{ userInfo.userName}}<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item class="user-info">
                        <img src="https://www.baidu.com/link?url=c0wzItscxg8jFf0vb0ctAU2r57a-FJQjDXaNLtS9-8THBrI3GMNY1Eo4vVIox8HbszsHSIMLsabWqIrVQd7YXUU8fsm4dShKsIw7j-Mf6hSch97Hpwk1SuzO0q77zkE1U6A6vEmp9_HqOClL6QEO4gY8CoSUQ1M6V88ir26WItY2m2bJPiQD05yUVH6IawxUpr9A5DL40RSxi9EumiScY-J-6FS7uY7TwVJXmTQqO_m4qjmCsP1rRSTTnVxQEVah-z-hicEKPzPfYOyO40fwva-dMxtlNiwns7xvrgPX5PBU7bRbW1vVo1fQB9bMkjHghF7KvIMhoSmTB72ZHUw--M0ebhpd4Exo6hcdCxRZ7F0vAx5hGv0iIqEO-dELvK3iTbpHNFVslnwEKll2Pl7W6TbAAsMkP0fqTSVveWkLwXlH0mfXU4dY9QACXmCAUJoBjj6hoimxxm-AWesBbkwAm1IZHL6G3LWMP3TjRTs6QMEBDvvZTi2P4HfWDC8niSgoogz_-cybObkNyCOJ91g9GMIIR_r9UkSh2079_o79BCW6dCGHwoLg790TZfPNDWvm4nuFtBXwHCSrpaxgKFG3qWraIzebb1Lfyb-FyeAHj9g-oumE0v9hTW74kD41G0arCz2A5v9klxNOIHvQV0HnJa&click_t=1611568724019&s_info=1903_977&wd=&eqid=db34f6870001b5ef00000005600e964f" />
                        <div class="user-info-text">
                        <div class="user-info-text-name">{{ userInfo.userName }}</div>
                        <!-- <div class="user-info-text-website">welife-visitor</div> -->
                        </div>
                    </el-dropdown-item>
                    <el-dropdown-item class="link-about">
                        <div class="right-btn-link-block login-out-block" @click="loginOut">
                        <span>关于系统</span>
                        </div>
                    </el-dropdown-item>
                    <el-dropdown-item class="login-out">
                        <div class="right-btn-link-block login-out-block" @click="loginOut">
                        <span>退出登录</span>
                        </div>
                    </el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </div>
      </div>
  </div>
</template>

<script>
import {mapGetters} from "vuex";
import changeTheme from "@/util/publicStyle.js"
import themeCat from "@/components/colorCat.vue"
export default {
    data(){
        return{
            themeValue:false,
            username:""
        }
    },
    computed:{
        ...mapGetters(['topMenu',"userInfo"])
    },
    mounted(){
        // console.log(this.userInfo)
    },
    methods:{
        getMenu(item){
            this.$store.dispatch("GetMenu", item.id)
        },
        openMenu(){
            let item = this.menu[0].children[0]
            this.$router.push({name:item.name})
        },
        loginOut(){}
    },
    watch:{
        themeValue(val){
            let theme = val ? "dark" : "white";
            changeTheme(theme)
        }
    },
    components:{
        themeCat
    }
}
</script>

<style lang="scss">
.layout-top{
    padding:0 20px;
    height:100%;
    display: flex;
    background: $bgColor;
    color:$textColor;
    &-left{
        i{
            line-height: 64px;
            font-size: 25px !important;
        }
    }
    &-title{
        flex:1;
        .top-menu{
            display: flex;
            &-item{
                line-height: 64px;
                padding: 0 10px;
                color:$textColor;
                cursor: pointer;
            }
        }
    }
    &-right{
        width: 400px;
        display: flex;
        align-items: center;
        justify-content: space-between;
        &-theme{
            position: relative;
            padding-right:20px;
            &-switch{
                width:30px !important;
                height:15px !important;
                position: absolute !important;
                top:calc(50% - 15px);
                right:5px;
                .el-switch__core{
                    border:3px solid #d1d5da !important;
                    border-color: #d1d5da !important;
                    background-color: #ffffff !important;
                    width:30px !important;
                    height:15px !important;
                    &::after{
                        top: -6px;
                        left: -3px;
                        width: 20px;
                        height: 20px;
                        background: black ;
                    }
                }     
            }
            .is-checked{
                .el-switch__core{
                    border:3px solid #3c1e70 !important;
                    border-color: #3c1e70 !important;
                    background-color: #271052 !important;
                    &::after{
                        background:#6e40c9;
                    }
                }
            }  
        }
        &-user{
            // margin-left:20px;
            display: flex;
            align-items: center;
            p{
                margin-right:10px;
            }
        }
    }
}
.el-dropdown-menu{
    .user-info {
        display: flex;
        justify-content: space-around;
        align-items: center;
        height: 50px;
        border-bottom: 1px solid rgba(255, 255, 255, 0.1);
        box-sizing: border-box;
        img {
            width: 48px;
            height: 48px;
            border-radius: 50%;
        }
        &-text {
            &-name {
                font-size: 14px;
                font-family: PingFangSC-Regular, PingFang SC;
                font-weight: 400;
                line-height: 17px;
                white-space: nowrap;
                overflow: hidden;
                /* line-break: anywhere; */
                text-overflow: ellipsis;
                max-width: 100px;
            }
        }
    }
}
</style>