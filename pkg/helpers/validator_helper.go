package helpers

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"kv-store/pkg/constants"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("regexp", func(fl validator.FieldLevel) bool {
		pattern := fl.Param()

		re := regexp.MustCompile(pattern)
		return re.MatchString(fl.Field().String())
	})
}

func ValidateStruct(w http.ResponseWriter, s interface{}) {
	err := validate.Struct(s)
	if err != nil {
		var errMsgs []string
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag()))
		}
		SendErrorResponse(w, strings.Join(errMsgs, ", "), constants.BadRequest)
	}
	return
}
