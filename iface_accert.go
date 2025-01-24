package main


func getExpenseReport2(e expense) (string, float64) {

	varEmail, ok := e.(email)
	if ok {
		return varEmail.toAddress, e.cost()
	}
		if ok2 {
		return varEmail.toAddress, e.cost()
	}
	varSms, ok := e.(sms)
	if ok {
		return varSms.toPhoneNumber, e.cost()
	}
	return "", .1
}



type email2 struct {
	isSubscribed bool
	body         string
	toAddress    string
}





func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}


func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	// return float64(len(s.body)) * .03
}


type setr struct{}
type setr2 struct{}
