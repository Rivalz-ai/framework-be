package event

import (
	//"time"
	"context"
	//"github.com/google/uuid"
	e "github.com/Rivalz-ai/framework-be/framework/base/error"
)

type ConsumeFn = func(e Event, ctx context.Context) error
type LogItemSuccessFn = func(e Event) error

// kafka, consider remove bellow
type WriteLogConsumeFn = func(e Event) error
type RePushFn = func(event Event, ctx context.Context) *e.Error

//event internal for workstream
/*
type Event struct {
	EventHeader            interface{} `json:"event_header,omitempty"` // ClientID, UserID
	EventID                uuid.UUID   `json:"event_id,omitempty"`
	EventName              string      `json:"event_name,omitempty" ` //dev set
	EventData              interface{} `json:"event_data,omitempty" ` //dev set
	Uid                    string      `json:"uid,omitempty" `        //dev set
	SourceID               string      `json:"source_id,omitempty"`
	WorkerID               int         `json:"worker_id,omitempty"`
	Flow                   *Queue      `json:"flow,omitempty" `
	Counter                int         `json:"counter,omitempty"`
	PushlishTime           time.Time   `json:"publish_time,omitempty"`
	ConsumeTime            time.Time   `json:"consume_time,omitempty"`
	FinishTime             time.Time   `json:"finish_time,omitempty"`
	ProcessingTime         float64     `json:"processing_time,omitempty"`
	Logs                   string      `json:"logs,omitempty"`
	ProcessedFlow          string      `json:"processed_flow,omitempty"`
	IgnoreUid              bool        `json:"ignore_uid,omitempty"`
	Transaction_id         uuid.UUID   `json:"transaction_id,omitempty"`
	Transaction_start_time int64       `json:"transaction_start_time,omitempty"` //timestamp
	IgnoreLogItemSuccess   bool        `json:"ignore_log_item_success,omitempty"`
}
*/
type Event struct {
	EventHeader interface{} `json:"event_header,omitempty"` // ClientID, UserID => for Websocket only
	EventName   string      `json:"event_name,omitempty" `  //dev set
	EventData   interface{} `json:"event_data,omitempty" `  //dev set
	Uid         string      `json:"uid,omitempty" `         //dev set
}
type Queue []string

func (self *Queue) Push(x string) {
	*self = append(*self, x)
}

func (self *Queue) Pop() string {
	h := *self
	var el string
	l := len(h)
	el, *self = h[0], h[1:l]
	return el
}

func NewQueue() *Queue {
	return &Queue{}
}
