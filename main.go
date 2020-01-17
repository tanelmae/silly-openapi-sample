package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/tanelmae/silly-openapi-sample/internal/service"
	"github.com/tanelmae/silly-openapi-sample/pkg/gen"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	swagger, err := gen.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	s := service.New()

	// This is how you set up a basic Echo router
	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.Logger())

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	e.Use(middleware.OapiRequestValidator(swagger))

	/*
		oapi-codegen uses kin-openapi for validating the requests:
		https://github.com/getkin/kin-openapi

		By default it supports content types:
		  - "text/plain"
		  - "application/json"
		  - "application/x-www-form-urlencoded"
		  - "multipart/form-data"
		  - "application/octet-stream"

		openapi3filter.FileBodyDecoder validates request body for
		"application/octet-stream" but can be reused to add other
		MIME types for binary payloads.
		ImageBodyDecoder is the custom decoder used to check
		payload for "image/jpeg" content type
	*/

	openapi3filter.RegisterBodyDecoder("image/jpeg", ImageBodyDecoder)
	openapi3filter.RegisterBodyDecoder("image/png", openapi3filter.FileBodyDecoder)
	// We now register our service with generated server
	gen.RegisterHandlers(e, s)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}

// Used to check that payload actually looks like "image/jpeg"
func ImageBodyDecoder(body io.Reader, header http.Header, schema *openapi3.SchemaRef, encFn openapi3filter.EncodingFn) (interface{}, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, &openapi3filter.ParseError{Kind: openapi3filter.KindInvalidFormat, Cause: err}
	}

	if http.DetectContentType(data) != "image/jpeg" {
		return "", fmt.Errorf("Recieved payload doesn't match the content-type")
	}

	return string(data), nil
}
