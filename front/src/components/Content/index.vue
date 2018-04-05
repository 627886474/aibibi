<template>
  <div>
    <p>CONTNTE</p>
    <span>{{onmsg}}</span>
  </div>
</template>
<script>
  export default{
    name: 'mycontent',
    data(){
      return {
        onmsg:'fadf第一次发布'
      }
    },
    mounted(){
      this.$message('asfssf')
    }
  }
</script>
