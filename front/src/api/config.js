/*
* @Author: jensen
* @Date:   2018-03-30 10:09:51
* @Last Modified by:   jensen
* @Last Modified time: 2018-03-31 18:58:41
*/


import axios from 'axios'
import Promise from 'es6-promise'
import { CONST_HEADER }  from './common'


// get env  and set env ip
const isPro = process.env.NODE_ENV == 'production'
const ip = {
  proUrl: '//192.168.0.106:8000/api/v1',
  devUrl: '//192.168.0.106:8000/api/v1',
}

const baseUrl = isPro ? ip.proUrl : ip.devUrl


// config base
axios.create({
	baseUrl,
	timeout: isPro? 3*1000 : 10*1000
})

// cross domain request  should carry cookie
// axios.default.withCredentials = true

// intercepte request
axios
	.interceptors
	.request
	.use( config => {
		Object.assign(config.headers, {'Content-Type': 'application/json;charset=UTF-8'}, CONST_HEADER())
		return config
	}, error => {
		return Promise.inject(error)
	})

// intercepte response
axios
	.interceptors
	.response
	.use( response => {
    console.log(response)
		return Promise.resolve(response)
	}, error => {
		return Promise.inject(error)
	})
