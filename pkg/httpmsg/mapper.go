package httpmsg

import (
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	"net/http"
)

func Error(err error) (message string, code int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)
		msg := re.Message()
		code := mapKindToHttpStatusCode(re.Kind())
		if code >= 500 {
			msg = errmsg.ErrorMsgSomethingWentWrong
		}
		return msg, code

	default:
		return err.Error(), http.StatusBadRequest
	}
}
func mapKindToHttpStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
