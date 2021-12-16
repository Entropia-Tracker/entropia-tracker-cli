package misc

import (
  "encoding/json"
)

type Event struct {
  Event   string             `json:"event"`
  Date    string             `json:"date"`
  Channel string             `json:"channel"`
  Values  *map[string]string `json:"values"`
}

// String representation of the Event
func (e *Event) String() (string, bool) {
  res, err := e.JSON()
  if err != nil {
    return "", false
  }

  return res, true
}

// JSON encode the Event
func (e *Event) JSON() (result string, err error) {
  s, err := json.Marshal(e)
  result = string(s)
  return
}

func NewEvent(event string) *Event {
  values := make(map[string]string)
  return &Event{Event: event, Values: &values}
}
