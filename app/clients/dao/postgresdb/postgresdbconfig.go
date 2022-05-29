package postgresdb

import (
	"database/sql"
	"errors"

	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

// CreateDatabaseConnection estabilish a connection with the database
func CreateDatabaseConnection() (*sql.DB, error) {
	log.Printf("Starting creating database connection")
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Error loading .env file")
		return nil, errors.New("database connection failed: unable to load environments")
	}

	url := os.Getenv("POSTGRESQL_CONNECTION_STRING")

	log.Printf("Opening connection with database...")
	// open the connection
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, errors.New("database connection failed")
	}

	log.Printf("Database connection established")
	return db, nil
}

// CreateBankAccountTable creates bank_account table on the database
func CreateBankAccountTable() {
	log.Printf("Start creating bank account table...")

	db, err := CreateDatabaseConnection()
	if err != nil {
		log.Print(err)
	}
	defer log.Printf("Bank account table created successfully")
	defer db.Close()

	sqlStatement := `CREATE TABLE IF NOT EXISTS public.bank_account (
		agency_number integer NOT NULL
		, agency_verification_code integer
		, account_number integer NOT NULL 
		, account_verification_code integer NOT NULL
		, owner varchar(40) NOT NULL
		, document varchar(14) NOT NULL
		, birthdate date NOT NULL
		, account_id varchar(36) NOT NULL
		, account_password varchar(65) NOT NULL
		, balance numeric DEFAULT 0
		, CONSTRAINT bapk PRIMARY KEY (account_id)
	)`

	log.Printf("Executing query...")
	_, errCreating := db.Query(sqlStatement)

	if errCreating != nil {
		log.Fatalf("Unable to execute the query. %v", errCreating)
	}

	createBankAccountTableIndex(db)
}

// createBankAccountTableIndex creates indexes for bank_account table
func createBankAccountTableIndex(db *sql.DB) {
	log.Println("Starting creating bank account table index...")
	q := `
		CREATE INDEX IF NOT EXISTS bank_account_idx
		ON public.bank_account (
			agency_number
			, account_number
			, account_verification_code
			, document
		)

	`
	log.Println("Executing query...")
	_, err := db.Query(q)

	if err != nil {
		log.Fatalf("Error while creating bank account table index. %v", err)
	}
	log.Println("Bank account index created successfully")
}

// CreateTransactionsTable creates transactions table on the database
func CreateTransactionsTable() {
	log.Printf("Start creating transactionsn table...")

	db, err := CreateDatabaseConnection()
	if err != nil {
		log.Print(err)
	}
	defer log.Printf("Transactions table created successfully")
	defer db.Close()

	sqlStatement := `CREATE TABLE IF NOT EXISTS public.transactions (
		transaction_id varchar(36) NOT NULL
		, transaction_type varchar(20) NOT NULL
		, agency_number integer NOT NULL
		, agency_verification_code integer
		, account_number integer NOT NULL 
		, account_verification_code integer NOT NULL
		, document varchar(14) NOT NULL
		, account_id varchar(36) NOT NULL
		, date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
		, value numeric DEFAULT 0
		, CONSTRAINT tpk PRIMARY KEY (transaction_id)
	)`

	log.Printf("Executing query...")
	_, errCreating := db.Query(sqlStatement)

	if errCreating != nil {
		log.Fatalf("Unable to execute the query. %v", errCreating)
	}

	createTransactionsTableIndex(db)
}

// createTransactionsTableIndex creates indexes for transactions table
func createTransactionsTableIndex(db *sql.DB) {
	log.Println("Starting creating transactions table index...")
	q := `
		CREATE INDEX IF NOT EXISTS transactions_idx
		ON public.transactions (
			account_id
		)

	`
	log.Println("Executing query...")
	_, err := db.Query(q)

	if err != nil {
		log.Fatalf("Error while creating transactions table index. %v", err)
	}
	log.Println("Transactions index created successfully")
}
