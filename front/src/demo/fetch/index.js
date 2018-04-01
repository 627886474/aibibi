/*
* @Author: jensen
* @Date:   2018-04-01 15:55:21
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-01 16:10:42
*/
function queueEvents(){
  let i=0,
      status = '',
      events = [],
      resolves = []
  let loop = async () => {
    if( status === 'pending') return
    status = 'pending'
    for ( let n=0 ; n<events.length;  n++) {
      resolves[n] (await events[n].fun(events[n].req))
    }
    console.log('done')
    status = ''
    i = 0
    events.length = 0
  }
  return (fun, req) =>{
    events.push({ fun, req })
    return new Promise(resolve =>{
      resolves[i++] = resolve
      loop()
    })
  }
}


const req = (name = 'tj') => {
  return new Request(`https://api.github.com/users/${name}/repos`, {
    mode: 'cors',
    method: 'GET',
    headers: new Headers({'Content-Type': 'application/json'})
  })
}

const h = queueEvents()
const f = (args) => {
  window.fetch(args)
}

let arr = [
    h(f, req('tj')).then(res => console.log(res))
    h(f, req('substack')).then(res => console.log(res))
    h(f, req('yyx990803')).then(res => console.log(res))
]

console.log('continue.....')
