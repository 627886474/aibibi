<template lang='pug'>
  div._warpper-login._flex._flex-column
    div._login-body._flex._flex-h-end
      div._form-wrapper
        section._form-header
          label 用户登录
        section._form-body
          div._input-item
            label 账 号
            el-input(v-model='account' placeholder='请输入账号/邮箱/手机号')
          div._input-item
            label 密 码
            el-input(v-model='psw' placeholder='请输入密码')
          div._input-item(v-show='false')
            label 验证码
            el-input(v-model='psw' placeholder='请输入验证码')
            img(src='fsdf')
        section._form-btn
          el-button(type='primary' @click.stop='login') 登 录
        section._form-forget
          router-link(to='/register' tag='span') 快速注册
          router-link(to='/forget' tag='span') 忘记密码？
</template>
<script type="text/javascript">
  export default {
    data() {
      return {
        account: '',
        psw: ''
      }
    },
    methods: {
      async login() {
        // input check
        const emailReg = /^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$/
        if(!this.account) return this.$message('请输入账号')
        if(!emailReg.test(this.account)) return this.$message('请输入正确的账号')
        if(!this.psw) return this.$message('请输入密码')

        let res = await this.$api.user.login({
          sigin_input: this.account || '科',                 // [string]  用户名
          pass_word: this.psw || '123456',               // [string]  密码
        })

        // error
        if(res.errNo) return this.$message({
          message:res.msg
        })

        // ok
        this.$router.push('/overview')
      }
    }
  }
</script>
<style lang='scss'>
._warpper-login {
  ._login-header, ._login-footer {
    text-align: center;
    height: 80px;
    line-height: 80px;
  }
  ._login-body {
    flex: 1;
    background: linear-gradient(to right, #3399ff, #2AABFC);
    ._form-wrapper {
      width: 360px;
      border-radius: 8px;
      box-shadow: 1px 1px 0 1px rgba(0,0,0,.3);
      background-color: #FFF;
      padding:  30px 48px;
      margin: auto 10vw;
      text-align: center;
      ._form-header{
        label {
          font-weight: 500;
          &::after {
            width: 100%;
            content: '';
            height: 2px;
            display: inline-block;
            border-radius: 100%;
            background-color: #2AABFC;
          }
        }
      }
      ._form-body{
        padding: 15px 0;
        ._input-item{
          display: flex;
          height: 40px;
          line-height: 40px;
          margin: 20px 0;
          label {
            font-size: 14px;
            color: #666;
            width: 50px;
            text-align: right;
          }
          .el-input{
            flex: 1;
            margin-left: 8px;
          }
          img {
            margin-left: 20px;
            width: 60px;
          }
        }
      }
      ._form-btn{
        .el-button{
          width: 100%;
        }
      }
      ._form-forget{
        margin: 20px 0;
        font-size: 14px;
        color: #999;
        cursor: pointer;
        display: flex;
        justify-content: space-between;
        span{
          &:hover {
            text-decoration: underline;
            color: #2AABFC;
          }
        }

      }
    }
  }
}


</style>
