package add

func Map(a map[string]interface{}, b interface{}) (map[string]interface{}, error) {
	switch bt := b.(type) {
	case map[string]interface{}:
		for k, v := range bt {
			a[k] = v
		}
		return a, nil
	}

	return nil, Cant(a, b)
}
