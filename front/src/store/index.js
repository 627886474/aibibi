/*
* @Author: jensen
* @Date:   2018-03-28 14:32:33
* @Last Modified by:   yinhuajin
* @Last Modified time: 2018-03-29 10:19:42
*/

import vue from "vue"
import vuex from "vuex"
import state from	"./state"
import * as mutations from "./mutation"

vue.use(vuex)

const store = new vuex.Store({
	state,
	mutations
})


export default store

 
