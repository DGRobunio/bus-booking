<template>
  <div class="purchase">
    <Navigator :user="user" />
    <div v-if="bus !== null" class="purchase-page">
      <div class="purchase-info">
        <div class="container">
          <table class="table table-bordered table-hover">
            <thead class="thead-light">
            <tr>
              <th colspan="2">订单详情</th>
            </tr>
            </thead>
            <tbody>
            <tr>
              <th>班车牌号</th>
              <td>{{bus.license}}</td>
            </tr>
            <tr>
              <th>出发地</th>
              <td>{{bus.departure}}</td>
            </tr>
            <tr>
              <th>目的地</th>
              <td>{{bus.destination}}</td>
            </tr>
            <tr>
              <th>出发时间</th>
              <td>{{bus.beginAt}}</td>
            </tr>
            <tr>
              <th>到达时间</th>
              <td>{{bus.endAt}}</td>
            </tr>
            <tr>
              <th>票价</th>
              <td>{{bus.price}} 元</td>
            </tr>
            <tr>
              <th>可用余额</th>
              <td>{{user.balance}} 元</td>
            </tr>
            </tbody>
          </table>
          <div class="mb-4">
            <router-link :to="'/bus/' + this.$route.params.busID" :user="user" class="btn btn-secondary">取消</router-link>
            <button v-if="user.balance - bus.price < 0 || bus.status !== 1" class="btn btn-primary disabled" aria-disabled="true">确定</button>
            <button v-else v-on:click="confirmOnClick" class="btn btn-primary">确定</button>
            <div v-if="user.balance - bus.price < 0"  class="alert alert-danger">"余额不足，请充值！"</div>
          </div>
          <div class="mb-4">
            <div v-if="tip.status === 'success'" class="alert alert-success">{{tip.message}}</div>
            <div v-if="tip.status === 'fail'" class="alert alert-danger">{{tip.message}}</div>
          </div>
        </div>
      </div>
    </div>
    <div v-else class="container">
      <div class="alert alert-danger">数据错误！</div>
    </div>
    <Foot />
  </div>
</template>

<script>
  import Navigator from '../components/Navigator'
  import Foot from '../components/Foot'

  export default {
    name: 'Purchase',
    components: {
      Navigator,
      Foot
    },
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
        bus: {
          busID: null,
          license: null,
          totalSeats: null,
          emptySeats: null,
          departure: null,
          destination: null,
          beginAt: null,
          endAt: null,
          price: null,
          info: null,
          weekly: null,
          status: null
        },
        tip: {
          status: null,
          message: null
        },
        postData: {
          busID: this.$route.params.busID
        }
      }
    },
    methods: {
      confirmOnClick () {
        const self = this
        $.post(api + 'order', self.postData).then(function (response) {
          if(response.status === 200) {
            self.tip.status = 'success'
            self.tip.message = '购票成功！正在转跳至订单列表...'
            setTimeout(() => {
              self.$emit('update')
              self.$router.push('/orders')
            }, 2000)
          }
        }).catch(function (error) {
          console.log(error)
          self.tip.status = 'fail'
          self.tip.message = '购票失败！'
          setTimeout(() => {
            self.$router.push('/')
          }, 1000)
        })
      }
    },
    beforeMount () {
      const self = this
      if(self.user.isAdmin) {
        self.$router.push('/')
      } else {
        $.get(api + 'bus/' + self.$route.params.busID).then(function (response) {
          self.bus = response.data.bus
        })
      }
    }
  }
</script>

<style scoped>
  .purchase-page {
    height: 100%;
    display: -ms-flexbox;
    display: flex;
    -ms-flex-align: center;
    align-items: center;
    text-align: center;
    padding-top: 40px;
    padding-bottom: 40px;
  }

  .purchase-info {
    width: 100%;
    max-width: 480px;
    padding: 15px;
    margin: auto;
  }

  .table {
    margin-bottom: 40px;
  }

  .purchase-page .btn {
    margin-top: 10px;
    margin-bottom: 10px;
  }
</style>
