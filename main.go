package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// ============================== DB variables ==============================

// MySQLDBHandler as type struct
type MySQLDBHandler struct {
	Conn *sqlx.DB
}

// Record as type struct
type Record struct {
	ID        string
	Data      string
	CreatedAt time.Time `db:"created_at"`
}

// RecordRequest as type struct
type RecordRequest struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

// RecordResponse as type struct
type RecordResponse struct {
	ID        string `json:"id"`
	Data      string `json:"data"`
	CreatedAt int64  `json:"createdAt"`
}

var (
	mysqlDBHandler *MySQLDBHandler
	recordTable    string = "records"
)

// HTTPResponseVM base http viewmodel for http rest responses
type HTTPResponseVM struct {
	Status    int         `json:"-"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode interface{} `json:"errorCode,omitempty"`
	Data      interface{} `json:"data"`
}

// initialize main function
func main() {
	port := ":8080"
	fmt.Println("Starting Server....")

	// initialize mysql db handler
	mysqlDBHandler = &MySQLDBHandler{}

	// connect to database
	err := mysqlDBHandler.Connect("127.0.0.1", "3306", "nuxify_training", "root", "1234")
	if err != nil {
		panic(err)
	}

	// initialize http router
	router := chi.NewRouter()

	// initialize middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api", func(router chi.Router) {
		// routes for version
		router.Route("/v1", func(router chi.Router) {
			// routes for record
			router.Route("/record", func(router chi.Router) {
				router.Post("/", CreateRecordHandler)
				// router.Get("/{id}", GetRecordByIDHandler)
			})
		})
	})
	fmt.Println("Server is listening on " + port)
	log.Fatal(http.ListenAndServe(port, router))
}

// CreateRecordHandler creates a new user resource
func CreateRecordHandler(w http.ResponseWriter, r *http.Request) {
	var request RecordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Invalid payload sent.",
		}

		response.JSON(w)
		return
	}

	// verify content must not empty
	if len(request.Data) == 0 {
		response := HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Data input cannot be empty.",
		}

		response.JSON(w)
		return
	}

	record := Record{
		ID:   request.ID,
		Data: request.Data,
	}

	// insert to database
	_, err := InsertDataRepository(record)
	if err != nil {
		if err.Error() == "DUPLICATE_ID" {
			response := HTTPResponseVM{
				Status:  http.StatusConflict,
				Success: false,
				Message: "Duplicate id.",
			}

			response.JSON(w)
			return
		}

		response := HTTPResponseVM{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := &HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully created user.",
		Data: &RecordResponse{
			ID:        record.ID,
			Data:      record.Data,
			CreatedAt: time.Now().Unix(),
		},
	}

	response.JSON(w)
}

// ============================== Repositories ==============================
// ============================== record repository ==============================

// InsertDataRepository insert a user data
func InsertDataRepository(data Record) (Record, error) {
	var record Record
	stmt := fmt.Sprintf("INSERT INTO %s (id, data) VALUES (:id,:data)", recordTable)
	_, err := mysqlDBHandler.Execute(stmt, data)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return Record{}, errors.New("DUPLICATE_ID")
		}
		return Record{}, errors.New("DATABASE_ERROR")
	}

	return record, nil
}

// ============================== MySQL Helper ==============================

// Connect opens a new connection to the mysql interface
func (h *MySQLDBHandler) Connect(host, port, database, username, password string) error {
	conn, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database))
	if err != nil {
		return err
	}

	h.Conn = conn

	err = conn.Ping()
	if err != nil {
		connErr := fmt.Errorf("[SERVER] Error connecting to the database! %s", err.Error())

		return connErr
	}

	fmt.Println("[SERVER] Database connected successfully")

	return nil
}

// Execute executes the mysql statement following NamedExec
// It requires a valid sql statement and its struct
func (h *MySQLDBHandler) Execute(stmt string, model interface{}) (sql.Result, error) {
	res, err := h.Conn.NamedExec(stmt, model)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query selects rows given by the sql statement
// It requires the statement, the model to bind the statement, and the target bind model for the results
func (h *MySQLDBHandler) Query(qstmt string, model interface{}, bindModel interface{}) error {
	nstmt, err := h.Conn.PrepareNamed(qstmt)
	if err != nil {
		return err
	}
	defer nstmt.Close()

	err = nstmt.Select(bindModel, model)
	if err != nil {
		return err
	}

	return nil
}

// ============================== HTTP Helper ==============================

// JSON converts http responsewriter to json
func (response *HTTPResponseVM) JSON(w http.ResponseWriter) {
	if response.Data == nil {
		response.Data = map[string]interface{}{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	_ = json.NewEncoder(w).Encode(response)
}