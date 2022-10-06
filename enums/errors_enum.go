package enums

const (
	ErrorDBDuplicatedKeyMsg          = "error_duplicate_key"
	ErrorRequestBodyCode             = "error_request_parameter"
	ErrorUserNameNotEmptyMsg         = "error_username_could_not_be_empty"
	ErrorUserNameNotEmptyCode        = "error_username_not_empty"
	ErrorUserNameExistsMsg           = "error_username_rep_already_exists"
	ErrorUserNameExistsCode          = "error_username_exists"
	ErrorUserNameNotFoundMsg         = "error_username_rep_not_found"
	ErrorUserNameNotFoundCode        = "error_username_not_found"
	ErrorInsertLocationCode          = "error_inserting_location"
	ErrorGetByUserNameCode           = "error_getting_user_by_username"
	ErrorUpdateLocationCode          = "error_updating_location"
	ErrorGetByLatitudeLongitudeRange = "error_getting_location_by_latitude_longitude_range"
)
