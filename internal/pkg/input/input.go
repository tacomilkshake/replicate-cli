package input

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/danielgtaylor/shorthand"
)

func GetPipedArgs() (string, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	if info.Mode()&os.ModeNamedPipe != 0 {
		reader := bufio.NewReader(os.Stdin)
		var output []rune

		for {
			input, _, err := reader.ReadRune()
			if err != nil && err == io.EOF {
				break
			}
			output = append(output, input)
		}

		return string(output), nil
	} else {
		return "", nil
	}
}

func ConvertShorthandToInterface(input []string) (map[string]interface{}, error) {
	result, err := shorthand.GetInput(input)
	if err != nil {
		return nil, errors.New("specify input in shorthand json format, see --help for details")
	}
	return result, nil
}

func ConvertInterfaceToJsonString(input map[string]interface{}) (string, error) {
	marshalled, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	shorthand := string(marshalled)
	return shorthand, nil
}
