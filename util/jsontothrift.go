// Package util
package util

import (
	"bytes"
	"encoding/json"
	"html/template"
	"math"
	"reflect"
	"regexp"
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

	// sort
	// pretty json
	var out bytes.Buffer
	err = json.Indent(&out, []byte(jsonBytes), "", "    ")
	if err != nil {
		return nil, err
	}
	r, _ := regexp.Compile(`\"\w+\":`)
	keySorts := r.FindAllString(out.String(), -1)
	for i := 0; i < len(keySorts); i++ {
		keySorts[i] = keySorts[i][1 : len(keySorts[i])-2]
	}
	sortMap := make(map[string]int, len(keySorts))
	for i, key := range keySorts {
		sortMap[key] = i
	}
	ts.SortFields = sortMap
	sort.Sort(ts)

	var buf bytes.Buffer

	tpl := template.New("clientTpl").Funcs(funcMap)
	tpl = template.Must(tpl.Parse(thriftStructTmpl))
	err = tpl.ExecuteTemplate(&buf, "clientTpl", ts)
	if err != nil {
		return nil, err
	}
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
		return "not support type " + reflect.ValueOf(v).Kind().String()
	}
}

type ThriftStruct struct {
	StructName string
	Fields     []*Declare
	SortFields map[string]int
}

func (ts *ThriftStruct) Len() int {
	return len(ts.Fields)
}

func (ts *ThriftStruct) Less(i, j int) bool {
	sa, sb := ts.SortFields[ts.Fields[i].Name], ts.SortFields[ts.Fields[j].Name]
	return sa < sb
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
{{range $i, $field := .Fields}}	{{inc $i}}: required {{$field.Type}} {{$field.Name}},
{{end}}}
`
