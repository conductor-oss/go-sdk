package integration_tests

import (
	"context"
	"testing"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestEventHandlerCreate(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	ctx := context.Background()

	uuid := uuid.New().String()
	eventHandlerName := "TEST_GO_EVENT_HANDLER_" + uuid
	eventName := "TEST_GO_EVENT_" + uuid

	initEventHandler := model.EventHandler{
		Name:  eventHandlerName,
		Event: eventName,
		Actions: []model.Action{
			{
				Action: "start_workflow",
			},
		},
		Tags: []model.Tag{
			{
				Key:   "TestTag",
				Type:  "TestType",
				Value: "TestValue",
			},
		},
	}
	_, err := testdata.EventHandlerClient.AddEventHandler(ctx, initEventHandler)
	require.NoError(t, err)

	events, _, err := testdata.EventHandlerClient.GetEventHandlersForEvent(ctx, eventName, &client.EventResourceApiGetEventHandlersForEventOpts{
		ActiveOnly: optional.NewBool(false),
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(events))

	// Check received event handler
	receivedEventHandler := events[0]

	require.Equal(t, eventHandlerName, receivedEventHandler.Name)
	require.Equal(t, eventName, receivedEventHandler.Event)

	// Check received actions
	require.Equal(t, len(receivedEventHandler.Actions), len(initEventHandler.Actions))
	receivedAction := receivedEventHandler.Actions[0]
	require.Equal(t, receivedAction.Action, initEventHandler.Actions[0].Action)

	// Check received tags
	require.Equal(t, len(receivedEventHandler.Tags), len(initEventHandler.Tags))
	receivedTag := receivedEventHandler.Tags[0]
	require.Equal(t, receivedTag.Key, initEventHandler.Tags[0].Key)
	require.Equal(t, receivedTag.Type, initEventHandler.Tags[0].Type)
	require.Equal(t, receivedTag.Value, initEventHandler.Tags[0].Value)

	t.Cleanup(func() {
		_, err := testdata.EventHandlerClient.RemoveEventHandler(ctx, eventHandlerName)
		require.NoError(t, err)
	})
}
