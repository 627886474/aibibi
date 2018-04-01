/*
* @Author: jensen
* @Date:   2018-03-31 19:24:54
* @Last Modified by:   jensen
* @Last Modified time: 2018-03-31 19:50:02
*/

import View from './src'

View.install =  function(Vue){
  Vue.component(View.name, View)
}

export default View
