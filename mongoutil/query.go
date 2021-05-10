package mongoutil

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindQuery is a convenient holder for a query and commonly used options
// like sort, offset and limit
type FindQuery struct {
	Query  Query
	Sort   interface{}
	Offset int
	Limit  int
}

func order(sortDescending bool) int {
	// Defualt to ascending
	direction := 1
	if sortDescending {
		direction = -1
	}

	return direction
}

func NewFindQuery(q Query, sortField string, sortDescending bool, offset, limit int) FindQuery {
	fq := FindQuery{
		Query:  q,
		Offset: offset,
		Limit:  limit,
	}

	// Sorting
	if sortField != "" {
		direction := order(sortDescending)
		fq.Sort = NewQuery(sortField, direction)
	}

	return fq
}

func (fq FindQuery) FindOptions() *options.FindOptions {
	opts := options.Find()

	if fq.Offset != 0 {
		opts.SetSkip(int64(fq.Offset))
	}
	if fq.Limit != 0 {
		opts.SetLimit(int64(fq.Limit))
	}

	if fq.Sort != nil {
		opts.SetSort(fq.Sort)
	}

	return opts
}

// Query is a MongoDB query implemented using bson.M for ease of manipulation
// Note that bson.M is almost the same as bson.D but does not maintain order
// Normally, in a query, this should not be important
type Query bson.M

type Queries []Query

func NewBlankQuery() Query {
	return Query(bson.M{})
}

func NewQuery(fieldName string, val interface{}) Query {
	return Query(bson.M{fieldName: val})
}

func (q Query) NewFindQuery() FindQuery {
	return NewFindQuery(q, "", false, 0, 0)
}

func (q Query) AddFilter(fieldName string, val interface{}) Query {
	q[fieldName] = val
	return q
}

func (q Query) MergeQuery(incoming Query) Query {
	// Incoming takes priority
	for fieldName, val := range incoming {
		q[fieldName] = val
	}

	return q
}

func (q Query) AddSubQuery(fieldNameA, fieldNameB string, val interface{}) Query {
	q[fieldNameA] = NewQuery(fieldNameA, val)
	return q
}

func (qs Queries) ToQuery() Query {
	nq := NewBlankQuery()

	for _, q := range qs {
		for fieldName, val := range q {
			nq.AddFilter(fieldName, val)
		}
	}

	return nq
}
