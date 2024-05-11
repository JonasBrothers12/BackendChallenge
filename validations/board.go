package validations

import (
	"fmt"
	"strconv"
	//"strings"
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
	/*
	var sumdigit10,sumdigit11 int = 0,0
	var realdigit10,realdigit11 int
	var cpf_int[14] int
	*/
	_ ,err := strconv.Atoi(taxid)
	if len(taxid) != 11 {
		return fmt.Errorf("invalid cpf number of digits")
	}
	if err != nil {
		return fmt.Errorf("invalid cpf digits")
	}
	/*
	cpf_string := strings.Split(taxid,"")
	for i := 0; i < 11; i++ {
		cpf_int[i],_ = strconv.Atoi(cpf_string[i])
		if i < 9{
			sumdigit10 += cpf_int[i]*(10-i)
		}
		if 0 < i && i < 10{
			sumdigit11 += cpf_int[i]*(11-i)			
	    }
	}
	if ((sumdigit10%11 == 0) || (sumdigit10%11 == 1)){
		realdigit10 = 0
	}else{
		realdigit10 = 11-(sumdigit10%11)
	}
	if ((sumdigit11%11 == 0) || (sumdigit11%11 == 1)){
		realdigit11 = 0
	}else{
		realdigit11 = 11-(sumdigit11%11)
	}
	if (realdigit10 != cpf_int[9]) || (realdigit11 != cpf_int[10]) {
		return fmt.Errorf("invalid cpf digits")
	}
	*/
	return nil
}





