package lets

import "encoding/json"

type any = interface{}

func Bind(source any, v any) {
	jsonSource, err := json.Marshal(source)
	if err != nil {
		LogE("Bind: %s", err.Error())
	}

	err = json.Unmarshal(jsonSource, v)
	if err != nil {
		LogE("Bind: %s", err.Error())
	}
}
