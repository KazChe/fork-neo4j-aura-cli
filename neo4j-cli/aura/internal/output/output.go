package output

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/neo4j/cli/common/clicfg"
	"github.com/neo4j/cli/neo4j-cli/aura/internal/api"
)

func PrintBodyMap(cmd *cobra.Command, cfg *clicfg.Config, values api.ResponseData, fields []string) {
	outputType := cfg.Aura.Output()

	switch output := outputType; output {
	case "json":
		bytes, err := json.MarshalIndent(values, "", "\t")
		if err != nil {
			panic(err)
		}
		cmd.Println(string(bytes))
	case "text":
		printPlainText(cmd, values, fields)
	case "table", "default":
		printTable(cmd, values, fields)
	default:
		// This is in case the value is unknown
		cmd.Println(values)
	}
}

func PrintBody(cmd *cobra.Command, cfg *clicfg.Config, body []byte, fields []string) {
	if len(body) == 0 {
		return
	}
	values := api.ParseBody(body)

	PrintBodyMap(cmd, cfg, values, fields)
}

func printTable(cmd *cobra.Command, responseData api.ResponseData, fields []string) {
	t := table.NewWriter()

	header := table.Row{}
	for _, f := range fields {
		header = append(header, f)
	}

	t.AppendHeader(header)
	for _, v := range responseData.AsArray() {
		row := table.Row{}
		for _, f := range fields {
			formattedValue := v[f]

			if v[f] == nil {
				formattedValue = ""
			}

			if reflect.TypeOf(formattedValue).Kind() == reflect.Slice {
				marshaledSlice, _ := json.MarshalIndent(formattedValue, "", "  ")
				formattedValue = string(marshaledSlice)
			}

			row = append(row, formattedValue)
		}
		t.AppendRow(row)
	}

	t.SetStyle(table.StyleLight)
	cmd.Println(t.Render())
}

func printPlainText(cmd *cobra.Command, responseData api.ResponseData, fields []string) {
	for _, item := range responseData.AsArray() {
		values := make([]string, 0)
		for _, field := range fields {
			value := item[field]
			if value == nil {
				values = append(values, "")
				continue
			}
			values = append(values, fmt.Sprintf("%v", value))
		}

		cmd.Println(strings.Join(values, "\t"))
	}
}
