package Integrations

import (
	"encoding/json"
	"fmt"
	"github.com/palantir/stacktrace"
	"gitlab.com/cloudnuro-mvp/cloudnurodao/Models"
	"gitlab.com/cloudnuro-mvp/framework/Utils"
	"strings"
)

/*
	Creates query parameters based on the map provided. The resulting
	key value pairs are not URL encoded by design, as golang will
	automatically encode them when they're put into the URL. This is
	designed so that all param values are on the same key, unlike the
	spec which designates an array.
*/
func createQueryParams(params map[string]string) string {
	str := "?"
	first := true // Not sure the idiomatic way to do this, as there's no index to work with.
	for key, value := range params {
		before := "&"
		if first {
			first = false
			before = ""
		}
		str += fmt.Sprintf("%v%v=%v", before, key, value)
	}
	return str
}

/*
	Replaces all the parameters in the Models.AppConfig with the
	values in the provided parameters map. It then returns the
	Models.AppConfig with the updated values, wherever they
	are.
*/
func replace(payload Models.AppConfig, params map[string]string) (*Models.AppConfig, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	// We remove the step of marshaling to a JsonObject to use the other method
	// If your replacement breaks json there's a bigger problem at hand
	data = []byte(Utilities.ReplaceParamsInString(string(data), params))
	var resp Models.AppConfig
	err = json.Unmarshal(data, &resp)
	return &resp, err
}

// TODO: Fix with hateoas when I have an example
/*
	Runs the specified integration. This will return the data from
	the URL, parsed to whatever format the AppConfig specifies.
	Any parameters in the AppConfig are in the params tag, and
	are replaced before sending. If the JWT key is non-null it
	is checked and the JWT key returned.
*/
func RunIntegration(payload *Models.AppConfig, params map[string]string) ([]byte, error) {
	// If JWT is nil, then it's a normal auth.
	if payload.JWT != nil {
		token, err := GetJwtToken(payload, params)
		return []byte(token), err
	}

	// Replace any variables that might be important
	payload, err := replace(*payload, params)

	if err != nil {
		return nil, stacktrace.Propagate(err, "error marshalling payload for replacement")
	}

	// Make the request to the URL
	resp, err := Utilities.Request(payload.URL+createQueryParams(payload.QueryParams), payload.Method, payload.Body, payload.Headers)

	if err != nil {
		return nil, stacktrace.Propagate(err, "Error requesting during integration")
	}

	// Read the body.
	responseBody, err := Utilities.ReadBody(resp)

	if err != nil {
		return nil, stacktrace.Propagate(err, "Error reading body from request")
	}

	// Parse the response into an object or an array based on the type hint
	respType := payload.Response.Type
	var response Utilities.JsonStructure
	if strings.ToLower(respType) == "object" {
		response, err = Utilities.MakeObj(string(responseBody))
	} else {
		response, err = Utilities.MakeArr(string(responseBody))
	}

	if err != nil {
		return nil, stacktrace.Propagate(err, "Error decoding response for parsing")
	}

	// Parse the response for the data we care about
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, stacktrace.Propagate(err, "error marshalling payload for parsing")
	}

	obj, err := Utilities.MakeObj(string(data))
	if err != nil {
		return nil, stacktrace.Propagate(err, "error making object")
	}

	return []byte(fmt.Sprintf("%v", Utilities.Parse(response, obj))), nil
}
