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

# Contribute

There's not much to do but if you find a bug or whatever and feel like fixing it then make a pull request.