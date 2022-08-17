<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/welcome' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>数据安全</el-breadcrumb-item>
      <el-breadcrumb-item>存储阶段</el-breadcrumb-item>
      <el-breadcrumb-item>数据资产</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片区 -->
    <el-card>
      <!-- 筛选区域 -->
      <el-tabs v-model="inventoryQueryInfo.GroupName" type="card" @tab-click="handleClick">
        <el-tab-pane label="All" name="All"></el-tab-pane>
        <el-tab-pane label="APP" name="APP"></el-tab-pane>
        <el-tab-pane label="DAP" name="DAP"></el-tab-pane>
        <el-tab-pane label="FMC" name="FMC"></el-tab-pane>
        <el-tab-pane label="ISDP" name="ISDP"></el-tab-pane>
      </el-tabs>
      <!-- 搜索区域 -->
      <el-row :gutter="15">
        <el-col :span="4">
          <el-select v-model="inventoryQueryInfo.InstanceType" clearable placeholder="筛选实例类型" @change="handleClick">
            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-col>
        <el-col :span="8">
          <el-input placeholder="请输入想查询的实例名" v-model="inventoryQueryInfo.InstanceName" clearable @clear="getInventorylist" @keyup.enter.native="getInventorylist" />
        </el-col>
        <el-col :span="4">
          <el-button @click="getInventorylist" type="primary">搜 索</el-button>
          <el-button @click="clearAll">重 置</el-button>
        </el-col>
      </el-row>
      <!-- 数据返回区域 -->
      <el-table :data="inventorylist" :default-sort="{props: 'TotalCount', order: 'descending'}" stripe style="width: 100%">
        <el-table-column type="expand">
          <template slot-scope="props">
            <!-- 折叠区域 -->
            <el-form label-position="left" inline class="table-expand">
              <!-- 创建时间 -->
              <el-form-item>
                <el-popover placement="left" width="200" trigger="hover" content="创建该数据资产实例的时间">
                  <label slot="reference">创建时间</label>
                </el-popover>
                <span>{{ props.row.CreationTime }}</span>
              </el-form-item>
              <!-- 实例描述 -->
              <el-form-item>
                <el-popover placement="left" width="200" trigger="hover" content="数据资产实例的描述信息">
                  <label slot="reference">实例描述</label>
                </el-popover>
                <span>{{ props.row.InstanceDescription }}</span>
              </el-form-item>
              <!-- 最近一次扫描时间 -->
              <el-form-item>
                <el-popover placement="left" width="200" trigger="hover" content="最近一次扫描数据资产实例的完成时间">
                  <label slot="reference">最近一次扫描时间</label>
                </el-popover>
                <span>{{ props.row.LastFinishTime }}</span>
              </el-form-item>
              <!-- 负责人 -->
              <el-form-item>
                <el-popover placement="left" width="200" trigger="hover" content="拥有该数据资产实例的阿里云账号">
                  <label slot="reference">负责人</label>
                </el-popover>
                <span>{{ props.row.Owner }}</span>
              </el-form-item>
              <!-- 类型 -->
              <el-form-item>
                <el-popover placement="left" width="200" trigger="hover" content="数据资产实例所属产品的名称，包括MaxCompute、OSS、RDS等">
                  <label slot="reference">类型</label>
                </el-popover>
                <span>{{ props.row.ProductCode }}</span>
              </el-form-item>
              <!-- 防护状态 -->
              <el-form-item>
                <el-popover placement="left" width="200" trigger="hover" content="数据资产实例的防护状态">
                  <label slot="reference">防护状态</label>
                </el-popover>
                <span>{{ props.row.Protection }}</span>
              </el-form-item>
              <!-- 风险等级 -->
              <el-form-item>
                <el-popover placement="left" width="200" trigger="hover" content="数据资产实例的风险等级名称">
                  <label slot="reference">风险等级</label>
                </el-popover>
                <el-tag v-if="props.row.RiskLevelName === 'S4'" type="danger">S4</el-tag>
                <el-tag v-else-if="props.row.RiskLevelName === 'S3'" type="warning">S3</el-tag>
                <el-tag v-else-if="props.row.RiskLevelName === 'S2'">S2</el-tag>
                <el-tag v-else-if="props.row.RiskLevelName === 'S1'" type="success">S1</el-tag>
                <el-tag v-else type="info">N/A</el-tag>
              </el-form-item>
              <!-- 是否包含敏感数据 -->
              <el-form-item>
                <el-popover placement="left" width="200" trigger="hover" content="数据资产实例中是否包含敏感数据">
                  <label slot="reference">是否包含敏感数据</label>
                </el-popover>
                <span>{{ props.row.Sensitive }}</span>
              </el-form-item>
              <!-- 敏感数据总数 -->
              <el-form-item>
                <el-popover placement="left" width="200" trigger="hover" content="数据资产实例中包含的敏感数据总数。例如：当数据资产为RDS时，表示该实例中数据库的敏感总表数">
                  <label slot="reference">敏感数据表总数</label>
                </el-popover>
                <span>{{ props.row.SensitiveCount }}</span>
              </el-form-item>
              <!-- 数据总数 -->
              <el-form-item>
                <el-popover placement="left" width="200" trigger="hover" content="数据资产实例中的数据总数。例如：当数据资产为RDS时，表示该实例中数据库的总表数">
                  <label slot="reference">数据表总数</label>
                </el-popover>
                <span>{{ props.row.TotalCount }}</span>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column label="项目组" prop="GroupName" width="100px" sortable/>
        <el-table-column label="名称" prop="Name" sortable/>
        <el-table-column label="类型" prop="ProductCode" sortable/>
        <el-table-column label="数据表总数" prop="TotalCount" sortable/>
        <el-table-column label="风险等级" prop="RiskLevelName" sortable>
          <template slot-scope="scope">
            <el-tag v-if="scope.row.RiskLevelName === 'S4'" type="danger">S4</el-tag>
            <el-tag v-else-if="scope.row.RiskLevelName === 'S3'" type="warning">S3</el-tag>
            <el-tag v-else-if="scope.row.RiskLevelName === 'S2'">S2</el-tag>
            <el-tag v-else-if="scope.row.RiskLevelName === 'S1'" type="success">S1</el-tag>
            <el-tag v-else type="info">N/A</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="数据库详情" align="center">
          <template slot-scope="scope">
            <el-link v-if="scope.row.ProductCode==='RDS'" icon="el-icon-view" type="primary" @click="handelDrawer(scope.row.Name)">查看</el-link>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页区域 -->
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="inventoryQueryInfo.PageNum"
        :page-sizes="[5, 10, 20, 50]"
        :page-size="inventoryQueryInfo.PageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total=total>
      </el-pagination>
    </el-card>
    <!-- 抽屉区域 -->
    <el-drawer :visible.sync="table" :before-close="handleDrawerClose" size="60%" >
      <!-- 表格 -->
      <el-table :data="tablelist" :default-sort="{props: 'TotalCount', order: 'descending'}" stripe fit>
        <el-table-column label="表 名" prop="Name" sortable></el-table-column>
        <el-table-column label="字段数" prop="TotalCount" sortable></el-table-column>
        <el-table-column label="风险等级" prop="RiskLevelName" sortable>
          <template slot-scope="scope">
            <el-tag v-if="scope.row.RiskLevelName === 'S4'" type="danger">S4</el-tag>
            <el-tag v-else-if="scope.row.RiskLevelName === 'S3'" type="warning">S3</el-tag>
            <el-tag v-else-if="scope.row.RiskLevelName === 'S2'">S2</el-tag>
            <el-tag v-else-if="scope.row.RiskLevelName === 'S1'" type="success">S1</el-tag>
            <el-tag v-else type="info">N/A</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="列详情">
          <template slot-scope="scope">
            <el-link icon="el-icon-view" type="primary"  @click="handelInnerDrawer(scope.row.Name)">查看</el-link>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页区域 -->
      <el-pagination
        hide-on-single-page = true
        @size-change="handleTableSizeChange"
        @current-change="handleTableCurrentChange"
        :current-page="tableQueryInfo.PageNum"
        :page-sizes="[5, 10, 20, 50]"
        :page-size="tableQueryInfo.PageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total=tabletotal>
      </el-pagination>
      <el-drawer :visible.sync="column" :before-close="handleInnerDrawerClose"  :append-to-body="true" size="50%">
        <!-- 表格 -->
        <el-table :data="columnlist" stripe fit>
          <el-table-column label="字段名" prop="Name" sortable></el-table-column>
          <el-table-column label="数据类型" prop="DataType" sortable></el-table-column>
          <el-table-column label="风险等级" prop="RiskLevelName" sortable>
            <template slot-scope="scope">
              <el-tag v-if="scope.row.RiskLevelName === 'S4'" type="danger">S4</el-tag>
              <el-tag v-else-if="scope.row.RiskLevelName === 'S3'" type="warning">S3</el-tag>
              <el-tag v-else-if="scope.row.RiskLevelName === 'S2'">S2</el-tag>
              <el-tag v-else-if="scope.row.RiskLevelName === 'S1'" type="success">S1</el-tag>
              <el-tag v-else type="info">N/A</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="字段样例">
            <template>
              <el-link icon="el-icon-view" type="primary">查看</el-link>
            </template>
          </el-table-column>
        </el-table>
        <!-- 分页区域 -->
        <el-pagination
          hide-on-single-page = true
          @size-change="handleColumnSizeChange"
          @current-change="handleColumnCurrentChange"
          :current-page="columnQueryInfo.PageNum"
          :page-sizes="[5, 10, 20, 50]"
          :page-size="columnQueryInfo.PageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total=columntotal>
        </el-pagination>
      </el-drawer>
    </el-drawer>
  </div>
</template>

<script>
export default {
  data () {
    return {
      table: false,
      column: false,
      input: '',
      // 获取数据资产列表
      inventoryQueryInfo: {
        GroupName: 'All',
        InstanceName: '',
        InstanceType: '',
        InstanceGroup: '',
        PageNum: 1,
        PageSize: 10
      },
      // 获取数据资产表列表
      tableQueryInfo: {
        InstanceName: '',
        GroupName: '',
        PageNum: 1,
        PageSize: 10
      },
      // 获取资产 字段 列表
      columnQueryInfo: {
        TableName: '',
        PageNum: 1,
        PageSize: 10
      },
      // 返回资产数据
      inventoryInfo: {
        UUID: '',
        CreationTime: '',
        DepartName: '',
        InstanceId: 0,
        InstanceDescription: '',
        Labelsec: false,
        LastFinishTime: '',
        Name: '',
        OdpsRiskLevelName: '',
        Owner: '',
        ProductCode: '',
        ProductId: 0,
        Protection: false,
        RiskLevelId: 0,
        RiskLevelName: '',
        RuleName: '',
        Sensitive: true,
        SensitiveCount: 0,
        TenantName: '',
        TotalCount: 0
      },
      // 抽屉 表 数据
      tableInfo: {
        UUID: '',
        GroupName: '',
        CreationTime: '',
        Id: 0,
        InstanceDescription: '',
        InstanceId: 0,
        InstanceName: '',
        Name: '',
        Owner: '',
        ProductCode: '',
        ProductId: '',
        RiskLevelId: 0,
        RiskLevelName: '',
        Sensitive: true,
        SensitiveCount: 0,
        SensitiveRatio: '',
        TenantName: '',
        TotalCount: 0
      },
      // 抽屉 列 数据
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
        { value: 'OSS', label: 'OSS' },
        { value: 'RDS', label: 'RDS' },
        { value: 'MaxCompute', label: 'MaxCompute' }
      ],
      inventorylist: [],
      tablelist: [],
      columnlist: [],
      total: 0,
      tabletotal: 0,
      columntotal: 0
    }
  },
  created () {
    this.getInventorylist()
  },
  methods: {
    async getInventorylist () {
      const { data: res } = await this.$http.post('inventory', this.inventoryQueryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('获取资产列表失败')
      this.inventorylist = res.Res_Data
      this.total = res.Inventory_Total
    },
    async getTablelist () {
      const { data: res } = await this.$http.post('tables', this.tableQueryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('获取列表失败')
      this.tablelist = res.Res_Data
      this.tabletotal = res.Tables_Total
    },
    async getColumnlist () {
      const { data: res } = await this.$http.post('tables/column', this.columnQueryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('获取列表失败')
      this.columnlist = res.Res_Data
      this.columntotal = res.Column_Total
    },
    // 监听
    handleSizeChange (newSize) {
      this.inventoryQueryInfo.pageSize = newSize
      this.getInventorylist()
    },
    handleCurrentChange (newPage) {
      this.inventoryQueryInfo.pageNum = newPage
      this.getInventorylist()
    },
    handleTableSizeChange (newSize) {
      this.tableQueryInfo.pageSize = newSize
      this.getTablelist()
    },
    handleTableCurrentChange (newPage) {
      this.tableQueryInfo.pageNum = newPage
      this.getTablelist()
    },
    handleColumnSizeChange (newSize) {
      this.columnQueryInfo.pageSize = newSize
      this.getColumnlist()
    },
    handleColumnCurrentChange (newPage) {
      this.columnQueryInfo.pageNum = newPage
      this.getColumnlist()
    },
    handleClick () {
      this.getInventorylist()
    },
    handelDrawer (instanceName) {
      this.table = true
      this.tableQueryInfo.InstanceName = instanceName
      this.getTablelist()
    },
    handelInnerDrawer (tableName) {
      this.column = true
      const temp = tableName.indexOf('.')
      this.columnQueryInfo.TableName = tableName.substr(temp + 1, tableName.length - temp - 1)
      this.getColumnlist()
    },
    clearAll () {
      this.inventoryQueryInfo.GroupName = 'All'
      this.inventoryQueryInfo.InstanceName = ''
      this.inventoryQueryInfo.InstanceType = 'All'
      this.getInventorylist()
    },
    handleDrawerClose () {
      this.tableQueryInfo.InstanceName = ''
      this.tableQueryInfo.PageNum = 1
      this.tableQueryInfo.PageSize = 10
      this.table = false
    },
    handleInnerDrawerClose () {
      this.columnQueryInfo.TableName = ''
      this.columnQueryInfo.PageNum = 1
      this.columnQueryInfo.PageSize = 10
      this.column = false
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
