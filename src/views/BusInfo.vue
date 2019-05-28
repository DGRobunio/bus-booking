<template>
  <div class="bus-info">
    <Navigator :user="user" />
    <div class="container">
      <div class="row">
        <div class="col">
          <h2 class="col-sm-8">{{bus.license}}班车详情</h2>
          <dl class="row">
            <dt class="col-sm-3">出发时间</dt>
            <dd class="col-sm-3">{{bus.beginAt}}</dd>
            <dt class="col-sm-3">抵达时间</dt>
            <dd class="col-sm-3">{{bus.endAt}}</dd>
            <dt class="col-sm-3">出发地</dt>
            <dd class="col-sm-3">{{bus.departure}}</dd>
            <dt class="col-sm-3">目的地</dt>
            <dd class="col-sm-3">{{bus.destination}}</dd>
            <dt class="col-sm-3">票务价格</dt>
            <dd class="col-sm-3">{{bus.price}} 元</dd>
            <dt class="col-sm-3">每周运营时间</dt>
            <dd class="col-sm-3">{{weeklyString}}</dd>
            <dt class="col-sm-2">剩余座位</dt>
            <dd class="col-sm-2">{{bus.emptySeats}} / {{bus.totalSeats}} </dd>
            <dt class="col-sm-3">信息介绍</dt>
            <dd class="col-sm-5">{{bus.info}}</dd>
          </dl>
          <hr/>
          <router-link v-if="user.userID === ''" :to="'/login'" class="btn btn-primary">立即订票</router-link>
          <router-link v-else-if="user.userID !== '' && bus.status === 1" :to="'/purchase/' + busID" :user="user" class="btn btn-primary">立即订票</router-link>
          <button v-else class="btn btn-primary disabled" aria-disabled="true">立即订票</button>
          <button v-if="!favor && user.userID !==''" @click="favoriteChange" class="btn btn-primary"> 收藏</button>
          <button v-else-if="favor && user.userID !==''" @click="favoriteChange" class="btn btn-secondary"> 取消收藏</button>
          <router-link v-if="user.isAdmin" :to="'/adminupdatebus/' + busID" :user="user" class="btn btn-warning">更新信息</router-link>
          <hr/>
          <h5>乘客评价(共{{Object.keys(comment).length}}条)</h5>
          <div class="alert alert-light" v-if="!commentFlag">暂无评价。</div>
          <div v-else>
            <div class="card" v-for="oneComment in comment" :key="oneComment.commentID">
              <div class="card-header">
                <div class="float-left">
                  <p>日期：{{oneComment.commentAt}}</p>
                </div>
                <div class="float-right">
                  <p>评分: {{oneComment.stars}}/5</p>
                </div>
              </div>
              <div class="card-body">
                <p class="card-text">{{oneComment.content}}</p>
                <div class="card" v-if="oneComment.isReplied" >
                  <p>管理员回复：{{oneComment.contentReplied}}</p>
                </div>
              </div>
            </div>
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
    name: 'BusInfo',
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
        busID: this.$route.params.busID,
        favor: false,
        commentFlag: false,
        weeklyString: null,
        bus: {
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
        },
        comment: {
          oneComment: {
            commentID: null,
            content: null,
            commentAt: null,
            stars: null,
            isReplied: null,
            contentReplied: null
          }
        },
        postData: {
          busID: this.$route.params.busID
        }
      }
    },
    methods: {
      favoriteChange () {
        const self = this
        if (self.favor) {
          $.delete(api + 'favorite?busID=' + self.busID).then(function (response) {
            if (response.status === 200)
            {
              self.favor = false
            }
          })
        } else {
          $.post(api + 'favorite', self.postData).then(function (response) {
            if (response.status === 200)
            {
              self.favor = true
            }
          })
        }
      },
      stringfyWeekly (weekly) {
        const self = this
        const week = ["周日","周一","周二","周三","周四","周五","周六"]
        self.weeklyString = ''
        var tmp = weekly
        for (var i=0; i<7; i++) {
          if(tmp >= 2**(6-i)) {
            self.weeklyString += week[i]
            tmp -= 2**(6-i)
            if (tmp) {
              self.weeklyString += ','
            }
          }
        }
      }
    },
    beforeMount () {
      const self = this
      $.get(api + 'bus/' + self.busID).then(function (response) {
        if (response.status === 200) {
          self.bus = response.data.bus
          self.stringfyWeekly(self.bus.weekly)
        }
      })
      $.get(api + 'comment/' + self.busID).then(function (response) {
        if (response.status === 200) {
          self.comment = response.data.comment
          if (Object.keys(self.comment).length !== 0) {
            self.commentFlag = true
          }
        }
      })
      $.get(api + 'favorite').then(function (response) {
        if (response.status === 200) {
          response.data.bus.forEach((item) => {
            if (item.busID === self.busID) {
              self.favor = true
            }
          })
        }
      })
    }
  }
</script>

<style scoped>

</style>
