// Package util
package util

import (
	"bytes"
	"encoding/json"
	"html/template"
	"math"
	"reflect"
	"sort"
)

func JsonToThrift(jsonBytes string) ([]byte, error) {
	mapJson := make(map[string]interface{})
	decoder := json.NewDecoder(bytes.NewBufferString(jsonBytes))
	decoder.UseNumber()
	err := decoder.Decode(&mapJson)
	if err != nil {
		return nil, err
	}

	ts := NewThriftStruct()

	for k, v := range mapJson {
		ts.Fields = append(ts.Fields, &Declare{getThriftType(v), k})
	}
	sort.Sort(ts)

	var buf bytes.Buffer

	tpl := template.New("clientTpl").Funcs(funcMap)
	tpl = template.Must(tpl.Parse(thriftStructTmpl))
	err = tpl.ExecuteTemplate(&buf, "clientTpl", ts)

	return buf.Bytes(), nil
}

func getThriftType(v interface{}) string {
	switch v.(type) {
	case bool:
		return "bool"
	case int8, uint8:
		return "byte"
	case string:
		return "string"
	case json.Number:
		vv, _ := v.(json.Number).Int64()
		if vv < math.MaxUint8 {
			return "byte"
		} else if vv < math.MaxInt16 {
			return "i16"
		} else if vv < math.MaxUint32 {
			return "i32"
		}
		return "i64"
	default:
		_ = v.(int)
		return "not support type " + reflect.ValueOf(v).Kind().String()
	}
}

type ThriftStruct struct {
	StructName string
	Fields     []*Declare
}

func (ts *ThriftStruct) Len() int {
	return len(ts.Fields)
}

func (ts *ThriftStruct) Less(i, j int) bool {
	return ts.Fields[i].Name < ts.Fields[j].Name
}

func (ts *ThriftStruct) Swap(i, j int) {
	ts.Fields[i], ts.Fields[j] = ts.Fields[j], ts.Fields[i]
}

type Declare struct {
	Type string
	Name string
}

func NewThriftStruct() *ThriftStruct {
	return &ThriftStruct{
		StructName: "Foo",
	}
}

var funcMap = template.FuncMap{
	// The name "inc" is what the function will be called in the template text.
	"inc": func(i int) int {
		return i + 1
	},
}

var thriftStructTmpl = `
struct {{.StructName}} {
{{range $i, $field := .Fields}}
    {{inc $i}}: required {{$field.Type}} {{$field.Name}},{{end}}
}
`
