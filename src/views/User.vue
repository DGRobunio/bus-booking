<template>
  <div class="user">
    <Navigator :user="user" />
    <div class="user-page">
      <div class="user-info">
        <div class="container">
          <img class="mb-4" src="../assets/logo.png" alt="" width="72"
               height="72">
          <table class="table table-bordered table-hover">
            <thead class="thead-light">
            <tr>
              <th colspan="2">用户信息</th>
            </tr>
            </thead>
            <tbody>
            <tr>
              <th>用户ID</th>
              <td>{{userInfo.userID}}</td>
            </tr>
            <tr>
              <th>手机号</th>
              <td>{{userInfo.account}}</td>
            </tr>
            <tr>
              <th>余额</th>
              <td>{{userInfo.balance}} 元</td>
            </tr>
            </tbody>
          </table>
          <div class="mb-4">
            <button @click="changePasswordOnClick" class="btn btn-primary">修改密码</button>
            <router-link to="/orders" class="btn btn-primary">我的订单</router-link>
            <button @click="billingOnClick" class="btn btn-primary">充值余额</button>
          </div>
          <div class="mb-4">
            <ChangePassword @reload="reload" v-if="changePasswordFlag === true" :user="user" />
          </div>
          <div class="mb-4">
            <Billing @reload="reload" v-if="billingFlag === true" :user="user" />
          </div>
        </div>
      </div>
    </div>
    <Foot />
  </div>
</template>

<script>
  import Navigator from '../components/Navigator'
  import Billing from '../components/Billing'
  import ChangePassword from '../components/ChangePassword'
  import Foot from '../components/Foot'

  export default {
    name: 'User',
    components: {
      Navigator,
      Billing,
      ChangePassword,
      Foot
    },
    props: {
      user: null
    },
    data () {
      return {
        billingFlag: false,
        changePasswordFlag: false,
        userInfo: {
          userID: null,
          account: null,
          balance: null,
          isAdmin: null
        }
      }
    },
    methods: {
      billingOnClick () {
        const self = this
        if(self.billingFlag) {
          self.billingFlag = false
        } else {
          self.billingFlag = true
          self.changePasswordFlag = false
        }
      },
      changePasswordOnClick () {
        const self = this
        if (self.changePasswordFlag) {
          self.changePasswordFlag = false
        } else {
          self.changePasswordFlag = true
          self.billingFlag = false
        }
      },
      reload () {
        console.log("reload")
        const self = this
        self.$emit('update')
        $.get(api + 'user').then(function (response) {
          self.userInfo = response.data.user
        })
        self.billingFlag = false
        self.changePasswordFlag = false
      }
    },
    beforeMount () {
      const self = this
      $.get(api + 'user').then(function (response) {
        self.userInfo = response.data.user
      })
    }
  }
</script>

<style scoped>
  .user-page {
    height: 100%;
    display: -ms-flexbox;
    display: flex;
    -ms-flex-align: center;
    align-items: center;
    text-align: center;
    padding-top: 40px;
    padding-bottom: 40px;
  }

  .user-info {
    width: 100%;
    max-width: 480px;
    padding: 15px;
    margin: auto;
  }

  .table {
    margin-bottom: 40px;
  }

  .user-page .btn {
    margin-top: 10px;
    margin-bottom: 10px;
  }
</style>
