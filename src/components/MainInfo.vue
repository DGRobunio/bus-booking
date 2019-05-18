<template>
  <div class="main-info">
    <div class="container">
      <div class="row">
        <div class="col">
          <div class="alert alert-danger" v-if="Object.keys(bus).length === 0 || bus.oneBus.busID === null">当前没有班车，如果信息有误请联系管理员。</div>
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
                  </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'MainInfo',
    props: {
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
    },
    methods: {
      moreInfo (busID) {
        this.$router.push('/bus/' + busID)
      }
    }
  }
</script>

<style scoped>

</style>
