/*
* @Author: jensen
* @Date:   2018-03-30 11:53:42
* @Last Modified by:   jensen
* @Last Modified time: 2018-03-30 11:56:53
*/

import axios  from 'axios'

export default {
	getTree(args = {
			// treePath: '',     // [string] 目录
	}){
		return axios.post('http://127.0.0.1:8080/api/rendering/file/operate/getUserDirFile', args)
	}
}