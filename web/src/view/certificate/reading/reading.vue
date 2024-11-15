<template>
  <div class="gva-search-box">
    <el-form
        ref="searchForm"
        :inline="true"
        :model="searchInfo"
    >
      <el-form-item label="姓名">
        <el-input
            v-model="searchInfo.name"
            placeholder="姓名"
        />
      </el-form-item>
      <el-form-item label="证书编号">
        <el-input
            v-model="searchInfo.number"
            placeholder="证书编号"
        />
      </el-form-item>
      <el-form-item label="手机号">
        <el-input
            v-model="searchInfo.mobile"
            placeholder="手机号"
        />
      </el-form-item>
      <el-form-item label="身份证号">
        <el-input
            v-model="searchInfo.idCard"
            placeholder="身份证号"
        />
      </el-form-item>
      <el-form-item>
        <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
        >
          查询
        </el-button>
        <el-button
            icon="refresh"
            @click="onReset"
        >
          重置
        </el-button>
      </el-form-item>
    </el-form>
  </div>
  <div class="gva-table-box">
    <div class="gva-btn-list">
      <el-button
          type="primary"
          icon="plus"
          @click="addReading"
      >新增证书</el-button>
    </div>
    <el-table
        :data="tableData"
        row-key="ID"
    >
      <el-table-column
          align="left"
          label="头像"
          min-width="75"
      >
        <template #default="scope">
          <CustomPic
              style="margin-top:8px"
              :pic-src="scope.row.profile"
          />
        </template>
      </el-table-column>
      <el-table-column
          align="left"
          label="ID"
          min-width="90"
          prop="ID"
      />
      <el-table-column
          align="left"
          label="姓名"
          min-width="150"
          prop="name"
      />
      <el-table-column
          align="left"
          label="证书编码"
          min-width="150"
          prop="number"
      />
      <el-table-column
          align="left"
          label="手机号"
          min-width="150"
          prop="mobile"
      />
      <el-table-column
          align="left"
          label="身份证号"
          min-width="180"
          prop="idCard"
      />
      <el-table-column
          align="left"
          label="所属组"
          min-width="150"
          prop="group"
      >
        <template #default="{ row }">
          <!-- 使用group值查找对应的label -->
          {{ getGroupLabel(row.group) }}
        </template>
      </el-table-column>
      <el-table-column
          align="left"
          label="推荐人"
          min-width="200"
          prop="referrer"
      >
      </el-table-column>
      <el-table-column
          label="操作"
          min-width="240"
          fixed="right"
      >
        <template #default="scope">
          <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteReadingFunc(scope.row)"
          >删除</el-button>
          <el-button
              type="primary"
              link
              icon="edit"
              @click="openEdit(scope.row)"
          >编辑</el-button>
          <el-button
              type="primary"
              link
              icon="edit"
              @click="download(scope.row)"
          >下载证书</el-button>
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
  <el-drawer
      v-model="addReadingDialog"
      size="60%"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
  >
    <template #header>
      <div class="flex justify-between items-center">
        <span class="text-lg">新增证书</span>
        <div>
          <el-button @click="closeAddReadingDialog">取 消</el-button>
          <el-button
              type="primary"
              @click="enterAddUserDialog"
          >确 定</el-button>
        </div>
      </div>
    </template>

    <el-form
        ref="readingForm"
        :rules="rules"
        :model="readingInfo"
        label-width="80px"
    >
      <el-form-item
          label="姓名"
          prop="name"
      >
        <el-input v-model="readingInfo.name" :disabled="dialogFlag === 'edit'" />
      </el-form-item>
      <el-form-item
          label="证书编号"
          prop="number"
      >
        <el-input v-model="readingInfo.number" :disabled="dialogFlag === 'edit'" />
      </el-form-item>
      <el-form-item
          label="手机号"
          prop="mobile"
      >
        <el-input v-model="readingInfo.mobile" />
      </el-form-item>
      <el-form-item
          label="身份证号"
          prop="idCard"
      >
        <el-input v-model="readingInfo.idCard" :disabled="dialogFlag === 'edit'" />
      </el-form-item>
      <el-form-item
          label="所属组"
          prop="group"
      >
        <el-select
            v-model="readingInfo.group"
            style="width: 100%"
            placeholder="选择所属组"
        >
          <el-option
              v-for="item in groupOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
          />
        </el-select>
<!--        <el-input v-model="readingInfo.group" />-->
      </el-form-item>
      <el-form-item
          label="推荐人"
          prop="referrer"
      >
        <el-input v-model="readingInfo.referrer" />

      </el-form-item>
      <el-form-item
          label="备注"
          prop="remark"
      >
        <el-input v-model="readingInfo.remark" />
      </el-form-item>
      <el-form-item
          label="状态"
          prop="disabled"
      >
        <el-switch
            v-model="readingInfo.status"
            inline-prompt
            :active-value="1"
            :inactive-value="2"
        />
      </el-form-item>
      <el-form-item
          label="头像"
          label-width="80px"
      >
        <SelectImage
            v-model="readingInfo.profile"
        />
      </el-form-item>
      <el-form-item
          label="证书文件"
          label-width="80px"
          prop="file"
      >
        <SelectImage
            v-model="readingInfo.file"
        />
      </el-form-item>
    </el-form>
  </el-drawer>
</template>

<script setup>
import {ref} from "vue";
import {getReadingList, createReading, editReading, deleteReading} from "@/api/certificate";
import CustomPic from "@/components/customPic/index.vue";
import SelectImage from "@/components/selectImage/selectImage.vue";
import {ElMessage, ElMessageBox} from "element-plus";

defineOptions({name: 'Reading'})

const searchInfo = ref({
  name: '',
  number: '',
  mobile: '',
  idCard: ''
})

const onSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = {
    name: '',
    number: '',
    mobile: '',
    idCard: ''
  }
  getTableData()
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 弹窗相关
const readingInfo = ref({
  name: '',
  number: '',
  mobile: '',
  idCard: '',
  group: '',
  referrer: '',
  remark: '',
  file: '',
  profile: '',
  status: 1,
})

const dialogFlag = ref('add')
const addReadingDialog = ref(false)

const addReading = () => {
  dialogFlag.value = 'add'
  addReadingDialog.value = true
}

const closeAddReadingDialog = () => {
  readingForm.value.resetFields()
  addReadingDialog.value = false
}

const deleteReadingFunc = async(row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteReading({ id: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      await getTableData()
    }
  })
}

const rules = ref({
  name: [
    { required: true, message: '请输入姓名', trigger: 'blur' },
    { min: 5, message: '最低2位字符', trigger: 'blur' }
  ],
  number: [
    { required: true, message: '请输入整数编码', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur' }
  ],
  mobile: [
    { required: true, pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/, message: '请输入正确的手机号', trigger: 'blur' }
  ],
  idCard: [
    { required: true, message: '请输入身份证号', trigger: 'blur' },
  ],
  file: [
    { required: true, message: '请上传证书文件', trigger: 'blur' },
  ],
  group: [
    { required: true, message: '请选择所属组', trigger: 'blur' }
  ]
})

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  readingInfo.value = JSON.parse(JSON.stringify(row))
  addReadingDialog.value = true
}

const download = (row) => {
  const baseUrl = import.meta.env.VITE_BASE_PATH + ":" + import.meta.env.VITE_CLI_PORT + import.meta.env.VITE_BASE_API + '/';

  const fileUrl = baseUrl + row.file
  // 创建一个 <a> 标签，模拟点击下载
  const link = document.createElement('a');
  link.href = fileUrl;
  link.download = row.number || '证书'; // 设置下载的文件名，使用证书编号或者其他名称
  document.body.appendChild(link); // 将 link 标签添加到 DOM 中
  link.click(); // 模拟点击
  document.body.removeChild(link); // 下载后移除 <a> 标签
}

const groupOptions = ref([
    {value: '1', label: '一年级'},
    {value: '2', label: '二年级'},
    {value: '3', label: '三年级'},
    {value: '4', label: '四年级'},
    {value: '5', label: '五年级'},
    {value: '6', label: '六年级'},
    {value: '7', label: '七年级'},
    {value: '8', label: '八年级'},
    {value: '9', label: '九年级'},
])

function getGroupLabel(groupValue) {
  const group = this.groupOptions.find(option => option.value === groupValue);
  return group ? group.label : '未知组'; // 如果没有找到对应项，返回'未知组'
}

const readingForm = ref(null)
const enterAddUserDialog = async() => {
  readingForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...readingInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await createReading(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeAddReadingDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await editReading(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeAddReadingDialog()
        }
      }
    }
  })
}

// 查询
const getTableData = async() => {
  const table = await getReadingList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const initPage = async() => {
  getTableData()
}

initPage()

</script>
<style scoped lang="scss">

</style>