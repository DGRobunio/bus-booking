<template>
  <div class="main-info">
    <Navigator :user="user" />
    <div class="container">
      <div class="row">
        <div class="col">
          <div class="alert alert-danger" v-if="Object.keys(codes).length === 0">当前暂时没有充值码。</div>
          <div v-else>
            <table class="table">
              <thead>
                <tr>
                  <th scope="col">充值码</th>
                  <th scope="col">使用者ID</th>
                  <th scope="col">充值码状态</th>
                  <th scope="col">下发充值码</th>
                </tr>
              </thead>
              <tbody>
                  <tr v-for="oneCode in codes" :key="oneCode.codeID">
                    <th scope="row">{{oneCode.codecode}}</th>
                    <td>{{oneCode.userID}}</td>
                    <td v-if="oneCode.status === 0">未使用</td>
                    <td v-else-if="oneCode.status === 1">已使用</td>
                    <td v-else-if="oneCode.status === -1">未发放</td>
                    <td v-else>未知</td>
                    <td>
                      <button v-if="oneCode.status === -1" class="btn btn-primary" @click="updateCode(oneCode.codeID)">下发充值码</button>
                      <button v-else class="btn btn-primary disabled" aria-disabled="true">已下发</button>
                    </td>
                    <!-- <button class="btn btn-lg btn-primary btn-block " type="updateCode">下发充值码</button> -->
                  </tr>
              </tbody>
            </table>
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
    name: 'CodesList',
    components: {
      Navigator,
      Foot
    },
    props: {
      user: null,
    },
    data(){
      return{
        codes: {
          oneCode: {
            codeID: null,
            status :null,
            codecode: null,
            userID: null,
            userAT: null
          }
        },
        tip: {
          type: null,
          message: null
        }
      }
    },
    methods: {
      updateCode(codeID) {
        const self = this
        let putData = {codeID: codeID}
        console.log(putData)
        $.put(api + 'code/' + codeID, putData).then(function (response) {
          if (response.status === 200) {
            $.get(api + 'code').then(function (response) {
              if (response.status === 200) {
                self.codes = response.data.codes
              }
            })
          }
        })
      }
    },
    beforeMount () {
      const self = this
      if (self.user.userID === '' || !self.user.isAdmin) {
        self.$router.push('/404')
      } else {
        $.get(api + 'code').then(function (response) {
          if (response.status === 200) {
            self.codes = response.data.codes
          }
        })
      }
    }
  }
</script>

<style scoped>

</style>
