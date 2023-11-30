package main

import (
	"os"

	"github.com/pb33f/libopenapi"
	validator "github.com/pb33f/libopenapi-validator"
)

func main() {
	openapiSpec, err := os.ReadFile("openapi.spec.yml")

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
