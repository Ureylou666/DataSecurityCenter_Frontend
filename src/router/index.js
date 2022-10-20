import Vue from 'vue'
import VueRouter from 'vue-router'
import AdminView from '../views/AdminView'
import HomePage from '@/components/homePage'
import DataInventory from '@/components/dataSecurity/Storage/DataInventory'
import DataDetaillist from '@/components/dataSecurity/Storage/DataField'
import DataRules from '@/components/dataSecurity/Storage/DataRules'
import GroupSetting from '@/components/settings/GroupSetting'
Vue.use(VueRouter)

const routes = [
  {
    path: '/admin',
    component: AdminView,
    redirect: '/admin/welcome',
    children: [
      { path: '/admin/welcome', component: HomePage, meta: { title: 'Home Page' } },
      { path: '/admin/PrivacyCenter', component: HomePage, meta: { title: 'Privacy Center' } },
      { path: '/admin/Inventory', component: DataInventory, meta: { title: 'Data Inventory' } },
      { path: '/admin/Detaillist', component: DataDetaillist, meta: { title: 'Data List' } },
      { path: '/admin/Rules', component: DataRules, meta: { title: 'Classification Rules' } },
      { path: '/admin/Group', component: GroupSetting, meta: { title: 'Dev Group' } }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
