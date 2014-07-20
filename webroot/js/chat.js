var ws = new WebSocket("ws://" + location.host + "/ws")
var area = document.querySelector(".message-area")
var input = document.querySelector(".input-area")
input.addEventListener('keydown', function(e) {
    if (e.which === 13) {
        var message = input.value
        input.value = ''

        ws.send(message)
    }
})
ws.onmessage = function(e) {
    var message = e.data
    var div = document.createElement("div")
    div.textContent = message
    area.appendChild(div)
}
