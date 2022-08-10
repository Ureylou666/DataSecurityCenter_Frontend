<template>
  <el-container>
    <el-header>
      <common-header></common-header>
    </el-header>
    <el-container>
      <el-aside width="auto">
        <common-aside></common-aside>
      </el-aside>
      <el-main width="auto">
        <!-- 路由占位符 -->
        <router-view></router-view>
      </el-main>
    </el-container>
    <el-footer>
      <common-Footer></common-Footer>
    </el-footer>
  </el-container>
</template>

<script>
import CommonHeader from '@/components/dataSecurity/Common/Header'
import CommonAside from '@/components/dataSecurity/Common/Aside'
import CommonFooter from '@/components/dataSecurity/Common/Footer'
export default {
  name: 'AdminView',
  components: { CommonHeader, CommonAside, CommonFooter },
  data () {
    return {}
  }
}
</script>
