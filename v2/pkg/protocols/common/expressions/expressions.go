package expressions

import (
	"regexp"

	"github.com/Knetic/govaluate"
	"github.com/onaio/nuclei/v2/pkg/operators/common/dsl"
	"github.com/onaio/nuclei/v2/pkg/protocols/common/generators"
	"github.com/onaio/nuclei/v2/pkg/protocols/common/replacer"
)

var templateExpressionRegex = regexp.MustCompile(`(?m)\{\{[^}]+\}\}["'\)\}]*`)

// Evaluate checks if the match contains a dynamic variable, for each
// found one we will check if it's an expression and can
// be compiled, it will be evaluated and the results will be returned.
//
// The provided keys from finalValues will be used as variable names
// for substitution inside the expression.
func Evaluate(data string, base map[string]interface{}) (string, error) {
	data = replacer.Replace(data, base)

	dynamicValues := make(map[string]interface{})
	for _, match := range templateExpressionRegex.FindAllString(data, -1) {
		expr := generators.TrimDelimiters(match)

		compiled, err := govaluate.NewEvaluableExpressionWithFunctions(expr, dsl.HelperFunctions())
		if err != nil {
			continue
		}
		result, err := compiled.Evaluate(base)
		if err != nil {
			continue
		}
		dynamicValues[expr] = result
	}
	// Replacer dynamic values if any in raw request and parse  it
	return replacer.Replace(data, dynamicValues), nil
}

// EvaluateByte checks if the match contains a dynamic variable, for each
// found one we will check if it's an expression and can
// be compiled, it will be evaluated and the results will be returned.
//
// The provided keys from finalValues will be used as variable names
// for substitution inside the expression.
func EvaluateByte(data []byte, base map[string]interface{}) ([]byte, error) {
	final := replacer.Replace(string(data), base)

	dynamicValues := make(map[string]interface{})
	for _, match := range templateExpressionRegex.FindAllString(final, -1) {
		expr := generators.TrimDelimiters(match)

		compiled, err := govaluate.NewEvaluableExpressionWithFunctions(expr, dsl.HelperFunctions())
		if err != nil {
			continue
		}
		result, err := compiled.Evaluate(base)
		if err != nil {
			continue
		}
		dynamicValues[expr] = result
	}
	// Replacer dynamic values if any in raw request and parse  it
	return []byte(replacer.Replace(final, dynamicValues)), nil
}
