package main


func canTransitionPayment(currentStatus, newStatus string) bool {
	if currentStatus == "PENDING"   {
		return newStatus == "PROCESSING" || newStatus == "FAILED"
	}
	if currentStatus == "PROCESSING" {
		return newStatus == "SUCCESS" || newStatus == "FAILED"
		
		
	}
	if currentStatus == "SUCCESS" || currentStatus == "FAILED" {
		return false
		
		
	}
	return false

}

func main (){
	currentStatus := "PENDIiNG"
	newStatus := "SUCCESS"

	allowed := canTransitionPayment(currentStatus, newStatus)

	if  !allowed {
	
	}
}