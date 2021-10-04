<template>
  <div class="app-container">
    <!-- 1.创建用户 -->
    <div class="filter-container">
      <span v-if="loginType === 'ldap'">
        <el-button v-waves class="filter-item" type="primary" icon="el-icon-edit" @click="handleCreateUserByLdap"> 创建用户
        </el-button>
      </span>
      <span v-else>
        <el-button v-waves class="filter-item" type="primary" icon="el-icon-edit" @click="handleCreateUserBySSO"> 创建用户
        </el-button>
      </span>

      <el-dialog
        title="创建用户"
        :visible.sync="createDialogVisible">
        <el-form ref="resultForm" :model="tempUser" label-position="right" label-width="105px" style="width: 90%;">
          <el-form-item label="用户名">
            <el-input v-model="tempUser.userName" placeholder="用户名" />
          </el-form-item>
          <el-form-item label="昵称">
            <el-input v-model="tempUser.nickName" placeholder="昵称" />
          </el-form-item>
          <el-form-item label="邮箱">
            <el-input v-model="tempUser.userMail" placeholder="邮箱" />
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="createDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="doCreateUser()">确 定</el-button>
        </span>
      </el-dialog>
    </div>
    <!-- 2.列表 -->
    <el-table border :data="tableData" style="width: 100%">
      <el-table-column label="创建时间" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.createdAt | dateFormatFilter }}</span>
        </template>
      </el-table-column>
      <el-table-column label="用户名" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.userName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="昵称" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.nickName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="上级" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.parentName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.state | stateMap }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="280">
        <template slot-scope="scope">
          <el-button v-if="scope.row.state === 1 && !scope.row.self" size="mini" @click="handleAllowedUser(scope.row)">启用</el-button>
          <el-button v-if="scope.row.state !== 1 && !scope.row.self" size="mini" @click="handleForbiddenUser(scope.row)">禁用</el-button>
          <el-button size="mini" @click="handleQueryPrivileges(scope.row)">查看权限</el-button>
          <el-button v-if="scope.row.state !== 1 && !scope.row.self" size="mini" @click="handleEditPrivileges(scope.row)">修改权限</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 权限弹出框 -->
    <div>
      <!-- 1.1查看权限 -->
      <el-dialog :visible.sync="privilegesDialogVisible">
        <el-tree :data="treeData" :props="defaultProps" />
      </el-dialog>
      <!-- 1.2修改权限 -->
      <el-dialog :visible.sync="editDialogVisible">
        <el-tree ref="tree" :data="editData" default-expand-all node-key="name" show-checkbox :props="defaultProps" />
        <el-tree ref="projectTree" :data="editProjectData" default-expand-all node-key="name" show-checkbox :props="defaultProps" />
        <el-tree ref="funcTree" :data="editFuncData" default-expand-all node-key="name" show-checkbox :props="defaultProps" />

        <span slot="footer" class="dialog-footer">
          <el-button @click="editDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="doEditPrivileges()">确 定</el-button>
        </span>
      </el-dialog>

      <!-- 2.1查看项目权限 -->
      <el-dialog :visible.sync="projectPrivilegesDialog" width="400px">
        <el-table :data="projectTreeData" border>
          <el-table-column prop="cnName" label="项目名称" />
        </el-table>
      </el-dialog>
      <!-- 2.2修改项目权限 -->
      <el-dialog :visible.sync="projectEditDialog">
        <el-tree ref="projectTree" :data="projectEditData" default-expand-all node-key="id" show-checkbox :props="defaultProjectProps" />

        <span slot="footer" class="dialog-footer">
          <el-button @click="projectEditDialog = false">取 消</el-button>
          <el-button type="primary" @click="doProjectEditPrivileges()">确 定</el-button>
        </span>
      </el-dialog>

      <!-- 3.1查看高级权限 -->
      <el-dialog :visible.sync="highPrivilegesDialog">
        <el-tree :data="highTreeData" :props="defaultHighProps" />
      </el-dialog>
      <!-- 修3.2改高级权限 -->
      <el-dialog :visible.sync="highEditDialog">
        <el-tree ref="highTree" :data="highEditData" default-expand-all node-key="path" show-checkbox :props="defaultHighProps" />

        <span slot="footer" class="dialog-footer">
          <el-button @click="highEditDialog = false">取 消</el-button>
          <el-button type="primary" @click="doHighEditPrivileges()">确 定</el-button>
        </span>
      </el-dialog>
    </div>

  </div>
</template>

<script>
import { createUser, getSubordinateList, forbiddenUser, allowedUser, changePrivileges, changeProjectPrivileges, getHighPrivileges, changeHighPrivileges } from '@/api/user'
import projectApi from '@/api/project'
import waves from '@/directive/waves'
import { Message } from 'element-ui'
import { isEmpty, contains } from '../../utils/validate'

let _this = null
const projectPrivilege = 'projectPrivilege'
const funcPrivilege = 'funcPrivilege'
export default {
  name: 'UserManager',
  directives: { waves },
  filters: {
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
      loginType: config.loginType,
      treeData: [],
      editData: [],
      editProjectData: [],
      editFuncData: [],
      tempPrivileges: {},
      editDialogVisible: false,
      privilegesDialogVisible: false,
      defaultProps: {
        children: 'children',
        label: 'displayName'
      },
      projectPrivilegesDialog: false,
      projectEditDialog: false,
      projectTreeData: [],
      projectEditData: [],
      projectTempPrivileges: {},
      defaultProjectProps: {
        label: 'cnName'
      },

      highPrivilegesDialog: false,
      highEditDialog: false,
      highTreeData: [],
      highEditData: [],
      highTempPrivileges: {},
      defaultHighProps: {
        children: 'children',
        label: 'intro'
      },

      tempUser: {
        userName: '',
        nickName: '',
        userMail: ''
      },
      createDialogVisible: false,
      tableData: [],
      projectList: [],
      highPrivilegeList: []
    }
  },
  created() {
    _this = this
    _this.querySubordinateList()
    // _this.fetchProjectData()
  },
  methods: {
    // 查询下属信息
    querySubordinateList() {
      getSubordinateList().then(response => {
        if (response.state === 1) {
          _this.tableData = response.data
          console.log('tableData=======>>>', this.tableData)
        }
      })
    },
    fetchProjectData() {
      projectApi.getList().then(res => {
        if (res.state === 1) {
          this.projectList = res.data
          console.log('this.projectList===>', this.projectList)
        }
      })

      getHighPrivileges().then(res => {
        if (res.state === 1) {
          this.highPrivilegeList = res.data
          console.log('this.highPrivilegeList===>', this.highPrivilegeList)
        }
      })
    },
    handleCreateUserByLdap() {
      _this.tempUser = {}
      _this.createDialogVisible = true
    },
    handleCreateUserBySSO() {
      this.$confirm('创建用户将跳转至SSO页面执行操作').then(() => {
        window.open(config.ssoDomain + 'user/usermanager')
      }).catch(_ => {})
    },
    doCreateUser() {
      if (isEmpty(_this.tempUser.userName) || isEmpty(_this.tempUser.nickName) || isEmpty(_this.tempUser.userMail)) {
        Message({
          message: '请填写完整参数',
          type: 'error',
          dangerouslyUseHTMLString: true
        })
      } else {
        const param = Object.assign({}, _this.tempUser)
        createUser(param).then(response => {
          if (response.state === 1) {
            Message({
              message: '操作成功',
              type: 'success',
              dangerouslyUseHTMLString: true
            })
            _this.tempUser = {}
            _this.createDialogVisible = false

            const timer = setInterval(() => {
              _this.querySubordinateList()
              clearInterval(timer)
            }, 500)
          }
        })
      }
    },
    // 启用
    handleAllowedUser(val) {
      const param = Object.assign({}, val)
      this.$confirm('确认要启用用户？').then(res => {
        allowedUser(param.userName).then(response => {
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
          const timer = setInterval(() => {
            _this.querySubordinateList()
            clearInterval(timer)
          }, 500)
        })
      }).catch(_ => {})
    },
    // 禁用
    handleForbiddenUser(val) {
      const param = Object.assign({}, val)
      console.log('param---->', param)
      this.$confirm('确认要禁用用户？').then(res => {
        forbiddenUser(param.userName).then(response => {
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
          const timer = setInterval(() => {
            _this.querySubordinateList()
            clearInterval(timer)
          }, 500)
        })
      }).catch(_ => {})
    },
    // 1.权限
    handleQueryPrivileges(val) {
      _this.treeData = val.privileges
      _this.privilegesDialogVisible = true
    },
    handleEditPrivileges(val) {
      _this.editProjectData = []
      _this.editFuncData = []
      _this.tempPrivileges = {}
      _this.tempPrivileges.userName = val.userName
      _this.editData = []

      _this.tableData.forEach(item => {
        if (item.self) {
          _this.editData = item.privileges
        }
      })
      console.log(' _this.editData===>', _this.editData)

      const removeIndex = []
      _this.editData.forEach((item, i) => {
        if (item.name === projectPrivilege) {
          _this.editProjectData.push(item)
          removeIndex.push(i)
        }
        if (item.name === funcPrivilege) {
          _this.editFuncData.push(item)
          removeIndex.push(i)
        }
        if (item.name !== projectPrivilege && item.name !== funcPrivilege) {
          item.children = []
        }
      })
      console.log('remove', removeIndex)
      const temp = []
      _this.editData.forEach((e, index) => {
        if (!contains(removeIndex, index)) {
          temp.push(e)
        }
      })
      _this.editData = temp
      console.log('val.privileges===>', val.privileges)
      _this.$nextTick(() => {
        if (val.privileges) {
          const list = []

          val.privileges.forEach(item => {
            list.push(item)
          })
          console.log('list====>', list)
          _this.$refs.tree.setCheckedNodes(list)
          _this.$refs.projectTree.setCheckedNodes(val.privileges[val.privileges.length - 2].children)
          _this.$refs.funcTree.setCheckedNodes(val.privileges[val.privileges.length - 1].children)
        }
      })
      _this.editDialogVisible = true
    },
    doEditPrivileges() {
      const privileges = this.$refs.tree.getCheckedKeys()
      const projectPrivileges = this.$refs.projectTree.getCheckedKeys()
      const funcPrivileges = this.$refs.funcTree.getCheckedKeys()
      if (projectPrivileges.length > 1 && projectPrivileges[0] === projectPrivilege) {
        projectPrivileges.splice(0, 1)
      }
      if (funcPrivileges.length > 1 && funcPrivileges[0] === funcPrivilege) {
        funcPrivileges.splice(0, 1)
      }
      _this.tempPrivileges.privileges = privileges
      _this.tempPrivileges.projectIds = projectPrivileges
      _this.tempPrivileges.funcs = funcPrivileges

      this.$confirm('确认要修改权限？').then(() => {
        changePrivileges(_this.tempPrivileges).then(response => {
          if (response.state !== 1) {
            Message({
              message: response.data.msg || '未知错误',
              type: 'error',
              dangerouslyUseHTMLString: true
            })
            return
          } else {
            Message({
              message: '操作成功',
              type: 'success',
              dangerouslyUseHTMLString: true
            })
            _this.editDialogVisible = false
            _this.querySubordinateList()
          }
        }).catch(e => {
          Message({
            message: e || '未知错误',
            type: 'error',
            dangerouslyUseHTMLString: true
          })
          return
        })
      })
    },
    // 2.项目权限
    handleQueryProjectPrivileges(val) {
      _this.projectTreeData = val.projectPrivileges
      _this.projectPrivilegesDialog = true
    },
    handleEditProjectPrivileges(val) {
      _this.projectTempPrivileges = {}
      _this.projectTempPrivileges.userName = val.userName
      _this.projectEditData = []

      _this.tableData.forEach(item => {
        if (item.self) {
          _this.projectEditData = item.projectPrivileges
        }
      })
      console.log(' _this.projectEditData===>', _this.projectEditData)

      console.log('val.privileges===>', val.projectPrivileges)
      _this.$nextTick(() => {
        if (val.projectPrivileges) {
          const list = []
          val.projectPrivileges.forEach(item => {
            list.push(item)
          })
          console.log('list====>', list)
          _this.$refs.projectTree.setCheckedNodes(list)
        }
      })
      _this.projectEditDialog = true
    },
    doProjectEditPrivileges() {
      const privileges = this.$refs.projectTree.getCheckedKeys()
      _this.projectTempPrivileges.projectPrivileges = privileges
      console.log('projectPrivileges===>', privileges)

      this.$confirm('确认要修改权限？').then(() => {
        changeProjectPrivileges(_this.projectTempPrivileges.userName, _this.projectTempPrivileges.projectPrivileges).then(response => {
          if (response.state !== 1) {
            Message({
              message: response.data.msg || '未知错误',
              type: 'error',
              dangerouslyUseHTMLString: true
            })
            return
          } else {
            Message({
              message: '操作成功',
              type: 'success',
              dangerouslyUseHTMLString: true
            })

            _this.projectEditDialog = false
            _this.querySubordinateList()
          }
        })
      }).catch(_ => {})
    },
    // 3.高级权限
    handleQueryHighPrivileges(val) {
      _this.highTreeData = val.apiPrivileges
      console.log('highTreeData===>', _this.highTreeData)
      _this.highPrivilegesDialog = true
    },
    handleEditHighPrivileges(val) {
      _this.highTempPrivileges = {}
      _this.highTempPrivileges.userName = val.userName
      _this.highEditData = []

      _this.tableData.forEach(item => {
        if (item.self) {
          _this.highEditData = item.apiPrivileges
        }
      })
      console.log(' _this.highEditData===>', _this.highEditData)

      console.log('val.privileges===>', val.apiPrivileges)
      _this.$nextTick(() => {
        if (val.apiPrivileges) {
          const list = []
          val.apiPrivileges.forEach(item => {
            if (item.children.length > 0) {
              item.children.forEach(child => {
                list.push(child)
              })
            } else {
              list.push(item)
            }
          })

          console.log('list====>', list)
          _this.$refs.highTree.setCheckedNodes(list)
        }
      })
      _this.highEditDialog = true
    },
    doHighEditPrivileges() {
      const privileges = this.$refs.highTree.getCheckedKeys()
      console.log('privileges===>', privileges)

      let privilegesParam = ''
      privileges.forEach(privilege => {
        privilegesParam += '/root/' + privilege + ','
      })
      const highPrivileges = privilegesParam.substring(0, privilegesParam.length - 1)
      console.log('highPrivileges====>', highPrivileges)

      this.$confirm('确认要修改权限？').then(() => {
        changeHighPrivileges(_this.highTempPrivileges.userName, highPrivileges).then(response => {
          if (response.state !== 1) {
            Message({
              message: response.data.msg || '未知错误',
              type: 'error',
              dangerouslyUseHTMLString: true
            })
            return
          } else {
            Message({
              message: '操作成功',
              type: 'success',
              dangerouslyUseHTMLString: true
            })
            _this.highEditDialog = false
            _this.querySubordinateList()
          }
        })
      }).catch(_ => {})
    }
  }
}
</script>
