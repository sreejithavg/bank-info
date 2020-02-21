package pkg

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)



func CsvReader(filePath string) []PersonalDetails {
	desc, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(desc))
	reader.Comma = ';'
	var bankDetails []PersonalDetails
	var person PersonalDetails
	count := 0
	for {
		line, error := reader.Read()
		if count == 0 {
			count++
			continue
		}

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		person.Age, _ = strconv.ParseInt(line[0], 10, 64)
		person.Job = line[1]
		person.Marital = line[2]
		person.Education = line[3]
		person.Default = line[4]
		person.Balance, _ = strconv.ParseFloat(line[5], 64)
		person.Housing = line[6]
		person.Loan = line[7]
		person.Job = line[1]
		person.Contact = line[8]
		person.Day, _ = strconv.ParseInt(line[9], 10, 64)
		person.Month = line[10]
		person.Duration, _ = strconv.ParseInt(line[11], 10, 64)
		person.Campaign, _ = strconv.ParseInt(line[12], 10, 64)
		person.PDays, _ = strconv.ParseInt(line[13], 10, 64)
		person.Previous, _ = strconv.ParseInt(line[14], 10, 64)
		person.POutcome = line[15]
		person.Y = line[16]
		bankDetails = append(bankDetails, person)
	}
	return bankDetails
}
func errorHandler(w http.ResponseWriter,err error,status int)  {
	errResponse:=Response{
		StatusCode:        status,
		StatusDescription: "failed",
		Message:           err.Error(),
		Data:              nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errResponse)
}
