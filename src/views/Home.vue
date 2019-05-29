<template>
  <div class="home">
    <Navigator :user="user" />
    <MainInfo :bus="bus" />
    <PieChart :id="'pieChart'" :option="option" v-if="showChart === true" style="width:244px;height:244px" />
    <Foot/>
  </div>
</template>

<script>
import Navigator from '../components/Navigator'
import MainInfo from '../components/MainInfo'
import PieChart from '../components/PieChart'
import Foot from '../components/Foot'

export default {
  name: 'Home',
  components: {
    Navigator,
    MainInfo,
    PieChart,
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
      },
      showChart:false,
      option: {
        title: {text: '班车运营收入占比'},
        tooltip: {
          pointFormat: '总收入：{point.y} 元'
        },
        plotOptions: {
          pie: {
            allowPointSelect: true,
            cursor: 'pointer',
            dataLabels: {
              enabled: true
            },
            showInLegend: true
          }
        },
        series: [{
          type: 'pie',
          data: [],
        }]
      }
    }
  },
  methods: {
    reload () {
      const self = this
      $.get(api + 'bus').then(function (response) {
        if (response.status === 200) {
          self.bus = response.data.bus
        }
      })
    }
  },
  mounted () {
    const self = this
    setTimeout(() => {
      self.showChart = self.user.isAdmin
    },500)
  },
  beforeMount () {
    const self = this
    self.$emit('update')
    $.get(api + 'bus').then(function (response) {
      if (response.status === 200) {
        self.bus = response.data.bus
      }
    })
    if(self.user.isAdmin) {
      $.get(api + 'figure').then(function (response) {
        if (response.status === 200) {
          response.data.figure.forEach((item) => {
            self.option.series[0].data.push({name: item.busID, y: eval(item.income)})
          })
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
