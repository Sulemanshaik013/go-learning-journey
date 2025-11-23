package main

import "fmt"

type Transaction struct {
	amount      float64
	category    string
	description string
	txType      string //income or expense
}

var transactions []Transaction = make([]Transaction, 0)

func addTransaction(amount float64, category, desc, txType string) {
	tx := Transaction{
		amount:      amount,
		category:    category,
		description: desc,
		txType:      txType,
	}
	transactions = append(transactions, tx)
	fmt.Println("Added transction of type", txType)
}

func caculateTotals() (float64, float64, float64) {
	var totalIncome, totalExpense float64
	// for i := 0; i < len(transactions); i++ {
	// 	if transactions[i].txType == "income" {
	// 		totalIncome += transactions[i].amount
	// 	} else if transactions[i].txType == "expense" {
	// 		totalExpense += transactions[i].amount
	// 	}
	// }
	for _, tx := range transactions {
		switch tx.txType {
		case "income":
			totalIncome += tx.amount
		case "expense":
			totalExpense += tx.amount
		default:
		}
	}
	return totalIncome, totalExpense, totalIncome - totalExpense
}

func getCategoryspendings() map[string]float64 {
	categoryMap := make(map[string]float64)
	for _, tx := range transactions {
		if tx.txType == "expense" {
			categoryMap[tx.category] += tx.amount
		}
	}
	return categoryMap
}

func displaySummay(totalIncome, totalExpense, netBalance float64,
	categorySpending map[string]float64) {
	fmt.Println("\nFinancial Summary")

	for idx, tx := range transactions {
		fmt.Printf("  Transaction %d - Type: %s Category: %s Description: %s Amount: $%.3f \n",
			idx+1, tx.txType, tx.category, tx.description, tx.amount)
	}
	fmt.Printf("\nTotal Income %.2f , Total Expenses %.2f , Net Balance %.2f\n",
		totalIncome, totalExpense, netBalance)

	fmt.Println("\nSpending by Category")
	for category, amount := range categorySpending {
		fmt.Printf("  %s: $%.2f\n", category, amount)
	}
}

func findHighestspending(categoryMap map[string]float64) (string, float64) {
	var highestCategory string
	var highestamount float64
	for category, amount := range categoryMap {
		if amount > highestamount {
			highestamount = amount
			highestCategory = category
		}
	}
	return highestCategory, highestamount
}

func checkBudget(totalIncome, totalExpense float64) {
	if totalIncome < totalExpense {
		fmt.Printf("Over spending $%.2f\n", totalExpense-totalIncome)

	} else {
		fmt.Printf("Total Saving $%.2f\n", totalIncome-totalExpense)
	}
}

func spendingPercentage(totalIncome, totalExpense float64) float32 {
	if totalIncome == 0 {
		return 0
	}
	return (float32(totalExpense/totalIncome) * 100)
}

func main() {

	fmt.Println("Budget Tracker Started\n")

	addTransaction(800.00, "salary", "Monthly Salary", "income")
	addTransaction(90.32, "food", "Restaurent", "expense")
	addTransaction(180.00, "entertainment", "Movie Tickets", "expense")
	addTransaction(200.00, "freelance", "side project payment", "income")
	addTransaction(280.76, "utilities", "EB Bill", "expense")
	addTransaction(120.16, "transport", "petrol", "expense")
	addTransaction(193, "food", "Groceries", "expense")

	totalIncome, totalExpenses, netbalance := caculateTotals()
	categorySpending := getCategoryspendings()

	displaySummay(totalIncome, totalExpenses, netbalance, categorySpending)
	highestCategory, Highestamount := findHighestspending(categorySpending)
	fmt.Printf("Highest Spending category : %s - $%.3f\n", highestCategory, Highestamount)
	checkBudget(totalIncome, totalExpenses)

	spendingPercentage := spendingPercentage(totalIncome, totalExpenses)
	fmt.Printf("You Spend %.2f%% of your income\n", spendingPercentage)

}
