<template>
  <div class="billing">
    <form @submit.prevent="submit" class="billing-form">
      <div class="form-row">
        <label for="key_input" class="sr-only">充值码</label>
        <input v-model="key" type="text" id="key_input" class="form-control col-sm-9" placeholder="在此输入充值码">
        <button class="btn btn-primary btn-block col-sm-3" type="submit">提交</button>
      </div>
      <div v-if="tip.status === 'success'" class="alert alert-success">
        {{tip.message}}
      </div>
      <div v-if="tip.status === 'warn'" class="alert alert-warning">
        {{tip.message}}
      </div>
      <div v-if="tip.status === 'fail'" class="alert alert-danger">
        {{tip.message}}
      </div>
    </form>
  </div>
</template>

<script>
  export default {
    name: 'Billing',
    props: {
      user: null
    },
    data () {
      return {
        key: null,
        tip: {
          status: null,
          message: null
        }
      }
    },
    methods: {
      submit () {
        const self = this
        $.put(api + 'billing', self.key).then(function (response) {
          if (response.status === 200) {
            if (response.data.user.balance > self.user.balance) {
              self.tip.status = 'success'
              self.tip.message = '充值成功！两秒内刷新页面'
              setTimeout(() => {
                self.key = null
                self.$emit('update')
                self.$emit('reload')
              })
            } else {
              self.tip.status ='fail'
              self.tip.message ='充值失败！请确认充值码是否有误！'
            }
          } else {
            self.tip.status = 'warn'
            self.tip.message ='服务器繁忙！请稍后再试！'
          }
        })
      }
    }
  }
</script>

<style scoped>
  .billing-form {
    width: 100%;
    max-width: 960px;
    padding: 15px;
    margin: auto;
  }
</style>
