package opers

// func Add(c *ast.Context, a, b interface{}) (interface{}, error) {
// 	switch at := a.(type) {
// 	case ast.Adder:
// 		return at.Add(c, b)
// 	case int:
// 		switch bt := b.(type) {
// 		case ast.Inter:
// 			return at + bt.Int(), nil
// 		case ast.Floater:
// 			return float64(at) + bt.Float(), nil
// 		case int:
// 			return at + bt, nil
// 		case float64:
// 			return float64(at) + bt, nil
// 		}
// 	case float64:
// 	case []interface{}:
// 	}
//
// 	return nil, fmt.Errorf("can't add %T and %T", a, b)
// }
