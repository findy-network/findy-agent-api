package tools

import (
	"fmt"
	"strconv"
	"time"

	"github.com/findy-network/findy-agent-api/graph/model"
)

func AddEvent() {
	created := time.Now().Unix()
	createdStr := strconv.FormatInt(created, 10)
	newEvent := InternalEvent{
		ID:           createdStr,
		Description:  "New event " + createdStr,
		ProtocolType: model.ProtocolTypeConnection,
		Type:         model.EventTypeNotification,
		CreatedMs:    created,
	}
	Events = append(Events, newEvent)
	cursor := CreateCursor(newEvent.CreatedMs, model.Event{})
	fmt.Println("Add event " + createdStr)
	for _, observer := range eventAddedObserver {
		observer <- &model.EventEdge{
			Cursor: cursor,
			Node: &model.Event{
				ID:          createdStr,
				Description: newEvent.Description,
				Protocol:    newEvent.ProtocolType,
				Type:        newEvent.Type,
				CreatedMs:   createdStr,
				// TODO: pairwise
			},
		}
	}
}

var Connections = []InternalPairwise{
	{"6dce3142-7e89-4f01-afb6-58a8504514e9", "nUMEIAICXTfpnFWosPpWNsnZX", "xqqchhxDuSmHfkHANjXiBIdag", "http://DGqSavn.net/EQSPkrY.php", "Miss Nayeli West", true, 393728727, 317558948},
	{"3e008a8f-ee44-49c1-964b-3bdfb4959bf3", "fGKUNxuYAVsrsuWgSXjLHqvcd", "EYVqCBUiSYXPolcVYdayEKtND", "https://ZUJkmkU.biz/hvwyFnU.html", "Princess Estrella Pouros", false, 255630516, 444709818},
	{"f2a882f3-30bd-44d6-b7b0-dc0db2fc17f8", "FLMfCjeUgkCfTjFksNqAXEafL", "dXQkLYIxsKoKFSGSmFqwruZdV", "https://pdKCNwj.biz/", "Dr. Edna Cremin", true, 1100900063, 1040768512},
	{"638b2835-fe34-4e43-a873-107a36b982d4", "DduSkPCYCwRLbMiHUhmloDwNs", "PWKJyfUYyhhkclTNsGtTaIjkv", "https://pUBYlen.biz/", "Prof. Laurie Roob", true, 451962065, 1078624229},
	{"474924dc-da50-41ff-ba6e-d71fcc3855f0", "fwYaWroDPqkVWiNUwfgNGHHWi", "MBGGfxrwiDvCEDSolouFDwAww", "https://NjlQnZq.com/", "Lady Ciara Gutkowski", false, 104503498, 1271671482},
}

var Events = []InternalEvent{
	{"14dcdb1e-db90-4b14-8a15-1c987c61eddd", "Sit perferendis consequatur voluptatem aut accusantium.", model.ProtocolTypeConnection, model.EventTypeNotification, "6dce3142-7e89-4f01-afb6-58a8504514e9", 506357851},
	{"47b96e6e-5b39-473e-b18d-dbf9393c8356", "Perferendis sit voluptatem consequatur accusantium aut.", model.ProtocolTypeBasicMessage, model.EventTypeAction, "638b2835-fe34-4e43-a873-107a36b982d4", 623558063},
	{"46293bbe-a991-428a-97d2-7d85a89da508", "Accusantium aut perferendis consequatur sit voluptatem.", model.ProtocolTypeProof, model.EventTypeNotification, "6dce3142-7e89-4f01-afb6-58a8504514e9", 666799773},
	{"b3068fea-b23a-4090-b2f9-463d41bb5890", "Consequatur accusantium aut perferendis voluptatem sit.", model.ProtocolTypeNone, model.EventTypeAction, "3e008a8f-ee44-49c1-964b-3bdfb4959bf3", 1286168468},
	{"3ee9c105-dbca-405f-b702-90b3e93bec29", "Consequatur sit voluptatem aut perferendis accusantium.", model.ProtocolTypeCredential, model.EventTypeNotification, "6dce3142-7e89-4f01-afb6-58a8504514e9", 1432851877},
}
