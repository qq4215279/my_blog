<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" :loading="reloading" @click="handleRefreshDBCache">
        刷新数据库缓存
      </el-button>
    </div>

    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 20%;"
    >
      <el-table-column label="数据库版本" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.cacheVersion }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getCacheInfo, refreshDBCache } from '@/api/system'
import waves from '@/directive/waves' // waves directive
import { formatFileSize } from '@/utils'

export default {
  name: 'hotfix-cdn',
  directives: { waves },
  filters: {
    // 格式化包类型
    formatPackType(type) {
      if (type == 0) {
        return '整包'
      } else {
        return '小包'
      }
    },
    // 格式化文件大小
    formatFileSize(fileSize) {
      return formatFileSize(fileSize)
    }
  },
  data() {
    return {
      tableKey: 0,
      cacheVersion: '',
      list: [],
      listLoading: false,
      reloading: false
    }
  },
  created() {
    this.getInfo()
  },
  methods: {
    getInfo() {
      this.listLoading = true
      getCacheInfo().then(response => {
        this.list = []
        this.list.push({ cacheVersion: response.data.cacheVersion })
        this.listLoading = false
      })
    },
    handleRefreshDBCache() {
      this.reloading = true
      refreshDBCache().then(response => {
        if (response.state === 1) {
          this.getInfo()
          this.reloading = false
        }
      })
    }
  }
}
</script>
