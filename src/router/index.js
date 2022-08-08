import Vue from 'vue'
import VueRouter from 'vue-router'
import AdminView from '../views/AdminView'
import DscBoard from '@/components/Dashboard/DscBoard'
import DataInventory from '@/components/Storage/DataInventory'
import DataDetaillist from '@/components/Storage/DataDetaillist'
import DataRules from '@/components/Storage/DataRules'
Vue.use(VueRouter)

const routes = [
  {
    path: '/admin',
    component: AdminView,
    redirect: '/welcome',
    children: [
      { path: '/welcome', component: DscBoard },
      { path: '/Inventory', component: DataInventory },
      { path: '/Detaillist', component: DataDetaillist },
      { path: '/Rules', component: DataRules }
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router
