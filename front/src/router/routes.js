/*
* @Author: jensen
* @Date:   2018-03-28 14:46:22
* @Last Modified by:   yinhuajin
* @Last Modified time: 2018-03-29 11:33:02
*/

const routes = [
	{
		path:"/",
		name:"test",
		component: ()=>import("@/views/test")
	},
	{
		path:"/on",
		name:"on",
		component: ()=>import("@/views/test/on")
	},
	{
		path:"/user",
		name:"user",
		components: require("@/views/user")
	},
	{
		path:"/ensure",
		name:"ensure",
		components: require.ensure([],(require)=>{require("@/views/ensure")})
	}
]

export default routes