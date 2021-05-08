package mongoutil

type Operator string

const (
	// Operators
	//https://docs.mongodb.com/manual/reference/operator/query/
	//Comparison
	// For comparison of different BSON type values, see the specified BSON comparison order.

	OpEq  Operator = "$eq"  //Matches values that are equal to a specified value.
	OpGt  Operator = "$gt"  //Matches values that are greater than a specified value.
	OpGte Operator = "$gte" //Matches values that are greater than or equal to a specified value.
	OpIn  Operator = "$in"  //Matches any of the values specified in an array.
	OpLt  Operator = "$lt"  //Matches values that are less than a specified value.
	OpLte Operator = "$lte" //Matches values that are less than or equal to a specified value.
	OpNe  Operator = "$ne"  //Matches all values that are not equal to a specified value.
	OpNin Operator = "$nin" //Matches none of the values specified in an array.

	// Logical

	OpAnd Operator = "$and" //Joins query clauses with a logical AND returns all documents that match the conditions of both clauses.
	OpNot Operator = "$not" // Inverts the effect of a query expression and returns documents that do not match the query expression.
	OpNor Operator = "$nor" // Joins query clauses with a logical NOR returns all documents that fail to match both clauses.
	OpOr  Operator = "$or"  // Joins query clauses with a logical OR returns all documents that match the conditions of either clause.

	// Element

	OpExists Operator = "$exists" // Matches documents that have the specified field.
	OpType   Operator = "$type"   // Selects documents if a field is of the specified type.

	// Evaluation

	OpExpr       Operator = "$expr"       // Allows use of aggregation expressions within the query language.
	OpJsonSchema Operator = "$jsonSchema" // Validate documents against the given JSON Schema.
	OpMod        Operator = "$mod"        // Performs a modulo operation on the value of a field and selects documents with a specified result.
	OpRegex      Operator = "$regex"      // Selects documents where values match a specified regular expression.
	OpText       Operator = "$text"       // Performs text search.
	OpWhere      Operator = "$where"      // Matches documents that satisfy a JavaScript expression.

	// Geospatial

	OpGeoIntersects Operator = "$geoIntersects" // Selects geometries that intersect with a GeoJSON geometry. The 2dsphere index supports $geoIntersects.
	OpGeoWithin     Operator = "$geoWithin"     //Selects geometries within a bounding GeoJSON geometry. The 2dsphere and 2d indexes support $geoWithin.
	OpNear          Operator = "$near"          //Returns geospatial objects in proximity to a point. Requires a geospatial index. The 2dsphere and 2d indexes support $near.
	OpNearSphere    Operator = "$nearSphere"    // Returns geospatial objects in proximity to a point on a sphere. Requires a geospatial index. The 2dsphere and 2d indexes support $nearSphere.

	// Array

	OpAll       Operator = "$all"       // Matches arrays that contain all elements specified in the query.
	OpElemMatch Operator = "$elemMatch" // Selects documents if element in the array field matches all the specified $elemMatch conditions.
	OpSize      Operator = "$size"      // Selects documents if the array field is a specified size.

	// Bitwise

	OpBitsAllClear Operator = "$bitsAllClear" // Matches numeric or binary values in which a set of bit positions all have a value of 0.
	OpBitsAllSet   Operator = "$bitsAllSet"   // Matches numeric or binary values in which a set of bit positions all have a value of 1.
	OpBitsAnyClear Operator = "$bitsAnyClear" // Matches numeric or binary values in which any bit from a set of bit positions has a value of 0.
	OpBitsAnySet   Operator = "$bitsAnySet"   // Matches numeric or binary values in which any bit from a set of bit positions has a value of 1.

	// Projection Operators

	OpDollar    Operator = "$"          // Projects the first element in an array that matches the query condition.
	OpElemmatch Operator = "$elemMatch" // Projects the first element in an array that matches the specified $elemMatch condition.
	OpMeta      Operator = "$meta"      // Projects the document's score assigned during $text operation.
	OpSlice     Operator = "$slice"     // Limits the number of elements projected from an array. Supports skip and limit slices.

	// Miscellaneous Operators

	OpComment Operator = "$comment" // Adds a comment to a query predicate.
	OpRand    Operator = "$rand"    // Generates a random float between 0 and 1.
)
