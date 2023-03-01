<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/admin' }">Home Page</el-breadcrumb-item>
      <el-breadcrumb-item>Data Security</el-breadcrumb-item>
      <el-breadcrumb-item>Configuration</el-breadcrumb-item>
      <el-breadcrumb-item>Category</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片区 -->
    <el-card>
      <el-row :gutter="15">
        <el-col :span="4">
          <el-input placeholder="Input CategoryName" v-model="queryInfo.CategoryName" clearable @clear="getCategoryList" @keyup.enter.native="getCategoryList" />
        </el-col>
        <el-col :span="4">
          <el-button @click="getCategoryList" type="primary" icon="el-icon-search">Search</el-button>
          <el-button type="danger" icon="el-icon-plus">New</el-button>
        </el-col>
      </el-row>
      <!-- 数据返回区域 -->
      <el-table :data="CategoryList" stripe style="width: 50%">
        <el-table-column type="index" width="50"/>
        <el-table-column label="CategoryName" align="left" prop="CategoryName"/>
        <el-table-column label="Level" align="center" prop="level" sortable/>
        <el-table-column label="Parent Category" align="center" prop="ParentCategory"/>
        <el-table-column label="Comments" align="center" prop="Comments"/>
      </el-table>
      <!-- 分页区域 -->
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="queryInfo.PageNum"
        :page-sizes="[5, 10, 20, 50]"
        :page-size="queryInfo.PageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total=total>
      </el-pagination>
    </el-card>
  </div>
</template>

<script>
export default {
  data () {
    return {
      input: '',
      // 获取数据资产列表
      queryInfo: {
        CategoryName: '',
        PageNum: 1,
        PageSize: 10
      },
      // 返回资产数据
      CategoryInfo: {
        UUID: '',
        CategoryName: '',
        Level: '',
        ParentCategory: ''
      },
      CategoryList: [],
      total: 0
    }
  },
  created () {
    this.getCategoryList()
  },
  methods: {
    async getCategoryList () {
      const { data: res } = await this.$http.post('Category/List', this.queryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('获取列表失败')
      this.CategoryList = res.Data
      this.total = res.categoryTotal
    },
    // 监听
    handleSizeChange (newSize) {
      this.queryInfo.pageSize = newSize
      this.getCategoryList()
    },
    handleCurrentChange (newPage) {
      this.queryInfo.pageNum = newPage
      this.getCategoryList()
    },
    handleClick () {
      this.getCategoryList()
    },
    clearAll () {
      this.queryInfo.CategoryName = ''
      this.getCategoryList()
    }
  },
  name: 'ConfigCategory'
}
</script>

<style scoped>

</style>
