package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var x food
	x.preptime = 404
	if order == "burger" {
		x.preptime = 15
	}
	if order == "chips" {
		x.preptime = 10
	}
	if order == "nuggets" {
		x.preptime = 12
	}
	return x.preptime
}
