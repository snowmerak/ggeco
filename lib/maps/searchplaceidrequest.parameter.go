package maps

type SearchPlaceIdRequestOptional func(*SearchPlaceIdRequest) *SearchPlaceIdRequest

func ApplySearchPlaceIdRequest(defaultValue SearchPlaceIdRequest, fn ...SearchPlaceIdRequestOptional) *SearchPlaceIdRequest {
	param := &defaultValue
	for _, f := range fn {
		param = f(param)
	}
	return param
}
