<template>
  <div class="register">
    <Navigator :user="user" />
    <div class="register-page">
      <form @submit.prevent="submit" class="register-form">
        <img class="mb-4" src="../assets/logo.png" alt="" width="72" height="72">
        <h1 class="h3 mb-3 font-weight-normal">Please Input your information.</h1>
        <label for="account" class="sr-only">用户名</label>
        <input v-model="register.account" type="text" id="account" class="form-control" placeholder="手机号" required autofocus>
        <label for="password" class="sr-only">密码</label>
        <input v-model="register.password" type="password" id="password" class="form-control" placeholder="密码" required>
        <label for="confirmPassword" class="sr-only">Confirm password</label>
        <input v-model="register.confirmPassword" type="password" id="confirmPassword" class="form-control" placeholder="确认密码" required>
        <div v-if="tip.status === 'success'" class="alert alert-success">
          {{tip.message}}
        </div>
        <div v-if="tip.status === 'warn'" class="alert alert-warning">
          {{tip.message}}
        </div>
        <div v-if="tip.status === 'fail'" class="alert alert-danger">
          {{tip.message}}
        </div>
        <button class="btn btn-lg btn-primary btn-block" type="submit">注册</button>
      </form>
    </div>
    <Foot />
  </div>
</template>

<script>
  import Navigator from '../components/Navigator'
  import Foot from '../components/Foot'

  export default {
    name: 'Register',
    components: {
      Navigator,
      Foot
    },
    props: {
      user: {
        userID: null,
        isAdmin: null,
        account: null,
        balance: null
      }
    },
    data () {
      return {
        register: {
          account: null,
          password: null,
          confirmPassword: null
        },
        registerInfo: {
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
      submit () {
        const self = this
        if(self.register.password !== self.register.confirmPassword)
        {
          self.tip.status = 'warn'
          self.tip.message = '密码不一致！'
          self.register.password = null
          self.register.confirmPassword=null
          setTimeout(() => {
            self.tip.status = null
            self.tip.message = null
          }, 2000)
        } else {
          self.registerInfo.password = md5(self.register.password)
          self.registerInfo.account = self.register.account
          $.post(api + 'user', self.registerInfo).then(function (response) {
            if (response.data.code === 200) {
              self.tip.status = 'success'
              self.tip.message = '注册成功!'
              self.$emit('update')
              setTimeout(() => {
                self.tip.message = '两秒内跳转...'
              }, 1000)
              setTimeout(() => {
                self.$router.push('/')
              }, 2000)
            } else {
              self.tip.status = 'fail'
              self.tip.message = '注册失败！'
              self.register.account = null
              self.register.password = null
              self.register.confirmPassword = null
              setTimeout(() => {
                self.tip.status = null
                self.tip.message = null
              }, 2000)
            }
          })
        }
      }
    }
  }
</script>

<style scoped>
  .register-page {
    height: 100%;
    display: -ms-flexbox;
    display: flex;
    -ms-flex-align: center;
    align-items: center;
    text-align: center;
    padding-top: 120px;
    padding-bottom: 120px;
  }

  .register-form {
    width: 100%;
    max-width: 480px;
    padding: 15px;
    margin: auto;
  }

  .register-form .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 16px;
    font-size: 16px;
  }

  .register-form .form-control:focus {
    z-index: 2;
  }

  .form-control input {
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
  }

  .register-form input[type="text"] {
    margin-top: 40px;
  }

  .register-form input[type="tel"] {
    margin-bottom: 40px;
  }

  .register-form .alert {
    margin-bottom: 40px;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
  }
</style>
