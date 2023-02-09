package go_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customers(name) VALUES('Arthur Shelby')"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT * FROM customers"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close() // close when finish executing

	// iterate all data
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name) // using pointer, so id and name values above will reflect after the method is called
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id)
		fmt.Println("name:", name)
	}

	fmt.Println("Success")
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `SELECT
		id,
		name,
		email,
		balance,
		rating,
		birth_date,
		is_married,
		created_at
		FROM customers`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close() // close when finish executing

	// iterate all data
	for rows.Next() {
		var id int
		var name string
		var email sql.NullString
		var balance sql.NullInt32
		var rating sql.NullFloat64
		var birthDate sql.NullTime
		var isMarried sql.NullBool
		var createdAt time.Time

		err := rows.Scan(
			&id,
			&name,
			&email,
			&balance,
			&rating,
			&birthDate,
			&isMarried,
			&createdAt,
		)

		if err != nil {
			panic(err)
		}

		fmt.Println("=================")
		fmt.Println("id:", id)
		fmt.Println("name:", name)
		if email.Valid == true {
			fmt.Println("email:", email.String)
		}
		if balance.Valid == true {
			fmt.Println("balance:", balance.Int32)
		}
		if rating.Valid == true {
			fmt.Println("rating:", rating.Float64)
		}
		if birthDate.Valid == true {
			fmt.Println("birth date:", birthDate.Time)
		}
		if isMarried.Valid == true {
			fmt.Println("is married:", isMarried.Bool)
		}
		fmt.Println("createdAt:", createdAt)
	}

	fmt.Println("Success")
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username := "admin"
	username := "admin'; #" // exploited
	password := "asdmin"

	query := "SELECT username FROM users WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// this mean to check if record exist
	if rows.Next() {
		var username string

		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Login success", username)
	} else {
		fmt.Println("Gagal")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username := "admin"
	username := "admin'; #" // exploited
	password := "asdmin"

	// Solve SQL Injection
	query := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// this mean to check if record exist
	if rows.Next() {
		var username string

		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Login success", username)
	} else {
		fmt.Println("Gagal")
	}
}

func TestExecSqlWithParameters(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "harry"
	password := "harry"

	query := "INSERT INTO users(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}
