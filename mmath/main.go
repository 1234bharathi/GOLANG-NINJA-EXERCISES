package mmath
func Average(x []float64) float64{
	var avg float64
	for _,no := range x {
		avg +=no 		
	}
	return avg/2
}
