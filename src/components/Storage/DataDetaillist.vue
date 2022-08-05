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
      <!-- 动态加载 筛选 -->
      <el-cascader :options="options" :show-all-levels="false" clearable filterable :props="props"></el-cascader>
    </el-card>
  </div>
</template>

<script>
let id = 0
export default {
  data () {
    return {
      props: {
        lazy: true,
        lazyLoad (node, resolve) {
          const { level } = node
          setTimeout(() => {
            const nodes = Array.from({ length: level + 1 })
              .map(item => ({
                value: ++id,
                label: `选项${id}`,
                leaf: level >= 2
              }))
            // 通过调用resolve将子节点数据返回，通知组件数据加载完成
            resolve(nodes)
          }, 1000)
        }
      }
    }
  },
  name: 'DataDetaillist'
}
</script>

<style scoped>

</style>
