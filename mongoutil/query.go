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
		fq.Sort = NewFilterQuery(sortField, direction)
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

func NewQuery() Query {
	return Query(bson.M{})
}

func NewFilterQuery(fieldName string, val interface{}) Query {
	return NewFilter(fieldName, val).ToQuery()
}

type Filter struct {
	FieldName string
	Val       interface{}
}

func NewFilter(fieldName string, val interface{}) Filter {
	return Filter{
		FieldName: fieldName,
		Val:       val,
	}
}

type Filters []Filter

func (f Filter) ToQuery() Query {
	return Query(bson.M{f.FieldName: f.Val})
}

func (q Query) AddFilter(f Filter) {
	q[f.FieldName] = f.Val
}

func (fs Filters) ToQuery() Query {
	q := NewQuery()

	for _, f := range fs {
		q.AddFilter(f)
	}

	return q
}

func (fs Filters) ToQueryArray() (res []Query) {
	for _, f := range fs {
		res = append(res, f.ToQuery())
	}

	return res
}
