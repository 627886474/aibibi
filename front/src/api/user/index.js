/*
* @Author: jensen
* @Date:   2018-03-30 11:26:00
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-04 11:15:07
*/

import http from './../config'


export default {
	register( args = {
		// name: '科',                        // [string]  用户名
    // pass_word: '123456',               // [string]  密码
    // mobile: '13712340004',             // [string]  手机号
	}){
		return http.post( '/user/register', args)
	}
}
