i
func getExpenseReport2(e expense) (string, float64) {

	varEmail, ok := e.(email)
	if ok {
		return varEmail.toAddress, e.cost()
	}

			return float64(len(e.body)) * 1

		return varEmail.toAddress, e.cost()





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


func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}
		return float64(len(e.body)) * .01

		return float64(len(e.body)) * .01

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	
func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
return 1
