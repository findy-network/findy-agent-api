package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/bxcodec/faker/v3"
	"github.com/findy-network/findy-agent-api/tools"
	"github.com/lainio/err2"
)

func printObject(s reflect.Value, t reflect.Type) {
	fmt.Printf("{")
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if !strings.HasPrefix(t.Field(i).Name, "Skip") {

			if i != 0 {
				fmt.Printf(",")
			}
			if f.Type().String() == "string" {
				fmt.Printf("\"%s\"", f.Interface())
			} else if f.Type().String() == "int64" {
				fmt.Printf("%d", f.Interface())
			} else if f.Type().String() == "bool" {
				fmt.Printf("%t", f.Interface())
			} else {
				fmt.Printf("%s", f.Interface())
			}

		}
	}
	fmt.Println("},")
}

func fakeConnections(count int) (c *[]tools.InternalPairwise, err error) {
	defer err2.Return(&err)
	conns := make([]tools.InternalPairwise, count)
	c = &conns
	for i := 0; i < count; i++ {
		conn := tools.InternalPairwise{}
		err2.Check(faker.FakeData(&conn))
		conns[i] = conn
	}
	return
}

func main() {
	defer err2.Catch(func(err error) {
		fmt.Println("ERROR:", err)
	})

	c, err := fakeConnections(5)
	err2.Check(err)

	fmt.Println("var connections = []tools.InternalPairwise{")
	for i := 0; i < len(*c); i++ {
		fmt.Printf("	")
		printObject(reflect.ValueOf(&(*c)[i]).Elem(), reflect.TypeOf(tools.InternalPairwise{}))
	}
	fmt.Println("}")
}
