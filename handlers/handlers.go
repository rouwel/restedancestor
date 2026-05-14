// Package handlers is used to separate the handlers from other functions of the API.
package handlers

import (

	//"log"
	"net/http"
	//"strings"
	//"unicode"
	//"restedancestor/database"
	//"restedancestor/quotes"

	//"github.com/julienschmidt/httprouter"
)
type Quote struct {
	ID int `json:"id_quote"`
	Content string `json:"content"`
	Score int `json:"score"`
	UUID string `json:"uuid"`
}

func GetQuotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := DB.Query("SELECT id_quote, content, score, uuid FROM quotes LIMIT 10")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var quotes []Quote
	for rows.Next() {
		var q Quote
		if err := rows.Scan(&q.ID, &q.Content, &q.Score, &q.UUID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		quotes = append(quotes, q)
	}

	if quotes == nil {
		quotes = []Quote{}
	}

	json.NewEncoder(w).Encode(quotes)
}

func GetQuoteByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var q Quote
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
//var repo = quotes.NewRepository(database.NewDb())

// Random handles the '/random' endpoint
//func Random(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

//	q := repo.Random()
//	err := writeResponse(w, q)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

// All handles the '/all' endpoint
//func All(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

//	err := writeResponse(w, repo.All())
//	if err != nil {
//		log.Fatal(err)
//	}
//}

// Search handles the '/search' endpoint
//func Search(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

//	word := strings.ToLower(p.ByName("word"))
//	qs := repo.AllByWord(word)

//	if len(qs) != 0 {
//		err := writeResponse(w, qs)
//		if err != nil {
//			log.Fatal(err)
//		}
//	} else {
//		err := writeNotFound(w, word)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//}

// Senile handles the '/senile' endpoint
//func Senile(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

//	q1 := repo.Random()
//	q2 := repo.Random()

//	quoteArray := strings.Split(q1.Quote, " ")
//	quoteArray1 := strings.Split(q2.Quote, " ")
//	var quote string

//	if len(quoteArray) < len(quoteArray1) {
//		quote = stringModifier(quoteArray, quoteArray1)
//	} else {
//		quote = stringModifier(quoteArray1, quoteArray)
//	}

//	joinedQuote := quotes.Quote{Quote: quote}
//	err := writeResponse(w, joinedQuote)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

//func Length(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

//	log.Println("testing length")
//	word := p.ByName("len")

//	for _, r := range word {
//		err := unicode.IsLetter(r)
//		if err {
//			log.Fatal("Not a number")
//			return
//		}
//	}

	//length, _ := strconv.ParseUint(word, 10, 32)
	//qs := repo.AllByLengthLessThanOrEqual(length)
	//qs := uint64(length)

	//if len(qs) != 0 {
	//	err := writeResponse(w, qs)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//} else {
	//	err := writeNotFound(w, word)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
//}

// Find handles the '/uuid/:uuid/find' endpoint
//func Find(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

//	uuidToSearch := p.ByName("uuid")

//	q := repo.FindByUUID(uuidToSearch)
//	if q == nil {
//		err := writeNotFound(w, uuidToSearch)
//		if err != nil {
//			log.Fatal(err)
//		}
//		return
//	}

//	err := writeResponse(w, q)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

// Like handles the '/uuid/:uuid/like' endpoint
//func Like(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

//	uuidToSearch := p.ByName("uuid")

//	if err := repo.IncrementsScore(uuidToSearch); err != nil {
//		err = writeNotFound(w, uuidToSearch)
//		if err != nil {
//			log.Fatal(err)
//		}
//		return
//	}
//}

// Dislike handles the '/uuid/:uuid/dislike' endpoint
//func Dislike(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

//	uuidToSearch := p.ByName("uuid")

//	if err := repo.DecrementsScore(uuidToSearch); err != nil {
//		err = writeNotFound(w, uuidToSearch)
//		if err != nil {
//			log.Fatal(err)
//		}
//		return
//	}
//}

// Top handles the '/top' endpoint
//func Top(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

//	err := writeResponse(w, repo.Preferred())
//	if err != nil {
//		log.Fatal(err)
//	}
//}
