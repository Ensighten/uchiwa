package daemon

import (
	"testing"

	"github.com/sensu/uchiwa/uchiwa/structs"
	"github.com/stretchr/testify/assert"
)

func TestBuildEvents(t *testing.T) {
	data := structs.Data{
		Events: []interface{}{
			map[string]interface{}{
				"dc": "us-east-1",
				"client": map[string]interface{}{
					"name": "narwhal",
				},
				"check": map[string]interface{}{
					"name": "check-horn",
				},
			},
			map[string]interface{}{
				"dc": "us-east-1",
				"client": map[string]interface{}{
					"name": "triceratops",
				},
				"check": map[string]interface{}{
					"name": "check-horn",
				},
			},
			map[string]interface{}{
				"dc": "us-east-1",
				"client": map[string]interface{}{
					"name": "triceratops",
				},
				"check": map[string]interface{}{
					"name": "check-dinosaur",
				},
			},
			map[string]interface{}{
				"dc": "us-east-1",
				"client": map[string]interface{}{
					"name": "rhino",
				},
				"check": map[string]interface{}{
					"name": "check-horn",
				},
			},
			map[string]interface{}{
				"dc": "us-east-1",
				"client": map[string]interface{}{
					"name": "unicorn",
				},
				"check": map[string]interface{}{
					"name": "check-horn",
				},
			},
			map[string]interface{}{
				"dc": "us-east-1",
				"client": map[string]interface{}{
					"name": "unicorn",
				},
				"check": map[string]interface{}{
					"name": "check-rainbow",
				},
			},
		},
		Silenced: []interface{}{
			map[string]interface{}{
				"dc": "us-east-1",
				"id": "client:unicorn:check-rainbow",
			},
			map[string]interface{}{
				"dc": "us-east-1",
				"id": "client:rhino:*",
			},
			map[string]interface{}{
				"dc": "us-east-1",
				"id": "*:check-dinosaur",
			},
		},
	}
	d := Daemon{Data: &data}
	d.buildEvents()

	event0 := d.Data.Events[0].(map[string]interface{})
	eventClient0 := event0["client"].(map[string]interface{})
	event1 := d.Data.Events[1].(map[string]interface{})
	eventClient1 := event1["client"].(map[string]interface{})
	event2 := d.Data.Events[2].(map[string]interface{})
	eventClient2 := event2["client"].(map[string]interface{})
	event3 := d.Data.Events[3].(map[string]interface{})
	eventClient3 := event3["client"].(map[string]interface{})
	event4 := d.Data.Events[4].(map[string]interface{})
	eventClient4 := event4["client"].(map[string]interface{})
	event5 := d.Data.Events[5].(map[string]interface{})
	eventClient5 := event5["client"].(map[string]interface{})

	// us-east-1/narwhal/check-horn: clear
	assert.Equal(t, event0["_id"], "us-east-1/narwhal/check-horn")
	assert.Equal(t, event0["silenced"], false)
	assert.Equal(t, event0["silenced_by"], []string(nil))
	assert.Equal(t, eventClient0["silenced"], false)

	// us-east-1/triceratops/check-horn: clear
	assert.Equal(t, event1["_id"], "us-east-1/triceratops/check-horn")
	assert.Equal(t, event1["silenced"], false)
	assert.Equal(t, event1["silenced_by"], []string(nil))
	assert.Equal(t, eventClient1["silenced"], false)

	// *:check-dinosaur: silent
	assert.Equal(t, event2["_id"], "us-east-1/triceratops/check-dinosaur")
	assert.Equal(t, event2["silenced"], true)
	assert.Equal(t, event2["silenced_by"], []string{"*:check-dinosaur"})
	assert.Equal(t, eventClient2["silenced"], false)

	// client:rhino:*: silent
	assert.Equal(t, event3["_id"], "us-east-1/rhino/check-horn")
	assert.Equal(t, event3["silenced"], true)
	assert.Equal(t, event3["silenced_by"], []string{"client:rhino:*"})
	assert.Equal(t, eventClient3["silenced"], true)

	// us-east-1/unicorn/check-horn: clear
	assert.Equal(t, event4["_id"], "us-east-1/unicorn/check-horn")
	assert.Equal(t, event4["silenced"], false)
	assert.Equal(t, event4["silenced_by"], []string(nil))
	assert.Equal(t, eventClient4["silenced"], false)

	// client:unicorn:check-rainbow: silent
	assert.Equal(t, event5["_id"], "us-east-1/unicorn/check-rainbow")
	assert.Equal(t, event5["silenced"], true)
	assert.Equal(t, event5["silenced_by"], []string{"client:unicorn:check-rainbow"})
	assert.Equal(t, eventClient5["silenced"], false)
}
