import Vue from 'vue'
import VueRouter from 'vue-router'
import AdminView from '../views/AdminView'
import HomePage from '@/components/homePage'
import DataInventory from '@/components/dataSecurity/Identification/DataInventory'
import DataDetailList from '@/components/dataSecurity/Identification/DataField'
import ConfigRules from '@/components/dataSecurity/Configuration/ConfigRules'
import ConfigCategory from '@/components/dataSecurity/Configuration/ConfigCategory'
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
      { path: '/admin/DetailList', component: DataDetailList, meta: { title: 'Data List' } },
      { path: '/admin/ConfigRules', component: ConfigRules, meta: { title: 'Config Rules' } },
      { path: '/admin/ConfigCategory', component: ConfigCategory, meta: { title: 'Config Category' } }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
