<template>
  <div class="header">
    <div class="header-left">
      <span>CN Digital Security</span>
    </div>
    <div class="header-right">
      <el-button type="danger" class="button" size="small">Logout</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CommonHeader',
  methods: {
    LogOut () {
      window.sessionStorage.clear('token')
      this.$router.push('/')
    }
  }
}
</script>

<style scoped>
.header {
  height: 100%;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.header-left {
  width: 20%;
  height: 100%;
  display: flex;
  align-items: center;
  font-size: 20px;
}
span {
  margin-left: 15%;
}
.header-right{
  height: 100%;
  width: 75%;
  display: flex;
  align-items: center;
  justify-content: flex-end;
}
.button {
  margin-left: 15px;
  margin-right: 3%;
}
</style>
