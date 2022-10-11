package enums

const (
	ErrorDBDuplicatedKeyMsg                   = "duplicate key"
	ErrorRequestBodyCode                      = "error request parameter"
	ErrorUserNameNotEmptyMsg                  = "error username could not be empty"
	ErrorUserNameNotEmptyCode                 = "error username not empty"
	ErrorUserNameExistsMsg                    = "error username %s already exists"
	ErrorUserNameExistsCode                   = "error username exists"
	ErrorUserNameNotFoundMsg                  = "error username %s data not found"
	ErrorUserNameNotFoundCode                 = "error username data not found"
	ErrorInsertLocationCode                   = "error inserting location"
	ErrorGetDistanceTraveledByUserNameCode    = "error getting distance by username"
	ErrorGetDistanceTraveledByUserNameMsg     = "error getting distance by username %s"
	ErrorGetLastLocationHistoryByUserNameCode = "error getting location history by username"
	ErrorGetLastLocationHistoryByUserNameMsg  = "error getting location history by username %s"
	ErrorUpdateLocationCode                   = "error updating location"
	ErrorGetByLatitudeLongitudeRangeMsg       = "error getting location by latitude longitude range"
	ErrorRadiusNotNegativeMsg                 = "error radius could not be less than 0"
)
