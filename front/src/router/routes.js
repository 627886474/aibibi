/*
* @Author: jensen
* @Date:   2018-03-28 14:46:22
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-05 16:13:52
*/

const routes = [
	{
		path:"/test",
		name:"test",
		component: ()=>import("@/views/test")
	},
	{
		path:"/brother",
		name:"brother",
		component: ()=>import("@/views/test/brother")
	},

  // 首页
  {
    path:"/overview",
    name:"overview",
    component: () => import("@/views/overview")
  },


  // 用户模块
	{
		path:"/user",
		name:"user",
		component: () => import("@/views/user")
	},
  {
    path:"/register",
    name:"register",
    component: () => import("@/views/user/register")
  },
  {
    path:"/login",
    name:"login",
    component: () => import("@/views/user/login")
  },
  {
    path:"/forget",
    name:"forget",
    component: () => import("@/views/user/forget")
  },

  // redirect
  {
    path: '*',
    redirect: '/login'
  }
]

export default routes
