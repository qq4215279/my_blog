<!--
  - Copyright 2018-2020,上海哈里奥网络科技有限公司
  - All Right Reserved.
  -->
<template>
  <div class="app-container">
    <!-- 1.添加环境变量 -->
    <div class="filter-container">
      <el-button type="primary" class="filter-item" icon="el-icon-add" @click="clickAddEnvironment">添加</el-button>
      <el-dialog title="添加环境变量" :visible.sync="dialogEnvironment" width="60%" :close-on-click-modal="false"
                 :before-close="handleClose">
        <el-form ref="form" :model="pojo" label-width="100px">
          <span v-if="operationType === 'update'">
            <el-form-item label="变量类型">
              <el-select v-model="pojo.type" placeholder="选择类型" style="width: 200px" disabled />
            </el-form-item>
            <el-form-item label="变量key">
              <el-input v-model="pojo.key" placeholder="请输入内容" disabled />
            </el-form-item>
          </span>
          <span v-if="operationType === 'create'">
            <el-form-item label="变量类型">
              <el-select v-model="pojo.type" placeholder="选择类型" style="width: 200px">
                <el-option v-for="(item,index) in envTypes" :key="index" :label="item" :value="item" />
              </el-select>
            </el-form-item>
            <el-form-item label="变量key">
              <el-input v-model="pojo.key" placeholder="请输入内容" />
            </el-form-item>
          </span>

          <el-form-item label="变量值">
            <el-input v-model="pojo.value" placeholder="请输入内容" />
          </el-form-item>
          <el-form-item label="描述">
            <el-input v-model="pojo.intro" placeholder="请输入内容" type="textarea" :autosize="{ minRows: 2, maxRows: 4}" />
          </el-form-item>

        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="handleCancel">取 消</el-button>
          <el-button type="primary" @click="handleEnvironmentSave()">确 定</el-button>
        </span>
      </el-dialog>
    </div>
    <!-- 2.展示 -->
    <el-table :data="envList" :row-class-name="tableRowClassName">
      <el-table-column width="60px">
        <template scope="scope">
          {{ scope.$index+1 }}
        </template>
      </el-table-column>
      <el-table-column prop="type" label="变量类型" />
      <el-table-column prop="key" label="变量key" />
      <el-table-column prop="value" label="变量值" width="300px" />
      <el-table-column prop="intro" label="描述" width="300px" />

      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button type="text" size="small" @click="clickUpdateEnvironment(scope.row.key)">修改</el-button>
          <el-button type="text" size="small" @click="clickDeleteEnvironment(scope.row.key)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

  </div>
</template>
<script>
import envApi from '@/api/env'
import { isEmpty } from '@/utils/validate'

export default {
  data() {
    return {
      dialogEnvironment: false,
      operationType: '',
      pojo: {
        type: '',
        key: '',
        value: '',
        intro: ''
      },
      envTypes: [],
      envList: []
    }
  },
  created() {
    this.fetchData()
    this.fetchEnvTypesData()
  },
  methods: {
    fetchData() {
      envApi.getList().then(res => {
        if (res.state === 1) {
          this.envList = res.data
          console.log('this.envList======>', this.envList)
        }
      })
    },
    // 获取环境变量类型
    fetchEnvTypesData() {
      envApi.getEnvType().then(res => {
        if (res.state === 1) {
          this.envTypes = res.data
        }
      })
    },
    clickAddEnvironment() {
      console.log('pojo--->', this.pojo)
      this.dialogEnvironment = true
      this.operationType = 'create'
    },
    clickUpdateEnvironment(key) {
      this.dialogEnvironment = true
      this.operationType = 'update'
      this.envList.forEach(env => {
        if (env.key === key) {
          this.pojo = env
        }
      })
    },
    handleEnvironmentSave() {
      console.log('this.operationType', this.operationType)
      if (this.operationType === 'create') {
        // 添加
        console.log('新增---------')
        envApi.create(this.pojo).then(res => {
          if (res.state === 1) {
            this.$message({
              type: 'success',
              message: res.data
            })
            this.fetchData()
            this.dialogEnvironment = false
            this.pojo = {
              type: '',
              key: '',
              value: '',
              intro: ''
            }
          }
        })
      } else {
        // 修改
        console.log('修改---------')
        envApi.update(this.pojo).then(res => {
          if (res.state === 1) {
            this.$message({
              type: 'success',
              message: res.data
            })
            this.fetchData()
            this.dialogEnvironment = false
            this.pojo = {
              type: '',
              key: '',
              value: '',
              intro: ''
            }
          }
        })
      }
    },
    clickDeleteEnvironment(key) {
      if (!isEmpty(key)) {
        this.$confirm('确认删除吗', {
          confirmButtonText: '确定',
          cancelButtonText: '取消'
        }).then(() => {
          envApi.delete(key).then(res => {
            if (res.state === 1) {
              this.$message({
                type: 'success',
                message: res.data
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
    handleCancel() {
      this.dialogEnvironment = false
      this.pojo = {
        type: '',
        key: '',
        value: '',
        intro: ''
      }
    },
    handleClose(done) {
      this.$confirm('关闭将清除表单数据，确认关闭？').then(_ => {
        this.pojo = {
          type: '',
          key: '',
          value: '',
          intro: ''
        }
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
