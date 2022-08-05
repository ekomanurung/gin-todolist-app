package model

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code   int      `json:"code" example:"400"`
	Errors []Errors `json:"errors"`
}

type Errors struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(c *gin.Context, err error) {
	var v validator.ValidationErrors
	if errors.As(err, &v) {
		out := make([]Errors, len(v))
		for i, fe := range v {
			out[i] = Errors{Field: fe.Field(), Message: getErrorMessage(fe)}
		}

		c.JSON(http.StatusBadRequest, Response{
			Code:   http.StatusBadRequest,
			Errors: out,
		})
	} else {
		c.JSON(http.StatusBadRequest,
			Response{
				Code:   http.StatusBadRequest,
				Errors: []Errors{{Message: err.Error()}},
			})
	}
}

//GetErrorMessage add another binding tag here
func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than  " + fe.Param()
	case "gte":
		return "Should be Greater than " + fe.Param()
	default:
		return "Unknown error"
	}
}
