package main

import (
	"flag"
	"os"

	"github.com/pb33f/libopenapi"
	validator "github.com/pb33f/libopenapi-validator"
)

func main() {
	specPtr := flag.String("spec", "openapi.spec.yml", "the openapi spec to be validated")

	flag.Parse()

	openapiSpec, err := os.ReadFile(*specPtr)

	if err != nil {
		panic(err)
	}

	doc, docErrs := libopenapi.NewDocument(openapiSpec)

	if docErrs != nil {
		panic(docErrs)
	}

	_, validatorErrs := validator.NewValidator(doc)

	if validatorErrs != nil {
		panic(validatorErrs)
	}

}
