<!--
  - Copyright 2018-2020,上海哈里奥网络科技有限公司
  - All Right Reserved.
  -->
<template>
  <div class="app-container">
    <div class="filter-container">
      <!-- 1.添加秘钥 -->
      <el-button type="primary" class="filter-item" icon="el-icon-add" @click="clickAddKey">添加</el-button>
      <el-button type="primary" class="filter-item" icon="el-icon-add" @click="clickAddSSHKey">生成SSH秘钥</el-button>
      <el-dialog title="添加秘钥" :visible.sync="dialogKey" width="60%" :close-on-click-modal="false"
                 :before-close="handleClose">
        <el-form ref="form" :model="pojo" label-width="100px">
          <el-form-item label="钥匙名称">
            <el-input v-model="pojo.name" placeholder="请输入内容" />
          </el-form-item>
          <el-form-item label="用户名">
            <el-input v-model="pojo.user" placeholder="请输入内容" />
          </el-form-item>
          <el-form-item label="密码">
            <el-input v-model="pojo.password" show-password placeholder="请输入密码" />
          </el-form-item>
          <el-form-item label="确认密码">
            <el-input v-model="pojo.rePassword" show-password placeholder="请再次输入密码" />
          </el-form-item>
          <el-form-item label="描述">
            <el-input v-model="pojo.intro" placeholder="请输入内容" type="textarea" :autosize="{ minRows: 2, maxRows: 4}" />
          </el-form-item>

        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="handleCancel">取 消</el-button>
          <el-button type="primary" @click="handleKeySave()">确 定</el-button>
        </span>
      </el-dialog>

      <el-dialog title="添加SSH秘钥" :visible.sync="dialogSSHKey" width="60%" :close-on-click-modal="false"
                 :before-close="handleClose">
        <el-form ref="form" :model="pojo" label-width="100px">
          <el-form-item label="名称">
            <el-input v-model="pojo.name" placeholder="请输入ssh秘钥名称" />
          </el-form-item>
          <el-form-item label="邮箱">
            <el-input v-model="pojo.email" placeholder="请输入邮箱" />
          </el-form-item>
          <el-form-item label="描述">
            <el-input v-model="pojo.intro" placeholder="请输入描述" type="textarea" :autosize="{ minRows: 2, maxRows: 4}" />
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="handleCancel">取 消</el-button>
          <el-button type="primary" @click="handleAddSSHKey()">确 定</el-button>
        </span>
      </el-dialog>
    </div>

    <!-- 2. -->
    <el-table :data="keyList" :row-class-name="tableRowClassName">
      <el-table-column label="序号">
        <template scope="scope">
          {{ scope.$index+1 }}
        </template>
      </el-table-column>
      <el-table-column prop="key" label="钥匙名称" />
      <el-table-column prop="user" label="用户名" />
      <el-table-column prop="intro" label="描述" />

      <el-table-column fixed="right" label="操作" width="100">
        <template slot-scope="scope">
          <el-button v-if="scope.row.type !== 'ssh'" type="text" size="small" @click="clickDeleteKey(scope.row.id)">删除</el-button>
          <el-button v-if="scope.row.type === 'ssh'" type="text" size="small" @click="clickDeleteSSHKey(scope.row.id)">删除SSHKey</el-button>
          <el-button v-if="scope.row.type === 'ssh'" type="text" size="small" @click="downloadSSHKey(scope.row.id)">下载公钥</el-button>
        </template>
      </el-table-column>
    </el-table>

  </div>
</template>
<script>
import keyApi from '@/api/key'
import { isEmpty } from '@/utils/validate'

export default {
  data() {
    return {
      dialogKey: false,
      dialogSSHKey: false,
      pojo: {
        name: '',
        user: '',
        password: '',
        rePassword: '',
        intro: '',
        email: ''
      },
      keyList: []
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      keyApi.getList().then(res => {
        if (res.state === 1) {
          this.keyList = res.data
          console.log('this.keyList', this.keyList)
        }
      })
    },
    clickAddKey() {
      this.dialogKey = true
    },
    clickAddSSHKey() {
      this.dialogSSHKey = true
    },
    handleKeySave() {
      // 添加
      if (this.pojo.password !== this.pojo.rePassword) {
        this.$message({
          type: 'warning',
          message: '两次密码输入不一致'
        })
      } else {
        keyApi.create(this.pojo).then(res => {
          if (res.state === 1) {
            this.$message({
              type: 'success',
              message: '创建成功'
            })
            this.fetchData()
            this.dialogKey = false
            this.pojo = {
              name: '',
              user: '',
              password: '',
              rePassword: '',
              intro: ''
            }
          }
        })
      }
    },
    handleAddSSHKey() {
      // 添加
      if (isEmpty(this.pojo.email)) {
        this.$message({
          type: 'warning',
          message: '邮箱不能为空'
        })
        return
      }
      keyApi.createSSH(this.pojo).then(res => {
        if (res.state === 1) {
          this.$message({
            type: 'success',
            message: '创建成功'
          })
          this.fetchData()
          this.dialogSSHKey = false
          this.pojo = {}
        }
      })
    },
    clickDeleteKey(id) {
      if (!isEmpty(id)) {
        this.$confirm('确认删除吗', {
          confirmButtonText: '确定',
          cancelButtonText: '取消'
        }).then(() => {
          keyApi.delete(id).then(res => {
            if (res.state === 1) {
              this.$message({
                type: 'success',
                message: '删除成功'
              })
              this.fetchData()
            }
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
      }
    },
    clickDeleteSSHKey(id) {
      if (isEmpty(id)) {
        return
      }
      this.$confirm('确认删除吗', {
        confirmButtonText: '确定',
        cancelButtonText: '取消'
      }).then(() => {
        keyApi.deleteSSHKey(id).then(res => {
          if (res.state === 1) {
            this.$message({
              type: 'success',
              message: '删除成功'
            })
            this.fetchData()
          }
        })
      })
    },
    downloadSSHKey(id) {
      window.open(config.rootUrl + '/key/downloadSSHKey?id=' + id)
    },
    // ===>
    handleCancel() {
      this.dialogKey = false
      this.dialogSSHKey = false
      this.pojo = {}
    },
    handleClose(done) {
      this.$confirm('确认关闭？').then(_ => {
        this.pojo = {}
        done()
      }).catch(_ => {
      })
    },
    tableRowClassName({ row, rowIndex }) {
      if ((rowIndex % 4) === 1) {
        return 'warning-row'
      } else if ((rowIndex % 4) === 3) {
        return 'success-row'
      }
      return ''
    }
  }
}

</script>
<style>
  .el-table .warning-row {
    background: oldlace;
  }

  .el-table .success-row {
    background: #f0f9eb;
  }
</style>
