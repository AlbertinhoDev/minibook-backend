package handlers

import (
	"minibook-backend/internal/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Transaction представляет структуру транзакции
type Transaction struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
	Type     string    `json:"type"`
	Date     time.Time `json:"date"`
}

// CreateTransaction создаёт новую транзакцию
func CreateTransaction(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var transaction Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Установка UserID из контекста
	transaction.UserID = userID.(int)
	transaction.Date = time.Now()

	// Валидация
	if transaction.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amount must be positive"})
		return
	}
	if transaction.Type != "income" && transaction.Type != "expense" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type must be 'income' or 'expense'"})
		return
	}

	// Сохранение в БД
	err := db.DB.QueryRow(
		"INSERT INTO transactions (user_id, amount, category, type, date) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		transaction.UserID, transaction.Amount, transaction.Category, transaction.Type, transaction.Date,
	).Scan(&transaction.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

// GetTransactions возвращает список транзакций пользователя
func GetTransactions(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	rows, err := db.DB.Query(
		"SELECT id, user_id, amount, category, type, date FROM transactions WHERE user_id = $1 ORDER BY date DESC",
		userID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var t Transaction
		err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.Category, &t.Type, &t.Date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		transactions = append(transactions, t)
	}

	c.JSON(http.StatusOK, transactions)
}
