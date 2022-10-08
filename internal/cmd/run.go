package cmd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jamiesteven/replicate-cli/internal/pkg/files"
	"github.com/jamiesteven/replicate-cli/internal/pkg/input"
	"github.com/jamiesteven/replicate-cli/internal/pkg/models"
	"github.com/jamiesteven/replicate-cli/internal/pkg/output"
	"github.com/jamiesteven/replicate-cli/internal/pkg/predictions"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var Async bool

var runCmd = &cobra.Command{
	Use:   "run [model/version] [input]",
	Short: "Run a model on replicate.com",
	Long: `Run a model on replicate.com.
	
For model/version, provide a model name, or a specific version id.

For input, provide input parameters in Shorthand JSON format, separated by comma. Examples below:
	replicate run stability-ai/stable-diffusion prompt:photo of a taco milkshake, width:768, height:768
	replicate run jingyunliang/swinir, image:https://picsum.photos/512/512
	replicate run sczhou/codeformer, image:/path/to/image.jpg

Images can be inputted as a URL or a local path. If a local path is provided, the image will be uploaded to replicate.com.

Chain commands together:
	replicate run stability-ai/stable-diffusion prompt:photo of a smiling person | replicate run sczhou/codeformer codeformer_fidelity:0.6, image:`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Improve error check with better feedback
		if len(args) <= 1 {
			return errors.New("must specify two arguments: 'replicate-cli run [model/version] [input]'")
		}

		// Start spinner
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.FinalMSG = ""
		s.Start()

		// Append piped args if they exist
		pipedArgs, err := input.GetPipedArgs()
		if err != nil {
			return err
		}
		args = append(args, pipedArgs)

		// If there are multiple args and the second to last one ends with a colon, concatenate the last two args, joining the piped input with the last arg
		// This allows you to pipe the output of one replicate command to another
		if strings.HasSuffix(args[len(args)-1], ":") {
			// Combine the last two args, in the 2nd to last position
			args[len(args)-2] = args[len(args)-2] + args[len(args)-1]

			// Delete the final, now redundant, arg
			args = args[:len(args)-1]
		}

		version, err := models.GetVersionId(args[0], TokenKey)
		if err != nil {
			return err
		}

		// Convert shorthand input to JSON map interface
		inputI, err := input.ConvertShorthandToInterface(args[1:])
		if err != nil {
			// TODO: Create error that shows formatting of inputs in shorthand
			return err
		}

		// Check for image input, and if so, evaluate if URL or local path, setting image field accordingly
		if val, ok := inputI["image"].(string); ok {
			image, err := files.GetFile(val)
			if err != nil {
				return err
			}
			inputI["image"] = image
		}

		// Convert (possibly) updated JSON map interface to JSON string
		inputS, err := input.ConvertInterfaceToJsonString(inputI)
		if err != nil {
			return err
		}

		// Prepare body JSON as string, including inputS
		bodyS := `{"version":"` + version + `","input":` + inputS + `}`

		// Start prediction
		predictionId, err := predictions.StartPrediction(bodyS, TokenKey)
		if err != nil {
			return err
		}

		// If async boolean flag is set, print prediction URL and exit
		if Async {
			// Stop spinner
			s.Stop()

			// Print prediction URL
			fmt.Println("https://replicate.com/p/" + predictionId)
			return nil
		}

		// Otherwise, loop every 2 seconds while waiting for prediction to complete
		for {
			// Sleep for 2 seconds
			time.Sleep(2 * time.Second)

			// Check prediction status
			response, err := predictions.CheckPrediction(predictionId, TokenKey)
			if err != nil {
				return err
			}
			status := gjson.Get(response, "status").String()

			if status == "succeeded" {
				s.Stop()
				outputString := output.FormatOutput(response)
				fmt.Println(outputString)
				break
			} else if status == "failed" {
				return errors.New("prediction failed")
			}
		}

		// Return no error
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Async, "async", "v", false, "run model asynchronously")
	rootCmd.AddCommand(runCmd)
}
