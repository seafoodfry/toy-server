# toy-server

An example server that repeats what you say.

---

Welcome to the toy server!

This is an HTTP server that accepts POST requests with a JSON body to the
`/api/echo` endpoint.
The response will be the body of the request sent with the field `echoed:
true`.
For example, the request
```
curl -H "Content-Type: application/json" -X POST -d '{"user": "abc", "data": "abc-123"}' http:localhost:8080/api/echo
```

Will result in the response
```json
{
  "user": "abc",
  "data": "abc-123",
  "echoed": true
}
```

If the request's body includes the top-level field `echoed` and this
has the value `true`, then the server will respond with the following response
```
HTTP/1.1 400 Bad Request

{"error": "request already had 'echoed: true'"}
```

However, the server accepts both POST and PUT requests.


## Development

We have some cool scripts in the [./hack](./hack) directory to improve the
quality of your code.
To run them do:
```
make validate
```


