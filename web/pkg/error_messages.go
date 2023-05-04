package pkg

import (
	"strconv"
	"strings"
)

type ErrorMessageMap map[string]string

var ErrorMessages = ErrorMessageMap{
	"required":   "The {0} field is required.",
	"email":      "The {0} field must be a valid email address.",
	"min":        "The {0} field must be at least {1} characters.",
	"max":        "The {0} field must not exceed {1} characters.",
	"unique":     "The {0} field must be unique.",
	"confirmed":  "The {0} confirmation does not match.",
	"alpha":      "The {0} field may only contain letters.",
	"alpha_num":  "The {0} field may only contain letters and numbers.",
	"alpha_dash": "The {0} field may only contain letters, numbers, dashes and underscores.",
	"numeric":    "The {0} field must be a number.",
	"integer":    "The {0} field must be an integer.",
	"ip":         "The {0} field must be a valid IP address.",
}

func ParseErrorMessage(message string, args ...string) string {
	for i, arg := range args {
		message = strings.Replace(message, "{"+strconv.Itoa(i)+"}", arg, -1)
	}

	return message
}
