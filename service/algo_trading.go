package service

import (
	"database/sql"
	"fmt"
	"log"

)

// User represents a row in the Users table
type AlgoUser struct {
	ID                   int
	InitialInvestmentAmt float64
	InitialInvestmentTm  string
}

// Algorithm represents a row in the Algorithms table
type Algorithm struct {
	ID          int
	Name        string
	Description string
	Parameters  string
}

// UserAlgorithmBudget represents a row in the User_Algorithm_Budgets table
type UserAlgorithmBudget struct {
	ID        int
	UserID    int
	Algorithm int
	BudgetAmt float64
}

// Trade represents a row in the Trades table
type Trade struct {
	ID            int
	UserID        int
	AlgorithmID   int
	ExecutionPrice float64
	Amount        float64
	Time          string
	TradeType     string
}

// Investment represents a row in the Investments table
type Investment struct {
	ID         int
	UserID     int
	AlgorithmID int
	CurrentVal  float64
	StartTime   string
}

// InvestmentLimit represents a row in the Investment_Limits table
type InvestmentLimit struct {
	ID                   int
	UserID               int
	AlgorithmID          int
	TotalAmtAllowed      float64
	TotalAmtInvested     float64
	RemainingAmtToInvest float64
}

// AlgorithmPerformance represents a row in the Algorithm_Performance table
type AlgorithmPerformance struct {
	ID            int
	AlgorithmID   int
	TotalReturn   float64
	TotalTrades   int
	ProfitLoss    float64
	PLPercentage  float64
}

// UserPerformance represents a row in the User_Performance table
type UserPerformance struct {
	ID            int
	UserID        int
	AlgorithmID   int
	TotalReturn   float64
	TotalTrades   int
	ProfitLoss    float64
	PLPercentage float64
}

// TradeSummary represents a row in the Trades_Summary table
type TradeSummary struct {
ID int
UserID int
AlgorithmID int
AvgExecutionPrice float64
TotalAmount float64
ProfitLoss float64
TradesCount int
}

func main() {
	// Connect to the database
	db, err := sql.Open("postgres", "user=username password=password dbname=trading_db sslmode=disable")
	if err != nil {
	log.Fatal(err)
	}
	defer db.Close()

	// Insert a new user into the Users table
	newUser := &AlgoUser{
		ID:                   1,
		InitialInvestmentAmt: 1000.00,
		InitialInvestmentTm:  "2022-01-01 12:00:00",
	}
	insertUserStmt, err := db.Prepare("INSERT INTO Users (user_id, initial_investment_amount, initial_investment_time) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = insertUserStmt.Exec(newUser.ID, newUser.InitialInvestmentAmt, newUser.InitialInvestmentTm)
	if err != nil {
		log.Fatal(err)
	}

	// Insert a new algorithm into the Algorithms table
	newAlgorithm := &Algorithm{
		ID:          1,
		Name:        "Trend Following Algorithm",
		Description: "An algorithm that follows trends in the market.",
		Parameters:  "{'lookback_period': 10}",
	}
	insertAlgorithmStmt, err := db.Prepare("INSERT INTO Algorithms (algorithm_id, algorithm_name, algorithm_description, algorithm_parameters) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = insertAlgorithmStmt.Exec(newAlgorithm.ID, newAlgorithm.Name, newAlgorithm.Description, newAlgorithm.Parameters)
	if err != nil {
		log.Fatal(err)
	}

	// Insert a new budget for a user and algorithm into the User_Algorithm_Budgets table
	newBudget := &UserAlgorithmBudget{
		ID:        1,
		UserID:    1,
		Algorithm: 1,
		BudgetAmt: 1000.00,
	}
	insertBudgetStmt, err := db.Prepare("INSERT INTO User_Algorithm_Budgets (budget_id, user_id, algorithm_id, budget_amount) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = insertBudgetStmt.Exec(newBudget.ID, newBudget.UserID, newBudget.Algorithm, newBudget.BudgetAmt)
	if err != nil {
		log.Fatal(err)
	}

	// Insert a new trade into the Trades table
	newTrade := &Trade{
		ID:            1,
		UserID:        1,
		AlgorithmID:   1,
		ExecutionPrice: 100.00,
		Amount:        10.00,
		Time:          "2022-01-01 12:00:00",
		TradeType:     "buy",
	}
	insertTradeStmt, err := db.Prepare("INSERT INTO Trades (trade_id, user_id, algorithm_id, execution_price, amount, time, trade_type) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = insertTradeStmt.Exec(newTrade.ID, newTrade.UserID, newTrade.AlgorithmID, newTrade.ExecutionPrice, newTrade.Amount, newTrade.Time, newTrade.TradeType)
	if err != nil {
		log.Fatal(err)
	}

	// Insert a new investment into the Investments table
	newInvestment := &Investment{
		ID:         1,
		UserID:     1,
		AlgorithmID: 1,
		CurrentVal:  1000.00,
		StartTime:   "2022-01-01 12:00:00",
	}
	insertInvestmentStmt, err := db.Prepare("INSERT INTO Investments (investment_id, user_id, algorithm_id, current_value, start_time) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = insertInvestmentStmt.Exec(newInvestment.ID, newInvestment.UserID, newInvestment.AlgorithmID, newInvestment.CurrentVal, newInvestment.StartTime)
	if err != nil {
		log.Fatal(err)
	}

	// Insert a new investment limit into the Investment_Limits table
	newLimit := &InvestmentLimit{
		ID:                   1,
		UserID:               1,
		AlgorithmID:          1,
		TotalAmtAllowed:      1000.00,
		TotalAmtInvested:     500.00,
		RemainingAmtToInvest: 500.00,
	}
	insertLimitStmt, err := db.Prepare("INSERT INTO Investment_Limits (limit_id, user_id, algorithm_id, total_amount_allowed, total_amount_invested, remaining_amount) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = insertLimitStmt.Exec(newLimit.ID, newLimit.UserID, newLimit.AlgorithmID, newLimit.TotalAmtAllowed, newLimit.TotalAmtInvested, newLimit.RemainingAmtToInvest)
	if err != nil {
	log.Fatal(err)
	}
	
	// Insert a new algorithm performance into the Algorithm_Performance table
	newAlgorithmPerformance := &AlgorithmPerformance{
		ID:            1,
		AlgorithmID:   1,
		TotalReturn:   1000.00,
		TotalTrades:   100,
		ProfitLoss:    500.00,
		PLPercentage:  0.5,
	}
	insertAlgorithmPerformanceStmt, err := db.Prepare("INSERT INTO Algorithm_Performance (performance_id, algorithm_id, total_return, total_trades, profit_loss, profit_loss_percentage) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = insertAlgorithmPerformanceStmt.Exec(newAlgorithmPerformance.ID, newAlgorithmPerformance.AlgorithmID, newAlgorithmPerformance.TotalReturn, newAlgorithmPerformance.TotalTrades, newAlgorithmPerformance.ProfitLoss, newAlgorithmPerformance.PLPercentage)
	if err != nil {
		log.Fatal(err)
	}

	// Insert a new user performance into the User_Performance table
	newUserPerformance := &UserPerformance{
		ID:            1,
		UserID:        1,
		AlgorithmID:   1,
		TotalReturn:   1000.00,
		TotalTrades:   100,
		ProfitLoss:    500.00,
		PLPercentage:  0.5,
	}
	insertUserPerformanceStmt, err := db.Prepare("INSERT INTO User_Performance (performance_id, user_id, algorithm_id, total_return, total_trades, profit_loss, profit_loss_percentage) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = insertUserPerformanceStmt.Exec(newUserPerformance.ID, newUserPerformance.UserID, newUserPerformance.AlgorithmID, newUserPerformance.TotalReturn, newUserPerformance.TotalTrades, newUserPerformance.ProfitLoss, newUserPerformance.PLPercentage)
	if err != nil {
		log.Fatal(err)
	}

	// Insert a new trade summary into the Trades_Summary table
	newTradeSummary := &TradeSummary{
		ID:                1,
		UserID:            1,
		AlgorithmID:       1,
		AvgExecutionPrice: 100.00,
		TotalAmount:       1000.00,
		ProfitLoss:        500.00,
		TradesCount:       100,
	}
	insertTradeSummaryStmt, err := db.Prepare("INSERT INTO Trades_Summary (summary_id, user_id, algorithm_id, average_execution_price, total_amount, profit_loss, trades_count) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
	log.Fatal(err)
	}
	_, err = insertTradeSummaryStmt.Exec(newTradeSummary.ID, newTradeSummary.UserID, newTradeSummary.AlgorithmID, newTradeSummary.AvgExecutionPrice, newTradeSummary.TotalAmount, newTradeSummary.ProfitLoss, newTradeSummary.TradesCount)
	if err != nil {
	log.Fatal(err)
	}

	// Query the Users table for a specific user
	var queriedUser AlgoUser
	err = db.QueryRow("SELECT * FROM Users WHERE user_id=$1", 1).Scan(&queriedUser.ID, &queriedUser.InitialInvestmentAmt, &queriedUser.InitialInvestmentTm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(queriedUser)

	// Query the Algorithms table for a specific algorithm
	var queriedAlgorithm Algorithm
	err = db.QueryRow("SELECT * FROM Algorithms WHERE algorithm_id=$1", 1).Scan(&queriedAlgorithm.ID, &queriedAlgorithm.Name, &queriedAlgorithm.Description, &queriedAlgorithm.Parameters)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(queriedAlgorithm)

	// Query the User_Algorithm_Budgets table for a specific budget
	var queriedBudget UserAlgorithmBudget
	err = db.QueryRow("SELECT * FROM User_Algorithm_Budgets WHERE budget_id=$1", 1).Scan(&queriedBudget.ID, &queriedBudget.UserID, &queriedBudget.Algorithm, &queriedBudget.BudgetAmt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(queriedBudget)

	// Query the Trades table for a specific trade
	var queriedTrade Trade
	err = db.QueryRow("SELECT * FROM Trades WHERE trade_id=$1", 1).Scan(&queriedTrade.ID, &queriedTrade.UserID, &queriedTrade.AlgorithmID, &queriedTrade.ExecutionPrice, &queriedTrade.Amount, &queriedTrade.Time, &queriedTrade.TradeType)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(queriedTrade)

	// Query the Investments table for a specific investment
	var queriedInvestment Investment
	err = db.QueryRow("SELECT * FROM Investments WHERE investment_id=$1", 1).Scan(&queriedInvestment.ID, &queriedInvestment.UserID, &queriedInvestment.AlgorithmID, &queriedInvestment.CurrentVal, &queriedInvestment.StartTime)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(queriedInvestment)

	// Query the Investment_Limits table for a specific limit
	var queriedLimit InvestmentLimit
	err = db.QueryRow("SELECT * FROM Investment_Limits WHERE limit_id=$1", 1).Scan(&queriedLimit.ID, &queriedLimit.UserID, &queriedLimit.AlgorithmID, &queriedLimit.TotalAmtAllowed, &queriedLimit.TotalAmtInvested, &queriedLimit.RemainingAmtToInvest)
	if err != nil {
	log.Fatal(err)
	}
	fmt.Println(queriedLimit)

	// Query the Algorithm_Performance table for a specific algorithm performance
	var queriedAlgorithmPerformance AlgorithmPerformance
	err = db.QueryRow("SELECT * FROM Algorithm_Performance WHERE performance_id=$1", 1).Scan(&queriedAlgorithmPerformance.ID, &queriedAlgorithmPerformance.AlgorithmID, &queriedAlgorithmPerformance.TotalReturn, &queriedAlgorithmPerformance.TotalTrades, &queriedAlgorithmPerformance.ProfitLoss, &queriedAlgorithmPerformance.PLPercentage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(queriedAlgorithmPerformance)

	// Query the User_Performance table for a specific user performance
	var queriedUserPerformance UserPerformance
	err = db.QueryRow("SELECT * FROM User_Performance WHERE performance_id=$1", 1).Scan(&queriedUserPerformance.ID, &queriedUserPerformance.UserID, &queriedUserPerformance.AlgorithmID, &queriedUserPerformance.TotalReturn, &queriedUserPerformance.TotalTrades, &queriedUserPerformance.ProfitLoss, &queriedUserPerformance.PLPercentage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(queriedUserPerformance)

	// Query the Trades_Summary table for a specific trade summary
	var queriedTradeSummary TradeSummary
	err = db.QueryRow("SELECT * FROM Trades_Summary WHERE summary_id=$1", 1).Scan(&queriedTradeSummary.ID, &queriedTradeSummary.UserID, &queriedTradeSummary.AlgorithmID, &queriedTradeSummary.AvgExecutionPrice, &queriedTradeSummary.TotalAmount, &queriedTradeSummary.ProfitLoss, &queriedTradeSummary.TradesCount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(queriedTradeSummary)

	// Close the database connection
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}


	


	
