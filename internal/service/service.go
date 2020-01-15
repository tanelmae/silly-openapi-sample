package service

import (
	"fmt"
	"github.com/Pallinder/sillyname-go"
	"github.com/labstack/echo/v4"
	"github.com/sourcekris/goTypos"
	"github.com/tanelmae/silly-openapi-sample/pkg/gen"
	"log"
	"math/rand"
	"net/http"
)

func New() Server {
	name := sillyname.GenerateStupidName()
	log.Printf("Server created with name %s", name)
	return Server{name: name}
}

type Server struct {
	name string
}

func (s Server) Hello(ctx echo.Context, params gen.HelloParams) error {
	log.Printf("%+v", params)

	// In OpenAPI yaml this parameter has 'required: false'
	// So we can handle the default here.
	if params.Name == nil {
		log.Println("Uh-oh, no name parameter found")
		n := sillyname.GenerateStupidName()
		params.Name = &n
		log.Printf("Calling him %s", *params.Name)
	}

	greeting := fmt.Sprintf("Hello, %s!", misspell(*params.Name))
	intro := fmt.Sprintf("My name is %s.", s.name)
	r := gen.HelloResp{
		Greeting:     &greeting,
		Introduction: &intro,
	}

	return ctx.JSON(http.StatusOK, r)
}

func (s Server) HelloPath(ctx echo.Context, name string) error {
	// For this parameter it is set 'required: true'
	// So no need to check if the value is set here

	greeting := fmt.Sprintf("Hello, %s!", misspell(name))
	intro := fmt.Sprintf("My name is %s.", s.name)
	r := gen.HelloResp{
		Greeting:     &greeting,
		Introduction: &intro,
	}

	return ctx.JSON(http.StatusOK, r)
}

func (s Server) Nameupload(ctx echo.Context) error {
	// Load struct values from the payload
	var req gen.HelloReq
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}

	greeting := fmt.Sprintf("Hello, %s!", misspell(req.Name))
	intro := fmt.Sprintf("My name is %s.", s.name)
	r := gen.HelloResp{
		Greeting:     &greeting,
		Introduction: &intro,
	}

	return ctx.JSON(http.StatusOK, r)
}

func misspell(name string) string {

	t := typos.NewTypos(name)

	switch rand.Intn(5) {
	case 0:
		t.SkipLetter()
	case 1:
		t.DoubleLetter()
	case 2:
		t.WrongVowel()
	case 3:
		t.ReverseLetter()
	case 4:
		t.WrongKey()
	}
	return t.Typos[0]
}
