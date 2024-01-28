package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	// 公募基金
	makeProto("fund", "公募基金", "https://tushare.pro/document/2?doc_id=18")
	// 期货
	makeProto("fut", "公募基金", "https://tushare.pro/document/2?doc_id=134")
}

// makeProto
func makeProto(name string, nameZh string, url string) {
	// 读取列表
	entries, err := os.ReadDir("./xlsx/" + name)
	if err != nil {
		panic(err)
	}

	protoContent := fmt.Sprintf(`syntax = "proto3";
package api.v1;
option go_package = "github.com/isfk/tushare/gen/api/v1;v1";

/*
	- %s
	- %s
*/

`, nameZh, url)

	for _, file := range entries {
		if strings.Contains(file.Name(), ".xlsx") {
			f, err := excelize.OpenFile("./xlsx/" + name + "/" + file.Name())
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()

			// info
			apiName, err := f.GetCellValue("info", "A1")
			if err != nil {
				fmt.Println(err)
				continue
			}
			ApiName, err := f.GetCellValue("info", "A2")
			if err != nil {
				fmt.Println(err)
				continue
			}
			apiNameZh, err := f.GetCellValue("info", "A3")
			if err != nil {
				fmt.Println(err)
				continue
			}

			// request
			rows, err := f.GetRows("request")
			if err != nil {
				fmt.Println(err)
				continue
			}
			requestLines := ""
			if len(rows) > 0 {
				requestLines += "\tstring limit = 1 [json_name = \"limit\"];\n\tstring offset = 2 [json_name = \"offset\"];\n"
				for i, row := range rows {
					rowSlice := []string{}
					rowSlice = append(rowSlice, row...)
					requestLines += line(rowSlice, i+2)
				}
			}
			// response
			rows, err = f.GetRows("response")
			if err != nil {
				fmt.Println(err)
				continue
			}

			responseLines := ""
			for i, row := range rows {
				rowSlice := []string{}
				rowSlice = append(rowSlice, row...)
				responseLines += line(rowSlice, i)
			}

			responseMessage := fmt.Sprintf(`// %v|%v
message %v {
%v}
message %vRequest {
%v}
message %vResponse {
	repeated %v list = 1;
}

`, apiNameZh, apiName,
				ApiName,
				responseLines,
				ApiName,
				requestLines,
				ApiName,
				ApiName,
			)

			protoContent += responseMessage
		}
	}

	err = os.WriteFile("../proto/api/v1/"+name+".proto", []byte(protoContent), 0o666)
	if err != nil {
		panic(err)
	}
}

func line(rowSlice []string, number int) string {
	if len(rowSlice) != 3 {
		return ""
	}

	field := rowSlice[0]
	kind := rowSlice[1]
	comment := rowSlice[2]

	switch kind {
	case "str":
		kind = "string"
	case "int":
		kind = "int64"
	case "float":
		kind = "float"
	default:
		kind = "string"
	}
	// ${kind} ${field} = ${number}+1 [json_name = "${field}"]; // ${comment}
	return fmt.Sprintf("\t%s %s = %d [json_name = \"%s\"]; // %s\n", kind, field, number+1, field, comment)
}
