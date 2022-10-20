<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/welcome' }">Home Page</el-breadcrumb-item>
      <el-breadcrumb-item>Data Security</el-breadcrumb-item>
      <el-breadcrumb-item>Identification</el-breadcrumb-item>
      <el-breadcrumb-item>Inventory</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片区 -->
    <el-card>
      <!-- 筛选区域 -->
      <el-tabs type="card" v-model="cloudAccountQueryInfo.GroupName" @tab-click="handleClick">
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
      <!-- 数据返回区域 -->
      <div style="margin-left:1%; margin-right:1%">
        <el-row>
          <el-col :span="7" v-for="(item) in cloudAccountList" :key="item.UUID" :offset="1">
            <el-card style="margin: 5px" shadow="hover">
              <div slot="header" style="alignment: top">
                <el-tag effect="plain" style="font-size: 15px">{{ item.DisplayName }}</el-tag>
                <el-button style="float: right;" type="plain" icon="el-icon-search" size="medium" round @click="handleRDSSearch(item.AccountId)"></el-button>
              </div>
              <div>AccountID : {{ item.AccountId }}</div>
              <div>Group : {{ item.GroupName }}</div>
              <div v-if ="item.JoinTime">Create Time : {{ item.JoinTime.slice(0,10) }}</div>
              <div>RDS Count : <i style="font-weight: bolder; font-size: large; color: #E6A23C"> {{ item.RDSCount }} </i></div>
            </el-card>
          </el-col>
        </el-row>
      </div>
      <el-divider></el-divider>
      <!-- CloudAccount 分页区域 -->
      <el-pagination
        @size-change="handleCloudAccountSizeChange"
        @current-change="handleCloudAccountCurrentChange"
        :current-page="cloudAccountQueryInfo.PageNum"
        :page-sizes="[3, 9, 21, 48]"
        :page-size="cloudAccountQueryInfo.PageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total=AccountTotal>
      </el-pagination>
      <!-- RDS Dialog区域 -->
      <el-dialog title="RDS Inventory List" width="80%" :visible.sync="RDSTableVisible" style="font-weight: bolder; font-size: larger">
        <el-table :data="RDSInventoryList" border stripe fit>
          <el-table-column property="RDSInstanceID" label="InstanceID"></el-table-column>
          <el-table-column property="RDSInstanceDescription" label="Description"></el-table-column>
          <el-table-column label="RDSEngine">
            <template slot-scope="scope">
              <el-tag type="plain">{{ scope.row.RDSEngine }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column property="RegionId" label="Region"></el-table-column>
          <el-table-column property="CreationTime" label="CreationTime"></el-table-column>
          <el-table-column width="50px">
            <template slot-scope="scope">
              <el-button style="float: inside" type="plain" icon="el-icon-view" circle @click="handleDatabaseSearch(scope.row.RDSInstanceID)"></el-button>
            </template>
          </el-table-column>
        </el-table>
        <!-- RDSInventory 分页区域 -->
        <el-pagination
          @size-change="handleRDSSizeChange"
          @current-change="handleRDSCurrentChange"
          :current-page="RDSQueryInfo.PageNum"
          :page-sizes="[1, 5, 10, 20]"
          :page-size="RDSQueryInfo.PageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total=RDSListTotal>
        </el-pagination>
      </el-dialog>
      <!-- Database Dialog区域 -->
      <el-dialog title="Database List" width="60%" :visible.sync="DatabaseVisible" >
        <el-table :data="DatabaseList" border stripe fit style="font-weight: bolder; font-size: larger">
          <el-table-column property="DatabaseName" label="Database"></el-table-column>
          <el-table-column property="DatabaseDescription" label="Description"></el-table-column>
          <el-table-column label="RDSEngine">
            <template slot-scope="scope">
              <el-tag type="plain">{{ scope.row.DatabaseEngine }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="Sensitive">
            <template slot-scope="scope">{{ scope.row.Sensitive }}</template>
          </el-table-column>
          <el-table-column width="50px">
            <template slot-scope="scope">
              <el-button style="float: right;" type="plain" icon="el-icon-view" circle @click="handleTableSearch(scope.row.RDSInstanceID, scope.row.DatabaseName)"></el-button>
            </template>
          </el-table-column>
        </el-table>
        <!-- DatabaseInventory 分页区域 -->
        <el-pagination
          @size-change="handleDatabaseSizeChange"
          @current-change="handleDatabaseCurrentChange"
          :current-page="DatabaseQueryInfo.PageNum"
          :page-sizes="[1, 5, 10, 20]"
          :page-size="DatabaseQueryInfo.PageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total=DatabaseListTotal>
        </el-pagination>
      </el-dialog>
      <!-- Table Drawer区域 -->
      <el-drawer title="Table List" :visible.sync="TableVisible" size="40%" direction="ltr" style="font-weight: bolder; font-size: larger">
        <el-table :data="TableList" border stripe fit style="margin: 5px">
          <el-table-column type="index"></el-table-column>
          <el-table-column property="TableName" label="Table Name"></el-table-column>
          <el-table-column property="Sensitive" label="Sensitive"></el-table-column>
          <el-table-column property="RiskLevelId" label="Risk Level"></el-table-column>
          <el-table-column width="60px">
            <template slot-scope="scope">
              <el-button style="float: right;" type="plain" icon="el-icon-view" circle @click="handleColumnSearch(scope.row.InstanceId, scope.row.DatabaseName, scope.row.TableName)"></el-button>
            </template>
          </el-table-column>
        </el-table>
        <!-- Table List 分页区域 -->
        <el-pagination
          @size-change="handleTableSizeChange"
          @current-change="handleTableCurrentChange"
          :current-page="TableQueryInfo.PageNum"
          :page-sizes="[5, 10, 20, 50]"
          :page-size="TableQueryInfo.PageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total=TableListTotal>
        </el-pagination>
      </el-drawer>
      <!-- Column Drawer区域 -->
      <el-drawer title="Column List" :visible.sync="ColumnVisible" size="70%" direction="rtl" style="font-weight: bolder; font-size: larger">
        <el-table :data="ColumnList" border stripe fit style="margin: 5px">
          <el-table-column type="index"></el-table-column>
          <el-table-column property="ColumnName" label="Column Name"></el-table-column>
          <el-table-column property="DataType" label="Data Type"></el-table-column>
          <el-table-column property="RuleName" label="Rule Name"></el-table-column>
          <el-table-column property="CategoryName" label="Category Name"></el-table-column>
          <el-table-column property="SensLevelName" label="Sens Level Name"></el-table-column>
          <el-table-column label="Sample" align="center">
            <template slot-scope="scope">
              <el-popover placement="top-start" title="Sample Data" trigger="hover">
                <p>{{ scope.row.SampleData }}</p>
                <i slot="reference" class="el-icon-search"></i>
              </el-popover>
            </template>
          </el-table-column>
        </el-table>
        <!-- Column List 分页区域 -->
        <el-pagination
          @size-change="handleColumnSizeChange"
          @current-change="handleColumnCurrentChange"
          :current-page="ColumnQueryInfo.PageNum"
          :page-sizes="[5, 10, 20, 50]"
          :page-size="ColumnQueryInfo.PageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total=ColumnListTotal>
        </el-pagination>
      </el-drawer>
    </el-card>
  </div>
</template>

<script>
export default {
  data () {
    return {
      input: '',
      // 获取云账号列表
      cloudAccountQueryInfo: {
        GroupName: 'All',
        PageNum: 1,
        PageSize: 9
      },
      // 获取RDS列表
      RDSQueryInfo: {
        AccountID: '',
        PageNum: 1,
        PageSize: 5
      },
      // 获取Database列表
      DatabaseQueryInfo: {
        InstanceID: '',
        PageNum: 1,
        PageSize: 5
      },
      // 获取数据资产表列表
      TableQueryInfo: {
        InstanceName: '',
        DatabaseName: '',
        PageNum: 1,
        PageSize: 10
      },
      // 获取资产 字段 列表
      ColumnQueryInfo: {
        InstanceName: '',
        DatabaseName: '',
        TableName: '',
        PageNum: 1,
        PageSize: 5
      },
      cloudAccountList: [],
      cloudAccountID: '',
      AccountTotal: 0,
      RDSTableVisible: false,
      RDSInventoryList: [],
      RDSListTotal: 0,
      DatabaseVisible: false,
      DatabaseList: [],
      DatabaseListTotal: 0,
      TableVisible: false,
      TableList: [],
      TableListTotal: 0,
      ColumnVisible: false,
      ColumnList: [],
      ColumnListTotal: 0
    }
  },
  created () {
    this.getCloudAccountlist()
  },
  methods: {
    initCloudQuery () {
      this.cloudAccountQueryInfo = {
        GroupName: 'All',
        PageNum: 1,
        PageSize: 9
      }
    },
    initRDSQuery (cloudAccountID) {
      this.RDSQueryInfo = {
        AccountID: cloudAccountID,
        PageSize: 5,
        PageNum: 1
      }
    },
    initDatabaseQuery (instanceID) {
      this.DatabaseQueryInfo = {
        InstanceID: instanceID,
        PageSize: 5,
        PageNum: 1
      }
    },
    initTableQuery (InstanceID, DatabaseName) {
      this.TableQueryInfo = {
        InstanceID: InstanceID,
        DatabaseName: DatabaseName,
        PageSize: 10,
        PageNum: 1
      }
    },
    initColumnQuery (InstanceID, DatabaseName, TableName) {
      this.ColumnQueryInfo = {
        InstanceName: InstanceID,
        DatabaseName: DatabaseName,
        TableName: TableName,
        PageSize: 5,
        PageNum: 1
      }
    },
    async getCloudAccountlist () {
      const { data: res } = await this.$http.post('Inventory/CloudAccount', this.cloudAccountQueryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('Error in query cloud account !')
      this.cloudAccountList = res.Res_Data
      this.AccountTotal = res.AccountListTotal
    },
    async getRDSInventorylist () {
      const { data: res } = await this.$http.post('Inventory/RDSList', this.RDSQueryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('Error in query RDS inventory !')
      this.RDSInventoryList = res.Res_Data
      this.RDSListTotal = res.rdsInventoryTotal
    },
    async getDatabaseInventorylist () {
      const { data: res } = await this.$http.post('Inventory/Database', this.DatabaseQueryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('Error in query Database List !')
      this.DatabaseList = res.Res_Data
      this.DatabaseListTotal = res.DatabaseTotal
    },
    async getTableList () {
      const { data: res } = await this.$http.post('Inventory/Table', this.TableQueryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('Error in query Database List !')
      this.TableList = res.Res_Data
      this.TableListTotal = res.TableListTotal
    },
    async getColumnList () {
      const { data: res } = await this.$http.post('Inventory/Column', this.ColumnQueryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('Error in query Database List !')
      this.ColumnList = res.Res_Data
      this.ColumnListTotal = res.Column_Total
    },
    // 监听
    handleCloudAccountSizeChange (newSize) {
      this.cloudAccountQueryInfo.pageSize = newSize
      this.getCloudAccountlist()
    },
    handleCloudAccountCurrentChange (newPage) {
      this.cloudAccountQueryInfo.pageNum = newPage
      this.getCloudAccountlist()
    },
    handleRDSSizeChange (newSize) {
      this.RDSQueryInfo.pageSize = newSize
      this.getRDSInventorylist()
    },
    handleRDSCurrentChange (newPage) {
      this.RDSQueryInfo.pageNum = newPage
      this.getRDSInventorylist()
    },
    handleDatabaseSizeChange (newSize) {
      this.DatabaseQueryInfo.pageSize = newSize
      this.getDatabaseInventorylist()
    },
    handleDatabaseCurrentChange (newPage) {
      this.DatabaseQueryInfo.pageNum = newPage
      this.getDatabaseInventorylist()
    },
    handleTableSizeChange (newSize) {
      this.TableQueryInfo.pageSize = newSize
      this.getTableList()
    },
    handleTableCurrentChange (newPage) {
      this.TableQueryInfo.pageNum = newPage
      this.getTableList()
    },
    handleColumnSizeChange (newSize) {
      this.ColumnQueryInfo.pageSize = newSize
      this.getColumnList()
    },
    handleColumnCurrentChange (newPage) {
      this.ColumnQueryInfo.pageNum = newPage
      this.getColumnList()
    },
    handleClick () {
      this.getCloudAccountlist()
    },
    handleRDSSearch (cloudAccountID) {
      this.RDSTableVisible = true
      this.initRDSQuery(cloudAccountID)
      this.getRDSInventorylist()
    },
    handleDatabaseSearch (InstanceID) {
      this.DatabaseVisible = true
      this.initDatabaseQuery(InstanceID)
      this.getDatabaseInventorylist()
    },
    handleTableSearch (InstanceID, DatabaseName) {
      this.TableVisible = true
      this.initTableQuery(InstanceID, DatabaseName)
      this.getTableList()
    },
    handleColumnSearch (InstanceID, DatabaseName, TableName) {
      this.ColumnVisible = true
      this.initColumnQuery(InstanceID, DatabaseName, TableName)
      this.getColumnList()
    }
  },
  name: 'DataInventory'
}
</script>

<style scoped>
.table-expand {
  font-size: 0;
}
.table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}
label {
  font-weight: bolder;
  color: #414141;
  padding: 0 12px 0 0;
}
</style>
