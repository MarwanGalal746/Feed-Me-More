package errs

import "errors"

var (
	ErrDb                      = errors.New("unexpected database error")
	ErrNoRowsFound             = errors.New("no values found")
	ErrServerErr               = errors.New("internal server error")
	ErrCannotDelete            = errors.New("business Manager can't be deleted")
	ErrCannotUpdate            = errors.New("business Manager can't be upated")
	ErrInvalidPassword         = errors.New("invalid password")
	ErrInvalidToken            = errors.New("invalid password")
	ErrDuplicateValue          = errors.New("this value already exists")
	ErrDrinkNotFound           = errors.New("this drink is not found")
	ErrDuplicateDrinkForUser   = errors.New("the user has this drink already")
	ErrDessertNotFound         = errors.New("this dessert is not found")
	ErrDuplicateDessertForUser = errors.New("the user has this dessert already")
	ErrNoDrinksFound           = errors.New("no drinks found")
	ErrNoDessertsFound         = errors.New("no desserts found")
	ErrEmailMissing            = errors.New("email is missing")
	ErrInvalidEmail            = errors.New("invalid email")
	ErrNoSandwichesFound       = errors.New("no sandwiches found")
	ErrSandwichNotFound        = errors.New("this sandwich is not found")
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewResponse(message string, status int) *Response {
	return &Response{Message: message, Status: status}
}
