package server

import (
	"errors"
	"sync"
)

var ErrOffsetNotFound = errors.New("offset not found")

type Record struct {
	Value  []byte `json:value`
	Offset uint64 `json:offset`
}

type log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *log {
	return &log{}
}

func (c *log) append(record Record) (uint64, error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *log) read(offset uint64) (Record, error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}
