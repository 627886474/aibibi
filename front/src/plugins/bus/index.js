/*
* @Author: jensen
* @Date:   2018-03-30 12:14:09
* @Last Modified by:   jensen
* @Last Modified time: 2018-03-31 02:05:15
*/
import Vue from 'vue'

export default {
	install(Vue) {
		Vue.prototype.$e = new Vue()
	}
}
