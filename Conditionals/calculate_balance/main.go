/*
First, set finalCost to the bulkMessageCost. If the user is premium, then the discountRate should be applied to the finalCost. For example, a discountRate of 0.10 means the discounted price per message would be 90% of the original price.

Next, if the user has enough money in their accountBalance:

Process the payment by deducting the finalCost from their accountBalance
Print the purchaseSuccessMessage
Otherwise, just print the insufficientFundMessage.
*/
package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed .Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful ."
	var accountBalance float64 = 100
	var bulkMessageCost float64 = 75
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	if isPremiumUser {
		finalCost = bulkMessageCost - (bulkMessageCost * discountRate)
	} else {
		finalCost = bulkMessageCost
	}
	if finalCost <= accountBalance {
		accountBalance = accountBalance - finalCost
		fmt.Println(purchaseSuccessMessage)
	} else {
		fmt.Println(insufficientFundMessage)
	}
	fmt.Println("Account balance:", accountBalance)

}
