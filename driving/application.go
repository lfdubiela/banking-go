package driving

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/lfdubiela/banking-go/driven/repository"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lfdubiela/banking-go/driving/handlers"
	"github.com/lfdubiela/banking-go/driving/middlewares"
)

type Application struct {
	db     *sql.DB
	router *mux.Router
}

func Start(addr string) {
	a := &Application{}

	a.setUpDb()
	a.run(addr)

	defer a.db.Close()
}

func (a *Application) setUpDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&timeout=10s&writeTimeout=10s&readTimeout10s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 30)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	a.db = db
}

func (a Application) run(addr string) {
	a.router = mux.NewRouter()

	a.registerHandlers()

	fmt.Println("Listening to port " + addr)

	server := &http.Server{
		Handler:      a.router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func (a *Application) registerHandlers() {
	//Defining middlewares
	a.router.Use(middlewares.Logging)
	a.router.Use(middlewares.ContentApplicationJson)

	//Repository dependencies
	accountRepository := repository.NewAccountRepository(a.db)
	transactionRepository := repository.NewTransactionRepository(a.db)

	createAccount := handlers.NewCreateAccount(accountRepository)
	findAccount := handlers.NewRetrieveAccount(accountRepository)
	createTransaction := handlers.NewCreateTransaction(accountRepository, transactionRepository)

	//Registering routes
	a.router.HandleFunc("/accounts", createAccount.Handler).Methods("POST")
	a.router.HandleFunc("/accounts/{id:[0-9]+}", findAccount.Handler).Methods("GET")
	a.router.HandleFunc("/transactions", createTransaction.Handler).Methods("POST")
}
