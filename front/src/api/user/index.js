/*
* @Author: jensen
* @Date:   2018-03-30 11:26:00
* @Last Modified by:   jensen
* @Last Modified time: 2018-03-30 11:38:04
*/

import axios from 'axios'


export default {
	/**
	 * @params  
	 */
	register( args = {
		// tel: '13651415201'  // [string]  手机号
	}){
		return axios.post( 'url', args)
	}
}