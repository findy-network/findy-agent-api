package faker

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/bxcodec/faker/v3"
	"github.com/findy-network/findy-agent-api/tools/data"
	"github.com/lainio/err2"
)

var getConnectionId = func() string {
	return ""
}

func init() {
	_ = faker.AddProvider("eventPairwiseId", func(v reflect.Value) (interface{}, error) {
		return getConnectionId(), nil
	})
}

func FakeEvents(
	count int,
	connections *data.Items,
) (events []data.InternalEvent, err error) {

	getConnectionId = func() string {
		return connections.RandomID()
	}

	events = make([]data.InternalEvent, count)
	for i := 0; i < count; i++ {
		event := data.InternalEvent{}
		err2.Check(faker.FakeData(&event))
		events[i] = event
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].CreatedMs < events[j].CreatedMs
	})
	return
}

func fakeAndPrintEvents(
	count int,
	connections []data.InternalPairwise,
) {
	var err error
	defer err2.Annotate("fakeAndPrintEvents", &err)

	events, err := FakeEvents(count, data.State.Connections)

	fmt.Println("\nvar Events = []InternalEvent{")
	for i := 0; i < len(events); i++ {
		fmt.Printf("	")
		printObject(&events[i], events[i])
	}
	fmt.Println("}")

}
