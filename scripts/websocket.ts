const ws = new WebSocket('ws://localhost:1323/ws')

ws.onopen = function () {
  console.log('Connected')
}
ws.onmessage = function (evt) {
  const out = document.getElementById('#count')
  if (out) {
    out.innerHTML = evt.data
  }
}

ws.onclose = function (this, ev) {
  console.log('closing', { ev })
}
ws.onerror = (ev) => console.log('error', { ev })

const increment = () => {
  ws.send('increment')
}

;(window as any).increment = increment
