<template>
  <div class="bus-info">
    <Navigator :user="user" />

    <div class="container">
      <div class="row">
        <div class="col">
          <!-- <div @adminaddbus.prevent="adminaddbus" class="update-form"> -->
          <form @adminaddbus.prevent="adminaddbus" class="update-form">
          <table class="table table-bordered table-hover">
            <thead class="thead-light">
            <tr>
              <th colspan="2">添加线路</th>
            </tr>
            </thead>
            <tbody>
            <tr>
              <th>班车车牌号</th>
              <label for="license" class="sr-only">请按照格式更改班车车牌号！</label>
              <input v-model="bus.license" id="license" class="form-control" placeholder="请按照格式更改班车车牌号！" required>
            </tr>
            <tr>
              <th>出发时间</th>
              <label for="beginAt" class="sr-only">请输入出发时间！</label>
              <input v-model="bus.beginAt" id="beginAt" class="form-control" placeholder="请输入出发时间！" required>
            </tr>
            <tr>
              <th>抵达时间</th>
              <label for="endAt" class="sr-only">请输入抵达时间！</label>
              <input v-model="bus.endAt" id="endAt" class="form-control" placeholder="请输入抵达时间！" required>
            </tr>
            <tr>
              <th>出发地</th>
              <label for="departure" class="sr-only">请输入出发地！</label>
              <input v-model="bus.departure" id="departure" class="form-control" placeholder="请输入出发地！" required>
            </tr>
            <tr>
              <th>目的地</th>
              <label for="destination" class="sr-only">请输入目的地！</label>
              <input v-model="bus.destination" id="destination" class="form-control" placeholder="请输入目的地！" required>
            </tr>
            <tr>
              <th>票务价格</th>
              <label for="price" class="sr-only">请输入票务价格！</label>
              <input v-model="bus.price" id="price" class="form-control" placeholder="请输入票务价格！" required>
            </tr>
            <tr>
              <th>每周运营时间</th>
              <label for="weekly" class="sr-only">请输入每周运营时间！</label>
              <input v-model="bus.weekly" id="weekly" class="form-control" placeholder="请输入每周运营时间！" required>
            </tr>
            <tr>
              <th>剩余座位</th>
              <label for="totalSeats" class="sr-only">请输入总共座位数！</label>
              <input v-model="bus.totalSeats" id="totalSeats" class="form-control" placeholder="请输入总共座位数！" required>
            </tr>
            <tr>
              <th>信息介绍</th>
              <label for="info" class="sr-only">请输入信息介绍！</label>
              <input v-model="bus.info" id="info" class="form-control" placeholder="请输入信息介绍！" required>
            </tr>
            <tr>
              <th>班车状态</th>
              <label for="status" class="sr-only">请输入班车状态！（ 0 即将运营， 1 运营中， -1 已废弃）</label>
              <input v-model="bus.status" id="status" class="form-control" placeholder="请输入班车状态！" required>
            </tr>
            </tbody>
          </table>
          <div v-if="tip.type === 0" class="alert alert-warning form-control"> {{tip.message}} </div>
          <div v-if="tip.type === 1" class="alert alert-success form-control"> {{tip.message}} </div>
          <button class="btn btn-lg btn-primary " @click="adminaddbus(bus.license, bus.totalSeats, bus.departure, bus.destination, bus.beginAt, bus.endAt, bus.price, bus.info, bus.weekly, bus.status)">确认修改</button>
          </form>
          <!-- </div> -->
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
    name: 'AdminAddBus',
    components: {
      Navigator,
      Foot
    },
    props: {
      user: null,
    },
    data(){
      return{
        bus: {
          license: null,
          totalSeats: null,
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
          type: null,
          message: null
        }
      }
    },
    methods: {
      adminaddbus(license, totalSeats, departure, destination, beginAt, endAt, price, info, weekly, status){
        const self = this
        if (license === null || totalSeats === null || departure === null || destination === null || beginAt === null || endAt === null || price === null || info === null || weekly === null || status === null)
        {
          self.tip.type = 0
          self.tip.message = "请输入完整信息！"
        }
        else{
          $.put(api + 'bus', self.bus).then(function (response) {
            if (response.status === 200) {
              self.tip.type = 1
              self.tip.message = "修改成功！"
              setTimeout(() => {
                self.$emit('update')
              }, 1000)
            } else {
              self.tip.type = 0
              self.tip.message = "网络繁忙，请稍后再试！"
            }
          })
        }
      }
    },
    beforeMount () {
      const self = this
      if (self.user.userID === '' || !self.user.isAdmin) {
        self.$router.push('/404')
      } else {
        console.log(this.$route.params.busID)
      }
    }
  }

</script>

<style scoped>

</style>
