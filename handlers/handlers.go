// Package handlers is used to separate the handlers from other functions of the API.
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"restedancestor/quotes"

	"github.com/julienschmidt/httprouter"
)

// Database variable
var DB *sql.DB

// Handler that generates all the quotes from the database
func GetQuotesHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := DB.Query("SELECT id_quote, content, score, uuid FROM quotes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var quoteslist []quotes.Quote
	for rows.Next() {
		var q quotes.Quote
		if err := rows.Scan(&q.ID, &q.Content, &q.Score, &q.UUID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		quoteslist = append(quoteslist, q)
	}

	if quoteslist == nil {
		quoteslist = []quotes.Quote{}
	}

	json.NewEncoder(w).Encode(quoteslist)
}

// handler that generates the  quote based on the id given
func GetQuoteByIDHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var q quotes.Quote
	err = DB.QueryRow("SELECT id_quote, content, score, uuid FROM quotes WHERE id_quote = ?", id).
		Scan(&q.ID, &q.Content, &q.Score, &q.UUID)

	if err == sql.ErrNoRows {
		http.Error(w, "Quote not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(q)
}

// Handler that generates random quotes
func Random(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var q quotes.Quote

	err := DB.QueryRow("SELECT id_quote, content, score, uuid FROM quotes ORDER BY RANDOM() LIMIT 1").
		Scan(&q.ID, &q.Content, &q.Score, &q.UUID)

	if err == sql.ErrNoRows {
		http.Error(w, "Quote not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(q)
}
