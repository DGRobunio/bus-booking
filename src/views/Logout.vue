<template>
    <div></div>
</template>

<script>
    export default {
        name: "Logout",
      beforeMount() {
        const self = this
        $.get(api + 'logout').then(function (response) {
          console.log(response.data.code)
          if (response.status === 200) {
            self.$emit('update')
            self.$router.push('/')
          } else {
            self.$router.push('/')
          }
        })
      }
    }
</script>

<style scoped>

</style>
