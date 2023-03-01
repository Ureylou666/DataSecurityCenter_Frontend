<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/admin' }">Home Page</el-breadcrumb-item>
      <el-breadcrumb-item>Data Security</el-breadcrumb-item>
      <el-breadcrumb-item>Configuration</el-breadcrumb-item>
      <el-breadcrumb-item>Rules</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片区 -->
    <el-card>
      <el-row :gutter="15">
        <el-col :span="2">
          <el-select v-model="queryInfo.SensitiveGrade" clearable placeholder="Filter">
            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-col>
        <el-col :span="2">
          <el-select v-model="queryInfo.CategoryName" clearable placeholder="Filter">
            <el-option v-for="item in Category" :key="item.Value" :label="item.Label" :value="item.Value"/>
          </el-select>
        </el-col>
        <el-col :span="8">
          <el-input placeholder="Input RuleName" v-model="queryInfo.RuleName" clearable @clear="getRuleslist"/>
        </el-col>
        <el-col :span="8">
          <el-button @click="getRuleslist" type="primary" icon="el-icon-search">Search</el-button>
          <el-button @click="handleNewRule" type="danger" icon="el-icon-plus">New</el-button>
          <el-button @click="clearAll" type="info" icon="el-icon-refresh-left">Clear</el-button>
        </el-col>
      </el-row>
      <!-- 数据返回区域 -->
      <el-table :data="RuleList" stripe index style="width: 100%">
        <el-table-column type="index" width="50"/>
        <el-table-column label="RuleName" prop="RuleName" align="center" sortable/>
        <el-table-column label="Category" prop="CategoryName" align="center" sortable/>
        <el-table-column label="Severity" align="center" sortable>
          <template slot-scope="scope">
            <el-tag v-if="scope.row.SensitiveGrade === 'S4'" type="danger">S4</el-tag>
            <el-tag v-else-if="scope.row.SensitiveGrade === 'S3'" type="warning">S3</el-tag>
            <el-tag v-else-if="scope.row.SensitiveGrade === 'S2'">S2</el-tag>
            <el-tag v-else-if="scope.row.SensitiveGrade === 'S1'" type="success">S1</el-tag>
            <el-tag v-else type="info">N/A</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="Status" prop="Status" align="center" sortable>
          <template slot-scope="scope">
            <el-switch v-model="scope.row.Status" disabled></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="Action" align="center" width="200px">
          <template slot-scope="scope">
            <div>
              <el-button type="primary" size="small" icon="el-icon-search" circle @click=getRuleDetails(scope.row.UUID) />
              <el-button type="warning" size="small" icon="el-icon-s-promotion" circle @click=enforceRule(scope.row.UUID) />
              <el-button type="danger" size="small" icon="el-icon-delete" circle @click="deleteRule(scope.row.UUID)"/>
            </div>
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
      <!-- 详情区域 -->
      <el-dialog title="Details" :visible.sync="RuleDetailsVisible" width="800px">
        <el-form :model="RulesInfo" label-width="150px">
          <el-form-item label="Rule Name">
            <el-input v-model="RulesInfo.RuleName" style="width: 80%" disabled></el-input>
          </el-form-item>
          <el-form-item label="Reg Express">
            <el-input v-model="RulesInfo.RegExpress" style="width: 80%" disabled></el-input>
          </el-form-item>
          <el-form-item label="Preferred Column">
            <el-input v-model="RulesInfo.PreferColumn" style="width: 80%" disabled></el-input>
          </el-form-item>
          <el-form-item label="Category Name">
            <el-input v-model="RulesInfo.CategoryName" style="width: 50%" disabled></el-input>
          </el-form-item>
          <el-form-item label="Sensitive Grade">
            <el-tag v-if="RulesInfo.SensitiveGrade === 'S4'" type="danger">S4</el-tag>
            <el-tag v-else-if="RulesInfo.SensitiveGrade === 'S3'" type="warning">S3</el-tag>
            <el-tag v-else-if="RulesInfo.SensitiveGrade === 'S2'">S2</el-tag>
            <el-tag v-else-if="RulesInfo.SensitiveGrade === 'S1'" type="success">S1</el-tag>
            <el-tag v-else type="info">N/A</el-tag>
          </el-form-item>
          <el-form-item label="Comments">
            <el-input v-model="RulesInfo.Comments" style="width: 50%" disabled></el-input>
          </el-form-item>
        </el-form>
      </el-dialog>
      <!-- 新增区域 -->
      <el-dialog title="New Rule" :visible.sync="NewRuleVisible" width="800px">
        <el-form :model="RulesInfo" label-width="150px" label-position="left">
          <el-form-item label="Rule Name">
            <el-input v-model="RulesInfo.RuleName" style="width: 80%"></el-input>
          </el-form-item>
          <el-form-item label="Reg Express">
            <el-input v-model="RulesInfo.RegExpress" style="width: 80%"></el-input>
          </el-form-item>
          <el-form-item label="Preferred Column">
            <el-input v-model="RulesInfo.PreferColumn" style="width: 80%"></el-input>
          </el-form-item>
          <el-form-item label="Category Name">
            <el-select v-model="RulesInfo.CategoryUUID" style="width: 50%">
              <el-option v-for="item in Category" :key="item.Value" :label="item.Label" :value="item.UUID"/>
            </el-select>
          </el-form-item>
          <el-form-item label="Sensitive Grade">
            <el-select v-model="RulesInfo.SensitiveGrade" placeholder="Select Sensitive Grade" style="width: 50%">
              <el-option label="S4" value="S4"></el-option>
              <el-option label="S3" value="S3"></el-option>
              <el-option label="S2" value="S2"></el-option>
              <el-option label="S1" value="S1"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="Comments">
            <el-input v-model="RulesInfo.Comments" style="width: 50%"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="NewRule">Confirm</el-button>
            <el-button @click="clearAll">Cancel</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
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
        CategoryName: '',
        SensitiveGrade: '',
        PageNum: 1,
        PageSize: 10
      },
      OneRuleInfo: {
        UUID: ''
      },
      // 返回资产数据
      RulesInfo: {
        uuid: '',
        RuleName: '',
        PreferColumn: '',
        RegExpress: '',
        CategoryName: '',
        CategoryUUID: '',
        SensitiveGrade: '',
        Comments: '',
        Status: false
      },
      options: [
        { value: '', label: 'All' },
        { value: 'S4', label: 'S4' },
        { value: 'S3', label: 'S3' },
        { value: 'S2', label: 'S2' },
        { value: 'S1', label: 'S1' }
      ],
      RuleList: [],
      Category: [],
      RuleDetailsVisible: false,
      NewRuleVisible: false,
      total: 0
    }
  },
  created () {
    this.getRuleslist()
    this.getCategoryName()
  },
  methods: {
    async getCategoryName () {
      const { data: res } = await this.$http.get('Category/All',
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('获取列表失败')
      this.Category = res.Data
      console.log(this.Category)
    },
    async getRuleslist () {
      const { data: res } = await this.$http.post('Rules/List', this.queryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('获取列表失败')
      this.RuleList = res.Data
      this.total = res.rulesTotal
    },
    async getRuleDetails (RuleUUID) {
      const { data: res } = await this.$http.get('Rules/' + RuleUUID,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('获取规则详情失败')
      this.RulesInfo = res.Data
      this.RuleDetailsVisible = true
    },
    async NewRule () {
      const { data: res } = await this.$http.post('Rules/Create', this.RulesInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('新增规则失败')
      this.NewRuleVisible = false
      this.clearAll()
    },
    async deleteRule (RuleUUID) {
      this.OneRuleInfo.UUID = RuleUUID
      const { data: res } = await this.$http.post('Rules/Delete', this.OneRuleInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('删除规则失败')
      this.clearAll()
    },
    async enforceRule (RuleUUID) {
      this.OneRuleInfo.UUID = RuleUUID
      const { data: res } = await this.$http.post('Rules/Enforce', this.OneRuleInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('删除规则失败')
      this.clearAll()
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
    handleNewRule () {
      this.clearRuleInfo()
      this.NewRuleVisible = true
    },
    clearAll () {
      this.queryInfo.SensitiveGrade = ''
      this.queryInfo.CategoryName = ''
      this.queryInfo.RuleName = ''
      this.NewRuleVisible = false
      this.RuleDetailsVisible = false
      this.clearRuleInfo()
      this.getRuleslist()
    },
    clearRuleInfo () {
      this.RulesInfo.RuleName = ''
      this.RulesInfo.CategoryName = ''
      this.RulesInfo.Comments = ''
      this.RulesInfo.SensitiveGrade = ''
      this.RulesInfo.uuid = ''
      this.RulesInfo.RegExpress = ''
      this.RulesInfo.PreferColumn = ''
      this.RulesInfo.Status = false
      this.getRuleslist()
    }
  },
  name: 'ConfigRules'
}
</script>

<style scoped>

</style>
