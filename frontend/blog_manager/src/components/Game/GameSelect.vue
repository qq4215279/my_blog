<template>
  <el-select
    v-model="owner.gameId"
    placeholder="选择游戏"
    style="width: 200px"
    class="filter-item"
    clearable
    @change="handleChange">
    <el-option v-for="game in games" :key="game.id" :label="game.name" :value="game.gameId" />
  </el-select>
</template>

<script>
export default {
  name: 'GameSelect',
  props: [
    'gameId',
    'games'
  ],
  data() {
    return {
      owner: {
        gameId: this.gameId
      }
    }
  },
  methods: {
    handleChange: function(value) {
      this.$emit('gameChanged', value)
    },
    handleValueChange: function(value) {
      this.owner.gameId = value
    }
  }
}
</script>
