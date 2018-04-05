<template lang='pug'>
  div._warpper-login
    header._login-header logo
    div._login-body
      div._form-wrapper
        section._form-header
          label 用户登录
        section._form-body
          div._input-item
            label 账号
            el-input
          div._input-item
            label 密码
            el-input
        section._form-btn
          el-button(type='primary') 登录
        section._form-forget
          span 忘记密码？
    footer._login-footer
      div
        label copy right@2018
</template>
<script type="text/javascript">
  export default {
    data() {
      return {

      }
    }
  }
</script>
<style lang='scss'>
._warpper-login {
  height: 100%;
  background-color: #FFF;
  display: flex;
  flex-direction: column;
  ._login-header, ._login-footer {
    text-align: center;
    height: 80px;
    line-height: 80px;
  }
  ._login-body {
    flex: 1;
    background: linear-gradient(to right, #3399ff, #2AABFC);
    display: flex;
    justify-content: flex-end;
    ._form-wrapper {
      width: 320px;
      border-radius: 8px;
      box-shadow: 1px 1px 0 1px rgba(0,0,0,.3);
      background-color: #FFF;
      padding:  20px 30px;
      margin: 120px 60px auto;
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
          }
          .el-input{
            flex: 1;
            margin-left: 8px;
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
        text-align: right;
        font-size: 14px;
        color: #999;
      }
    }
  }
}


</style>
