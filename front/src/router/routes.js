/*
* @Author: jensen
* @Date:   2018-03-28 14:46:22
* @Last Modified by:   jensen
* @Last Modified time: 2018-03-31 18:49:00
*/

const routes = [
	{
		path:"/",
		name:"test",
		component: ()=>import("@/views/test")
	},
	{
		path:"/brother",
		name:"brother",
		component: ()=>import("@/views/test/brother")
	},

  // 用户模块
	{
		path:"/user",
		name:"user",
		components: require("@/views/user")
	},
  {
    path:"/register",
    name:"register",
    components: require("@/views/user/register")
  },

  // redirect
  {
    path: '*',
    redirect: '/'
  }
]

export default routes
