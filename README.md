# ScyllaGo

ScyllaGo is a very simple Go Module for the scylla.sh API
## Getting Started

### Installation 

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

```sh
go get github.com/0xjbb/scyllago
```

### Usage

Import the package into your project.

```go
import "github.com/0xjbb/scyllago"
```

### Examples

Below is a simple example to get you started.

```go
results, err := scyllago.Query("username:jb", 10, 0)

	if err != nil{
		log.Fatal(err)
	}

	for _, values := range results {
		fmt.Printf("%#v\n", values)
	}
```

# Contribute

There's not much to do but if you find a bug or whatever and feel like fixing it then make a pull request.