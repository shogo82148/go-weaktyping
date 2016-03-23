# go-weaktyping
go-weaktyping is a go package for unmashaling weakly-typed-JSON.

Here is an example of weakly-typed-JSON.

``` json
{
  "integer1": "123",
  "integer2": 456
}
```

If you want to umarshal this JSON into following Go struct, "encoding/json".Unmarshal will fail.

``` go
ptr := &struct {
  Integer1 int
  Integer2 int
}()
json.Unmarshal(in, ptr) // will fail
```

You can use weaktyping.Int instead of int for unmarshaling this JSON.

``` go
ptr := &struct {
  Integer1 weaktyping.Int
  Integer2 weaktyping.Int
}()
json.Unmarshal(in, ptr) // will succeed
```

See [godoc](https://godoc.org/github.com/shogo82148/go-weaktyping) for more detail.

## EXAMPLES

``` go
func ExampleInt_UnmarshalJSON() {
	ptr := &struct {
		Foo weaktyping.Int `json:"foo"`
	}{}

	if err := json.Unmarshal([]byte(`{"foo":123}`), ptr); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Foo:", ptr.Foo)
	if err := json.Unmarshal([]byte(`{"foo":"456"}`), ptr); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Foo:", ptr.Foo)

	// Output:
	// Foo: 123
	// Foo: 456
}
```

## LICENSE

This software is released under the MIT License, see LICENSE.md.
