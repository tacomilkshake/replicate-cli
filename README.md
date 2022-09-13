# replicate-cli
Command line interface for the Replicate API, powered by Go.

### Version 0.1.0 - Experimental Support

Run models and retrieves model version IDs. See TODO for proposed feature list. Currently only supports models that accept a single "image" input. **Note: Stable Diffusion and other text input models are not yet supported, but will be soon. Executable name may change to replicate, from replicate-cli.** 

## Requirements
* Go 1.19
* Replicate API Token, [get yours here](https://replicate.com/account).

## Install, Build & Run

### Install
```
go install github.com/jamiesteven/replicate-cli@latest
```

### Run
```
replicate-cli
```

### Uninstall
```
rm $GOPATH/bin/replicate-cli
```

## How to use

### Authentication
The Replicate API requires an API token. [Get one here.](https://replicate.com/account).

Set a ```REPLICATE_TOKEN``` environment variable, or use ```--token <your key here>```.
```
replicate-cli --token <your key here>
```

### Get model versions
```
replicate-cli versions jingyunliang/swinir
```

### Run a model

**Run using model name.** replicate-cli will use the latest version.
```
replicate-cli run jingyunliang/swinir [input]
```
Input currently supports a fully qualified domain-name.

**Run using version ID.**
```
replicate-cli run jingyunliang/swinir
```

## Development
Pull requests welcome! replicate-cli is built using [Cobra](https://cobra.dev).

---

**Copyright (c) 2022 Jamie Steven. Licensed Under Apache 2.0.**