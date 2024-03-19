package jsonwriter

import (
	"encoding/json"
	"io"
)

/*
table := tablewriter.NewWriter(w)
table.SetHeader(header)
table.SetAutoWrapText(false)
table.SetAutoFormatHeaders(true)
table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
table.SetAlignment(tablewriter.ALIGN_LEFT)
table.SetCenterSeparator("")
table.SetColumnSeparator("")
table.SetRowSeparator("")
table.SetHeaderLine(false)
table.SetBorder(false)
table.SetTablePadding("\t")
table.SetNoWhiteSpace(true)
table.AppendBulk(rows)
table.Render()
*/

func (doc *JsonDoc) Render() {
	b, err := json.Marshal(doc.data)
	if err != nil {
		panic(err)
	}
	_, err = doc.w.Write(b)
	if err != nil {
		panic(err)
	}
}

func (doc *JsonDoc) AppendBulk(rows [][]string) {
	doc.data.Data = rows
}

func (doc *JsonDoc) SetHeader(keys []string) {
	doc.data.Metadata.Header = keys
}

type JsonDoc struct {
	w    io.Writer
	data rawData
}

type rawData struct {
	Metadata metadata   `json:"metadata"`
	Data     [][]string `json:"data"`
}

type metadata struct {
	Header        []string `json:"header"`
	SchemaVersion string   `json:"schema-version"`
}

func NewWriter(w io.Writer) *JsonDoc {
	jsonDoc := new(JsonDoc)
	jsonDoc.setWriter(w)
	jsonDoc.data.Metadata.SchemaVersion = JSON_SCHEMA_VERSION
	return jsonDoc
}

func (doc *JsonDoc) setWriter(w io.Writer) {
	doc.w = w
}

const (
	JSON_SCHEMA_VERSION string = "0.1"
)
