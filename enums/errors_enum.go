package enums

const (
	ErrorDBDuplicatedKeyMsg          = "duplicate key"
	ErrorRequestBodyCode             = "error request parameter"
	ErrorUserNameNotEmptyMsg         = "error username could not be empty"
	ErrorUserNameNotEmptyCode        = "error username not empty"
	ErrorUserNameExistsMsg           = "error username already exists %s"
	ErrorUserNameExistsCode          = "error username exists"
	ErrorUserNameNotFoundMsg         = "error username not found %s"
	ErrorUserNameNotFoundCode        = "error username not found"
	ErrorInsertLocationCode          = "error inserting location"
	ErrorGetByUserNameCode           = "error getting user by username"
	ErrorUpdateLocationCode          = "error updating location"
	ErrorGetByLatitudeLongitudeRange = "error getting location by latitude longitude range"
)
