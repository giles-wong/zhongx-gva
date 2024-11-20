<template>
  <div class="gva-search-box">
    <el-form
        ref="searchForm"
        :inline="true"
        :model="searchInfo"
    >
      <el-form-item label="书名">
        <el-input
            v-model="searchInfo.name"
            placeholder="书名"
        />
      </el-form-item>
      <el-form-item label="标准书号">
        <el-input
            v-model="searchInfo.isbn"
            placeholder="标准书号"
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
          @click="addBooks"
      >新增图书</el-button>
    </div>
    <el-table
        :data="tableData"
        row-key="ID"
    >
      <el-table-column
          align="left"
          label="封面"
          min-width="75"
      >
        <template #default="scope">
          <CustomPic
              style="margin-top:8px"
              :pic-src="scope.row.coverImage"
          />
        </template>
      </el-table-column>
      <el-table-column
          align="left"
          label="ID"
          min-width="175"
          prop="bookId"
      />
      <el-table-column
          align="left"
          label="书名"
          min-width="190"
          prop="bookName"
      />
      <el-table-column
          align="left"
          label="国际书号"
          min-width="185"
          prop="isbn"
      />
      <el-table-column
          align="left"
          label="作者"
          min-width="150"
          prop="author"
      />
      <el-table-column
          align="left"
          label="定价"
          min-width="95"
          prop="price"
      />
      <el-table-column
          align="left"
          label="出版社"
          min-width="150"
          prop="publisher"
      />
      <el-table-column
          align="left"
          label="出版时间"
          min-width="130"
          prop="publicationDate"
      />
      <el-table-column
          label="操作"
          min-width="160"
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
      v-model="addBookDialog"
      size="60%"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
  >
    <template #header>
      <div class="flex justify-between items-center">
        <span class="text-lg" v-if="dialogFlag === 'add'">新增图书</span>
        <span class="text-lg" v-if="dialogFlag === 'edit'">编辑图书</span>
        <div>
          <el-button @click="closeAddBookDialog">取 消</el-button>
          <el-button
              type="primary"
              @click="enterAddBookDialog"
          >确 定</el-button>
        </div>
      </div>
    </template>

    <el-form
        ref="bookForm"
        :rules="rules"
        :model="bookInfo"
        label-width="80px"
    >
      <el-form-item label="ID" prop="bookId">
        <el-input v-model="bookInfo.bookId" disabled/>
      </el-form-item>
      <el-form-item label="书名" prop="bookName">
        <el-input v-model="bookInfo.bookName"/>
      </el-form-item>
      <el-form-item label="国际书号" prop="isbn">
        <el-input v-model="bookInfo.isbn"/>
      </el-form-item>
      <el-form-item label="定价" prop="price">
        <el-input v-model="bookInfo.price"/>
      </el-form-item>
      <el-form-item label="作者" prop="author">
        <el-input v-model="bookInfo.author"/>
      </el-form-item>
      <el-form-item label="出版社" prop="publisher">
        <el-input v-model="bookInfo.publisher"/>
      </el-form-item>
      <el-form-item label="出版时间" prop="publicationDate">
        <el-date-picker
            v-model="bookInfo.publicationDate"
            type="date"
            placeholder="选择日期"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
        />
      </el-form-item>
      <el-form-item
          label="图书封面"
          label-width="80px"
          prop="coverImage"
      >
        <SelectImage
            v-model="bookInfo.coverImage"
        />
      </el-form-item>
      <el-form-item label="图书简介" prop="summary">
        <el-input
            type="textarea"
            :rows="2"
            placeholder="请输入内容"
            v-model="bookInfo.summary"
        />
      </el-form-item>
      <el-form-item label="备注" prop="remark">
        <el-input v-model="bookInfo.remark"/>
      </el-form-item>
    </el-form>
  </el-drawer>
</template>

<script setup>
  import {ref} from "vue";
  import {createBook, editBook, getBooksList} from "@/api/books";
  import CustomPic from "@/components/customPic/index.vue";
  import SelectImage from "@/components/selectImage/selectImage.vue";
  import {ElMessage} from "element-plus";

  defineOptions({name: 'Books'})

  const searchInfo = ref({
    name: '',
    isbn: '',
  })
  // 筛选项提交
  const onSubmit = () => {
    page.value = 1
    getTableData()
  }
  // 筛选项重置
  const onReset = () => {
    searchInfo.value = {
      name: '',
      isbn: '',
    }
    getTableData()
  }
  // 弹窗相关

  const bookInfo = ref({
    bookId:0,
    bookName: '',
    isbn: '',
    author: '',
    publisher: '',
    publicationDate: '',
    price: '',
    coverImage: '',
    summary:'',
    remark: '',
    status: 1,
  })

  const dialogFlag = ref('add')
  const addBookDialog = ref(false)

  const addBooks = () => {
    dialogFlag.value = 'add'
    addBookDialog.value = true
  }

  const openEdit = (row) => {
    dialogFlag.value = 'edit'
    bookInfo.value = JSON.parse(JSON.stringify(row))
    addBookDialog.value = true
  }
  const bookForm = ref(null)
  const closeAddBookDialog = () => {
    bookForm.value.resetFields()
    bookForm.value.file = []
    bookForm.value = {}
    addBookDialog.value = false // 关闭弹窗
  }

  const enterAddBookDialog = async() => {
    bookForm.value.validate(async valid => {
      if (valid) {
        const req = {
          ...bookInfo.value
        }

        if (req.publicationDate === '') {
          delete req.publicationDate;
        }

        if (dialogFlag.value === 'add') {
          const res = await createBook(req)
          if (res.code === 0) {
            ElMessage({ type: 'success', message: '创建成功' })
            await getTableData()
            closeAddBookDialog()
          }
        }
        if (dialogFlag.value === 'edit') {
          const res = await editBook(req)
          if (res.code === 0) {
            ElMessage({ type: 'success', message: '编辑成功' })
            await getTableData()
            closeAddBookDialog()
          }
        }
      }
    })
  }

  const rules = ref({
    bookName: [
      { required: true, message: '请输入书名', trigger: 'blur' },
      { min: 2, message: '最低2位字符', trigger: 'blur' }
    ],
    isbn: [
      { required: true, message: '请输入国际书号', trigger: 'blur' },
      { min: 10, message: '最低10位字符', trigger: 'blur' }
    ],
    author: [
      { required: true, message: '请输入作者', trigger: 'blur' },
    ],
    publisher: [
      { required: true, message: '请输入出版社', trigger: 'blur' },
    ],
    price: [
      { required: true, message: '请输入单价', trigger: 'blur' },
      {
        validator: (rule, value, callback) => {
          const price = parseFloat(value);
          if (isNaN(price)) {
            callback(new Error('请输入有效的数字'));
          } else if (price <= 0) {
            callback(new Error('单价必须大于0'));
          } else if (!/^\d+(\.\d{1,2})?$/.test(value)) {
            callback(new Error('单价必须是最多两位小数的数字'));
          } else {
            callback();
          }
        },
        trigger: 'blur'
      }
    ],
    coverImage: [
      { required: true, message: '请输入封面', trigger: 'blur' },
    ],
  })

  // 页码初始化
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

  // 查询
  const getTableData = async() => {
    const table = await getBooksList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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