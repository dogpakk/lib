package mongoutil

const (
	UpdateOpCurrentDate = "$currentDate" //Sets the value of a field to current date, either as a Date or a Timestamp.
	UpdateOpInc         = "$inc"         //Increments the value of the field by the specified amount.
	UpdateOpMin         = "$min"         //Only updates the field if the specified value is less than the existing field value.
	UpdateOpMax         = "$max"         //Only updates the field if the specified value is greater than the existing field value.
	UpdateOpMul         = "$mul"         //Multiplies the value of the field by the specified amount.
	UpdateOpRename      = "$rename"      //Renames a field.
	UpdateOpSet         = "$set"         //Sets the value of a field in a document.
	UpdateOpSetOnInsert = "$setOnInsert" //Sets the value of a field if an update results in an insert of a document. Has no effect on update operations that modify existing documents.
	UpdateOpUnset       = "$unset"       //Removes the specified field from a document.

	// Array

	UpdateOpDollar       = "$"               //Acts as a placeholder to update the first element that matches the query condition.
	UpdateOpArrayAll     = "$[]"             //Acts as a placeholder to update all elements in an array for the documents that match the query condition.
	UpdateOpArrayFilters = "$[<identifier>]" //Acts as a placeholder to update all elements that match the arrayFilters condition for the documents that match the query condition.
	UpdateOpAddToSet     = "$addToSet"       //Adds elements to an array only if they do not already exist in the set.
	UpdateOpPop          = "$pop"            //Removes the first or last item of an array.
	UpdateOpPull         = "$pull"           //Removes all array elements that match a specified query.
	UpdateOpPush         = "$push"           //Adds an item to an array.
	UpdateOpPullAll      = "$pullAll"        //Removes all matching values from an array.

	// Modifiers
	UpdateOpEach     = "$each"     //Modifies the $push and $addToSet operators to append multiple items for array updates.
	UpdateOpPosition = "$position" //Modifies the $push operator to specify the position in the array to add elements.
	UpdateOpSlice    = "$slice"    //Modifies the $push operator to limit the size of updated arrays.
	UpdateOpSort     = "$sort"     //Modifies the $push operator to reorder documents stored in an array.

	// Bitwise
	UpdateOpBit = "$bit" //Performs bitwise AND, OR, and XOR updates of integer values.
)
