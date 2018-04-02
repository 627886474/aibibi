<template>
	<div>
		<p>{{page}}</p>
		<p>{{user}}</p>
	</div>
</template>

<script>
	import { mapState } from "vuex"
	export default {
		data(){
			return{
				page:"user center"
			}
		},
		computed: {
			...mapState({user: 'user'})
		},
		methods: {
			async getTree(){
				let res = await this.$api.test.getTree({
					treePath: '/'
				})
				console.log(res)
			}
		},
		created(){
			this.$store.commit("updateUser",'jensen')
			this.getTree()
		}
	}
</script>
<style lang='scss'>

</style>
