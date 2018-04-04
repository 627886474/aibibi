/*
* @Author: jensen
* @Date:   2018-03-30 10:09:51
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-04 11:15:27
*/


import axios from 'axios'
import Promise from 'es6-promise'
import { CONST_HEADER }  from './common'


// get env  and set env ip
const isPro = process.env.NODE_ENV == 'production'
const ip = {
  proUrl: '//47.104.195.175:8000/api/v1',
  devUrl: '//47.104.195.175:8000/api/v1',
}

const baseUrl = isPro ? ip.proUrl : ip.devUrl

// config base
const http = axios.create({
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

export default http
