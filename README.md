# jsonprinter
json printer for golang

## version 0.1 
Initial APIs to support

```
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
```
