<template>
  <div class="app-container">
    <!-- 查询表单 -->
    <div class="filter-container">
      <el-form ref="queryForm" :model="queryParams" :inline="true">
        <el-form-item label="球员名称" prop="name">
          <el-input
            v-model="queryParams.name"
            placeholder="请输入球员名称"
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
      <el-table-column label="球员名称" prop="name" align="center" />
      <el-table-column label="球员号码" prop="number" align="center" />
      <el-table-column label="描述" prop="description" align="center" />
      <el-table-column label="身高" prop="height" align="center" />
      <el-table-column label="体重" prop="weight" align="center" />
      <el-table-column label="位置" prop="position" align="center" />
      <el-table-column label="球队" prop="team_name" align="center" />
      <el-table-column label="状态" prop="status" align="center" />
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
        <el-form-item label="球员名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入球员名称" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" placeholder="请输入球员描述" />
        </el-form-item>
        <el-form-item label="球员号码" prop="number">
          <el-input v-model.number="form.number" placeholder="请输入球员号码" />
        </el-form-item>
        <el-form-item label="身高" prop="height">
          <el-input v-model.number="form.height" placeholder="请输入球员身高" />
        </el-form-item>
        <el-form-item label="体重" prop="weight">
          <el-input v-model.number="form.weight" placeholder="请输入球员体重" />
        </el-form-item>
        <el-form-item label="位置" prop="position">
          <el-input v-model="form.position" placeholder="请输入球员位置" />
        </el-form-item>
        <el-form-item label="球队" prop="team_id">
          <el-select v-model="form.team_id" placeholder="请选择所属球队">
            <el-option v-for="item in teamList" :key="item.id" :label="item.name" :value="item.id" />
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
import { list, submit } from '@/api/player'
import { fetchList } from '@/api/team'

export default {
  name: '',
  data() {
    return {
      teamList: [],
      tableData: [],
      total: 0,
      title: '球员模块',
      loading: false,
      queryParams: {
        name: ''
      },
      dialogTitle: '',
      dialogVisible: false,
      form: {
        id: 0,
        name: '',
        description: '',
        number: '',
        height: null,
        weight: null,
        position: '',
        team_id: null,
        status: 1
      },
      rules: {
        name: [{ required: true, message: '请输入球员名称', trigger: 'blur' }],
        description: [{ required: false, message: '请输入球员描述', trigger: 'blur' }],
        number: [{ required: true, message: '球员号码不能为空' }, { type: 'number', message: '球员号码必须为数字值' }],
        height: [{ required: false, message: '请输入球员身高', trigger: 'blur' }, { type: 'number', message: '请输入数字', trigger: 'blur' }],
        weight: [{ required: false, message: '请输入球员体重', trigger: 'blur' }, { type: 'number', message: '请输入数字', trigger: 'blur' }],
        position: [{ required: false, message: '请输入球员位置', trigger: 'blur' }],
        team_id: [{ required: false, message: '请选择所属球队', trigger: 'change' }],
        status: [{ required: true, message: '请选择状态', trigger: 'change' }]
      }
    }
  },
  mounted() {
    this.getList()
  },
  methods: {
    openHandle() {
      this.getTeamList()
    },
    getTeamList() {
      fetchList()
        .then((res) => {
          this.teamList = res.data
        })
        .catch(() => {
          this.$message({
            message: '获取球队列表失败',
            type: 'error'
          })
        })
    },
    getList() {
      this.loading = true
      list(this.queryParams)
        .then((res) => {
          this.loading = false
          this.tableData = res.data.players
          this.total = res.data.total
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
      this.dialogTitle = '新增球员'
      this.dialogVisible = true
    },
    handleEdit(row) {
      this.dialogTitle = '编辑球员'
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
      })
    },

    resetQuery() {
      this.queryParams.name = ''
      this.getList()
    }

  }
}
</script>
