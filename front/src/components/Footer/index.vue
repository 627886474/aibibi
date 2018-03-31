<template>
	<div>
		<p>footer</p>
		<span>{{onmsg}}</span>
	</div>
</template>
<script>
	export default{
		data(){
			return {
				onmsg:''
			}
		},
		mounted(){
			this.$e.$on("broad-brother", val=>{
				this.onmsg = val
			})
		}
	}
</script>