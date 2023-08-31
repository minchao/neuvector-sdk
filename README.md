# neuvector-sdk

Unofficial Golang SDK for NeuVector API.

## Usage

Use `go get` to install the latest version of the SDK.

```shell
go get -u github.com/minchao/neuvector-sdk
```

### Authentication

The NeuVector API supports two types of authentication: Username/Password and API key.
You can find the example code in [examples/auth_user](examples/auth_user/main.go) and [examples/auth_apikey](examples/auth_apikey/main.go).

## Development

### Generate SDK

Download the latest version of the NeuVector API specification.

```shell
make download-api-spec
```

Generate the SDK.

```bash
make generate
```
