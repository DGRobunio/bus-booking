<template>
  <div class="home">
    <Navigator :user="user" />
    <MainInfo :bus="bus" />
    <Foot/>
  </div>
</template>

<script>
import Navigator from '../components/Navigator'
import MainInfo from '../components/MainInfo'
import Foot from '../components/Foot'

export default {
  name: 'Home',
  components: {
    Navigator,
    MainInfo,
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
    reload () {
      const self = this
      $.get(api + 'bus').then(function (response) {
        if (response.status === 200) {
          self.bus = response.data.bus
        }
      })
    }
  },
  beforeMount () {
    const self = this
    $.get(api + 'bus').then(function (response) {
      if (response.status === 200) {
        self.bus = response.data.bus
      }
    })
  }
}
</script>

<style scoped>

</style>
