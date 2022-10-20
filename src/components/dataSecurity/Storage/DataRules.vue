<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/welcome' }">Home Page</el-breadcrumb-item>
      <el-breadcrumb-item>Data Security</el-breadcrumb-item>
      <el-breadcrumb-item>Identification</el-breadcrumb-item>
      <el-breadcrumb-item>Classification</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片区 -->
    <el-card>
      <el-row :gutter="15">
        <el-col :span="4">
          <el-select v-model="queryInfo.RiskLevelName" clearable placeholder="筛选敏感数据等级" @change="handleClick">
            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-col>
        <el-col :span="8">
          <el-input placeholder="请输入想查询的规则名" v-model="queryInfo.RuleName" clearable @clear="getRuleslist" @keyup.enter.native="getRuleslist" />
        </el-col>
        <el-col :span="4">
          <el-button @click="getRuleslist" type="primary">搜 索</el-button>
          <el-button @click="clearAll">重 置</el-button>
        </el-col>
      </el-row>
      <!-- 数据返回区域 -->
      <el-table :data="ruleslist" stripe style="width: 100%">
        <el-table-column label="数据等级" width="100px" align="center" sortable>
          <template slot-scope="scope">
            <el-tag v-if="scope.row.RiskLevelName === 'S4'" type="danger">S4</el-tag>
            <el-tag v-else-if="scope.row.RiskLevelName === 'S3'" type="warning">S3</el-tag>
            <el-tag v-else-if="scope.row.RiskLevelName === 'S2'">S2</el-tag>
            <el-tag v-else-if="scope.row.RiskLevelName === 'S1'" type="success">S1</el-tag>
            <el-tag v-else type="info">N/A</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="数据类别" prop="Description" align="center" sortable/>
        <el-table-column label="规则名" prop="Name" align="center" sortable/>
        <el-table-column label="匹配方法" prop="CategoryName" align="center" sortable/>
        <el-table-column label="规则状态" prop="Status" width="100px" align="center" sortable/>
        <el-table-column label="最后修改时间" prop="GmtModified" align="center" sortable/>
        <el-table-column label="修改人" prop="DisplayName" align="center" sortable/>
        <el-table-column label="操作" align="center" width="200px">
          <div>
            <el-button type="primary" size="small">编 辑</el-button>
            <el-button type="danger" size="small">删 除</el-button>
          </div>
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
      input: '',
      // 获取数据资产列表
      queryInfo: {
        RuleName: '',
        RiskLevelName: '',
        PageNum: 1,
        PageSize: 10
      },
      // 返回资产数据
      RulesInfo: {
        UUID: '',
        Category: 0,
        CategoryName: '',
        Content: '',
        ContentCategory: '',
        CustomType: 0,
        Description: '',
        DisplayName: '',
        GmtCreate: '',
        GmtModified: '',
        HitTotalCount: 0,
        LoginName: '',
        Name: '',
        RiskLevelId: 0,
        RiskLevelName: '',
        Status: 0
      },
      options: [
        { value: '', label: 'All' },
        { value: 'S4', label: 'S4' },
        { value: 'S3', label: 'S3' },
        { value: 'S2', label: 'S2' },
        { value: 'S1', label: 'S1' }
      ],
      ruleslist: [],
      total: 0
    }
  },
  created () {
    this.getRuleslist()
  },
  methods: {
    async getRuleslist () {
      const { data: res } = await this.$http.post('rules', this.queryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('获取列表失败')
      this.ruleslist = res.Res_Data
      this.total = res.Rules_Total
    },
    // 监听
    handleSizeChange (newSize) {
      this.queryInfo.pageSize = newSize
      this.getRuleslist()
    },
    handleCurrentChange (newPage) {
      this.queryInfo.pageNum = newPage
      this.getRuleslist()
    },
    handleClick () {
      this.getRuleslist()
    },
    clearAll () {
      this.queryInfo.RiskLevelName = 'All'
      this.queryInfo.RuleName = ''
      this.getRuleslist()
    }
  },
  name: 'DataRules'
}
</script>

<style scoped>

</style>
