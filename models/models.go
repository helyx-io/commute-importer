package models

import(
	"fmt"
	"time"
)

type JSONDate time.Time

func (t JSONDate) MarshalJSON() ([]byte, error) {
	timestamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02"))
	return []byte(timestamp), nil
}
