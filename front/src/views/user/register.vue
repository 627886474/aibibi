<template>
  <div>
    <button @click.stop='toReq'>click to requset</button>
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
    },
    queueEvents() {
      let i = 0,
        status = '',
        events = [],
        resolves = []
      let loop = async() => {
        this.count++
        if (status === 'pending') return
        status = 'pending'
        for (let n = 0; n < events.length; n++) {
          // resolve(obj)
          resolves[n](await events[n].fun(events[n].req))
        }
        console.log(`done${i}`)
        status = ''
        i = 0
        events.length = 0
      }
      return (fun, req) => {
        events.push({ fun, req })
        return new Promise(resolve => {
          resolves[i++] = resolve
          loop()
        })
      }
    },
    toReq() {
      const req = (name = 'tj') => {
        return new Request(`https://api.github.com/users/${name}/repos`, {
          mode: 'cors',
          method: 'GET',
          headers: new Headers({ 'Content-Type': 'application/json' })
        })
      }
      const f = (args) => { window.fetch(args) }
      let arr = [
        this.queueEvents()(f, req('tj'))
          .then(res => console.log(res)),
        this.queueEvents()(f, req('substack'))
          .then(res => console.log(res)),
        this.queueEvents()(f, req('yyx990803'))
          .then(res => console.log(res))
      ]
      console.log('continue.....')
    }
  },
  // hook
  async created() {
    // await this.register()
    // console.log('确认是否异步')
  }
}

</script>
