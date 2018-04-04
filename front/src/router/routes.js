/*
* @Author: jensen
* @Date:   2018-03-28 14:46:22
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-02 18:19:05
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

  // redirect
  {
    path: '*',
    redirect: '/register'
  }
]

export default routes
