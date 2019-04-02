<p align="center">
    <a href="https://github.com/adrien3d/lumen-cli">
        <img width="500px" src="https://raw.githubusercontent.com/adrien3d/lumen-cli/master/lumen-logo.png" />
    </a>
</p>

<h1 align="center">Lumen CLI</h1>

<p align="center">
    <a href="https://github.com/adrien3d/lumen-api/blob/master/LICENSE.md">
        <img alt="Go Report Card" src="https://img.shields.io/github/license/adrien3d/lumen-api.svg">
    </a>
    <a href="https://goreportcard.com/report/github.com/adrien3d/lumen-cli">
        <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/adrien3d/lumen-api">
    </a>
    <a href="https://godoc.org/github.com/adrien3d/lumen-cli">
        <img alt="latest release" src="https://godoc.org/github.com/adrien3d/lumen-api?status.svg">
    </a>
</p>


Lumen is a CLI that allows you tou generate API with [lumen-api](https://github.com/adrien3d/lumen-api) framework, generating models, controllers, router and store.


### Prerequisites

Define `export PATH=$PATH:$GOPATH/bin`

### Installing

`go get -u github.com/adrien3d/lumen`

`cd $GOPATH/src/github.com/adrien3d/lumen`

`go install`

## Usage

* To generate the project directory, just run `lumen new` with your namespace, for example `lumen new github.com/user/project`
* Generate a first model with `lumen model`

### Auto mode
* Select entities for each selected model and generate controller, store and router with `lumen generate`

### Manual mode
* Then, generate the corresponding controller with `lumen controller`
* Then, generate the corresponding store with `lumen store`
* Finally, generate the router with `lumen router`

## Built With

* [cobra](github.com/spf13/cobra) - A Commander for modern Go CLI interactions.
* [ishell](github.com/abiosoft/ishell) - Library for creating interactive cli applications.
* [lumen-api](github.com/adrien3d/lumen-api) - The boilerplate for GoLang api development

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
* **Adrien Chapelet** - *Initial work* - [IoThings](https://github.com/adrien3d)

