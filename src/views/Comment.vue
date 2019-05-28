<template>
  <div class="passage_comment">
    <Navigator :user="user" />
    <div v-if="user.isAdmin === false" class="container">
      <div class="row">
        <div class="col">
          <table class="table table-bordered table-hover">
            <thead class="thead-light">
            <tr>
              <th colspan="2">班车评价</th>
            </tr>
            </thead>
          </table>
          <form @submit.prevent="submit" class="col">
            <div class="form-group row">
              <legend class="col-form-label col-sm-3">评分:</legend>
              <div class="col-sm-9">
                <label class="form-check-label col-sm-2" v-for="(item,index) in radioData" :key="index">
                  <input @click="getRadioVal" type="radio" name="type" :value="index + 1" required>{{item.value}}星
                </label>
              </div>
            </div>
            <div class="form-group row">
              <label for="content" class="col-form-label col-sm-3">评价:</label>
              <div class="col-sm-9">
                <textarea v-model="comment.content" class="form-control" id="content" placeholder="请在此输入您的评价"></textarea>
              </div>
            </div>
            <button class="btn btn-primary col-sm-2" type="submit">提交</button>
            <div v-if="tip.type === 0" class="alert alert-warning form-control"> {{tip.message}} </div>
            <div v-if="tip.type === 1" class="alert alert-success form-control"> {{tip.message}} </div>
          </form>
        </div>
      </div>
    </div>
    <div v-else >
      <div class="container">
        <div class="row">
          <div class="col">
            <div class="alert alert-danger col-sm-12">页面不存在！</div>
            <router-link to="/" class="btn btn-secondary">返回主页</router-link>
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
    name: 'comment',
    components: {
      Navigator,
      Foot,
    },
    props: {
      user: null,
    },
    data (){
      return{
        stars: null,
        tip: {
          type: null,
          message: null
        },
        radioData:[
          {value: 1},
          {value: 2},
          {value: 3},
          {value: 4},
          {value: 5}
        ],
        comment: {
          orderID: this.$route.params.orderID,
          busID: null,
          content: null,
          stars: null
        }
      }
    },
    methods: {
      getRadioVal(index){
        const self = this
        self.stars = index.target.value;
        console.log(self.stars)
      },
      submit(){
        const self = this
        self.comment.stars = self.stars
        $.post(api + 'comment', self.comment).then(function (response) {
          if (response.status === 200){
            self.tip.type = 1
            self.tip.message = "提交成功！"
            self.$emit('update')
            setTimeout(() => {
              self.tip.message = '两秒内跳转...'
            }, 1000)
            setTimeout(() => {
              self.$router.push('/orders')
            }, 2000)
          }
          else{
            self.tip.type = 0
            self.tip.message = "网络繁忙！"
          }
        })
      }
    },
    beforeMount () {
      const self = this
      var flag = false
      $.get(api + 'order').then(function (response) {
        if(response.status === 200) {
          response.data.order.forEach((item) => {
            if (item.orderID === self.$route.params.orderID && item.status === 1) {
              self.comment.busID = item.bus.busID
              flag = true
            }
          })
          if(!flag) {
            self.$router.push('/404')
          }
        }
      })
    }
  }
</script>

<style scoped>

</style>
