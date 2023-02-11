package pb_kv

func (p *KeyValues) StrFieldValidation(maps map[string]string, out map[string]interface{}) {
	var (
		v  *StrKeyValue
		ok bool
	)
	if len(p.StrList) == 0 {
		return
	}
	for _, v = range p.StrList {
		if _, ok = maps[v.Key]; ok == false {
			continue
		}
		out[v.Key] = v.Value
	}
}

func (p *KeyValues) IntFieldValidation(maps map[string]string, out map[string]interface{}) {
	var (
		v  *IntKeyValue
		ok bool
	)
	if len(p.IntList) == 0 {
		return
	}
	for _, v = range p.IntList {
		if _, ok = maps[v.Key]; ok == false {
			continue
		}
		out[v.Key] = v.Value
	}
}

func (p *KeyValues) FloatFieldValidation(maps map[string]string, out map[string]interface{}) {
	var (
		v  *FloatKeyValue
		ok bool
	)
	if len(p.FloatList) == 0 {
		return
	}
	for _, v = range p.FloatList {
		if _, ok = maps[v.Key]; ok == false {
			continue
		}
		out[v.Key] = v.Value
	}
}
