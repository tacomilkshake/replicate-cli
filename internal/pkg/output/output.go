package output

import "github.com/tidwall/gjson"

func FormatOutput(output string) string {
	outputArray := gjson.Get(output, "output").Array()
	return outputArray[0].String()
}
