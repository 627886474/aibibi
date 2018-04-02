/*
* @Author: jensen
* @Date:   2018-03-30 10:09:51
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-02 10:32:11
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
	baseURL:baseUrl,
	timeout: isPro? 3*1000 : 10*1000,
	// headers:''
})

// cross domain request  should carry cookie
// axios.default.withCredentials = true

// intercepte request
axios
	.interceptors
	.request
	.use( config => {
		config.headers['Content-Type'] = 'application/json;charset=UTF-8'
		Object.assign(config.headers, CONST_HEADER())
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
    const res = response.data
    return res
	}, error => {
		return Promise.inject(error)
	})
