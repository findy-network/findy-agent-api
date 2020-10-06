package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"

	"github.com/bxcodec/faker/v3"
	"github.com/findy-network/findy-agent-api/tools"
	"github.com/lainio/err2"
)

func printObject(objectPtr interface{}, object interface{}) {
	t := reflect.TypeOf(object)
	s := reflect.ValueOf(objectPtr).Elem()
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

type fakeLastName struct {
	Name string `faker:"last_name"`
}

func fakeConnections(count int) (c *[]tools.InternalPairwise, err error) {
	defer err2.Return(&err)
	conns := make([]tools.InternalPairwise, count)
	c = &conns

	err = faker.AddProvider("organisationLabel", func(v reflect.Value) (interface{}, error) {
		orgs := []string{"Bank", "Ltd", "Agency", "Company", "United"}
		index := rand.Intn(len(orgs))
		f := fakeLastName{}
		faker.FakeData(&f)
		return f.Name + " " + orgs[index], nil
	})

	for i := 0; i < count; i++ {
		conn := tools.InternalPairwise{}
		err2.Check(faker.FakeData(&conn))
		conns[i] = conn
	}
	sort.Slice(conns, func(i, j int) bool {
		return conns[i].CreatedMs < conns[j].CreatedMs
	})
	fmt.Println("var Connections = []InternalPairwise{")
	for i := 0; i < len(conns); i++ {
		fmt.Printf("	")
		printObject(&(conns)[i], (conns)[i])
	}
	fmt.Println("}")
	return
}

func fakeEvents(count int) (e *[]tools.InternalEvent, err error) {
	defer err2.Return(&err)
	events := make([]tools.InternalEvent, count)
	e = &events
	for i := 0; i < count; i++ {
		event := tools.InternalEvent{}
		err2.Check(faker.FakeData(&event))
		events[i] = event
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].CreatedMs < events[j].CreatedMs
	})
	fmt.Println("\nvar Events = []InternalEvent{")
	for i := 0; i < len(events); i++ {
		fmt.Printf("	")
		printObject(&events[i], events[i])
	}
	fmt.Println("}")
	return
}

func main() {
	defer err2.Catch(func(err error) {
		fmt.Println("ERROR:", err)
	})

	connCount := 5

	conns, err := fakeConnections(connCount)
	err2.Check(err)

	err = faker.AddProvider("eventPairwiseId", func(v reflect.Value) (interface{}, error) {
		max := len(*conns) - 1
		index := rand.Intn(max)
		return (*conns)[index].ID, nil
	})
	err2.Check(err)

	_, err = fakeEvents(connCount * 10)
	err2.Check(err)
}
