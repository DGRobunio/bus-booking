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
      user: null
    }
  },
  beforeMount () {
    const self = this
    $.get(api).then(function (response) {
      self.user = response.data.user
    })
  },
  methods: {
    update () {
      const self = this
      $.get(api).then(function (response) {
        self.user = response.data.user
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
