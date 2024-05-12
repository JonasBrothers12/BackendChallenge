package validations

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
	var sumdigit10,sumdigit11 int = 0,0
	var realdigit10,realdigit11 int
	var cpf_int[11] int
	re := regexp.MustCompile(`^[0-9]{3}\.[0-9]{3}\.[0-9]{3}-[0-9]{2}$`)
	if !re.MatchString(taxid) || len(taxid) != 14{
		return fmt.Errorf("document number invalid, please type 14 characters in the format 000.000.000-00")
	}
	cpf_string := strings.Split(taxid,"")
	c := 0
	for a := 0; a < 14; a++{
		if (a != 3 && a != 7 && a != 11) {
			b ,err := strconv.Atoi(cpf_string[a])
			if err != nil {
				return fmt.Errorf("invalid cpf")
			}else{
				cpf_int[c] = b
				c++
			}
		}
	}
	for i := 0; i < 11; i++ {
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
		return fmt.Errorf("invalid cpf")
	}
	return nil
}





