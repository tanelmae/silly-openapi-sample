// Package gen provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package gen

// Error defines model for Error.
type Error struct {
	Message *string `json:"message,omitempty"`
}

// HelloReq defines model for HelloReq.
type HelloReq struct {
	Name string `json:"name"`
}

// HelloResp defines model for HelloResp.
type HelloResp struct {
	Greeting     *string `json:"greeting,omitempty"`
	Introduction *string `json:"introduction,omitempty"`
}

// HelloParams defines parameters for Hello.
type HelloParams struct {

	// Send client name to the server
	Name *string `json:"name,omitempty"`
}

// nameuploadJSONBody defines parameters for Nameupload.
type nameuploadJSONBody HelloReq

// NameuploadRequestBody defines body for Nameupload for application/json ContentType.
type NameuploadJSONRequestBody nameuploadJSONBody