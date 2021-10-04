<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-edit" @click="handleChangePassword">
        修改密码
      </el-button>
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-view" @click="handleQueryPrivileges">
        查看权限
      </el-button>
    </div>
    <el-form ref="form" :model="selfData" label-width="80px">
      <el-form-item label="用户名">
        {{ selfData.userName }}
      </el-form-item>
      <el-form-item label="昵称">
        {{ selfData.nickName }}
      </el-form-item>
      <el-form-item label="上级">
        {{ selfData.parentName }}
      </el-form-item>
      <el-form-item label="创建时间">
        {{ selfData.createTime | dataFormat }}
      </el-form-item>
      <el-form-item label="状态">
        {{ selfData.state | stateMap }}
      </el-form-item>
    </el-form>

    <el-dialog
      title="修改密码"
      :visible.sync="changeDialogVisible">
      <el-form ref="resultForm" :model="tempPassword" label-position="right" label-width="105px" style="width: 90%;">
        <el-form-item label="新密码">
          <el-input v-model="tempPassword.password" placeholder="新密码" show-password />
        </el-form-item>
        <el-form-item label="再次输入">
          <el-input v-model="tempPassword.checkPassword" placeholder="再次输入" show-password />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="changeDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="doChangePassword()">确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog :visible.sync="privilegesDialogVisible">
      <el-tree :data="treeData" :props="defaultProps" />
    </el-dialog>
  </div>
</template>

<script>
import { getSelfInfo, changePassword } from '@/api/privileges'
import waves from '@/directive/waves'
import { format } from '@/utils/date-format'
import { Message } from 'element-ui'
let _this = null

export default {
  name: 'SelfManager',
  directives: { waves },
  filters: {
    dataFormat: function(value) {
      return format(new Date(value), 'yyyy-MM-dd hh:mm:ss.S')
    },
    stateMap: function(value) {
      if (value === 0) {
        return '初始化'
      } else if (value === 1) {
        return '禁用'
      } else if (value === 2) {
        return '正常'
      } else {
        return '数据错误'
      }
    }
  },
  data() {
    return {
      treeData: [],
      privilegesDialogVisible: false,
      defaultProps: {
        children: 'children',
        label: 'displayName'
      },
      tempPassword: {},
      changeDialogVisible: false,
      selfData: {}
    }
  },
  created() {
    _this = this
    _this.querySelfInfo()
  },
  methods: {
    handleQueryPrivileges() {
      // _this.treeData = val.privileges
      _this.privilegesDialogVisible = true
    },
    handleChangePassword() {
      _this.changeDialogVisible = true
      _this.tempPassword = {}
    },
    doChangePassword() {
      const param = Object.assign({}, _this.tempPassword)
      if (param.password.length < 6) {
        Message({
          message: '密码长度不能小于六',
          type: 'error',
          dangerouslyUseHTMLString: true
        })
        return
      }
      if (param.password !== param.checkPassword) {
        Message({
          message: '两次密码不一样',
          type: 'error',
          dangerouslyUseHTMLString: true
        })
        return
      }
      this.$confirm('确认要修改密码？')
        .then(res => {
          changePassword(param).then(response => {
            if (response.state !== 1) {
              Message({
                message: response.data.msg || '未知错误',
                type: 'error',
                dangerouslyUseHTMLString: true
              })
            } else {
              Message({
                message: '操作成功',
                type: 'success',
                dangerouslyUseHTMLString: true
              })
            }
            _this.changeDialogVisible = false
          })
        })
        .catch(_ => {})
    },
    querySelfInfo() {
      getSelfInfo().then(response => {
        if (response.state !== 1) {
          Message({
            message: response.data.msg || '未知错误',
            type: 'error',
            dangerouslyUseHTMLString: true
          })
          return
        }
        _this.selfData = response.data
        if (_this.selfData.state === 0) {
          _this.changeDialogVisible = true
        }
        _this.treeData = _this.selfData.privileges
      })
    }
  }
}
</script>
