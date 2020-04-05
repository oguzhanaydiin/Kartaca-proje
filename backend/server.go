package main

import (
	"encoding/json"
	"net/http"

	databaseutil "./database"
)

type persons struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Mail       string `json:"mail"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}

type operation struct {
	Errorcode    int    `json:"errorcode"`
	Errormessage string `json:"errormessage"`
}

func process(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var personData persons
	var operationData operation

	decoder.Decode(&personData)

	var a string = personData.Password
	var b string = personData.Repassword

	if a != b {
		//fmt.Print("passwords are not same")
		operationData.Errorcode = 2
		operationData.Errormessage = "Passwords are not same"
	} else {
		var k int = databaseutil.Createperson(personData.Firstname, personData.Lastname, personData.Mail, personData.Password)
		operationData.Errorcode = k
		if k == 0 {
			operationData.Errormessage = "Succesful registeration. Please Log In!"
		} else if k == 1 {
			operationData.Errormessage = "You have to use unique e-mail for each account!"
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(operationData); err != nil {
		panic(err)
	}

}

/*func forposts(w http.ResponseWriter, request *http.Request) {
	c := databaseutil.Getposts()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(c); err != nil {
		panic(err)
	}
}*/

func main() {
	http.HandleFunc("/calc", process)
	http.ListenAndServe(":8090", nil)

	//http.HandleFunc("/forposts", forposts)
	//http.ListenAndServe(":8090", nil)
}

//Errorcodes:
// 2 -> passwords are not same
// 1 -> not unique mail
// 0 -> succesful operation
