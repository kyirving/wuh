<template>
  <div class="app-container">
    <!-- 查询表单 -->
    <div class="filter-container">
      <el-form ref="queryForm" :model="queryParams" :inline="true">
        <el-form-item label="联赛名称" prop="name">
          <el-input
            v-model="queryParams.username"
            placeholder="请输入用户名"
            clearable
            style="width: 200px;"
            @keyup.enter.native="handleQuery"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            icon="el-icon-search"
            @click="handleQuery"
          >
            搜索
          </el-button>
          <el-button
            icon="el-icon-refresh"
            @click="resetQuery"
          >
            重置
          </el-button>
          <el-button
            type="success"
            icon="el-icon-plus"
            @click="handleAdd"
          >
            新增
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 表格 -->
    <el-table
      v-loading="loading"
      :data="tableData"
      @selection-change="handleSelectionChange"
    >

      <el-table-column label="ID" prop="id" align="center" />
      <el-table-column label="name" prop="name" align="center" />
      <el-table-column label="描述" prop="description" align="center" />
      <el-table-column label="状态" prop="status" align="center" />
      <el-table-column label="创建时间" prop="created_at" align="center" />
      <el-table-column label="更新时间" prop="updated_at" align="center" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="{row}">
          <el-button
            type="text"
            icon="el-icon-edit"
            @click="handleEdit(row)"
          >
            编辑
          </el-button>
          <el-button
            type="text"
            icon="el-icon-delete"
            style="color: #f56c6c;"
            @click="handleDelete(row)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryParams.pageNum"
      :limit.sync="queryParams.pageSize"
      @pagination="getList"
    />

    <!-- 新增/编辑对话框 -->
    <el-dialog :title="dialogTitle" :visible.sync="dialogVisible" width="500px">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="联赛名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入联赛名称" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" placeholder="请输入联赛描述" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="0">正常</el-radio>
            <el-radio :label="1">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitForm">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import Pagination from '@/components/Pagination'

export default {
  name: '',
  components: { Pagination },
  data() {
    return {

      tableData: [],
      total: 0,
      title: '联赛模块',
      loading: false,
      queryParams: {
        name: ''
      },
      dialogTitle: '',
      dialogVisible: false,
      form: {
        name: '',
        description: ''
      },
      rules: {
        name: [{ required: true, message: '请输入联赛名称', trigger: 'blur' }],
        description: [{ required: true, message: '请输入联赛描述', trigger: 'blur' }],
        status: [{ required: true, message: '请选择状态', trigger: 'change' }]
      }
    }
  },
  mounted() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      this.$store.dispatch('league/getList')
        .then(() => {
          this.loading = false
          this.tableData = this.$store.state.league.list
          this.total = this.$store.state.league.total
        })
        .catch(() => {
          this.loading = false
        })
    },
    handleSelectionChange() {

    },
    getTitle() {
      return this.title
    },

    handleAdd() {
      this.dialogTitle = '新增联赛'
      this.dialogVisible = true
    },
    handleEdit(row) {
      this.dialogTitle = '编辑联赛'
      this.dialogVisible = true
      this.form = { ...row }
    },

    submitForm() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          this.$store.dispatch('league/submit', this.form)
            .then(() => {
              this.$message({
                message: '操作成功',
                type: 'success'
              })
              this.dialogVisible = false
              this.getList()
            })
            .catch(() => {
              this.$message({
                message: '操作失败',
                type: 'error'
              })
            })
        } else {
          this.$message({
            message: '请填写完整信息',
            type: 'warning'
          })
        }
      })
    }

  }
}
</script>
