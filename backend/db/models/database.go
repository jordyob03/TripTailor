package DBmodels

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

// ------------------------------------------------------------------------------------------------------------------------------
// ----------------------------------------------------- Database Commands ------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------

var DB *sql.DB

func InitDB(connStr string) error {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	if err := DB.Ping(); err != nil {
		return err
	}
	fmt.Println("Successfully connected to the database!")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

// ------------------------------------------------------------------------------------------------------------------------------
// ----------------------------------------------------- Table Commands ---------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------

func CreateTable(createTableSQL string) error {
	_, err := DB.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}
	log.Println("Table created successfully")
	return nil
}

func DeleteTable(tableName string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName)
	_, err := DB.Exec(query)
	if err != nil {
		return fmt.Errorf("error deleting table: %w", err)
	}
	log.Println("Table deleted successfully")
	return nil
}

// ------------------------------------------------------------------------------------------------------------------------------
// ----------------------------------------------------- Row Commands -----------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------

func AddRow(table string, data map[string]interface{}) error {
	columns := make([]string, 0, len(data))
	placeholders := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))

	for column, value := range data {
		columns = append(columns, column)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)+1))
		values = append(values, value)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	_, err := DB.Exec(query, values...)

	return err
}

func DeleteRow(table string, condition string, args ...interface{}) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", table, condition)
	_, err := DB.Exec(query, args...)
	return err
}

func UpdateRow(table string, data map[string]interface{}, condition string, args ...interface{}) error {
	setClauses := make([]string, 0, len(data))

	for column := range data {
		setClauses = append(setClauses, fmt.Sprintf("%s = $1", column))
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, strings.Join(setClauses, ", "), condition)

	values := make([]interface{}, 0, len(data)+len(args))
	for _, value := range data {
		values = append(values, value)
	}

	values = append(values, args...)
	_, err := DB.Exec(query, values...)

	return err
}

func GetRows(table string, condition string, args ...interface{}) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", table, condition)

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]map[string]interface{}, 0)
	columns, _ := rows.Columns()

	for rows.Next() {
		rowValues := make([]interface{}, len(columns))
		for i := range rowValues {
			var v interface{}
			rowValues[i] = &v
		}
		if err := rows.Scan(rowValues...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			rowMap[colName] = *(rowValues[i].(*interface{}))
		}
		results = append(results, rowMap)
	}

	return results, nil
}
