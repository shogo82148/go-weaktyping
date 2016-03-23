package weaktyping_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/shogo82148/go-weaktyping"
)

func ExampleIntUnmarshalJSON() {
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
