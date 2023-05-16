package exception

import "bridging-hl7/utils"

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func SendLogIfErorr(err error, line string) {
	if err != nil {
		utils.SendMessage("LINE " + line + "\n" + " Log Type: Error\n" + "Error: \n" + err.Error() + "\n")
	}
}
