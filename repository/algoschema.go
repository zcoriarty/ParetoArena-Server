package repository

import (
	"fmt"
	"log"

	"github.com/zcoriarty/Backend/config"
	

	"github.com/go-pg/pg/v9"
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

func CreateAlgoSchema(db *pg.DB, p *config.PostgresConfig) {

	// defer db.Close()

	// Create the AlgoUsers table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS AlgoUsers (
			user_id SERIAL PRIMARY KEY,
			initial_investment_amount NUMERIC(10,2) NOT NULL,
			initial_investment_time TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	// Create the Algorithms table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Algorithms (
			algorithm_id SERIAL PRIMARY KEY,
			algorithm_name TEXT NOT NULL,
			algorithm_description TEXT NOT NULL,
			algorithm_parameters TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the User_Algorithm_Budgets table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS User_Algorithm_Budgets (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES AlgoUsers (user_id),
			algorithm_id INTEGER NOT NULL REFERENCES Algorithms (algorithm_id),
			budget_amount NUMERIC(10,2) NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the Trades table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Trades (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES AlgoUsers (user_id),
			algorithm_id INTEGER NOT NULL REFERENCES Algorithms (algorithm_id),
			execution_price NUMERIC(10,2) NOT NULL,
			amount NUMERIC(10,2) NOT NULL,
			time TIMESTAMP NOT NULL,
			trade_type CHAR(4) NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the Investments table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Investments (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES AlgoUsers (user_id),
			algorithm_id INTEGER NOT NULL REFERENCES Algorithms (algorithm_id),
			current_value NUMERIC(10,2) NOT NULL,
			start_time TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the Investment_Limits table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Investment_Limits (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES AlgoUsers (user_id),
		algorithm_id INTEGER NOT NULL REFERENCES Algorithms (algorithm_id),
		total_amount_allowed NUMERIC(10,2) NOT NULL,
		total_amount_invested NUMERIC(10,2) NOT NULL,
		remaining_amount_to_invest NUMERIC(10,2) NOT NULL
	)
	`)
	if err != nil {
		log.Fatal(err)
	}
	// Create the Algorithm_Performance table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Algorithm_Performance (
		id SERIAL PRIMARY KEY,
		algorithm_id INTEGER NOT NULL REFERENCES Algorithms (algorithm_id),
		total_return NUMERIC(10,2) NOT NULL,
		total_trades INTEGER NOT NULL,
		profit_loss NUMERIC(10,2) NOT NULL,
		pl_percentage NUMERIC(10,2) NOT NULL
	)
	`)
	if err != nil {
	log.Fatal(err)
	}

	// Create the User_Performance table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS User_Performance (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES AlgoUsers (user_id),
		algorithm_id INTEGER NOT NULL REFERENCES Algorithms (algorithm_id),
		total_return NUMERIC(10,2) NOT NULL,
		total_trades INTEGER NOT NULL,
		profit_loss NUMERIC(10,2) NOT NULL,
		pl_percentage NUMERIC(10,2) NOT NULL
	)
	`)
	if err != nil {
	log.Fatal(err)
	}

	// Create the Trades_Summary table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Trades_Summary (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES AlgoUsers (user_id),
		algorithm_id INTEGER NOT NULL REFERENCES Algorithms (algorithm_id),
		avg_execution_price NUMERIC(10,2) NOT NULL,
		total_amount NUMERIC(10,2) NOT NULL,
		profit_loss NUMERIC(10,2) NOT NULL,
		trades_count INTEGER NOT NULL
	)
	`)
	if err != nil {
	log.Fatal(err)
	}

	fmt.Println("Created Algo schema")

	}