<template>
  <div class="main-info">
    <Navigator :user="user" />
    <div class="container">
      <div class="row">
        <div class="col">
          <div class="alert alert-light" v-if="Object.keys(bus).length === 0 || bus.oneBus.busID === null">暂无收藏</div>
          <div v-else>
            <table class="table">
              <thead>
              <tr>
                <th scope="col">车牌号</th>
                <th scope="col">出发地</th>
                <th scope="col">目的地</th>
                <th scope="col">剩余座位</th>
                <th scope="col">出发时间</th>
                <th scope="col">抵达时间</th>
                <th scope="col">价格</th>
                <th scope="col">状态</th>
                <th scope="col">详情</th>
                <th scope="col"></th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="oneBus in bus" :key="oneBus.busID">
                <th scope="row">{{oneBus.license}}</th>
                <td>{{oneBus.departure}}</td>
                <td>{{oneBus.destination}}</td>
                <td>{{oneBus.emptySeats}}/{{oneBus.totalSeats}}</td>
                <td>{{oneBus.beginAt}}</td>
                <td>{{oneBus.endAt}}</td>
                <td>{{oneBus.price}}</td>
                <td v-if="oneBus.status === 0">即将运营</td>
                <td v-else-if="oneBus.status === 1">运营中</td>
                <td v-else-if="oneBus.status === -1">已废弃</td>
                <td v-else>未知</td>
                <td><button class="btn btn-primary" :user="user" @click="moreInfo(oneBus.busID)">更多</button></td>
                <td><button class="btn btn-danger"  @click="deleteFavorite(oneBus.busID)">取消收藏</button></td>
              </tr>
              </tbody>
            </table>
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
    name: 'Favorites',
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
          oneBus: {
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
          }
        }
      }
    },
    methods: {
      moreInfo (busID) {
        this.$router.push('/bus/' + busID)
      },
      deleteFavorite(busID) {
        const self = this
        $.delete(api + 'favorite', busID).then(function (response) {
          if (response.data.code == 200)
          {
            self.favor = false
          }
        })
        $.get(api + 'favorite').then(function (response) {
          if (response.data.code === 200) {
            self.bus = response.data.bus
          }
        })
      }
    },
    beforeMount () {
      const self = this
      $.get(api + 'favorite').then(function (response) {
        if (response.data.code === 200) {
          self.bus = response.data.bus
        }
      })
    }
  }
</script>

<style scoped>

</style>
