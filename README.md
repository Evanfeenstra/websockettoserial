# websockettoserial

```
go build -ldflags=-s
```

receives webosocket on port 8000/:something and send serial

`./websockettoserial` to run

`chmod 770 websockettoserial` if you get a permission error


### javascript client

```
var ws = new WebSocket('ws://localhost:8000/cool')

ws.send(msg)
```

### Arduino code to receive a single integer

```
void setup() {
  Serial.begin(9600);
  pinMode(13, OUTPUT);
}

elapsedMillis timer;
unsigned int integerValue=0;
char incomingByte;

void loop() {
  
  if (Serial.available() > 0) {
    digitalWrite(13, HIGH);
    integerValue = 0;
    while(1) {
      incomingByte = Serial.read();
      if (incomingByte == '\n') break;
      if (incomingByte == -1) continue;
      integerValue *= 10;
      // convert ASCII to integer, add, and shift left 1 decimal place
      integerValue = ((incomingByte - 48) + integerValue);
    }
    Serial.println(integerValue);
    digitalWrite(13, LOW);
  }
}
```