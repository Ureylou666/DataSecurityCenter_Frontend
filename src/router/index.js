import Vue from 'vue'
import VueRouter from 'vue-router'
import AdminView from '../views/AdminView'
import HomePage from '@/components/homePage'
import DataInventory from '@/components/dataSecurity/Storage/DataInventory'
import DataDetaillist from '@/components/dataSecurity/Storage/DataDetaillist'
import DataRules from '@/components/dataSecurity/Storage/DataRules'
import GroupSetting from '@/components/settings/GroupSetting'
Vue.use(VueRouter)

const routes = [
  {
    path: '/admin',
    component: AdminView,
    redirect: '/admin/welcome',
    children: [
      { path: '/admin/welcome', component: HomePage, meta: { title: '首页' } },
      { path: '/admin/Inventory', component: DataInventory, meta: { title: '数据资产' } },
      { path: '/admin/Detaillist', component: DataDetaillist, meta: { title: '数据清单' } },
      { path: '/admin/Rules', component: DataRules, meta: { title: '分类分级规则' } },
      { path: '/admin/Group', component: GroupSetting, meta: { title: '项目配置' } }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
