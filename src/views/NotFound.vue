<template>
  <div class="not-found">
    <Navigator :user="user" />
    <div class="container">
      <div class="row">
        <div class="col">
          <div class="alert alert-danger col-sm-12">页面不存在！</div>
          <router-link to="/" class="btn btn-secondary">返回主页</router-link>
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
    name: 'NotFound',
    components: {
      Navigator,
      Foot
    },
    props: {
      user: null
    }
  }
</script>

<style scoped>

</style>
