package helpers

import(
	"github.com/astaxie/beego/validation"
	"fmt"
)
type Html struct{

}

func(this *Html) ShowValidationErrors(v *validation.Validation){
	for _, err := range v.Errors {
		fmt.Println(err.Key + " : " + err.Message)
	}
}