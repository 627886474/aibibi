<template>
	<div>
		<a href="">on page</a>
		<p>{{onmsg}}</p>
	</div>
</template>
<script>
	
	export	default {
		data(){
			return{
				onmsg: ""
			}
		},
		methods:{
			onevent(val){
				console.log(val)
				this.onmsg = val
			}
		},
		mounted(){
		}
	}
</script>