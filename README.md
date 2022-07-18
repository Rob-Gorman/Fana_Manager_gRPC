# Go Web Services
## (It's Express)
Like `import express from 'express'`, `app = express()`, and `app.use(routes)`
The patterns are really, really similar for Go web services.

I opted for `mux` router package over the std lib one, and to be honest I forget why.

But the `mux.router` is our express package. This is the thing that behaves very similarly to express
(I wrapped it in a custom `Server` struct, but don't let that confuse you. All the web stuff is done via the router)

Just like `app.use(routes)`, we assign behavior to specific paths by defining them in very similar ways.
```go
s.HandleFunc("/flags", s.H.GetAllFlags).Methods("GET")
```
When a GET request comes in, our router will send it to our `GetAllFlags` controller (called a handler here, but it could just had easily been called a controller--word never really struck me right)
(The actual handler has a lot of ORM functionality that is real tough to parse, but that's for another conversation)

So we have our router, our routes, and our controller//handler functions

lastly, we let the router listen:
```go
http.ListenAndServe(":6000", srv)
```