package errorhandler

import "encoding/json"

func (o *ErrorHandler) ToJson(v interface{}) string {
	if o.Err != nil {
		return ""
	}

	var jsonstr []byte
	jsonstr, o.Err = json.Marshal(v)
	if o.Err != nil {
		return ""
	}

	return string(jsonstr)
}


func (o *ErrorHandler) FromJson(jsonstr string) map[string]interface{} {
	if o.Err != nil {
		return nil
	}

	ret := make(map[string]interface{})
	o.Err = json.Unmarshal([]byte(jsonstr), &ret)
	return ret
}
