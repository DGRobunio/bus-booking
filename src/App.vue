<template>
  <div id="app">
    <router-view :user="user" @update="update" />
  </div>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      user: {
        userID: null,
        account: null,
        balance: null,
        isAdmin: null
      }
    }
  },
  beforeMount () {
    const self = this
    $.get(api).then(function (response) {
      if(response.status === 200) {
        self.user = response.data.user
      } else {
        self.user.userID = ''
      }
    })
  },
  methods: {
    update () {
      const self = this
      $.get(api).then(function (response) {
        if(response.status === 200) {
          self.user = response.data.user
        } else {
          self.user.userID = ''
        }
      })
    }
  }
}
</script>

<style>
#app {
  padding-top: 3.5rem;
}
</style>
