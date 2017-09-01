package dto

import (
  "time"
  "encoding/gob"
)

type SensorMessage struct {
  Name string
  Value float64
  TimeStamp time.Time
}


func init() {
  gob.Register(SensorMessage{})
}
