<template>
  <div class="orders">
    <Navigator :user="user" />
    <div class="container">
      <div class="row">
        <div v-if="!display.qrCodeFlag" class="col">
          <div v-if="Object.keys(order).length === 0 " class="alert alert-warning col-sm-12">暂无订单。</div>
          <div v-else class="accordion" id="order-list">
            <div class="card">
              <div class="card-header" id="unused">
                <button class="btn btn-lg" type="button" data-toggle="collapse" data-target="#unused-body" aria-expanded="true" aria-controls="unused-body">未使用订单</button>
              </div>
              <div id="unused-body" class="collapse show" aria-labelledby="unused" data-parent="#order-list">
                <div class="card-body">
                  <div class="list-group">
                    <div v-for="oneOrder in order" :key="oneOrder.orderID" >
                      <div @click="generateQRCode(oneOrder.orderID)" v-if="oneOrder.status === 0" class="list-group-item list-group-item-action">
                        订单编号:{{oneOrder.orderID}} 订票时间:{{oneOrder.orderAt}} 班车牌号:{{oneOrder.bus.license}} 出发时间:{{oneOrder.bus.beginAt}}
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="card">
              <div class="card-header" id="used">
                <button class="btn btn-lg" type="button" data-toggle="collapse" data-target="#used-body" aria-expanded="false" aria-controls="used-body">已使用订单</button>
              </div>
              <div id="used-body" class="collapse" aria-labelledby="used" data-parent="#order-list">
                <div class="card-body">
                  <div class="list-group">
                    <div v-for="oneOrder in order" :key="oneOrder.orderID" >
                      <div v-if="oneOrder.status === 1" class="list-group-item list-group-item-action">
                        <p>
                          订单编号:{{oneOrder.orderID}} 订票时间:{{oneOrder.orderAt}} 班车牌号:{{oneOrder.bus.license}} 出发时间:{{oneOrder.bus.beginAt}}
                        </p>
                        <router-link :to="'/comment/' + oneOrder.orderID">
                          前往评论
                        </router-link>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div v-else class="col">
          <div class="row justify-content-sm-center">
            <div class="col-sm-12 alert alert-light text-center">请出示二维码以检票：</div>
          </div>
          <div class="row justify-content-sm-center">
            <div class="m-auto">
              <div id="qrcode" ref="qrcode"/>
            </div>
          </div>
          <div class="row justify-content-sm-center">
            <div class="col-sm-12 alert alert-light text-center">或者使用以下订单号：{{display.id}}</div>
            <button @click="returnOrders" class="col-sm-4 btn btn-primary">返回</button>
          </div>
        </div>
      </div>
    </div>
    <Foot />
  </div>
</template>

<script>
  import Navigator from '../components/Navigator'
  import Foot from '../components/Foot'

  export default {
    name: 'Orders',
    components: {
      Navigator,
      Foot
    },
    props: {
      user: null,
    },
    data () {
      return {
        order: {
          oneOrder: {
            orderID: null,
            bus: null,
            orderAt: null,
            status: null
          }
        },
        display: {
          qrCodeFlag: false,
          id: null,
          config : {
            value: null,
            imagePath: '/assets/logo.png',
            size: 100,
            filter: 'color'
          }
        }

      }
    },
    methods: {
      returnOrders () {
        const self = this
        self.display.qrCodeFlag = false
      },
      generateQRCode (orderID) {
        const self = this
        self.display.qrCodeFlag = true
        self.display.id = orderID
        self.display.config.value = '{"orderID": '+ orderID + '}'
        // self.qrcode(self.display.config.value)
        setTimeout(() => {
          new QRCode(this.$refs.qrcode, self.display.config.value)
        }, 100)
      }
    },
    beforeMount () {
      const self = this
      self.display.qrCodeFlag = false
      $.get(api + 'order').then(function (response) {
        if (response.status === 200) {
          self.order = response.data.order
        }
      })
    }
  }
</script>

<style scoped>
  .center-div {
    height: 100%;
    display: -ms-flexbox;
    display: flex;
    -ms-flex-align: center;
    align-items: center;
    text-align: center;
    padding-top: 40px;
    padding-bottom: 40px;
    margin-left: auto;
    margin-right: auto;
  }
</style>
