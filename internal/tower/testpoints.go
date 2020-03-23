package tower

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

type MethodType int

const (
	Binary MethodType = iota
	Unary
	Custom
)

type MethodTypeMap map[string]MethodType

// MethodTypes for named constants inside templates
var MethodTypes = MethodTypeMap{
	"Binary": Binary,
	"Unary":  Unary,
	"Custom": Custom,
}

type Method struct {
	Name string
	Type MethodType
}

type FString = []string
type FStringPair = [2]FString

type TestPoint struct {
	In  FStringPair
	Out []FString
}

func GenerateTestOutputs(inputs []TestPoint, sagePath string, sagePrefixArgs ...string) (result []TestPoint, err error) {
	var args, outputs []string

	result = inputs

	// prepare input args for sage
	args = sagePrefixArgs
	for i := range result {
		for j := range result[i].In {
			args = append(args, result[i].In[j][:]...)
		}
	}

	// get outputs from sage
	var output bytes.Buffer
	cmd := exec.Command(sagePath, args...)
	cmd.Stdout = &output
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, errors.New("sage script failed")
	}
	outputs = strings.Split(output.String(), "\n")

	// uncomment for debugging
	// for i := range outputs {
	// 	fmt.Printf("\"%s\",\n", outputs[i])
	// }

	// uncomment to hard-code sage output
	// outputs = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}

	// fill result outputs
	for i := range result {
		for j := range result[i].Out {
			for k := range result[i].Out[j] {
				index := i
				index = index*len(result[i].Out) + j
				index = index*len(result[i].Out[j]) + k
				result[i].Out[j][k] = outputs[index]
			}
		}
	}

	return result, nil
}

const TestPoints = `
// Code generated by internal/tower DO NOT EDIT 
package {{.PackageName}}

{{ define "printTestPoint" }}
	{{- range $coord := .Coords -}}
		"{{$coord}}",
	{{- end -}}
{{ end }}

func init() {
	{{if ne (len .TestPoints) 0}}
		{{.Name}}TestPoints = make([]{{.Name}}TestPoint, {{len .TestPoints}})
	{{end}}
	{{range $i, $point := .TestPoints}}
		{{- range $j, $in := $point.In }}
			{{$.Name}}TestPoints[{{$i}}].in[{{$j}}].SetString({{ template "printTestPoint" dict "Coords" (index $point.In $j) }})
		{{- end }}
		{{- range $j, $out := $point.Out }}
			{{$.Name}}TestPoints[{{$i}}].out[{{$j}}].SetString({{ template "printTestPoint" dict "Coords" (index $point.Out $j) }})
		{{- end }}
	{{end}}

	{{if ne (len .TestPoints) 0}}
		// benchmark inputs should be randomly generated,
		// so use the final test point
		{{.Name}}BenchIn1.Set(&{{.Name}}TestPoints[len({{.Name}}TestPoints)-1].in[0])
		{{.Name}}BenchIn2.Set(&{{.Name}}TestPoints[len({{.Name}}TestPoints)-1].in[1])
	{{- end }}
}
`
