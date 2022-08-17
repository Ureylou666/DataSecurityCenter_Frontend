<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/welcome' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>系统管理</el-breadcrumb-item>
      <el-breadcrumb-item>项目配置</el-breadcrumb-item>
    </el-breadcrumb>
    <el-divider></el-divider>
    <!-- 卡片列表 -->
    <el-row :gutter="20">
      <el-col :span="6" v-for="GroupInfo in grouplist" :key="GroupInfo.UUID" :offset="1">
        <el-card>
          <div slot="header" class="clearfix">
            <span>{{ GroupInfo.GroupName }}</span>
            <el-popconfirm
              confirm-button-text='确认'
              confirm-button-type="danger"
              cancel-button-text='不用了'
              icon="el-icon-info"
              title="确定重置所有数据吗？"
            >
              <el-button style="float: right; padding: 3px 0" type="text" icon="el-icon-refresh" slot="reference">重置数据</el-button>
            </el-popconfirm>
          </div>
          <div class="text item">项目组负责人：{{ GroupInfo.GroupOwner }}</div>
          <div class="text item">负责人邮箱：{{ GroupInfo.OwnerMail }}</div>
          <div class="text item">SecurityPO：{{ GroupInfo.SecPo }}</div>
          <div class="text item">SecPO邮箱：{{ GroupInfo.SeoPoMail }}</div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  data () {
    return {
      queryInfo: {
        PageNum: 1,
        PageSize: 10
      },
      GroupInfo: {
        UUID: '',
        GroupName: '',
        GroupOwner: '',
        OwnerMail: '',
        SecPo: '',
        SeoPoMail: ''
      },
      initGroupData: {
        GroupName: ''
      },
      grouplist: [],
      total: 0
    }
  },
  created () {
    this.getGrouplist()
  },
  methods: {
    async getGrouplist () {
      const { data: res } = await this.$http.post('group', this.queryInfo,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('获取列表失败')
      this.grouplist = res.Res_Data
      this.total = res.Group_Total
    },
    async initData () {
      const { data: res } = await this.$http.post('group/initData', this.initGroupData,
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Code !== 200) return this.$message.error('更新失败')
      return this.$message.success('已更新成功')
    }
  },
  name: 'GroupSetting'
}
</script>

<style scoped>
.el-card {
  margin-bottom: 20px;
  border-radius: 30px;
}
.text {
  font-size: 14px;
}
.item {
  margin-bottom: 18px;
}
</style>
