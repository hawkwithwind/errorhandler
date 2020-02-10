package errorhandler

import "fmt"

func (o *ErrorHandler) ListValue(value interface{}) []interface{} {
	if o.Err != nil {
		return nil
	}

	switch value := value.(type) {
	case []interface{}:
		return value
	case nil:
		o.Err = fmt.Errorf("expect listvalue, but nil")
	default:
		o.Err = fmt.Errorf("expect listvalue but %T", value)
	}

	return nil
}

func (o *ErrorHandler) NilAs(value interface{}, defaultValue interface{}) interface{} {
	// this function won't produce error, for semantic consistency, this function didn't short circuit o.Err != nil
	
	if value == nil {
		return defaultValue
	} else {
		return value
	}
}



