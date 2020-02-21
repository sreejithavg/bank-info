package pkg

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitRouter()  *mux.Router {
	router:=mux.NewRouter()
	router.HandleFunc("/user/non_cellular", nonCellularUser).Methods("GET")
	router.HandleFunc("/user/get",userHandler).Methods("GET")
	router.HandleFunc("/user/filter",userFilterHandler).Methods("POST")
	return router
}

// will return the number of non-cellular users
func nonCellularUser(w http.ResponseWriter, r *http.Request) {
	var recordCount int
	for _, bankInfo := range BankDetails {
		if bankInfo.Contact != "cellular" {
			recordCount++
		}
	}
	response := Response{
		StatusCode:        http.StatusOK,
		StatusDescription: http.StatusText(http.StatusOK),
		Message:           "number of non-cellular user",
		Data:              recordCount,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err.Error())
		errorHandler(w, err, http.StatusInternalServerError)
		return
	}
	return
}

// will return users info after the october 15 th
func userHandler(w http.ResponseWriter, r *http.Request) {
	var bankDetails []PersonalDetails
	for _, bankInfo := range BankDetails {
		switch bankInfo.Month {
		case "oct":
			if bankInfo.Day > 15 {
				bankDetails = append(bankDetails, bankInfo)
			}
		case "nov", "dec":
			bankDetails = append(bankDetails, bankInfo)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Response{
		StatusCode:        http.StatusOK,
		StatusDescription: http.StatusText(http.StatusOK),
		Message:           "users after 15th october",
		Data: UserDetails{
			Records:     len(bankDetails),
			BankDetails: bankDetails,
		},
	})
	if err != nil {
		log.Println(err.Error())
		errorHandler(w, err, http.StatusInternalServerError)
		return
	}
}

//return the no of user records satisfy the request conditions
func userFilterHandler(w http.ResponseWriter, r *http.Request) {
	var request Request
	var record int
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Bad request binding failed")
		errorHandler(w, err, http.StatusBadRequest)
		return
	}
	validate :=validator.New()
	err=validate.Struct(request)
	if err != nil {
		log.Println("Bad request validation failed")
		errorHandler(w, err, http.StatusBadRequest)
		return
	}
	err=validate.VarWithValue(request.EndAge, request.StartAge, "gtfield")
	if err != nil {
		log.Println("Bad request validation failed")
		errorHandler(w, err, http.StatusBadRequest)
		return
	}
	for _, info := range BankDetails {
		if info.Age >= request.StartAge && info.Age <= request.EndAge && info.Marital == request.MartialStatus {
			record++
		}
	}
	response := Response{
		StatusCode:        http.StatusOK,
		StatusDescription: http.StatusText(http.StatusOK),
		Message:           "number of  users between specified age group and martial status",
		Data:              record,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err.Error())
		errorHandler(w, err, http.StatusInternalServerError)
		return
	}
}
