package dto_kv

type FloatKeyValue struct {
	Key   string  `json:"key,omitempty"`
	Value float64 `json:"value,omitempty"`
}

type IntKeyValue struct {
	Key   string `json:"key,omitempty"`
	Value int64  `json:"value,omitempty"`
}

type StrKeyValue struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type KeyValues struct {
	StrList   []*StrKeyValue   `json:"str_list,omitempty"`
	IntList   []*IntKeyValue   `json:"int_list,omitempty"`
	FloatList []*FloatKeyValue `json:"float_list,omitempty"`
}
