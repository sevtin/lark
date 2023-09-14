package do

type Result struct {
	Code int32
	Msg  string
	Err  error
}

func (r *Result) Set(Code int32, msg string) {
	r.Code = Code
	r.Msg = msg
}
