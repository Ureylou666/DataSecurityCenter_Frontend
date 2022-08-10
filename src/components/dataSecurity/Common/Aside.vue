<template>
  <el-aside width="200px">
    <el-menu text-color="#FFF" background-color="#484848" active-text-color="#ffd04b" :unique-opened="true"
             :collapse="isCollapse" :collapse-transition="false" router :default-active="activePath">
      <el-menu-item index="welcome"><i class="el-icon-s-home"/>首 页</el-menu-item>
      <el-submenu index="1">
        <template slot="title"><i class="el-icon-s-data"/>数据安全</template>
        <el-menu-item-group>
          <el-submenu index="1-1">
            <template slot="title">收集阶段</template>
            <el-menu-item>隐私政策</el-menu-item>
          </el-submenu>
          <el-submenu index="1-2">
            <template slot="title">存储阶段</template>
            <el-menu-item index="Inventory">数据资产</el-menu-item>
            <el-menu-item index="Detaillist">数据清单</el-menu-item>
            <el-menu-item index="Rules">数据分级</el-menu-item>
          </el-submenu>
          <el-submenu index="1-3">
            <template slot="title">处理阶段</template>
            <el-menu-item>安全监控</el-menu-item>
            <el-menu-item>数据分享</el-menu-item>
          </el-submenu>
          <el-submenu index="1-4">
            <template slot="title">销毁阶段</template>
            <el-menu-item>销毁请求</el-menu-item>
          </el-submenu>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="2">
        <template slot="title"><i class="el-icon-s-claim"></i>合规报表</template>
        <el-menu-item-group>
          <el-menu-item>数据安全法</el-menu-item>
          <el-menu-item >等级保护</el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="3">
        <template slot="title"><i class="el-icon-setting"></i>系统管理</template>
        <el-menu-item-group>
          <el-menu-item index="group">项目配置</el-menu-item>
          <el-menu-item>用户配置</el-menu-item>
          <el-menu-item>角色配置</el-menu-item>
        </el-menu-item-group>
      </el-submenu>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  name: 'CommonAside'
}
</script>

<style scoped>

</style>
