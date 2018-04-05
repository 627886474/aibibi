<template>
  <div id="app" class="_flex _flex-column">
    <header class="app-header _flex">
      <div class="app-logo">Logo</div>
      <nav class="app-nav _flex-item">
        <ul class="_flex">
          <li class="active"><router-link to='/overview'>首页</router-link></li>
          <li><router-link to='/overview'>社区</router-link></li>
          <li><router-link to='/overview'>合作</router-link></li>
          <li><router-link to='/overview'>关于我们</router-link></li>
        </ul>
      </nav>
    </header>
    <router-view class="_flex-item"/>
    <footer class="app-footer _center">版权所有</footer>
  </div>
</template>

<script>
import View  from './components/View'
export default {
  name: 'App'
}
</script>

<style lang='scss'>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  height: 100vh;
  width: 100vw;
  .app-header {
    height: 80px;
    line-height: 80px;
    background-color: #336666;
    color: #999;
    .app-logo{
      width: 30vw;
      text-align: center;
    }
    .app-nav {
      margin-left: 60px;
      li{
        position: relative;
        a{
          padding: 0 20px;
        }
        &:after {
          position: absolute;
          bottom: 0;
          left: 100%;
          content: '';
          height: 1px;
          display: inline-block;
          background-color: #fff;
          width: 0;
        }
        &:hover {
          background-color: rgba(0,0,0,0.3);
          color: #fff;
          &:after {
            width: 100%;
            left: 0;
            transition: all 0.5s linear;
          }
        }
        &:hover ~ li::after{
           left: 0;
        }
      }
    }
  }
  .app-footer {
    height: 40px;
    line-height: 40px;
    font-size: 13px;
  }
}
</style>
