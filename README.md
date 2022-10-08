# replicate-cli
Command line interface for the [Replicate API](https://replicate.com/docs/reference/http), powered by Go.

Run open source AI models from the command line using [Replicate.com's](https://replicate.com) hardware.

## Features

Version 0.2.2 - Expanded model support + chaining

* Run models
* Chaining: pass the output of one model to another replicate command
* URL or Local file support
* Get model info
* Get model input parameters
* Get model versions

## Requirements
* Go 1.19
* Replicate API Token, [get yours here](https://replicate.com/account).

## Installation

### Install
[go](https://go.dev/dl/) is required.
```
go install github.com/jamiesteven/replicate-cli/cmd/replicate@latest
```

### Run
```
replicate
```

### Uninstall
```
rm $GOPATH/bin/replicate
```

## Usage

### Authentication
The Replicate API requires an API token. [Get one here.](https://replicate.com/account)

Set a ```REPLICATE_TOKEN``` environment variable, or use ```--token <your key here>```.
```
replicate --token <your key here>
```

### Input Parameters
Input parameters use the [shorthand](https://github.com/danielgtaylor/shorthand) format, separated by comma.
```
image:https://picsum.photos/512/512, fidelity:0.4
```

### run: run a model
```
replicate run [model] [inputs]
```

#### Run a model with parameter input
```
replicate run stability-ai/stable-diffusion prompt:photo of a cat
```

#### Run a model using an image URL
```
replicate run jingyunliang/swinir image:https://picsum.photos/512/512
```

#### Run a model using using a local image
```
replicate run jingyunliang/swinir image:/path/to/file.png
```

### Chaining: run multiple AI models in succession
Pipe the output of one replicate command to another:
```
replicate run stability-ai/stable-diffusion prompt:photo of a smiling person \
| replicate run jingyunliang/swinir image: \
| replicate run sczhou/codeformer codeformer_fidelity:0.4, upscale:1, image:
```

### info: get model info
```
replicate info [model]

replicate info stability-ai/stable-diffusion
```

### inputs: get inputs for a specific model
```
replicate inputs [model] [?version]

replicate inputs jingyunliang/swinir

replicate inputs jingyunliang/swinir 9d91795e944f3a585fa83f749617fc75821bea8b323348f39cf84f8fd0cbc2f7
```

### versions: get versions for a specific model
```
replicate versions [model]

replicate versions jingyunliang/swinir
```

## Development
Pull requests welcome!

## Credits

replicate-cli uses the following packages:

* [cobra](https://github.com/spf13/cobra)
* [shorthand](https://github.com/danielgtaylor/shorthand)
* [resty](https://github.com/go-resty/resty)
* [gjson](https://github.com/tidwall/gjson)
* [tablewritter](https://github.com/olekukonko/tablewriter)
* [spinner](https://github.com/briandowns/spinner)

**Copyright (c) 2022 Jamie Steven. Licensed Under Apache 2.0.**