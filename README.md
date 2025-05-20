# grpc-ping-pong

A demonstration of gRPC for inter-service communication, implementing a simple ping-pong protocol.

# Protocol Rules
Send any message starting with "ping" → returns "pong" response
Other messages → returns error

# api
POST /pingpong
```
{
  "message": "ping hello world"
}
```

Success Response:
```
{
  "message": "pong hello world"
}
```

Error Response:
```
{
  "error": "string"
}
```