<template lang='pug'>
  .hello
    h1 {{msg}}
    - const buttons = []
    mixin menu-button(buttons)
      for item,index in buttons
        button=item+(index+1)
    + menu-button(buttons)
    span receive:{{parent}}
    Header(title="Header broadcast's title" @broad-parent="onChild")
    Footer 
    input(v-focus='6')
    span {{2222215.1465 | toMoney}}
</template>

<script >
import { Component, Prop, Vue, Watch } from 'vue-property-decorator'
import Header from "@/components/Header"
import Footer from "@/components/Footer"

export default  {
  components:{Header,Footer},
  data(){
    return {
      msg: 'sfsdf',
      parent:''
    }
  },
  methods: {
    onChild(val){
      this.parent = val
    }
  },
  mounted(){
    
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
