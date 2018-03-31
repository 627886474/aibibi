/*
* @Author: jensen
* @Date:   2018-03-27 09:53:38
* @Last Modified by:   yinhuajin
* @Last Modified time: 2018-03-27 18:05:08
*/


// 自执行函数方式
for (var i = 0; i < 10; i++) {
	(function(i){
		setTimeout(()=>{
			console.log(i)
		}, i*1000)
	})(i)
}



// es6  块级作用域
for(let i = 0; i< 10; i++){
	setTimeout(()=>{
		console.log(i)
	},i*1000)
}
