package helpers

import (
	"fmt"
	"net/http"
	"strings"

	"kv-store/internal/constants"
	"kv-store/pkg/logger"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("noSQLKeywords", noSQLKeywords)
}

func noSQLKeywords(fl validator.FieldLevel) bool {
	sqlKeywords := []string{"SELECT", "FROM", "WHERE", "DELETE", "UPDATE", "INSERT", "DROP", "CREATE", "ALTER", "TRUNCATE"}

	value := fl.Field().String()
	for _, keyword := range sqlKeywords {
		if strings.Contains(strings.ToUpper(value), keyword) {
			return false
		}
	}
	return true
}

func ValidateStruct(w http.ResponseWriter, s interface{}) {
	log := logger.GetLogger()
	log.Debug().Interface("struct", s).Msg("Validating Request Data")

	err := validate.Struct(s)
	if err != nil {
		log.Error().Err(err).Msg("Invalid Request Data")
		var errMsgs []string
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag()))
		}
		SendErrorResponse(w, strings.Join(errMsgs, ", "), constants.BadRequest)
		return
	}
	return
}
