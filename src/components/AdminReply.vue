<template>
  <form @submit.prevent="submit" class="reply-form">
    <div class="form-row">
      <label for="reply" class="col-sm-1">回复</label>
      <textarea v-model="Comment.contentReplied" id="reply" class="form-control col-sm-11" type="text" required>
      </textarea>
      <button class="btn btn-primary btn-block col-sm-2" type="submit">确认回复</button>
      <div v-if="tip.status === 'success'" class="alert alert-success col-sm-10">
        {{tip.message}}
      </div>
      <div v-if="tip.status === 'warn'" class="alert alert-warning col-sm-10">
        {{tip.message}}
      </div>
      <div v-if="tip.status === 'fail'" class="alert alert-danger col-sm-10">
        {{tip.message}}
      </div>
    </div>
  </form>
</template>

<script>
export default {
  name: 'submit',
  props: {
    commentID: null
  },
  data() {
    return {
      Comment: {
        commentID: null,
        contentReplied: null
      },
      tip: {
        status: null,
        message: null
      }
    }
  },
  methods: {
    submit (){
      const self = this
      self.Comment.commentID = self.commentID
      $.put(api + 'comment/' + self.commentID, self.Comment).then(function (response) {
          if (response.status === 200) {
            self.tip.status = 'success'
            self.tip.message = '回复成功！两秒内刷新页面'
            setTimeout(() => {
              self.tip.status = null
              self.tip.message = null
              self.$emit('refresh')
            }, 2000)
          }
        })
    }
  }

}
</script>

<style>
  .reply-form {
    width: 100%;
    /*max-width: 480px;*/
    padding: 15px;
    margin: auto;
  }

  .reply-form .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 10px;
  }

  .reply-form .form-control:focus {
    z-index: 2;
  }

  .form-control textarea {
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
  }

  .reply-form .alert {
    margin-bottom: 0;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
  }
</style>
