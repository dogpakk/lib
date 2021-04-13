package mongolist

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	// standard date format
	dateFormat = "2006-01-02T15:04:05.000Z"

	// mongo keywords
	mongoAnd   = "$and"
	mongoOr    = "$or"
	mongoMatch = "$match"
	mongoSkip  = "$skip"
	mongoLimit = "$limit"
	mongoSort  = "$sort"

	// filter operators
	filterOperatorEq              = "eq"
	filterOperatorEqOrNull        = "eqornull"
	filterOperatorNeq             = "neq"
	filterOperatorGt              = "gt"
	filterOperatorGte             = "gte"
	filterOperatorLt              = "lt"
	filterOperatorLte             = "lte"
	filterOperatorStartsWith      = "starts"
	filterOperatorContains        = "contains"
	filterOperatorEndsWith        = "ends"
	filterOperatorBefore          = "before"
	filterOperatorOnOrBefore      = "onorbefore"
	filterOperatorAfter           = "after"
	filterOperatorOnOrAfter       = "onorafter"
	filterOperatorEntityNullCheck = "entityNullCheck"
)

type Filter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type ListState struct {
	Filters          []Filter `json:"filters"`
	FilterCombineOr  bool     `json:"filterCombineOr"`
	Order            string   `json:"order"`
	OrderDescending  bool     `json:"orderDescending"`
	Limit            int      `json:"limit"`
	Offset           int      `json:"offset"`
	IncludeInactives bool     `json:"includeInactives"`
}

func BodyToPipelines(listState ListState) (bson.A, bson.A, error) {
	// Filtering first
	var mongoFilters []bson.M

	for _, filter := range listState.Filters {
		mongoFilter, err := createFilter(filter.Field, filter.Operator, filter.Value)
		if err != nil {
			return bson.A{}, bson.A{}, err
		} else {
			mongoFilters = append(mongoFilters, bson.M{
				filter.Field: mongoFilter,
			})
		}
	}

	filterCombine := mongoAnd
	if listState.FilterCombineOr {
		filterCombine = mongoOr
	}

	inactiveFilter := bson.M{}
	if !listState.IncludeInactives {
		inactiveFilter = bson.M{"inactive": bson.M{"$in": bson.A{false, nil}}}
	}

	basePipeline := bson.A{inactiveFilter}
	if len(mongoFilters) > 0 {
		basePipeline = append(basePipeline, bson.M{filterCombine: mongoFilters})
	}

	pipeline := bson.A{bson.M{mongoMatch: bson.M{"$and": basePipeline}}}

	// Sorting
	if listState.Order != "" {
		direction := order(listState.OrderDescending)
		pipeline = append(pipeline, bson.M{
			mongoSort: bson.M{
				listState.Order: direction,
			}})
	}

	// Order of offset and limit is important.  Offset first!!
	// Cache the 'noLimitPipeline' for counting later.
	noLimitPipeline := pipeline

	// Offset
	if listState.Offset > 0 {
		pipeline = append(pipeline, intAgg(listState.Offset, mongoSkip))
	}

	// Limit
	if listState.Limit > 0 {
		pipeline = append(pipeline, intAgg(listState.Limit, mongoLimit))
	}

	return pipeline, noLimitPipeline, nil
}

func intAgg(val int, mongoKey string) bson.M {
	return bson.M{
		mongoKey: val,
	}
}

func order(sortDescending bool) int {
	// Defualt to ascending
	direction := 1
	if sortDescending {
		direction = -1
	}

	return direction
}

func createFilter(field, operator string, value interface{}) (interface{}, error) {
	switch operator {
	case filterOperatorStartsWith:
		return bson.D{
			{"$regex", fmt.Sprintf("^%s", value)},
			{"$options", "i"},
		}, nil

	case filterOperatorContains:
		return bson.D{
			{"$regex", fmt.Sprintf("%s", value)},
			{"$options", "i"},
		}, nil
	case filterOperatorEndsWith:
		return bson.D{
			{"$regex", fmt.Sprintf("%s$", value)},
			{"$options", "i"},
		}, nil
	case filterOperatorNeq:
		return bson.M{
			"$ne": value,
		}, nil
	case filterOperatorEqOrNull:
		eqFilter, _ := createFilter(field, "eq", value)
		return bson.M{
			"$in": bson.A{eqFilter, nil, primitive.ObjectID{}},
		}, nil
	case filterOperatorGt:
		return bson.M{
			"$gt": value,
		}, nil

	case filterOperatorGte:
		return bson.M{
			"$gte": value,
		}, nil
	case filterOperatorLt:
		return bson.M{
			"$lt": value,
		}, nil
	case filterOperatorLte:
		return bson.M{
			"$lte": value,
		}, nil
	case filterOperatorBefore:
		t, err := time.Parse(dateFormat, value.(string))
		if err != nil {
			return bson.M{}, err
		}

		return bson.M{
			"$lt": t,
		}, nil
	case filterOperatorOnOrBefore:
		t, err := time.Parse(dateFormat, value.(string))
		if err != nil {
			return bson.M{}, err
		}

		return bson.M{
			"$lte": t,
		}, nil

	case filterOperatorAfter:
		t, err := time.Parse(dateFormat, value.(string))
		if err != nil {
			return bson.M{}, err
		}

		return bson.M{
			"$gt": t,
		}, nil

	case filterOperatorOnOrAfter:
		t, err := time.Parse(dateFormat, value.(string))
		if err != nil {
			return bson.M{}, err
		}

		return bson.M{
			"$gte": t,
		}, nil

	case filterOperatorEntityNullCheck:
		b := value.(bool)
		if !b {
			return bson.M{"$in": bson.A{primitive.ObjectID{}, nil}}, nil
		}
		return bson.M{"$gt": primitive.ObjectID{}}, nil

	default:
		// Bools are a special case, because fields are not always there in Mongo
		// when you say featured=true, that's what you mean, so no modification is needed
		// but when you say featured=false you actually mean featured = false OR null (missing)
		if b, isBool := value.(bool); isBool && !b {
			return bson.M{
				"$in": bson.A{false, nil},
			}, nil
		}

		// Try object ID
		if s, ok := value.(string); ok {
			if id, err := primitive.ObjectIDFromHex(s); err == nil {
				return id, nil
			}
		}
	}

	return value, nil
}
