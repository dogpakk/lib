package mongoutil

import "go.mongodb.org/mongo-driver/bson"

// FindQuery is a convenient holder for a query and commonly used options
// like sort, offset and limit
type FindQuery struct {
	Query  Query
	Sort   string
	Offset int
	Limit  int
}

func NewFindQuery(q Query, sort string, offset, limit int) FindQuery {
	return FindQuery{
		Query:  q,
		Sory:   sort,
		Offset: offset,
		Limit:  limit,
	}
}

// Query is a MongoDB query implemented using bson.M for ease of manipulation
// Note that bson.M is almost the same as bson.D but does not maintain order
// Normally, in a query, this should not be important
type Query bson.M

func NewQuery() Query {
	return Query(bson.M{})
}

type Filter struct {
	FieldName string
	Val       interface{}
}

func (f Filter) toQuery() Query {
	return Query(bson.M{f.FieldName: f.Val})
}

func (q Query) AddFilter(f Filter) {
	q[f.FieldName] = f.Val
}
