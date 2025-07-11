package utils

func InterfaceToString(params interface{}) string {
	if params == nil {
		return ""
	}

	return params.(string)
}

func InterfaceToInt(params interface{}) int {
	if params == nil {
		return 0
	}

	return params.(int)
}

func InterfaceToFloat(params interface{}) float64 {
	if params == nil {
		return 0.0
	}

	return params.(float64)
}
