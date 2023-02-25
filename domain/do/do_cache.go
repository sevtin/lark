package do

type KeysValues struct {
	Keys   []string      `json:"keys"`
	Values []interface{} `json:"values"`
}

type KeyMaps struct {
	Key  interface{}            `json:"key"`
	Maps map[string]interface{} `json:"maps"`
}

type KeyFieldValue struct {
	Key   interface{} `json:"key"`
	Field interface{} `json:"field"`
	Value interface{} `json:"value"`
}
