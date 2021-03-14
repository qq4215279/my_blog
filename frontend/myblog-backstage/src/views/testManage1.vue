<template>
  <div>test<span @click="jump">跳转</span></div>
</template>

<script>
export default {
    methods:{
        jump(){
            this.$router.push({
                path:"/menuManage/index",
                query:{
                    a:{a:1}
                }
            })
        }
    }
}
</script>

<style>

</style>