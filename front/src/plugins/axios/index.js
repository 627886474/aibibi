/*
* @Author: jensen
* @Date:   2018-03-29 17:22:13
* @Last Modified by:   jensen
* @Last Modified time: 2018-03-31 02:05:40
*/
import axios from 'axios'


export default {
	install(Vue) {
		Vue.prototype.$axios = axios
	}
}
