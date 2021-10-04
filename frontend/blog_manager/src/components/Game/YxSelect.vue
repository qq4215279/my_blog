<template>
  <el-select
    v-model="owner.yx"
    placeholder="选择联运"
    style="width: 200px"
    class="filter-item"
    clearable
    @change="handleChange">
    <el-option v-for="(item,key) of yxs" :key="key" :label="item.value" :value="item.value" />
  </el-select>
</template>

<script>
export default {
  name: 'YxSelect',
  props: [
    'yx',
    'yxs'
  ],
  data() {
    return {
      owner: {
        yx: this.yx
      }
    }
  },
  methods: {
    handleChange: function(value) {
      this.$emit('yxChanged', value)
    },
    handleValueChange: function(value) {
      this.owner.yx = value
    }
  }
}
</script>
