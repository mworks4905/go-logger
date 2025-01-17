package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	record.Offset = len(c.records)
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset int) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if offset >= len(c.records) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

type Record struct {
	Value  []byte `json:"value"`
	Offset int    `json:"offset"`
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
