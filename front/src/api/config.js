/*
* @Author: jensen
* @Date:   2018-03-30 10:09:51
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-05 13:37:38
*/


import axios from 'axios'
import Promise from 'es6-promise'
import { CONST_HEADER }  from './default'


// get env  and set env ip
const isPro = process.env.NODE_ENV == 'production'
const ip = {
  proUrl: '//47.104.195.175:8000/api/v1',
  devUrl: '//47.104.195.175:8000/api/v1',
}

// config defaults    --add property to  defaults
axios.defaults.baseURL = isPro? ip.proUrl : ip.devUrl
axios.defaults.timeout = isPro? 3*1000 : 10*1000
// axios.defaults.baseURL = 'http://172.16.4.7:8980'
// cross domain request  should carry cookie
// axios.defaults.withCredentials = true

// intercepte request
axios
	.interceptors
	.request
	.use( config => {
		config.headers['Content-Type'] = 'application/json;charset=UTF-8'
		Object.assign(config.headers, CONST_HEADER())
		return Promise.resolve(config)
	}, error => {
		return Promise.inject(error)
	})

// intercepte response
axios
	.interceptors
	.response
	.use( response => {
    let res = response.data
    return Promise.resolve(res)
	}, error => {
		return Promise.inject(error)
	})
