package configman

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func extractPathParams(urlPath string) (op, module, typeName, resourceName string, err error) {
	// Set the default operation to single
	op = "single"

	// Check if url has proper length
	arr := strings.Split(urlPath[1:], "/")
	if len(arr) > 5 || len(arr) < 4 {
		err = fmt.Errorf("invalid config url provided - %s", urlPath)
		return
	}

	// Check the operation type
	if len(arr) == 5 {
		op = "list"
	}

	// Set the other parameters
	module = arr[2]
	typeName = arr[3]
	resourceName = arr[4]
	return
}

func verifySpecSchema(typeDef *TypeDefinition, spec interface{}) ([]string, error) {
	// Skip verification if no json schema is supplied
	if typeDef.Schema == nil {
		return nil, nil
	}

	// Perform JSON schema validation
	schemaLoader := gojsonschema.NewGoLoader(typeDef.Schema)
	documentLoader := gojsonschema.NewGoLoader(spec)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, err
	}

	// Skip if no errros were found
	if result.Valid() {
		return nil, nil
	}

	// Send back all the errors
	arr := make([]string, len(result.Errors()))
	for i, err := range result.Errors() {
		arr[i] = err.String()
	}

	return arr, fmt.Errorf("json schema validation failed")
}

func verifyConfigParents(typeDef *TypeDefinition, parents map[string]string) error {
	// Simply return if object has no required parents
	if len(typeDef.RequiredParents) == 0 {
		return nil
	}

	// Send error if no parents are provided
	if parents == nil {
		return errors.New("resource doesn't have required parents")
	}

	// Check if all required parents are available
	for _, parent := range typeDef.RequiredParents {
		if _, p := parents[parent]; !p {
			return fmt.Errorf("parent '%s' not present in resource", parent)
		}
	}

	return nil
}

func prepareErrorResponseBody(err error, schemaErrors []string) interface{} {
	return map[string]interface{}{
		"error":        err.Error(),
		"schemaErrors": schemaErrors,
	}
}
