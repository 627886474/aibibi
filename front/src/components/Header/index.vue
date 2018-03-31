<template>
	<div>
		<p>{{title}}</p>
		<button @click="broadBrother">broadBrother  click</button><br>
		<button @click="broadParent">broadParent  click</button>
	</div>
</template>
<script>
	export default {
		data(){
			return {
				child:''
			}
		},
		props:{
			title:{
				type: String,
				default: "nothing"
			}
		},
		methods :{
			broadBrother(){
				this.$e.$emit("broad-brother", "i am  from   test  broad")
			},
			broadParent(){
				this.$emit("broad-parent", 'tell to  child')
			}
		},
		mounted(){

			
		}
	}
</script>