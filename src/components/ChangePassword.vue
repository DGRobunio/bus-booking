<template>
  <div class="change-password-component">
    <form @submit.prevent="submit" class="update-form">
      <label for="oldPass" class="sr-only">旧密码</label>
      <input v-model="formInfo.oldPass" type="password" id="oldPass" class="form-control" placeholder="旧密码" required>
      <label for="newPass" class="sr-only">新密码</label>
      <input v-model="formInfo.newPass" type="password" id="newPass" class="form-control" placeholder="新密码" required>
      <label for="repeatPass" class="sr-only">重复新密码</label>
      <input v-model="formInfo.repeatPass" type="password" id="repeatPass" class="form-control" placeholder="重复新密码" required>
      <div v-if="tip.type === 0" class="alert alert-warning form-control"> {{tip.message}} </div>
      <div v-if="tip.type === 1" class="alert alert-success form-control"> {{tip.message}} </div>
      <button class="btn btn-lg btn-primary btn-block " type="submit">修改密码</button>
    </form>
  </div>
</template>

<script>
  export default {
    name: 'ChangePassword',
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
        formInfo: {
          oldPass: null,
          newPass: null,
          repeatPass: null
        },
        tip: {
          type: null,
          message: null
        },
        passwordInfo : {
          oldPassword: null,
          newPassword: null
        }
      }
    },
    methods: {
      submit() {
        const self = this
        if (self.formInfo.oldPass === self.formInfo.newPass || self.formInfo.oldPass === self.formInfo.repeatPass) {
          self.tip.type = 0
          self.tip.message = "新旧密码不能相同！"
        } else if (self.formInfo.newPass === self.formInfo.repeatPass) {
          self.passwordInfo.oldPassword = md5(self.formInfo.oldPass)
          self.passwordInfo.newPassword = md5(self.formInfo.newPass)
          $.put(api + 'user', self.passwordInfo).then(function (response) {
            if (response.status === 200) {
              self.tip.type = 1
              self.tip.message = "修改成功！"
              setTimeout(() => {
                self.formInfo.oldPass = null
                self.formInfo.newPass = null
                self.formInfo.repeatPass = null
                self.tip.type = null
                self.$emit('relogin')
              }, 1000)
            }
          }).catch(function (error) {
            console.log(error)
            self.tip.type = 0
            self.tip.message = "原密码不符，修改失败！"
            setTimeout(() => {
              self.formInfo.oldPass = null
              self.formInfo.newPass = null
              self.formInfo.repeatPass = null
              self.tip.type = null
              self.tip.message = null
            }, 1000)
          })
        } else {
          self.tip.type = 0
          self.tip.message = "两次输入不一致！"
          self.formInfo.newPass = null
          self.formInfo.repeatPass = null
        }
      }
    }
  }
</script>

<style scoped>

  .update-form {
    width: 100%;
    max-width: 480px;
    padding: 15px;
    margin: auto;
  }

  .update-form .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 10px;
  }

  .update-form .form-control:focus {
    z-index: 2;
  }

  .form-control input {
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
  }

  .update-form .alert {
    margin-bottom: 0px;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
  }
</style>
