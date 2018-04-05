/*
* @Author: jensen
* @Date:   2018-03-30 12:12:31
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-05 13:40:48
*/

import Vue from 'vue'
import vender from  './global-vender'
import Tools from  './gloabl-tool'

Vue.use(vender)
Vue.use(new Tools())

export default {
	vender
}
