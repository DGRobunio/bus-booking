<template>
  <div class="login">
    <Navigator :user="user" />
    <div class="login-page">
      <form @submit.prevent="submit" class="login-form">
        <img class="mb-4" src="../assets/logo.png" alt="" width="72" height="72">
        <h1 class="h3 mb-3 font-weight-normal">请登录</h1>
        <label for="user_name" class="sr-only">账号(手机号)</label>
        <input v-model="login.account" type="text" id="user_name" class="form-control" placeholder="账号(手机号)" required autofocus>
        <label for="user_password" class="sr-only">密码</label>
        <input v-model="login.password" type="password" id="user_password" class="form-control" placeholder="密码" required>
        <div v-if="tip.status === 'success'" class="alert alert-success">{{tip.message}}</div>
        <div v-if="tip.status === 'fail'" class="alert alert-danger">{{tip.message}}</div>
        <button class="btn btn-lg btn-primary btn-block" type="submit">登录</button>
      </form>
    </div>
    <Foot />
  </div>
</template>

<script>
  import Navigator from '../components/Navigator'
  import Foot from '../components/Foot'

  export default {
    name: 'login',
    components: {
      Navigator,
      Foot
    },
    props: {
      user: null
    },
    data() {
      return {
        login: {
          account: null,
          password: null
        },
        tip: {
          status: null,
          message: null
        }
      }
    },
    methods: {
      submit() {
        const self = this
        self.login.password = md5(self.login.password)
        $.post(api + 'login', self.login).then(function (response) {
          if (response.data.code === "200")
          {
            self.tip.status = 'success'
            self.tip.message = '登录成功!'
            self.$emit('update')
            setTimeout(() => {
              self.tip.message = '将在两秒内转跳....'
            }, 1000)
            setTimeout(() => {
              self.$router.push('/')
            }, 2000)
          } else {
            self.tip.status = 'fail'
            self.tip.message = '登录失败！用户名或密码错误！'
            self.login.account = null
            self.login.password = null
            setTimeout(() => {
              self.tip.status = null
              self.tip.message = null
            }, 2000)
          }
        })
      }
    }
  }
</script>

<style scoped>
  .login-page {
    height: 100%;
    display: -ms-flexbox;
    display: flex;
    -ms-flex-align: center;
    align-items: center;
    text-align: center;
    padding-top: 120px;
    padding-bottom: 120px;
  }

  .login-form {
    width: 100%;
    max-width: 480px;
    padding: 15px;
    margin: auto;
  }

  .login-form .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 16px;
    font-size: 16px;
  }

  .login-form .form-control:focus {
    z-index: 2;
  }

  .login-form input[type="text"] {
    margin-top: 40px;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
  }

  .login-form input[type="password"] {
    margin-bottom: 40px;
    border-top-left-radius: 0;
    border-top-right-radius: 0;
  }

  .login-form .alert {
    margin-bottom: 40px;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
  }
</style>
