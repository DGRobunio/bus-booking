<template>
  <div class="navigator">
    <div class="nav navbar navbar-expand-md navbar-dark fixed-top bg-dark" role="navigation">
      <router-link to="/" class="navbar-brand">班车订票系统</router-link>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul v-if="user.userID === ''" class="navbar-nav ml-auto">
          <li class="nav-item mr-auto">
            <router-link to="/register" class="nav-link">注册</router-link>
          </li>
          <li class="nav-item mr-auto">
            <router-link to="/login" class="nav-link">登录</router-link>
          </li>
        </ul>
        <ul v-else class="nav bar-nav ml-auto">
          <li class="nav-item dropdown">
            <a class="nav-link text-light dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Welcome, {{user.userID}}! </a>
            <div class="dropdown-menu" aria-labelledby="dropdown">
              <router-link to="/user" class="dropdown-item">余额：{{user.balance}}元</router-link>
              <router-link to="/orders" class="dropdown-item">我的订单</router-link>
              <router-link to="/favorites" :user="user" class="dropdown-item">收藏夹</router-link>
            </div>
          </li>
          <li class="nav-item">
            <button v-on:click="logout" class=" btn btn-dark">登出</button>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Navigator',
  props: {
    user: {
      userID: null,
      account: null,
      balance: null,
      isAdmin: null
    }
  },
  data () {
    return {
      search: null
    }
  },
  methods: {
    logout() {
      $.get(api + 'logout').then(function (response) {
        console.log(response.data.code)
        if (response.status === 200) {
          cookie.deleteCookie('session')
        this.$emit('update')
        this.$router.push('/')
        }
      })
    }
  }
}
</script>
