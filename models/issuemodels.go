package models

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"
	"yuvraj/project/config"
)

type IssuedBook struct {
	UserID    int       `json:"user_id"`
	BookID    int       `json:"book_id"`
	IssueDate time.Time `json:"issue_date"`
	Duration  time.Time `json:"duration_date"`
}

func IssueBook(issuedBook IssuedBook) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()
	book, err := GetBookByID(issuedBook.BookID)
	if err != nil {
		return err
	}

	user, err := GetUserByID(issuedBook.UserID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("User not found")
	}
	_, err = db.Exec("INSERT INTO issue_book (user_id, book_id, issue_date, duration_date) VALUES ($1, $2, NOW(),NOW() + INTERVAL '7 days')", issuedBook.UserID, issuedBook.BookID)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE book SET quantity = quantity-1 WHERE id = $1", book.ID)
	if err != nil {
		return err
	}

	return nil
}
func GetIssuedBookByBookIDAndUserID(bookID string, userID int) (*IssuedBook, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		return nil, err
	}

	issuedBook := &IssuedBook{}
	err = db.QueryRow("SELECT user_id, book_id, issue_date, duration_date FROM issue_book WHERE book_id = $1 AND user_id = $2", bookIDInt, userID).Scan(&issuedBook.UserID, &issuedBook.BookID, &issuedBook.IssueDate, &issuedBook.Duration)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return issuedBook, nil
}

func ReturnBookFromUser(bookID int, userID int) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = GetIssuedBookByBookIDAndUserID(strconv.Itoa(bookID), userID)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM issue_book WHERE book_id = $1 AND user_id = $2", bookID, userID)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE book SET quantity = quantity + 1 WHERE id = $1", bookID)
	if err != nil {
		return err
	}

	return nil
}
func ReissueBook(bookID int, userID int) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	issuedBook, err := GetIssuedBookByBookIDAndUserID(strconv.Itoa(bookID), userID)
	if err != nil {
		return err
	}
	if issuedBook == nil {
		return errors.New("Book not issued to the user")
	}
	newDuration := time.Now().AddDate(0, 0, 7)
	_, err = db.Exec("UPDATE issue_book SET duration_date = $1 WHERE book_id = $2 AND user_id = $3", newDuration, bookID, userID)
	if err != nil {
		return err
	}

	return nil
}

func CalculateFine(userID int, bookID int) (int, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	query := `SELECT CASE WHEN NOW()::DATE > duration_date::DATE THEN (NOW()::DATE - duration_date::DATE) * 10 + fine	ELSE fine END AS total_fine FROM issue_book WHERE user_id = $1 AND book_id = $2`
	var fine int
	err = db.QueryRow(query, userID, bookID).Scan(&fine)
	if err != nil {
		log.Println("CalculateFine failed while querying:", err)
		return 0, err
	}

	return fine, nil
}

func UpdateFine(userID int, bookID int) error {
	db, err := config.GetDB()
	if err != nil {
		log.Println("GetDB failed while connecting to the database:", err)
		return err
	}
	defer db.Close()

	query := `UPDATE issue_book SET fine = 0 WHERE user_id = $1 AND book_id = $2 RETURNING id`

	_, err = db.Exec(query, userID, bookID)
	if err != nil {
		log.Println("UpdateFine failed while executing query:", err)
		return err
	}

	return nil
}
