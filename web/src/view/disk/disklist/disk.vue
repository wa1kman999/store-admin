<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="disk厂商">
          <el-input placeholder="请输入厂商" v-model="searchInfo.vendor" clearable @clear="this.getTableData"></el-input>
        </el-form-item>
        <el-form-item label="disk系列">
          <el-input placeholder="请输入系列" v-model="searchInfo.series" clearable @clear="this.getTableData"></el-input>
        </el-form-item>
        <el-form-item label="model号">
          <el-input placeholder="请输入model号" v-model="searchInfo.model_number" clearable @clear="this.getTableData"></el-input>
        </el-form-item>
        <el-form-item label="sn号">
          <el-input placeholder="请输入sn号" v-model="searchInfo.sn" clearable @clear="this.getTableData"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('addDisk')" type="primary">新增硬盘</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" @sort-change="sortChange" border stripe>
      <el-table-column label="厂商" min-width="100" prop="vendor" show-overflow-tooltip sortable="custom"></el-table-column>
      <el-table-column label="系列" min-width="100" prop="series" show-overflow-tooltip sortable="custom"></el-table-column>
      <el-table-column label="model号" min-width="150" prop="model_number" show-overflow-tooltip></el-table-column>
      <el-table-column label="sn号" prop="sn" min-width="200" show-overflow-tooltip></el-table-column>
      <el-table-column label="容量" min-width="60" prop="capacity"></el-table-column>
      <el-table-column label="接口类型" prop="interfaces" min-width="80">
          <template slot-scope="scope">
            <el-tag v-if="scope.row.interfaces === 1" type="" effect="plain">SAS</el-tag>
            <el-tag v-else-if="scope.row.interfaces === 2" type="danger" effect="plain">SATA</el-tag>
            <el-tag v-else type="warning" effect="plain">PCIe</el-tag>
          </template>
      </el-table-column>
      <el-table-column label="介质类型" prop="type" min-width="80">
          <template slot-scope="scope">
            <el-tag v-if="scope.row.type === 1" type="">HDD</el-tag>
            <el-tag v-else type="info">SDD</el-tag>
          </template>
      </el-table-column>
      <el-table-column label="占用状态" prop="status" min-width="90">
          <template slot-scope="scope">
            <el-tag v-if="scope.row.status === 1" type="success" effect="dark">闲置中</el-tag>
            <el-tag v-else-if="scope.row.status === 2" type="warning" effect="dark">占用中</el-tag>
            <el-tag v-else type="danger" effect="dark">已损坏</el-tag>
          </template>
      </el-table-column>
      <el-table-column label="地区" prop="region" min-width="60"></el-table-column>
      <el-table-column label="占用人" prop="user.nickName" min-width="80" show-overflow-tooltip></el-table-column>
      <el-table-column label="描述" prop="postscript" min-width="160" show-overflow-tooltip></el-table-column>
      <el-table-column fixed="right" label="操作" width="320" >
        <template slot-scope="scope">
          <el-button @click="editDisk(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button @click="deleteDisk(scope.row)" size="small" type="danger" icon="el-icon-delete">删除</el-button>
          <el-button @click="occupyDisk(scope.row)" size="small" type="warning" icon="el-icon-plus">借出</el-button>
          <el-button @click="releaseDisk(scope.row)" size="small" type="success" icon="el-icon-minus">归还</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 页码数 -->
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
    <!-- 新增硬盘弹窗 -->
    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="diskForm">
        <el-form-item label="厂商" prop="vendor">
          <el-input autocomplete="off" v-model="form.vendor"></el-input>
        </el-form-item>
        <el-form-item label="系列" prop="series">
          <el-input autocomplete="off" v-model="form.series"></el-input>
        </el-form-item>
        <el-form-item label="model号" prop="model_number">
          <el-input autocomplete="off" v-model="form.model_number"></el-input>
        </el-form-item>
        <el-form-item label="sn号" prop="sn">
          <el-input autocomplete="off" v-model="form.sn"></el-input>
        </el-form-item>
        <el-form-item label="扇区" prop="section">
          <el-input autocomplete="off" v-model="form.section"></el-input>
        </el-form-item>
        <el-form-item label="容量" prop="capacity">
          <el-input autocomplete="off" v-model="form.capacity"></el-input>
        </el-form-item>

        <el-form-item label="接口" prop="interfaces">
          <el-select placeholder="请选择接口" v-model="form.interfaces" clearable>
            <el-option
              v-for="item in interfacesOptions"
              :key="item.value"
              :label="`${item.label}`"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select placeholder="请选择" v-model="form.type" clearable>
            <el-option
              v-for="item in typeOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select placeholder="请选择" v-model="form.status" clearable>
            <el-option
              v-for="item in statusOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
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
      <!-- <div class="warning">新增硬盘需要在角色管理内配置权限才可使用</div> -->
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import { getDiskList, createDisk, getDiskById, updateDisk, deleteDisk } from '@/api/disk'
import infoList from '@/components/mixins/infoList'
import { mapGetters } from 'vuex'
import { toSQLLine } from '@/utils/stringFun'
const interfacesOptions = [
  {
    value: 1,
    label: 'SAS',
   },
  {
    value: 2,
    label: 'SATA',
  },
  {
    value: 3,
    label: 'PCIe',
   }
]
const typeOptions = [
  {
    value: 1,
    label: 'HDD',
   },
  {
    value: 2,
    label: 'SSD',
  },
]
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

export default {
  name: 'Api',
  mixins: [infoList],
  data() {
    return {
      listApi: getDiskList,
      dialogFormVisible: false,
      dialogTitle: '新增硬盘',
      form: {
        vendor: '',
        series: '',
        model_number: '',
        sn: '',
        section: '',
        capacity: '',
        interfaces: '',
        type: '',
        status: '',
        region: '',
        // healthy: '',
        postscript: '',
        user_id: '',
      },
      user_form: {
        user_id: '',
      },
      // methodOptions: methodOptions,
      interfacesOptions: interfacesOptions,
      typeOptions: typeOptions,
      statusOptions: statusOptions,
      regionOptions: regionOptions,
      type: '',
      rules: {
        vendor: [{ required: true, message: '请输入厂商', trigger: 'blur' }],
        series: [
          { required: true, message: '请输入系列', trigger: 'blur' }
        ],
        model_number: [
          { required: true, message: '请输入model号', trigger: 'blur' }
        ],
        sn: [
          { required: true, message: '请输入sn号', trigger: 'blur' }
        ],
        section: [
          { required: true, message: '请输入扇区大小', trigger: 'blur' }
        ],
        capacity: [
          { required: true, message: '请输入容量', trigger: 'blur' }
        ],
        interfaces: [
          { required: true, message: '请选择接口', trigger: 'blur' }
        ],
        type: [
          { required: true, message: '请选择类型', trigger: 'blur' }
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
      this.$refs.diskForm.resetFields()
      this.form= {
        vendor: '',
        series: '',
        model_number: '',
        sn: '',
        section: '',
        capacity: '',
        interfaces: '',
        type: '',
        status: '',
        region: '',
        // healthy: '',
        postscript: '',
        user_id: '',
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    openDialog(type) {
      switch (type) {
        case 'addDisk':
          this.dialogTitle = '新增Disk'
          break
        case 'edit':
          this.dialogTitle = '编辑Disk'
          break
        case 'occupy':
          this.dialogTitle = '占用Disk'
        default:
          break
      }
      this.type = type
      this.dialogFormVisible = true
    },
    async editDisk(row) {
      const res = await getDiskById({ id: row.ID })
      console.log(res,11111)
      this.form = res.data.disk
      this.openDialog('edit')
    },
    async occupyDisk(row) {
      this.$confirm('此操作将占用此硬盘，是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const result = await getDiskById({ id: row.ID })
          if (result.data.disk.status == 1) {
            result.data.disk.user_id = this.userInfo.ID
            result.data.disk.status = 2
            this.form = result.data.disk
            const res = await updateDisk(this.form)
            if (res.code == 0) {
              this.$message({
              type: 'success',
              message: '占用成功!',
            })
            this.getTableData()
            }
          } else {
              this.$message({
              type: 'warning',
              message: '无法占用!',
            })
          }      
        })
        .catch(() => {
        this.$message({
        type: 'info',
        message: '已取消占用'
        })
        })
    },
    async releaseDisk(row) {
      this.$confirm('此操作将释放此硬盘，是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const result = await getDiskById({ id: row.ID })
          result.data.disk.user_id = ''
          result.data.disk.status = 1
          this.form = result.data.disk
          const res = await updateDisk(this.form)
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
    async deleteDisk(row) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const res = await deleteDisk(row)
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
      this.$refs.diskForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'addDisk':
              {
                const res = await createDisk(this.form)
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
                const res = await updateDisk(this.form)
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
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>