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
curl -H "Content-Type: application/json" -X POST -d '{"user": "abc", "data": "abc-123"}' http://localhost:8080/api/echo
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

If you want to run the app, you can either compile it and run it or you can run
```
make run
```
This will build a container image for the app, and run it while exposing port
8080 for traffic.

## Metrics

This app is surfaces prometheus-style metrics.
To get the metrics simply send a GET request to the `/metrics` endpoint.

We additionally expose a custom metric for keeping track of the number of
requests to `/api/echo`
```
# HELP echo_requests_total Total http requests to our echo JSON api
# TYPE echo_requests_total counter
echo_requests_total{code="200"} 2
echo_requests_total{code="400"} 3
echo_requests_total{code="500"} 1
```

## Maintainability

This toy-server is to show case what an app can look like.

When working with apps like this one there are 2 SLIs that may be beneficial

1. Percente of non-200 requests to your API: non-200 responses are expected in
   any service and understanding your traffic patterns is very important. If
   non-200 responses spike that could be a signal that something went wrong.
2. Latency: most of the time, faster is better - you don't want your users
   having to wait for long times to get a response back. Measuring the time it
   takes to serve a request will help you provide a good UX (and it may help
   you prevent errors by investigating what type of requests cause increases in
   latency).
