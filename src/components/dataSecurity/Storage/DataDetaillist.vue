<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/welcome' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>数据安全</el-breadcrumb-item>
      <el-breadcrumb-item>存储阶段</el-breadcrumb-item>
      <el-breadcrumb-item>数据清单</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片区 -->
    <el-card>
      <!-- 筛选区域 -->
      <el-tabs v-model="queryInfo.GroupName" type="card" @tab-click="handleClick">
        <el-tab-pane label="All" name="All"></el-tab-pane>
        <el-tab-pane label="APP" name="APP"></el-tab-pane>
        <el-tab-pane label="DAP" name="DAP"></el-tab-pane>
        <el-tab-pane label="FMC" name="FMC"></el-tab-pane>
        <el-tab-pane label="ISDP" name="ISDP"></el-tab-pane>
      </el-tabs>
      <!-- 搜索区域 -->
      <el-row :gutter="15">
        <el-col :span="4">
          <el-select v-model="queryInfo.RiskLevelName" clearable placeholder="筛选敏感数据等级" @change="handleClick">
            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-col>
        <el-col :span="8">
          <el-input placeholder="请输入想查询的规则名" v-model="queryInfo.RuleName" clearable @clear="getColumnlist" @keyup.enter.native="getColumnlist" />
        </el-col>
        <el-col :span="4">
          <el-button @click="getColumnlist" type="primary">搜 索</el-button>
          <el-button @click="clearAll">重 置</el-button>
        </el-col>
      </el-row>
      <!-- 数据返回区域 -->
      <el-table :data="columnlist" stripe style="width: 100%">
        <el-table-column label="项目组" prop="GroupName" width="100px" sortable/>
        <el-table-column label="实例名" prop="InstanceName" sortable/>
        <el-table-column label="字段名称" prop="Name" sortable/>
        <el-table-column label="数据类型" prop="DataType" sortable/>
        <el-table-column label="命中规则" prop="RuleName" sortable/>
        <el-table-column label="数据等级" sortable>
          <template slot-scope="scope">
            <el-tag v-if="scope.row.RiskLevelName === 'S4'" type="danger">S4</el-tag>
            <el-tag v-else-if="scope.row.RiskLevelName === 'S3'" type="warning">S3</el-tag>
            <el-tag v-else-if="scope.row.RiskLevelName === 'S2'">S2</el-tag>
            <el-tag v-else-if="scope.row.RiskLevelName === 'S1'" type="success">S1</el-tag>
            <el-tag v-else type="info">N/A</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="数据表名" sortable>
          <template slot-scope="scope">
            <el-link type="primary" icon="el-icon-view">{{ scope.row.TableName }}</el-link>
          </template>
        </el-table-column>
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
      activeName: '',
      input: '',
      // 获取数据资产列表
      queryInfo: {
        GroupName: 'All',
        RiskLevelName: '',
        RuleName: '',
        PageNum: 1,
        PageSize: 10
      },
      // 返回资产数据
      columnInfo: {
        UUID: '',
        GroupName: '',
        CreationTime: '',
        DataType: '',
        Id: 0,
        InstanceId: '',
        InstanceName: '',
        Name: '',
        ProductCode: '',
        RevisionId: 0,
        RevisionStatus: 0,
        RiskLevelId: 0,
        RiskLevelName: '',
        RuleName: '',
        SensLevelName: '',
        Sensitive: '',
        TableId: '',
        TableName: '',
        TotalCount: 0
      },
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
      const { data: res } = await this.$http.post('column', this.queryInfo,
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
