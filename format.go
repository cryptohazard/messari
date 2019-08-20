package messari

import (
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

var tablewriter = table.NewWriter()

func init() {
	tablewriter.SetStyle(table.Style{
		Name: "myNewStyle",
		Box: table.BoxStyle{
			MiddleHorizontal: "-",
			MiddleSeparator:  "\t",
			MiddleVertical:   "\t",
		},
		Format: table.FormatOptions{
			Header: text.FormatUpper,
			Row:    text.FormatDefault,
		},
		Options: table.Options{
			SeparateColumns: true,
			SeparateHeader:  false,
		},
	})
}
