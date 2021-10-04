<template>
  <el-select
    v-model="owner.serverId"
    placeholder="选择服"
    style="width: 200px"
    class="filter-item"
    clearable
    @change="handleChange">
    <el-option v-for="(item,key) of servers" :key="key" :label="item.value" :value="item.key" />
  </el-select>
</template>

<script>
export default {
  name: 'ServerSelect',
  props: [
    'serverId',
    'servers'
  ],
  data() {
    return {
      owner: {
        serverId: this.serverId
      }
    }
  },
  methods: {
    handleChange: function(value) {
      this.$emit('serverChanged', value)
    },
    handleValueChange: function(value) {
      this.owner.serverId = value
    }
  }
}
</script>
