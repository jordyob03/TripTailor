package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/lib/pq"
)

func InitDB(DB *sql.DB, connStr string) error {
	if err := DB.Ping(); err != nil {
		return err
	}
	fmt.Println("Successfully connected to the database!")
	return nil
}

func CreateAllTables(DB *sql.DB) error {
	if err := CreateUserTable(DB); err != nil {
		return err
	}
	if err := CreateBoardTable(DB); err != nil {
		return err
	}
	if err := CreatePostTable(DB); err != nil {
		return err
	}
	if err := CreateItineraryTable(DB); err != nil {
		return err
	}
	if err := CreateEventTable(DB); err != nil {
		return err
	}
	if err := CreateImageTable(DB); err != nil {
		return err
	}
	if err := CreateCommentTable(DB); err != nil {
		return err
	}
	return nil
}

func DeleteAllTables(DB *sql.DB) error {
	dropTablesSQL := `
	DROP TABLE IF EXISTS users, boards, posts, itineraries, events, images, comments CASCADE;`

	_, err := DB.Exec(dropTablesSQL)
	if err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}

	return nil
}

func CloseDB(DB *sql.DB) {
	if DB != nil {
		DB.Close()
	}
}

func CreateTable(DB *sql.DB, createTableSQL string) error {
	_, err := DB.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}
	log.Println("Table created successfully")
	return nil
}

func DeleteTable(DB *sql.DB, tableName string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName)
	_, err := DB.Exec(query)
	if err != nil {
		return fmt.Errorf("error deleting table: %w", err)
	}
	log.Println("Table deleted successfully")
	return nil
}

func IntsToStrings(ints []int) []string {
	strings := make([]string, len(ints))
	for i, v := range ints {
		strings[i] = fmt.Sprintf("%d", v)
	}
	return strings
}

func StringsToInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, v := range strings {
		var err error
		ints[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func UpdateAttribute(DB *sql.DB, table string, identifierCol string, identifier interface{}, column string, value interface{}) error {
	updateSQL := fmt.Sprintf("UPDATE %s SET %s = $1 WHERE %s = $2", table, column, identifierCol)

	_, err := DB.Exec(updateSQL, value, identifier)
	if err != nil {
		return fmt.Errorf("error updating %s in table %s: %w", column, table, err)
	}
	return nil
}

func AddArrayAttribute(DB *sql.DB, table, identifierCol string, identifier interface{}, column string, values interface{}) error {
	var existingValues []string
	var existingIntValues []int

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = $1", column, table, identifierCol)

	var err error
	switch v := values.(type) {
	case []string:
		err = DB.QueryRow(query, identifier).Scan(pq.Array(&existingValues))
	case []int:
		err = DB.QueryRow(query, identifier).Scan(pq.Array(&existingIntValues))
	default:
		log.Printf("Unsupported value type: %T\n", v)
		return fmt.Errorf("unsupported value type")
	}

	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error fetching %s for %s: %v\n", column, identifier, err)
		return err
	}

	existingValueMap := make(map[interface{}]struct{})
	for _, val := range existingValues {
		existingValueMap[val] = struct{}{}
	}
	for _, val := range existingIntValues {
		existingValueMap[val] = struct{}{}
	}

	newValues := []interface{}{}
	switch v := values.(type) {
	case []string:
		for _, val := range v {
			if _, exists := existingValueMap[val]; !exists {
				newValues = append(newValues, val)
			}
		}
	case []int:
		for _, val := range v {
			if _, exists := existingValueMap[val]; !exists {
				newValues = append(newValues, val)
			}
		}
	}

	if len(newValues) == 0 {
		log.Printf("No new values to add for %s in column %s\n", identifier, column)
		return nil
	}

	updateSQL := fmt.Sprintf(`UPDATE %s SET %s = array_cat(%s, $1) WHERE %s = $2`, table, column, column, identifierCol)

	_, err = DB.Exec(updateSQL, pq.Array(newValues), identifier)
	if err != nil {
		log.Printf("Error adding values to %s for %s: %v\n", column, identifier, err)
		return err
	}

	return nil
}

func RemoveArrayAttribute(DB *sql.DB, table, identifierCol string, identifier interface{}, column string, values interface{}) error {
	switch v := values.(type) {
	case []string:
		for _, val := range v {
			removeValSQL := fmt.Sprintf(`UPDATE %s SET %s = array_remove(%s, $1) WHERE %s = $2`, table, column, column, identifierCol)

			_, err := DB.Exec(removeValSQL, val, identifier)
			if err != nil {
				log.Printf("Error removing %s '%s' for %s: %v\n", column, val, identifier, err)
				return err
			}
		}
	case []int:
		for _, val := range v {
			removeValSQL := fmt.Sprintf(`UPDATE %s SET %s = array_remove(%s, $1) WHERE %s = $2`, table, column, column, identifierCol)

			_, err := DB.Exec(removeValSQL, val, identifier)
			if err != nil {
				log.Printf("Error removing %s '%d' for %s: %v\n", column, val, identifier, err)
				return err
			}
		}
	default:
		log.Printf("Unsupported value type: %T\n", v)
		return fmt.Errorf("unsupported value type")
	}

	return nil
}
