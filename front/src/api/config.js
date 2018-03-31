/*
* @Author: jensen
* @Date:   2018-03-30 10:09:51
* @Last Modified by:   jensen
* @Last Modified time: 2018-03-31 02:39:48
*/


import axios from 'axios'
import Promise from 'es6-promise'
import { CONST_HEADER }  from './common'


// config base
axios.create({
	// baseUrl: '/',
	// timeout: 3*1000
})

// cross domain request  should carry cookie
// axios.default.withCredentials = true

// intercepte request
axios
	.interceptors
	.request
	.use( config => {
		// config.headers['Content-Type']: 'application/json'
		// Object.assign(config.headers, CONST_HEADER())
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
