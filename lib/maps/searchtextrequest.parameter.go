package maps

type SearchTextRequestOptional func(*SearchTextRequest) *SearchTextRequest

func ApplySearchTextRequest(defaultValue SearchTextRequest, fn ...SearchTextRequestOptional) *SearchTextRequest {
	param := &defaultValue
	for _, f := range fn {
		param = f(param)
	}
	return param
}
