package entity

type EsDsl struct {
	From   int       `json:"from,omitempty"`
	Size   int       `json:"size,omitempty"`
	Query  *Query    `json:"query,omitempty"`
	Source bool      `json:"_source,omitempty"`
	Fields []*Fields `json:"fields,omitempty"`
	Sort   []*Sort   `json:"sort,omitempty"`
}

func NewEsDsl() *EsDsl {
	dsl := &EsDsl{
		Query: &Query{Bool: &Bool{
			Must:               make([]*Must, 0),
			AdjustPureNegative: false,
			Boost:              1.0,
		}},
		Fields: make([]*Fields, 0),
		Sort:   make([]*Sort, 0)}
	return dsl
}

func (e *EsDsl) SetSize(size int) *EsDsl {
	e.Size = size
	return e
}

func (e *EsDsl) SetSource(source bool) *EsDsl {
	e.Source = source
	return e
}

func (e *EsDsl) SetFrom(from int) *EsDsl {
	e.From = from
	return e
}

func (e *EsDsl) SetMust(must *Must) {
	e.Query.Bool.Must = append(e.Query.Bool.Must, must)
}

func (e *EsDsl) SetSort(sort *Sort) {
	e.Sort = append(e.Sort, sort)
}

func (m *Must) SetFieldRange(key string, field *Field) {
	m.Range[key] = field
}

func (m *Must) SetWildcard(wildcard string, boost float64) {
	m.Wildcard = &Wildcard{BodyKeyword: &BodyKeyword{
		Wildcard: wildcard,
		Boost:    boost,
	}}
}

type Query struct {
	Bool *Bool `json:"bool,omitempty"`
}

type Bool struct {
	Must               []*Must `json:"must,omitempty"`
	AdjustPureNegative bool    `json:"adjust_pure_negative,omitempty"`
	Boost              float64 `json:"boost,omitempty"`
}

type Must struct {
	Wildcard *Wildcard              `json:"wildcard,omitempty"`
	Range    map[string]interface{} `json:"range,omitempty"`
	Bool     *Bool                  `json:"bool,omitempty"`
}

type Wildcard struct {
	BodyKeyword *BodyKeyword `json:"body.keyword"`
}

type BodyKeyword struct {
	Wildcard string  `json:"wildcard,omitempty"`
	Boost    float64 `json:"boost,omitempty"`
}

type Fields struct {
	Field string `json:"field,omitempty"`
}

type Sort struct {
	Doc *Doc `json:"_doc,omitempty"`
}

type Doc struct {
	Order string `json:"order,omitempty"`
}

type Field struct {
	From         int         `json:"from,omitempty"`
	To           interface{} `json:"to,omitempty"`
	IncludeLower bool        `json:"include_lower,omitempty"`
	IncludeUpper bool        `json:"include_upper,omitempty"`
	Boost        float64     `json:"boost,omitempty"`
}
