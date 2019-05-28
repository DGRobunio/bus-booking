<template>
  <div class="bus-info">
    <Navigator :user="user" />
    <div class="container">
      <div class="row">
        <div class="col">
          <table class="table table-bordered table-hover">
            <thead class="thead-light">
            <tr>
              <th v-if="!updateFlag" colspan="2">添加线路</th>
              <th v-if="updateFlag" colspan="2">更新线路</th>
            </tr>
            </thead>
          </table>
          <form @submit.prevent="updateBus" class="col">
            <div class="form-group row">
              <label for="license" class="col-sm-3 col-form-label" >班车牌号</label>
              <div class="col-sm-9">
                <input v-model="bus.license" type="text" class="form-control" id="license" required>
              </div>
            </div>
            <div class="form-group row">
              <label for="total-seats" class="col-sm-3 col-form-label" >总座位数</label>
              <div class="col-sm-9">
                <input v-model="bus.totalSeats" type="number" min="0" class="form-control" id="total-seats" required>
              </div>
            </div>
            <div class="form-group row">
              <label for="departure" class="col-sm-3 col-form-label" >出发地</label>
              <div class="col-sm-9">
                <input v-model="bus.departure" type="text" class="form-control" id="departure" required>
              </div>
            </div>
            <div class="form-group row">
              <label for="destination" class="col-sm-3 col-form-label" >目的地</label>
              <div class="col-sm-9">
                <input v-model="bus.destination" type="text" class="form-control" id="destination" required>
              </div>
            </div>
            <div class="form-group row">
              <label for="beginAt" class="col-sm-3 col-form-label" >出发时间</label>
              <div class="col-sm-9">
                <date-picker v-model="bus.beginAt" type="text" :config="options" class="form-control" id="beginAt" required> </date-picker>
              </div>
            </div>
            <div class="form-group row">
              <label for="endAt" class="col-sm-3 col-form-label" >抵达时间</label>
              <div class="col-sm-9">
                <date-picker v-model="bus.endAt" type="text" :config="options" class="form-control" id="endAt" required> </date-picker>
              </div>
            </div>
            <div class="form-group row">
              <label for="price" class="col-sm-3 col-form-label" >票价</label>
              <div class="col-sm-9">
                <div class="input-group">
                  <input v-model="bus.price" type="number" min="0" class="form-control" id="price" required>
                  <div class="input-group-append">
                    <div class="input-group-text">元</div>
                  </div>
                </div>
              </div>
            </div>
            <div class="form-group row">
              <label  class="col-sm-3 col-form-label" >每周运营时间</label>
              <div class="col-sm-9 m-auto">
                <div class="form-check form-check-inline">
                  <input class="form-check-input" type="checkbox" id="sun" @change="weeklyDataUpdate" v-model="weeklyData" value=64>
                  <label class="form-check-label" for="sun">周日</label>
                </div>
                <div class="form-check form-check-inline">
                  <input class="form-check-input" type="checkbox" id="mon" @change="weeklyDataUpdate" v-model="weeklyData" value=32>
                  <label class="form-check-label" for="mon">周一</label>
                </div>
                <div class="form-check form-check-inline">
                  <input class="form-check-input" type="checkbox" id="tue" @change="weeklyDataUpdate" v-model="weeklyData" value=16>
                  <label class="form-check-label" for="tue">周二</label>
                </div>
                <div class="form-check form-check-inline">
                  <input class="form-check-input" type="checkbox" id="wed" @change="weeklyDataUpdate" v-model="weeklyData" value=8>
                  <label class="form-check-label" for="wed">周三</label>
                </div>
                <div class="form-check form-check-inline">
                  <input class="form-check-input" type="checkbox" id="thu" @change="weeklyDataUpdate" v-model="weeklyData" value=4>
                  <label class="form-check-label" for="thu">周四</label>
                </div>
                <div class="form-check form-check-inline">
                  <input class="form-check-input" type="checkbox" id="fri" @change="weeklyDataUpdate" v-model="weeklyData" value=2>
                  <label class="form-check-label" for="fri">周五</label>
                </div>
                <div class="form-check form-check-inline">
                  <input class="form-check-input" type="checkbox" id="sat" @change="weeklyDataUpdate" v-model="weeklyData" value=1>
                  <label class="form-check-label" for="sat">周六</label>
                </div>
              </div>
            </div>
            <div class="form-group row">
              <label for="status" class="col-sm-3 col-form-label" >班车状态</label>
              <div class="col-sm-9">
                <select v-model="bus.status" class="form-control" id="status" required>
                  <option value=1  >已运营</option>
                  <option value=0  >即将运营</option>
                  <option value=-1 >已废弃</option>
                </select>
              </div>
            </div>
            <div class="form-group row">
              <label for="info" class="col-sm-3 col-form-label" >班车介绍</label>
              <div class="col-sm-9">
                <textarea v-model="bus.info" type="text" class="form-control" id="info" required></textarea>
              </div>
            </div>
            <div v-if="tip.type === 0" class="alert alert-danger form-control"> {{tip.message}} </div>
            <div v-if="tip.type === 1" class="alert alert-success form-control"> {{tip.message}} </div>
            <button class="btn btn-primary col-sm-3" type="submit">确认提交</button>
          </form>
        </div>
      </div>
    </div>
    <Foot />
  </div>
</template>

<script>
import Navigator from '../components/Navigator'
import Foot from '../components/Foot'
import datePicker from 'vue-bootstrap-datetimepicker'
import '@fortawesome/fontawesome-free/css/all.css'
import 'pc-bootstrap4-datetimepicker/build/css/bootstrap-datetimepicker.css'


  export default {
    name: 'AdminUpdateBus',
    components: {
      Navigator,
      datePicker,
      Foot
    },
    props: {
      user: null,
    },
    data(){
      return{
        updateFlag:null,
        bus: {
          license: null,
          totalSeats: null,
          departure: null,
          destination: null,
          beginAt: null,
          endAt: null,
          price: null,
          info: null,
          weekly: null,
          status: 1
        },
        tip: {
          type: null,
          message: null
        },
        weeklyData: [],
        options: {
          format: 'HH:mm:ss',
          useCurrent: false,
          locale: 'zh-cn',
          tooltips: {
            selectTime: ''
          }
        }
      }
    },
    methods: {
      updateBus () {
        const self = this
        if(self.bus.weekly === 0) {
          self.tip.status = 0
          self.tip.message = '请指定班车每周运营时间！'
          setTimeout(() => {
            self.tip.status = null
            self.tip.message = null
          }, 2000)
        } else {
          if(self.updateFlag) {
            $.put(api + 'bus/' + self.bus.busID, self.bus).then(function (response) {
              if (response.status === 200)
              {
                self.tip.status = 1
                self.tip.message = '修改成功!'
                setTimeout(() => {
                  self.$router.push('/')
                }, 1000)
              }
            }).catch(function (error) {
              console.log(error)
              self.tip.status = 0
              self.tip.message = '修改失败！请稍后重试！'
              setTimeout(() => {
                self.tip.status = null
                self.tip.message = null
              }, 2000)
            })
          } else {
            $.post(api + 'bus', self.bus).then(function (response) {
              if (response.status === 200)
              {
                self.tip.status = 1
                self.tip.message = '添加成功!'
                setTimeout(() => {
                  self.$router.push('/')
                }, 1000)
              }
            }).catch(function (error) {
              console.log(error)
              self.tip.status = 0
              self.tip.message = '添加失败！请稍后重试！'
              setTimeout(() => {
                self.tip.status = null
                self.tip.message = null
              }, 2000)
            })
          }
        }
      },
      weeklyDataUpdate() {
        const self = this
        self.bus.weekly = 0
        self.weeklyData.forEach((item) => {
          self.bus.weekly += Number(item)
        })
        console.log(self.bus.weekly)
      }
    },
    beforeMount () {
      const self = this
      if (self.user.userID === '' || !self.user.isAdmin) {
        self.$router.push('/404')
      } else {
        if(self.$route.params.busID) {
          self.updateFlag = true
          $.get(api + 'bus/' + self.$route.params.busID).then(function (response) {
            if(response.status === 200) {
              self.bus = response.data.bus
            }
          })
        } else {
          self.updateFlag = false
        }
      }
      self.weeklyDataUpdate()
    },
    created: function() {
      $.extend(true, $.fn.datetimepicker.defaults, {
        icons: {
          time: 'far fa-clock',
          date: 'far fa-calendar',
          up: 'fas fa-arrow-up',
          down: 'fas fa-arrow-down',
          previous: 'fas fa-chevron-left',
          next: 'fas fa-chevron-right',
          today: 'fas fa-calendar-check',
          clear: 'far fa-trash-alt',
          close: 'far fa-times-circle'
        }
      })
    }
  }

</script>

<style scoped>

</style>
