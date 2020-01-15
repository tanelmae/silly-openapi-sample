## Silly server with Go and OpenAPI spec


Uses OpenAPI client and server code generator [**oapi-codegen**](https://github.com/deepmap/oapi-codegen)
The generator supports both [Chi](https://github.com/go-chi/chi) and [Echo](https://github.com/labstack/echo) but just Echo is used in this sample.

Generate server, types and spec from OpenAPI yaml:
```
./codegen.sh
```
Generated code lives under `pkg/gen`

Under `internal/servce` there is implementation for the generated ServerInterface interface and that is where the "business logic" is defined.

Glue to tie the genegrated code and business logic is in `main.go`. While is seems to just expose opportunity to write bolierplate code it also allows settig up any [Echo middleware](https://echo.labstack.com/middleware) that might be needed.

#### Run the service:
```
go run main.go
```

Get with query parameter:
```
curl http://localhost:8080/hello?name=Bill
```
Get with parameter in path:
```
curl http://localhost:8080/hello/tim
```

Post request:
```
curl --header "Content-Type: application/json" --request POST --data '{"name":"Yoda"}' http://localhost:8080/upload
```

Invalid request
```
curl http://localhost:8080/hello?ss=ee
```