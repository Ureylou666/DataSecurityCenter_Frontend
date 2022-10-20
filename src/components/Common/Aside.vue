<template>
  <el-aside width="200px">
    <el-menu text-color="#FFF" background-color="#484848" active-text-color="#ffd04b" :unique-opened="true" router>
      <el-menu-item index="welcome"><i class="el-icon-s-home"/>Home Page</el-menu-item>
      <el-menu-item index="PrivcyCenter"><i class="el-icon-s-claim"/>Privicy Center</el-menu-item>
      <el-submenu index="2">
        <template slot="title"><i class="el-icon-s-data"/>Data Security</template>
        <el-menu-item-group>
          <el-submenu index="2-2">
            <template slot="title">Identification</template>
            <el-menu-item index="Inventory">Data Inventory</el-menu-item>
            <el-menu-item index="DetailList">Data Field</el-menu-item>
            <el-menu-item index="Rules">Classification</el-menu-item>
          </el-submenu>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="3">
        <template slot="title"><i class="el-icon-s-claim"></i>Compliance Report</template>
        <el-menu-item-group>
          <el-menu-item>DSL</el-menu-item>
          <el-menu-item >MLPS</el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="4">
        <template slot="title"><i class="el-icon-setting"></i>Configuration</template>
        <el-menu-item-group>
          <el-menu-item index="group">Dev Group</el-menu-item>
          <el-menu-item>Users</el-menu-item>
          <el-menu-item>Role</el-menu-item>
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
