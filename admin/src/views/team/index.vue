<template>
  <div class="app-container">
    <!-- 查询表单 -->
    <div class="filter-container">
      <el-form ref="queryForm" :model="queryParams" :inline="true">
        <el-form-item label="球队名称" prop="name">
          <el-input
            v-model="queryParams.name"
            placeholder="请输入球队名称"
            clearable
            style="width: 200px;"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            icon="el-icon-search"
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
      <el-table-column label="球队名称" prop="name" align="center" />
      <el-table-column label="所属联赛" prop="league_name" align="center" />
      <el-table-column label="描述" prop="description" align="center" />
      <el-table-column label="状态" prop="status" align="center" />
      <el-table-column label="创建时间" prop="create_time" align="center" />
      <el-table-column label="更新时间" prop="update_time" align="center" />
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

    <!-- 新增/编辑对话框 -->
    <el-dialog :title="dialogTitle" :visible.sync="dialogVisible" width="500px" @open="openHandle">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="球队名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入球队名称" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" placeholder="请输入球队描述" />
        </el-form-item>
        <el-form-item label="所属联赛" prop="league_id">
          <el-select v-model="form.league_id" placeholder="请选择所属联赛">
            <el-option v-for="item in leagueList" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">正常</el-radio>
            <el-radio :label="0">禁用</el-radio>
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
import { fetchList, submit } from '@/api/team'

export default {
  name: '',
  data() {
    return {
      leagueList: [],
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
        id: '',
        name: '',
        description: '',
        league_id: '',
        status: 1
      },
      rules: {
        name: [{ required: true, message: '请输入球队名称', trigger: 'blur' }],
        description: [{ required: false, message: '请输入球队描述', trigger: 'blur' }],
        league_id: [{ required: false, message: '请选择所属联赛', trigger: 'change' }],
        status: [{ required: true, message: '请选择状态', trigger: 'change' }]
      }
    }
  },
  mounted() {
    this.getList()
  },
  methods: {
    openHandle() {
      this.getLeagueList()
    },
    getLeagueList() {
      this.$store.dispatch('league/getList')
        .then(() => {
          this.leagueList = this.$store.state.league.list
        })
        .catch(() => {
          this.$message({
            message: '获取联赛列表失败',
            type: 'error'
          })
        })
    },
    getList() {
      this.loading = true
      fetchList(this.queryParams)
        .then((res) => {
          this.loading = false
          this.tableData = res.data
          this.total = res.total
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
        if (!valid) {
          this.$message({
            message: '请填写完整信息',
            type: 'warning'
          })

          return
        }

        submit(this.form)
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
      })
    },

    resetQuery() {
      this.queryParams.name = ''
      this.getList()
    }

  }
}
</script>
