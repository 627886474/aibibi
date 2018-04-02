<template>
  <div class="_view-wrapper">
    <header class="_view-header">
      <slot name='view-heder'></slot>
    </header>
    <content class="_view-content">
      <slot name='view-content'></slot>
    </content>
    <footer class="_view-footer">
      <slot name='view-footer'></slot>
    </footer>
  </div>
</template>

<script type="text/javascript">
  export default {
    name: 'View'
  }
</script>

<style lang='scss'>
  ._view-wrapper{
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
    ._view-header{
      height: 80px;
    }
    ._view-content{
      flex: 1;
    }
    ._view-footer{
      height: 80px;
    }
  }
</style>
