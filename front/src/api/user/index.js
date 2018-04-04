/*
* @Author: jensen
* @Date:   2018-03-30 11:26:00
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-04 13:29:42
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
		// name: '科',                        // [string]  用户名
    // pass_word: '123456',               // [string]  密码
    // mobile: '13712340004',             // [string]  手机号
	}){
		return axios.post( '/user/register', args)
	}
}
