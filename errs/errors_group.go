package errs

type HttpErrorGroup struct {
	storage map[string]*HttpError
}

func NewHttpErrorGroup() *HttpErrorGroup {
	return &HttpErrorGroup{
		storage: map[string]*HttpError{},
	}
}

func (e *HttpErrorGroup) AddError(key string, err HttpError) {
	e.storage[key] = &err
}

func (e *HttpErrorGroup) ByKey(key string) *HttpError {
	return e.storage[key]
}

///

var (
	Errors = NewHttpErrorGroup()
)

//
//var (
//	ErrUserNotFound = Errors.ByKey("user_not_found")
//)
//
//func test() error {
//	lang := "kz"
//	var userId int
//
//	//... some logic
//
//	if userId == 0 {
//		return ErrUserNotFound.InLanguage(lang)
//	}
//
//	userId += 1
//	return nil
//}
