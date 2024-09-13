package utils

import "encoding/json"

func JsonMarshaller(content interface{}) []byte {

	b, _ := json.Marshal(content)

	return b

}
