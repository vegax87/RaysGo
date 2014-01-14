package helpers

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

func ShowValidationErrors(v *validation.Validation) {
	for _, err := range v.Errors {
		fmt.Println(err.Key + " : " + err.Message)
	}
}
