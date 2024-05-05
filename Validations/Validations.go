package validations

import (
	"fmt"
)


func ValidateName(first string,last string) error{
	if len(first) < 3 {
		return fmt.Errorf("first name too short")
	}
	if len(last) < 3{
		return fmt.Errorf("last name too short")
	}
	return nil
}

func ValidateTaxID(taxid string) error{

	return nil
}

func Checkdigit(cpf string){
	
}



