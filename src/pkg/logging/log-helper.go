package logging

//mapToZapParams

func mapToZapParams(keys map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0, len(keys))

	for k, v := range keys {
		params = append(params, string(k))
		params = append(params, v)
	}

	return params
}

func mapToZeroParams(extra map[ExtraKey]interface{})map[string]interface{} {
	params := map[string]interface{}{}
	for k, v:= range extra{
		params[string(k)] = v
	}
	return params
}