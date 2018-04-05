/*
* @Author: jensen
* @Date:   2018-03-30 11:26:00
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-05 15:43:17
*/

import axios from 'axios'


export default {
	register( args = {
		// name: '科',                        // [string]  用户名
    // pass_word: '123456',               // [string]  密码
    // mobile: '13712340004',             // [string]  手机号
	}){
		return axios.post( '/user/register', args)
	},
	login( args = {
		// sigin_input: '科',                 // [string]  用户名
    // pass_word: '123456',               // [string]  密码
	}){
		return axios.post( '/user/login', args)
	}
}
