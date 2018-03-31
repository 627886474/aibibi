/*
* @Author: jensen
* @Date:   2018-03-29 15:54:35
* @Last Modified by:   jensen
* @Last Modified time: 2018-03-29 17:19:36
*/
import Vue from 'vue'

Vue.filter('toMoney' , val =>{
	if(typeof val != Number) val = Number(val)
	val = val.toFixed(2).toString().split('.')
	let reg = /(\d)(?=(?:\d{3})+$)/g;
  return  (val[0] || 0).toString().replace(reg, '$1,') + '.'+val[1];
})