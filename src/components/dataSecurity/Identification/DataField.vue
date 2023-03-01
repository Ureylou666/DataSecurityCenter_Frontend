<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/admin' }">Home</el-breadcrumb-item>
      <el-breadcrumb-item>Data Security</el-breadcrumb-item>
      <el-breadcrumb-item>Identification</el-breadcrumb-item>
      <el-breadcrumb-item>Data Field</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片区 -->
    <el-card>
      <!-- 筛选区域 -->
      <el-tabs v-model="queryInfo.GroupName" type="card" @tab-click="handleClick">
        <el-tab-pane label="All" name="All"></el-tab-pane>
        <el-tab-pane label="APP" name="APP"></el-tab-pane>
        <el-tab-pane label="DAP" name="DAP"></el-tab-pane>
        <el-tab-pane label="FMC" name="FMC"></el-tab-pane>
        <el-tab-pane label="O2O" name="O2O"></el-tab-pane>
        <el-tab-pane label="AID" name="AID"></el-tab-pane>
        <el-tab-pane label="CSC" name="CSC"></el-tab-pane>
        <el-tab-pane label="Promotion" name="Promotion"></el-tab-pane>
        <el-tab-pane label="ISDP" name="ISDP"></el-tab-pane>
      </el-tabs>
      <!-- 搜索区域 -->
      <el-row :gutter="15">
        <el-col :span="2">
          <el-select v-model="queryInfo.SensLevelName" clearable placeholder="Filter" @change="handleClick">
            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-col>
        <el-col :span="2">
          <el-select v-model="queryInfo.CategoryName" clearable placeholder="Filter" @change="handleClick">
            <el-option v-for="item in Category" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-col>
        <el-col :span="8">
          <el-input placeholder="Input query rule name" v-model="queryInfo.RuleName" clearable @clear="getColumnlist" @keyup.enter.native="getColumnlist" />
        </el-col>
        <el-col :span="8">
          <el-button @click="getColumnlist" type="primary">Search</el-button>
          <el-button @click="clearAll">Reset</el-button>
        </el-col>
      </el-row>
      <!-- 数据返回区域 -->
      <el-table :data="columnlist" stripe style="width: 100%">
        <el-table-column label="Group" prop="GroupName" width="100px" sortable/>
        <el-table-column label="Instance ID" prop="InstanceId" sortable/>
        <el-table-column label="Database" prop="DatabaseName" sortable/>
        <el-table-column label="Table" prop="TableName" sortable/>
        <el-table-column label="Column" prop="ColumnName" sortable/>
        <el-table-column label="Data Type" prop="DataType" sortable/>
        <el-table-column label="Category" prop="CategoryName" sortable/>
        <el-table-column label="Rule Name" prop="RuleName" sortable/>
        <el-table-column label="Severity" sortable>
          <template slot-scope="scope">
            <el-tag v-if="scope.row.SensLevelName === 'S4'" type="danger">S4</el-tag>
            <el-tag v-else-if="scope.row.SensLevelName === 'S3'" type="warning">S3</el-tag>
            <el-tag v-else-if="scope.row.SensLevelName === 'S2'">S2</el-tag>
            <el-tag v-else-if="scope.row.SensLevelName === 'S1'" type="success">S1</el-tag>
            <el-tag v-else type="info">N/A</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="Methods">
          <template slot-scope="scope2">
            <el-popover placement="left" title="Sample Data" trigger="click">
              <p>{{ scope2.row.SampleData }}</p>
              <el-button slot="reference" icon="el-icon-search" circle></el-button>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页区域 -->
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="queryInfo.PageNum"
        :page-sizes="[10, 20, 50, 100]"
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
      activeName: '',
      input: '',
      // 获取数据资产列表
      queryInfo: {
        GroupName: 'All',
        SensLevelName: '',
        RuleName: '',
        PageNum: 1,
        PageSize: 20
      },
      // 返回资产数据
      columnInfo: {
        UUID: '',
        GroupName: '',
        DataType: '',
        InstanceId: '',
        DatabaseName: '',
        TableName: '',
        ColumnName: '',
        RuleId: null,
        RuleName: '',
        SensLevelName: '',
        Sensitive: '',
        SampleData: '',
        TotalCount: 0
      },
      Category: [
        { value: 'StrictlyConfidential', label: 'StrictlyConfidential' },
        { value: 'Confidential', label: 'Confidential' },
        { value: 'Internal', label: 'Internal' },
        { value: 'Public', label: 'Public' }
      ],
      options: [
        { value: '', label: 'All' },
        { value: 'S4', label: 'S4' },
        { value: 'S3', label: 'S3' },
        { value: 'S2', label: 'S2' },
        { value: 'S1', label: 'S1' }
      ],
      columnlist: [],
      total: 0
    }
  },
  created () {
    this.getColumnlist()
  },
  methods: {
    async getColumnlist () {
      const { data: res } = await this.$http.post('DataFields', this.queryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('获取资产列表失败')
      this.columnlist = res.Res_Data
      this.total = res.Column_Total
    },
    // 监听
    handleSizeChange (newSize) {
      this.queryInfo.pageSize = newSize
      this.getColumnlist()
    },
    handleCurrentChange (newPage) {
      this.queryInfo.pageNum = newPage
      this.getColumnlist()
    },
    handleClick () {
      this.getColumnlist()
    },
    clearAll () {
      this.queryInfo.GroupName = 'All'
      this.queryInfo.RiskLevelName = 'All'
      this.queryInfo.RuleName = ''
      this.getColumnlist()
    }
  },
  name: 'DataDetaillist'
}
</script>

<style scoped>

</style>
