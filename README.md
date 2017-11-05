# websockettoserial

go build -ldflags=-s

receives webosocket on port 8000/:something and send serial

`./websockettoserial` to run

`chmod 770 websockettoserial` if you get a permission error


### javascript client

```
var ws = new WebSocket('ws://localhost:8000/cool')

// Received a message
ws.onmessage = function(e){
  var msg = JSON.parse(e.data)
  console.log(msg)
}
```