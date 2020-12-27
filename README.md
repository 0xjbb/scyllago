# scyllago
Simple Go Module for the scylla.sh API

# Examples

```go
results, err := scyllago.Query("username:jb", 10, 0)

	if err != nil{
		log.Fatal(err)
	}

	for _, values := range results {
		fmt.Printf("%#v\n", values)
	}
```