package stringutil

import (
	"bytes"
	"html/template"
)

// Insert parameters to template string
func Insert(str string, params interface{}) (string, error) {
	var res string
	tmpl, err := template.New("str").Parse(str)

	if err != nil {
		return res, err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, params)
	if err != nil {
		return res, err
	}

	res = buf.String()

	return res, nil
}
