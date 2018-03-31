/*
* @Author: jensen
* @Date:   2018-03-29 15:32:11
* @Last Modified by:   jensen
* @Last Modified time: 2018-03-29 18:40:24
*/


import Vue from "vue"

Vue.directive("focus", {
	inserted(el){
		el.focus()
	},
	bind(el, binding){
		el.addEventListener("keyup", ()=>{
			console.log('触发'+binding.value)
		})
	}
})