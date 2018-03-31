/*
* @Author: jensen
* @Date:   2018-03-30 10:09:51
* @Last Modified by:   jensen
* @Last Modified time: 2018-03-31 02:14:26
*/

import './config'
import user from './user'
import test from './test'

export default {
	install(Vue) {
		Vue.prototype.$api = {
			user,
			test
		}
	}
}
