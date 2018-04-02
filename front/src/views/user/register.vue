<template>
  <div>
    <p>to test oncopy</p>
    <el-input v-model='count' />
  </div>
</template>
<script type="text/javascript">
export default {
  data() {
    return {
      count:0
    }
  },
  methods: {

    async register() {
      let res = await this.$api.user.register({
        name: '科', // [string]  用户名
        pass_word: '123456', // [string]  密码
        mobile: '13712340004', // [string]  手机号
      })
      console.log('注册接口返回:' + res)
    }
    
  },
  // hook
  async created() {
    // await this.register()
    // console.log('确认是否异步')
    
    this.$message('asfssf')
  }
}

</script>
