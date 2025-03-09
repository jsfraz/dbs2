package utils

import (
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

// Vrací pole fizz.OperationOption popisující API endpoint.
//
//	@param summary
//	@param description
//	@param errorStatuses
//	@param useSecurity
//	@return []fizz.OperationOption
func CreateOperationOption(summary string, description string, useSecurity bool) []fizz.OperationOption {
	var option []fizz.OperationOption
	option = append(option, fizz.Summary(summary))
	if description != "" {
		option = append(option, fizz.Description(description))
	}
	if useSecurity {
		option = append(option, fizz.Security(&openapi.SecurityRequirement{
			"bearerAuth": []string{},
		}))
	}
	return option
}
