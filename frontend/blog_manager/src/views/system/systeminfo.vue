<template>
  <div class="app-container">
    <el-form label-width="100px" style="width:50%" v-loading="listLoading">
      <el-form-item label="系统版本:" class="el-table cell">
        <span>{{ sysInfo.sysVersion }}</span>
      </el-form-item>
      <el-form-item label="数据库版本:" class="el-table cell">
        <span>{{ sysInfo.dbVersion }}</span>
      </el-form-item>
      <el-form-item label="更新日志:" class="el-table cell">
      </el-form-item>
    </el-form>
    <div id="markdown-viewer" style="border:2px solid #dddddd; padding:5px"/>
  </div>
</template>

<script>
import { getEnvInfo } from '@/api/system'
import waves from '@/directive/waves' // waves directive
import MarkdownEditor from '@/components/MarkdownEditor'
import Editor from 'tui-editor'

export default {
  name: 'hotfix-cdn',
  directives: { waves },
  components: { MarkdownEditor },
  data() {
    return {
      listLoading: false,
      sysInfo: {
        sysVersion: '1.1.1',
        dbVersion: '1.1.1',
        updateInfo: '1.1.1\n-------\n1. 版本新增\n2. 版本新增'
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      getEnvInfo().then(response => {
        this.sysInfo.dbVersion = response.data.dbVersion
        this.sysInfo.sysVersion = response.data.sysVersion
        this.sysInfo.updateInfo = response.data.updateInfo
        this.listLoading = false
        const instance = Editor.factory({
          el: document.querySelector('#markdown-viewer'),
          viewer: true,
          height: '600px',
          initialValue: response.data.updateInfo
        });
      })
    }
  }
}
</script>