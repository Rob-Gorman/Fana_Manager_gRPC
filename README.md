# Fána Manager - gRPC

Iteration of [Fána](https://fana-io.github.io/) Manager serving a gRPC API in lieu of the previous RESTful API:

- `pb` package contains the protobuf files and request/response message types
- Static server wrapped in HTTP/2 handler to accommodate that API (HTTP/1.1 requests unaffected)
- Response methods on GORM models changed accordingly to return appropriate pb message types

---

The Fána Manager handles the following responsibilities:

- Serving the static content for the developer dashboard
- Serving the gRPC API to the developer dashboard
- Managing and executing operations on the PostgreSQL database
- Publishing appropriate update messages from those data operations to Redis pub/sub

---

In lieu of `Makefiles`, the following deployment scripts can be run via Bash:

Protofiles can be compiled by way of:

```bash
bash protobuf.sh
```

Deployment in local Docker network:

```bash
bash deploylocal.sh
```
