<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="类别">
          <el-input placeholder="请输入类别" v-model="searchInfo.category" clearable @clear="this.getTableData"></el-input>
        </el-form-item>
        <el-form-item label="IP">
          <el-input placeholder="请输入IP地址" v-model="searchInfo.ip" clearable @clear="this.getTableData"></el-input>
        </el-form-item>
        <el-form-item label="状态">
          <el-input placeholder="请选择状态" v-model="searchInfo.status" clearable @clear="this.getTableData"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('addServer')" type="primary">新增服务器</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" @sort-change="sortChange" border stripe>
      <el-table-column label="类别" min-width="100" prop="category" show-overflow-tooltip sortable="custom"></el-table-column>
      <el-table-column label="架构" min-width="80" prop="architecture" show-overflow-tooltip sortable="custom">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.architecture === 'x86'" type="">x86</el-tag>
          <el-tag v-else-if="scope.row.architecture === 'arm'" type="danger">arm</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="IP" prop="ip" min-width="100" show-overflow-tooltip></el-table-column>
      <el-table-column label="密码" prop="password" min-width="110" show-overflow-tooltip></el-table-column>
      <el-table-column label="背板" prop="backboard" min-width="200" show-overflow-tooltip></el-table-column>
      <el-table-column label="raid卡" prop="raid_card" min-width="200" show-overflow-tooltip></el-table-column>
      <el-table-column label="占用状态" prop="status" min-width="90">
          <template slot-scope="scope">
            <el-tag v-if="scope.row.status === 1" type="success" effect="dark">闲置中</el-tag>
            <el-tag v-else-if="scope.row.status === 2" type="warning" effect="dark">占用中</el-tag>
            <el-tag v-else type="danger" effect="dark">已损坏</el-tag>
          </template>
      </el-table-column>
      <el-table-column label="业务IP" prop="business_ip" min-width="100" show-overflow-tooltip></el-table-column>
      <el-table-column label="业务密码" prop="business_password" min-width="110" show-overflow-tooltip></el-table-column>
      <el-table-column label="sn号" prop="sn" min-width="200" show-overflow-tooltip></el-table-column>
      <el-table-column label="位置" prop="location" min-width="60"></el-table-column>
      <el-table-column label="地区" prop="region" min-width="60"></el-table-column>
      <el-table-column label="上次占用人" prop="last_user" min-width="100" show-overflow-tooltip></el-table-column>
        <!-- <template slot-scope="scope">
          <el-tag v-if="scope.row.last_user != ''" type="" effect="dark">{{ scope.row.last_user }}</el-tag>
          <el-tag v-else type="success" effect="dark">暂无</el-tag>
        </template> -->
      <el-table-column label="占用人" prop="user.nickName" min-width="90" show-overflow-tooltip>
        <template slot-scope="scope">
          <el-tag v-if="scope.row.user.nickName != ''" type="warning" effect="dark">{{ scope.row.user.nickName }}</el-tag>
          <el-tag v-else type="success" effect="dark">暂无</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="占用开始时间" prop="occupy_time" min-width="110" show-overflow-tooltip>
        <template slot-scope="scope">
          <div>{{ scope.row.occupy_time }}</div>
        </template>
      </el-table-column>
      <el-table-column label="预计释放时间" prop="release_time" min-width="110" show-overflow-tooltip>
        <template slot-scope="scope">
          <div>{{ scope.row.release_time }}</div>
        </template>
      </el-table-column>
      <el-table-column label="备注" prop="postscript" min-width="160" show-overflow-tooltip></el-table-column>
      <el-table-column fixed="right" label="操作" width="320" >
        <template slot-scope="scope">
          <el-button @click="editServer(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button @click="deleteServer(scope.row)" size="small" type="danger" icon="el-icon-delete">删除</el-button>
          <el-button @click="occupyServer(scope.row)" size="small" type="warning" icon="el-icon-plus">占用</el-button>
          <el-button @click="releaseServer(scope.row)" size="small" type="success" icon="el-icon-minus">释放</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 页码数 -->
    <div class="fl-left left-mg-lg">
      <div>用户ID：{{userInfo.ID}}</div>
      <div>用户昵称：{{userInfo.nickName}}</div>
      <div>用户组：{{userInfo.authority&&userInfo.authority.authorityName}}</div>
    </div>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
    <!-- 新增服务器弹窗 -->
    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="ServerForm">
        <el-form-item label="IP" prop="ip">
          <el-input autocomplete="off" v-model="form.ip"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input autocomplete="off" v-model="form.password"></el-input>
        </el-form-item>
        <el-form-item label="业务IP" prop="business_ip">
          <el-input autocomplete="off" v-model="form.business_ip"></el-input>
        </el-form-item>
        <el-form-item label="业务密码" prop="business_password">
          <el-input autocomplete="off" v-model="form.business_password"></el-input>
        </el-form-item>
        <el-form-item label="背板" prop="backboard">
          <el-input autocomplete="off" v-model="form.backboard"></el-input>
        </el-form-item>
        <el-form-item label="raid卡" prop="raid_card">
          <el-input autocomplete="off" v-model="form.raid_card"></el-input>
        </el-form-item>
        <el-form-item label="位置" prop="location">
          <el-input autocomplete="off" v-model="form.location"></el-input>
        </el-form-item>
        <el-form-item label="sn号" prop="sn">
          <el-input autocomplete="off" v-model="form.sn"></el-input>
        </el-form-item>
        <el-form-item label="类别" prop="category">
          <!-- <el-input autocomplete="off" v-model="form.category"></el-input> -->
          <el-select placeholder="请选择" v-model="form.category" clearable>
            <el-option
              v-for="item in categoryOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="架构" prop="architecture">
          <!-- <el-input autocomplete="off" v-model="form.architecture"></el-input> -->
          <el-select placeholder="请选择" v-model="form.architecture" clearable>
            <el-option
              v-for="item in architectureOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <!-- <el-form-item label="状态" prop="status">
          <el-select placeholder="请选择" v-model="form.status" clearable>
            <el-option
              v-for="item in statusOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item> -->
        <el-form-item label="地区" prop="region">
          <el-select placeholder="请选择" v-model="form.region" clearable>
            <el-option
              v-for="item in regionOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="描述" prop="postscript">
          <el-input autocomplete="off" v-model="form.postscript" type="textarea" :rows="2"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <!-- 占用服务器弹窗 -->
    <el-dialog :before-close="closeOccupyDialog" :title="OccupydialogTitle" :visible.sync="occupyFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="ServerForm">
        <el-form-item label="占用日期">
          <el-date-picker
          v-model="form.occupy_time"
          type="date"
          format="yyyy 年 MM 月 dd 日"
          value-format="yyyy-MM-dd"
          placeholder="选择日期">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="释放日期">
          <el-date-picker
          v-model="form.release_time"
          format="yyyy 年 MM 月 dd 日"
          value-format="yyyy-MM-dd"
          type="date"
          placeholder="选择日期">
          </el-date-picker>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeOccupyDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import { getServerList, createServer, getServerById, updateServer, deleteServer } from '@/api/server'
import infoList from '@/components/mixins/infoList'
import { mapGetters } from 'vuex'
import { toSQLLine } from '@/utils/stringFun'
const statusOptions = [
  {
    value: 1,
    label: '闲置中',
   },
  {
    value: 2,
    label: '占用中',
  },
  {
    value: 3,
    label: '已损坏',
   }
]
const regionOptions = [
  {
    value: '成都',
    label: '成都',
   },
  {
    value: '东莞',
    label: '东莞',
  },
  {
    value: '西安',
    label: '西安',
   }
]
const architectureOptions = [
  {
    value: 'x86',
    label: 'x86',
   },
  {
    value: 'arm',
    label: 'arm',
  }
]
const categoryOptions = [
  {
    value: '2288H V5',
    label: '2288H V5',
   },
  {
    value: '1288H V5',
    label: '1288H V5',
  },
    {
    value: '5288H V5',
    label: '5288H V5',
  }
]

export default {
  name: 'Api',
  mixins: [infoList],
  data() {
    return {
      listApi: getServerList,
      dialogFormVisible: false,
      occupyFormVisible: false,
      dialogTitle: '新增服务器',
      OccupydialogTitle: '占用服务器',
      form: {
        category: '',
        architecture: '',
        sn: '',
        location: '',
        status: '',
        region: '',
        release_time: '',
        postscript: '',
        user_id: '',
      },
      user_form: {
        user_id: '',
      },
      statusOptions: statusOptions,
      regionOptions: regionOptions,
      architectureOptions:architectureOptions,
      categoryOptions: categoryOptions,
      type: '',
      rules: {
        category: [{ required: true, message: '请输入类别', trigger: 'blur' }
        ],
        Architecture: [
          { required: true, message: '请输入架构', trigger: 'blur' }
        ],
        location: [
          { required: true, message: '请输入具体位置', trigger: 'blur' }
        ],
        sn: [
          { required: true, message: '请输入sn号', trigger: 'blur' }
        ],
        status: [
          { required: true, message: '请选择状态', trigger: 'blur' }
        ],
        region: [
          { required: true, message: '请选择地区', trigger: 'blur' }
        ],
        // user_id: [
        //   { required: true, message: '请选择用户', trigger: 'blur' }
        // ],
      }
    }
  },
   computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  methods: {
    // 排序
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
        this.searchInfo.desc = order == 'descending'
      }
      this.getTableData()
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    initForm() {
      this.$refs.ServerForm.resetFields()
      this.form= {
        category: '',
        ip: '',
        password: '',
        sn: '',
        backboard: '',
        raid_card: '',
        business_ip: '',
        business_password: '',
        status: '',
        release_time: '',
        postscript: '',
        user_id: 0,
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    closeOccupyDialog() {
      this.initForm()
      this.occupyFormVisible = false
    },
    openDialog(type) {
      switch (type) {
        case 'addServer':
          this.dialogTitle = '新增Server'
          break
        case 'edit':
          this.dialogTitle = '编辑Server'
          break
        case 'occupy':
          this.OccupydialogTitle = '占用Server'
        default:
          break
      }
      this.type = type
      this.dialogFormVisible = true
    },
    openOccupyDialog(type) {
      switch (type) {
        case 'occupy':
          this.OccupydialogTitle = '占用Server'
        default:
          break
      }
      this.type = type
      this.occupyFormVisible = true
    },
    async editServer(row) {
      const res = await getServerById({ id: row.ID })
      console.log(res,11111)
      this.form = res.data.server
      this.openDialog('edit')
    },
    // async occupyServer(row) {
    //   this.$confirm('此操作将占用此服务器，是否继续?', '提示', {
    //     confirmButtonText: '确定',
    //     cancelButtonText: '取消',
    //     type: 'warning'
    //   })
    //     .then(async () => {
    //       const result = await getServerById({ id: row.ID })
    //       if (result.data.disk.status == 1) {
    //         result.data.disk.user_id = this.userInfo.ID
    //         result.data.disk.status = 2
    //         this.form = result.data.disk
    //         const res = await updateDisk(this.form)
    //         if (res.code == 0) {
    //           this.$message({
    //           type: 'success',
    //           message: '占用成功!',
    //         })
    //         this.getTableData()
    //         }
    //       } else {
    //           this.$message({
    //           type: 'warning',
    //           message: '无法占用!',
    //         })
    //       }      
    //     })
    //     .catch(() => {
    //     this.$message({
    //     type: 'info',
    //     message: '已取消占用'
    //     })
    //     })
    // },
    async occupyServer(row) {
      const result = await getServerById({ id: row.ID })
      if (result.data.server.status == 1) {
        result.data.server.user_id = this.userInfo.ID
        result.data.server.status = 2
        this.form = result.data.server
        this.openOccupyDialog('occupy')
      } else {
        this.$message({
        type: 'warning',
        message: '无法占用!',
        })
      }
    },
    async releaseServer(row) {
      this.$confirm('此操作将释放此服务器，是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const result = await getServerById({ id: row.ID })
          console.log(result, 4444)
          result.data.server.user_id = ''
          result.data.server.status = 1
          result.data.server.occupy_time = ''
          result.data.server.release_time = ''
          result.data.server.last_user = this.userInfo.nickName
          this.form = result.data.server
          const res = await updateServer(this.form)
          console.log(res, 222)
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: '释放成功!',
            })
            this.getTableData()
          }
        })
        .catch(() => {
        this.$message({
        type: 'info',
        message: '已取消释放'
        })
        })
    },
    async deleteServer(row) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const res = await deleteServer(row)
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            this.getTableData()
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    async enterDialog() {
      this.$refs.ServerForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'addServer':
              {
                const res = await createServer(this.form)
                if (res.code == 0) {
                  this.$message({
                    type: 'success',
                    message: '添加成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeDialog()
              }

              break
            case 'edit':
              {
                const res = await updateServer(this.form)
                if (res.code == 0) {
                  this.$message({
                    type: 'success',
                    message: '编辑成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeDialog()
              }
              break
            case 'occupy':
              { 
                const res = await updateServer(this.form)
                if (res.code == 0) {
                  this.$message({
                    type: 'success',
                    message: '占用成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeOccupyDialog()
              }
              break
            default:
              {
                this.$message({
                  type: 'error',
                  message: '未知操作',
                  showClose: true
                })
              }
              break
          }
        }
      })
    }
  },
  filters: {
    // methodFiletr(value) {
    //   const target = methodOptions.filter(item => item.value === value)[0]
    //   // return target && `${target.label}(${target.value})`
    //   return target && `${target.label}`
    // },
    // tagTypeFiletr(value) {
    //   const target = methodOptions.filter(item => item.value === value)[0]
    //   return target && `${target.type}`
    // }
  },
  created(){
    this.getTableData()
  }
}
</script>
<style scoped lang="scss">