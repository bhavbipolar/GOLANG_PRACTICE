/*Complete the getMonthlyPrice function. It accepts a tier (string) as input and returns the monthly price for that tier in pennies. Here are the prices in dollars:

"basic" - $100.00
"premium" - $150.00
"enterprise" - $500.00
Convert the prices from dollars to pennies. If the given tier doesn't match any of the above, return 0 pennies.

Note: a dollar consists of 100 pennies.
*/

package main

func getMonthlyPrice(tier string) int {
	// ?
		if tier =="basic"{
			return 10000
		} else if tier =="premium"{
		return 15000
	} else if tier == "enterprise" {
		return 50000
	} else {
		return 0
	}
	
}