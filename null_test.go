package null_test

import (
	"encoding/json"
	"github.com/go-digo/null"
	"log"
	"testing"
)

func TestNull_New(t *testing.T) {
	a := null.New[string]("da")

	b := null.Nullable[string]()

	c := null.Undefined[string]()

	log.Println(a, b, c)
}

func TestNull_Struct_New(t *testing.T) {
	type Data struct {
		Name string
	}

	a := null.New[Data](Data{
		Name: "as",
	})

	b := null.Nullable[Data]()

	c := null.Undefined[Data]()

	log.Println(a, b, c)
}

func TestNull_Struct_New1(t *testing.T) {
	c := null.New[string]("da")
	c.Get()
}

func BenchmarkNew(b *testing.B) {
	c := null.Undefined[string]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Get()
	}
}

func TestNull_MarshalJSON(t *testing.T) {
	a := null.New[string]("da")
	jsona, err := a.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}

	b := null.Nullable[string]()
	jsonb, err := b.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}

	c := null.Undefined[string]()
	jsonv, err := c.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(string(jsona), string(jsonb), string(jsonv))
}

func TestNull_UnmarshalJSON(t *testing.T) {
	type A struct {
		Name null.Null[string] `json:"name,omitempty"`
	}
	c := null.Undefined[string]()
	a := A{Name: c}

	marshal, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
		return
	}

	log.Println(string(marshal))
}

func TestName(t *testing.T) {
	type A struct {
		Name null.Null[string] `json:"name,omitempty"`
	}
	var a A
	//err := json.Unmarshal([]byte(`{"name":"dada"}`), &a)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	a.Name.Set("dada")

	log.Println(a)
}
