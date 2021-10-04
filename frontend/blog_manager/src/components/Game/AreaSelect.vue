<template>
  <el-select
    v-model="owner.areaId"
    placeholder="选择地区"
    style="width: 200px"
    class="filter-item"
    clearable
    @change="handleChange">
    <el-option v-for="(item,key) of areas" :key="key" :label="item.value" :value="item.key" />
  </el-select>
</template>

<script>
export default {
  name: 'AreaSelect',
  props: [
    'areaId',
    'areas'
  ],
  data() {
    return {
      owner: {
        areaId: this.areaId
      }
    }
  },
  methods: {
    handleChange: function(value) {
      this.$emit('areaChanged', value)
    },
    handleValueChange: function(value) {
      this.owner.areaId = value
    }
  }
}
</script>
