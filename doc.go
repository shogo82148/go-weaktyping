/*
go-weaktyping is a go package for unmashaling weakly-typed-JSON.

Here is an example of weakly-typed-JSON.
  {
    "integer1": "123",
    "integer2": 456
  }

If you want to umarshal this JSON into following Go struct, "encoding/json".Unmarshal will fail.
  ptr := &struct {
    Integer1 int
    Integer2 int
  }()
  json.Unmarshal(in, ptr) // will fail

You can use weaktyping.Int instead of int for unmarshaling this JSON.
  ptr := &struct {
    Integer1 weaktyping.Int
    Integer2 weaktyping.Int
  }()
  json.Unmarshal(in, ptr) // will succeed
*/
package weaktyping
