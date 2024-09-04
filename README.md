# options

Functional Options Pattern utility for Go.

## Usage

```sh
go get github.com/tingtt/options
```

```go
import "github.com/tingtt/options"
```

- Simple one
	```go
	type Option struct {}

	func NewService(_options ...options.Applier[Option]) (Service, error) {
		option := options.Create(_options...) // Create option from functional options.

		service := new(service)

		// effect option to service ...

		return service, nil
	}
	```
- Create option with default value
	```go
	type Option struct {}

	func DefaultOption() Option {
		return Option{
			// write default fields...
		}
	}

	func NewService(_options ...options.Applier[Option]) (Service, error) {
		option := options.CreateWithDefault(DefaultOption(), _options...) // Create option from functional options.

		service := new(service)

		// effect option to service ...

		return service, nil
	}
	```

