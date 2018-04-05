/*
* @Author: jensen
* @Date:   2018-03-29 17:22:13
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-02 17:39:56
*/
import vue from 'vue'
import axios from 'axios'


export default {
	install(Vue) {
		Vue.$axios = Vue.prototype.$axios = axios
		Vue.$e = Vue.prototype.$e = new vue()
	}
}
