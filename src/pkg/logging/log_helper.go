package logging

func mapToZapParams(extras map[ExtraKey]interface{}) []interface{} {
	params := []interface{}{}
	for k, v := range extras {
		params = append(params, string(k))
		params = append(params, v)
	}
	return params
}
