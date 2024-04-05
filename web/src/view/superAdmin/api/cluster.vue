<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog('addApi')"
        >新增
        </el-button>
      </div>
      <el-table
        :data="tableData"
        @sort-change="sortChange"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="集群ip组"
          min-width="150"
          prop="ipGroup"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="集群名称"
          min-width="150"
          prop="name"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="集群描述"
          min-width="150"
          prop="name"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="默认配置文件"
          min-width="150"
          prop="config"
          sortable="custom"
        />
        <el-table-column
          align="left"
          fixed="right"
          label="操作"
          width="200"
        >
          <template #default="scope">
            <el-button
              icon="delete"

              type="primary"
              link
              @click="deleteApiFunc(scope.row)"
            >删除
            </el-button>
            <upload-common
              v-model:imageCommon="imageCommon"
              v-model:id="scope.row.ID"
              v-model:type="cluster"
              @on-success="getTableData"
            />
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>

    </div>

    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="dialogTitle"
    >
      <warning-bar title="新增cluster，需要在角色管理内配置权限才可使用" />
      <el-form
        ref="apiForm"
        :model="form"
        :rules="rules"
        label-width="150px"
      >
        <el-form-item
          label="主机组ip"
          prop="ipGroup"
        >
          <el-input
            v-model="form.ipGroup"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="集群名称"
          prop="name"
        >
          <el-input
            v-model="form.name"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="集群描述"
          prop="description"
        >
          <el-input
            v-model="form.description"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
            label="集群配置文件"
            prop="config"
        >
          <el-input
              v-model="form.config"
              autocomplete="off"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button
            type="primary"
            @click="enterDialog"
          >确 定
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { deleteApisByIds, freshCasbin, getApiById, updateApi } from '@/api/api'
import { createCluster, deleteCluster, getClusterList } from '@/api/cluster'
import { toSQLLine } from '@/utils/stringFun'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { VideoCameraFilled } from '@element-plus/icons-vue'
import { toDoc } from '@/utils/doc'
import UploadCommon from '@/components/upload/common.vue'

defineOptions({
  name: 'Api',
})

const methodFilter = (value) => {
  const target = methodOptions.value.filter(item => item.value === value)[0]
  return target && `${target.label}`
}
const imageCommon = ref('')
const cluster = 'cluster'
const apis = ref([])
const form = ref({
  name: '',
  namespace: '',
  description: '',
  ipGroup: '',
  config: '',
})
const methodOptions = ref([
  {
    value: 'POST',
    label: '创建',
    type: 'success'
  },
  {
    value: 'GET',
    label: '查看',
    type: ''
  },
  {
    value: 'PUT',
    label: '更新',
    type: 'warning'
  },
  {
    value: 'DELETE',
    label: '删除',
    type: 'danger'
  }
])

const type = ref('')
const rules = ref({
  ipGroup: [
    { required: true, message: '请输入主机组ip', trigger: 'blur' }
  ],
  name: [
    { required: true, message: '请输入cluster名称', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入cluster描述', trigger: 'blur' }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
}
// 搜索

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    if (prop === 'ID') {
      prop = 'id'
    }
    searchInfo.value.orderKey = toSQLLine(prop)
    searchInfo.value.desc = order === 'descending'
  }
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getClusterList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 批量操作
const handleSelectionChange = (val) => {
  apis.value = val
}

const deleteVisible = ref(false)
const onDelete = async() => {
  const ids = apis.value.map(item => item.ID)
  const res = await deleteApisByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: res.msg
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}
const freshVisible = ref(false)
const onFresh = async() => {
  const res = await freshCasbin()
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: res.msg
    })
  }
  freshVisible.value = false
}

// 弹窗相关
const apiForm = ref(null)
const initForm = () => {
  apiForm.value.resetFields()
  form.value = {
    method: '',
    description: ''
  }
}

const dialogTitle = ref('新增cluster')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
  switch (key) {
    case 'addApi':
      dialogTitle.value = '新增cluster'
      break
    case 'edit':
      dialogTitle.value = '编辑cluster'
      break
    default:
      break
  }
  type.value = key
  dialogFormVisible.value = true
}
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}

const editApiFunc = async(row) => {
  const res = await getApiById({ id: row.ID })
  form.value = res.data.api
  openDialog('edit')
}

const enterDialog = async() => {
  apiForm.value.validate(async valid => {
    if (valid) {
      switch (type.value) {
        case 'addApi': {
          const res = await createCluster(form.value)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '添加成功',
              showClose: true
            })
          }
          getTableData()
          closeDialog()
        }

          break
        case 'edit': {
          const res = await updateApi(form.value)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '编辑成功',
              showClose: true
            })
          }
          getTableData()
          closeDialog()
        }
          break
        default:
          // eslint-disable-next-line no-lone-blocks
          {
            ElMessage({
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

const deleteApiFunc = async(row) => {
  ElMessageBox.confirm('此操作将会删除该cluster以及配置文件?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      console.log(row)
      const res = await deleteCluster(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
}

</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
</style>
