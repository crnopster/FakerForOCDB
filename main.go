package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"syreclabs.com/go/faker"
)

func main() {
	insertCustomers()
}

func insertCustomers() {
	db, err := sql.Open("mysql", "root:1488@/OpenCart")
	if err != nil {
		log.Println(err.Error())
	}
	fname := faker.Name()
	fnumber := faker.Number()
	fphone := faker.PhoneNumber()
	fdate := faker.Date()
	finternet := faker.Internet()
	for i := 0; i < 600; i++ {
		customerGroupID := fnumber.NumberInt(1)
		storeID := fnumber.NumberInt(2)
		languageID := fnumber.NumberInt(2)
		firstName := fname.FirstName()
		lastName := fname.LastName()
		email := finternet.Email()
		telephone := fphone.PhoneNumber()
		fax := fphone.SubscriberNumber(6)
		password := finternet.Password(6, 20)
		salt := ("salted")
		customField := "i don't know, what this means"
		ip := finternet.IpV4Address()
		status := fnumber.NumberInt(1)
		safe := 1
		token := finternet.Password(50, 60)
		code := finternet.Password(1, 40)
		dateAdded := fdate.Birthday(0, 3)

		insForm, err := db.Prepare("INSERT INTO oc_customer(customer_group_id, store_id, language_id, firstname, lastname, email, telephone, fax, password, salt, custom_field, ip, status, safe, token, code, date_added) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			log.Println(err.Error())
		}
		r, err := insForm.Exec(
			customerGroupID,
			storeID,
			languageID,
			firstName,
			lastName,
			email,
			telephone,
			fax,
			password,
			salt,
			customField,
			ip,
			status,
			safe,
			token,
			code,
			dateAdded)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("added", r)
	}
	defer db.Close()
}
