<template>
  <div class="update-orders">
    <Navigator :user="user" />
    <div class="container">
      <div class="col">
        <div class="row">
          <div class="m-auto">
            <div>
              <h4 class="text m-auto" style="padding-top: 20px">请将二维码对准摄像头：</h4>
              <qrcode-stream style="width: 400px;height: 400px;" @decode="onDecode" @init="onInit"></qrcode-stream>
            </div>
          </div>
          <form @submit.prevent="submit" class="update-orders-form">
            <div class="form-row">
              <label for="orderID" class="sr-only">或请输入订单号</label>
              <input v-model="putData.orderID" type="text" id="orderID" class="col-sm-9" placeholder="或在此手动输入订单号" required>
              <button class="btn btn-primary col-sm-3" type="submit">确认单号</button>
            </div>
            <div v-if="tip.status === 'success'" class="alert alert-success">
              {{tip.message}}
            </div>
            <div v-if="tip.status === 'fail'" class="alert alert-danger">
              {{tip.message}}
            </div>
          </form>
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
    name: 'UpdateOrders',
    components: {
      Navigator,
      Foot
    },
    props: {
      user: null
    },
    data () {
      return {
        putData: {
          orderID: null,
          status: 1
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
        $.put(api + 'order/' + self.putData.orderID, self.putData).then(function (response) {
          if (response.status === 200) {
            self.tip.status = 'success'
            self.tip.message = '手动操作成功！两秒内刷新页面...'
            setTimeout(() => {
              self.tip.status = null
              self.tip.message = null
              self.putData.orderID = null
            }, 2000)
          }
        }).catch(function (error) {
          console.log(error)
          self.tip.status ='fail'
          self.tip.message ='操作失败！'
          setTimeout(() => {
            self.tip.status = null
            self.tip.message = null
            self.putData.orderID = null
          }, 2000)
        })
      },
      onDecode (content) {
        const self = this
        if(content !== undefined) {
          self.putData.orderID = eval(content.split('^')[0])
          self.submit()
        }
      },

      onInit (promise) {
        const self = this
        promise.then(() => {
          console.log('Successfully initilized! Ready for scanning now!')
        })
        .catch(error => {
          self.tip.status = 'fail'
          if (error.name === 'NotAllowedError') {
            self.tip.message = '需要相机权限'
          } else if (error.name === 'NotFoundError') {
            self.tip.message = '未找到设备'
          } else if (error.name === 'NotSupportedError') {
            self.tip.message = '当前页面不支持此操作'
          } else if (error.name === 'NotReadableError') {
            self.tip.message = '无法获取相机权限，请检查是否呗占用'
          } else if (error.name === 'OverconstrainedError') {
            self.tip.message = '无法匹配任何已有相机，请检查。'
          } else {
            self.tip.message = '未知错误: ' + error.message
          }
        })
      }
    }
  }
</script>

<style scoped>
  .update-orders-form {
    width: 100%;
    max-width: 960px;
    padding: 15px;
    margin: auto;
  }
</style>
