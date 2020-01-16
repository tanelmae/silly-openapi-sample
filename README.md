## Silly server with Go and OpenAPI spec


Uses OpenAPI client and server code generator [**oapi-codegen**](https://github.com/deepmap/oapi-codegen)
The generator supports both [Chi](https://github.com/go-chi/chi) and [Echo](https://github.com/labstack/echo) but just Echo is used in this sample.

Generate server, types and spec from OpenAPI yaml:
```
./codegen.sh
```
Generated code lives under `pkg/gen`

Under `internal/servce` there is implementation for the generated ServerInterface interface and that is where the "business logic" is defined.

Glue to tie the genegrated code and business logic is in `main.go`. This is not just to be able to write bolierplate code. It allows setting up any [Echo middleware](https://echo.labstack.com/middleware), adding support for extra content types (see the example) or other needed customization.

#### Run the service:
```
go run main.go
```

##### Example requsts
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
curl --request POST --header "Content-Type: application/json" --data '{"name":"Yoda"}' http://localhost:8080/upload
```

Invalid request
```
curl http://localhost:8080/hello?ss=ee
```

Upload image
```
curl --request POST --header 'Content-Type: image/jpeg' --data-binary @img/moon.jpg http://localhost:8080/image
curl --request POST --header 'Content-Type: image/png' --data-binary @img/space.png http://localhost:8080/image
```