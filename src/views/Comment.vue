<template>
  <div class="passage_comment">
    <Navigator :user="user" />
    <div v-if="user.isAdmin === false" class="container">
      <div class="row">
        <div class="col">
          <form @submit.prevent="submit" class="comment-form mb-4">
            <div class="input-group input-group-lg">
              <div class="input-group-prepend">
                <span class="input-group-text" >评论：</span>
              </div>
              <input v-model="addcomment.content" id="content" class="form-control" placeholder="评论：" required>
            </div>
            <div class="input-group input-group-lg">
              <div class="input-group-prepend">
                <span class="input-group-text" >评分：</span>
              </div>
              <div>
                <label v-for="(item,index) in radioData" :key="index">
                  <input @click="getRadioVal" type="radio" name="type" :value="item.value" required>{{item.value}}
                </label>
              </div>
            </div>
            <button class="btn btn-primary" type="submit">提交</button>
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
        startemp: null,
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
        addcomment: {
          busID: this.$route.params.busID,
          comment: null,
          stars: null
        }
      }
    },
    methods: {
      getRadioVal(index){
        const self = this
        self.startemp = index.target.value;
        console.log(self.startemp)
      },
      submit(){
        const self = this
        self.addcomment.stars = String.valueOf(self.startemp) + "星"
        $.post(api + 'comment', self.addcomment).then(function (response) {
          if (response.status == 200){
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
    }
  }
</script>

<style scoped>

</style>
