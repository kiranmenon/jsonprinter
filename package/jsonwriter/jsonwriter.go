package jsonwriter

import (
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
	return
}

func (doc *JsonDoc) AppendBulk(rows [][]string) {
	return
}

func (doc *JsonDoc) SetHeader(keys []string) {
	return
}

func (doc *JsonDoc) SetNoWhiteSpace(b bool) {
	return
}

func (doc *JsonDoc) SetTablePadding(s string) {
	return
}

func (doc *JsonDoc) SetBorder(b bool) {
	return
}

func (doc *JsonDoc) SetHeaderLine(b bool) {
	return
}

func (doc *JsonDoc) SetRowSeparator(s string) {
	return
}

func (doc *JsonDoc) SetColumnSeparator(s string) {
	return
}

func (doc *JsonDoc) SetCenterSeparator(s string) {
	return
}

const (
	ALIGN_LEFT = iota
)

func (doc *JsonDoc) SetAlignment(i uint) {
	return
}

func (doc *JsonDoc) SetHeaderAlignment(i uint) {
	return
}

func (doc *JsonDoc) SetAutoFormatHeaders(b bool) {
	return
}

func (doc *JsonDoc) SetAutoWrapText(b bool) {
	return
}

type JsonDoc struct {
	w io.Writer
}

func NewWriter(w io.Writer) *JsonDoc {
	jsonDoc := new(JsonDoc)
	jsonDoc.SetWriter(w)
	return jsonDoc
}

func (doc *JsonDoc) SetWriter(w io.Writer) {
	doc.w = w
}
