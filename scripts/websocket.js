// scripts/websocket.ts
var ws = new WebSocket("ws://localhost:1323/ws");
ws.onopen = function() {
  console.log("Connected");
};
ws.onmessage = function(evt) {
  const out = document.getElementById("#count");
  if (out) {
    out.innerHTML = evt.data;
  }
};
ws.onclose = function(ev) {
  console.log("closing", { ev });
};
ws.onerror = (ev) => console.log("error", { ev });
var increment = () => {
  ws.send("increment");
};
window.increment = increment;
