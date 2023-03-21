package appqueries

import (
	"cqrs-go/cmd/domain/bus"
	"cqrs-go/cmd/domain/params"
	"time"
)

type getStatementQuery struct {
	data      params.GetStatementQueryParams
	timestamp string
	name      string
}

var GetStatementQueryName = "GetStatementQuery"

func NewGetStatementQuery(data params.GetStatementQueryParams) bus.Query {
	return &getStatementQuery{
		data:      data,
		timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		name:      GetStatementQueryName,
	}
}

func (c *getStatementQuery) Name() string {
	return c.name
}

func (c *getStatementQuery) Type() any {
	return c
}

func (c *getStatementQuery) Data() any {
	return c.data
}

func (c *getStatementQuery) Timestamp() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}
