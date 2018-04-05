/*
* @Author: jensen
* @Date:   2018-04-02 14:13:47
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-05 13:38:56
*/

export default class Tools {
	install(Vue) {
		const statics = this.statics()
		Object.keys(statics).map(name => {
			Vue[name] = Vue.prototype[name] = statics[name]
		})
	}
	statics = () => ({
		// use more plugins
		useAll(...plugins) {
			plugins.map(item => {
				if(Array.isArray(item) && item.length>0)
					return item.map( _item => {
						this.use(_item)
					})
				this.use(item)
			})
		},

		// add more components
		components(...components) {
			components.map(item => {
				if(Array.isArray(item) && item.length>0)
					return item.map( _item => {
						if(!item.name) {
							return this.component(_item)
						}
						this.component(_item.name, _item)
					})
				if(!item.name) {
					return this.component(item)
				}
				this.component(item.name, item)
			})
		},

		// ajax queue
		queueEvent() {
			let i = 0
				,	status = ''
				, events = []
				,	resolves = []
			let loop = async () => {
				if(status === 'pending') return
				status = 'pending'
				for(let n = 0; n < events.length; n++ ) {
					resolves[n](await events[n].func(events[n].req))
				}
				status = ''
				i = 0
				// events.splice(0, events.length)
				events.length = 0

 			}
			return (func, req) => {
				events.push({func, req})
				// loop()
				return new Promise((resolve, reject) => {
					resolves[i++] = resolve
					loop()
				})
			}
		}
	})
}
