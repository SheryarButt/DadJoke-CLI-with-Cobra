# DadJoke-CLI-with-Cobra

## Prerequisites

- Make sure you have `Go` installed.
- Make sure you have `$GOBIN` set and added to your `PATH`.

## Downloading Cobra

- Download the latest version of `Cobra` using:
  - `go install github.com/spf13/cobra-cli@latest`
- Run the `cobra-cli` commad to verify installation.

## Usage with Go

### Initialize the CLI

- Create a new `go` project.
  - I'll name mine `DadJoke-CLI-with-Cobra`.

- Initialize the `go mod` using:
  - `go mod init github.com/sheryarbutt/DadJoke-CLI-with-Cobra`

- Initialize the `cobra` command using:
  - `cobra-cli init`

This will create a `main.go` file and a `cmd` directory in the project.

- Navgite to the `cmd` directory and update Long and Short descriptions in `root.go`.

- Verify proper initialization using:
  - `go run main.go`

### Create a new command

- Add new command called `random` using:
  - `cobra-cli add random`
- This will add a new `random.go` file to the `cmd` directory.
- Navgite to the `cmd` directory and update Long and Short descriptions in `random.go`.
- Add code to handle the `random` command in `random.go`.
  - I'm using `icanhazdadjoke.com` to get random dad jokes through their API.

## Runing the application

- Run the CLI application using:
  - `go run main.go`
- Get a random joke using:
  - `go run main.go random`
