package pkg

type PersonalDetails struct {
	Age       int64   `json:"age"`
	Job       string  `json:"job"`       //: type of job (categorical: "admin.","unknown","unemployed","management","housemaid","entrepreneur","student", "blue-collar","self-employed","retired","technician","services")
	Marital   string  `json:"marital"`   //: marital status (categorical: "married","divorced","single"; note: "divorced" means divorced or widowed)
	Education string  `json:"education"` //(categorical: "unknown","secondary","primary","tertiary")
	Default   string  `json:"default"`   //: has credit in default? (binary: "yes","no")
	Balance   float64 `json:"balance"`   //: average yearly balance, in euros (numeric)
	Housing   string  `json:"housing"`   //: has housing loan? (binary: "yes","no")
	Loan      string  `json:"loan"`      //: has personal loan? (binary: "yes","no")
	Contact   string  `json:"contact"`   // : contact communication type (categorical: "unknown","telephone","cellular")
	Day       int64   `json:"day"`       //: last contact day of the month (numeric)
	Month     string  `json:"month"`     //: last contact month of year (categorical: "jan", "feb", "mar", ..., "nov", "dec")
	Duration  int64   `json:"duration"`  //: last contact duration, in seconds (numeric)
	Campaign  int64   `json:"campaign"`  //: number of contacts performed during this campaign and for this client (numeric, includes last contact)
	PDays     int64   `json:"p_days"`    //: number of days that passed by after the client was last contacted from a previous campaign (numeric, -1 means client was not previously contacted)
	Previous  int64   `json:"previous"`  //: number of contacts performed before this campaign and for this client (numeric)
	POutcome  string  `json:"p_outcome"` //: outcome of the previous marketing campaign (categorical: "unknown","other","failure","success")  Output variable (desired target):
	Y         string  `json:"y"`         //- has the client subscribed a term deposit? (binary: "yes","no")
}
type Response struct {
	StatusCode        int         `json:"status_code"`
	StatusDescription string      `json:"status_description"`
	Message           string      `json:"message"`
	Data              interface{} `json:"data"`
}
type Request struct {
	StartAge      int64  `json:"start_age" validate:"required"`
	EndAge        int64  `json:"end_age" validate:"required"`
	MartialStatus string `json:"martial_status" validate:"required,oneof=married divorced single"`
}

var BankDetails []PersonalDetails

type UserDetails struct {
	Records     int               `json:"records"`
	BankDetails []PersonalDetails `json:"bank_details"`
}
