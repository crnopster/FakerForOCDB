package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"syreclabs.com/go/faker"
)

func main() {
	insertOrderProduct()
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := "1488"
	dbName := "OpenCart"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	if err != nil {
		log.Println(err.Error())
	}
	return db
}

func insertProduct() {

	db := dbConn()
	defer db.Close()
	rand.Seed(time.Now().UnixNano())

	insForm, err := db.Prepare("INSERT INTO oc_product(model, sku, upc, ean, jan, isbn, mpn, location, quantity, stock_status_id, manufacturer_id, shipping, price, tax_class_id, date_added, date_modified) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
	}

	fnumber := faker.Number()
	fdate := faker.Date()

	for i := 0; i < 20000; i++ {
		model := ("product model " + fnumber.Between(1, 10000))
		sku := fnumber.Number(64)
		upc := fnumber.Number(12)
		ean := fnumber.Number(14)
		jan := fnumber.Number(13)
		isbn := fnumber.Number(17)
		mpn := fnumber.Number(64)
		location := ("test " + fnumber.Number(1))
		quantity := rand.Intn(9999)
		stockStatusID := rand.Intn(9)
		manufacturerID := rand.Intn(50)
		shipping := rand.Intn(1)
		price := float32(rand.Intn(99999))
		taxClassID := rand.Intn(9)
		dateAdded := fdate.Birthday(3, 6)
		dateModified := fdate.Birthday(1, 2)

		resProduct, err := insForm.Exec(
			model,
			sku,
			upc,
			ean,
			jan,
			isbn,
			mpn,
			location,
			quantity,
			stockStatusID,
			manufacturerID,
			shipping,
			price,
			taxClassID,
			dateAdded,
			dateModified)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("product", resProduct)
	}
}

func insertProductDescription() {
	db := dbConn()
	defer db.Close()
	rand.Seed(time.Now().UnixNano())

	insForm, err := db.Prepare("INSERT INTO oc_product_description(product_id, language_id, name, description, tag, meta_title, meta_description, meta_keyword, meta_h1) VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
	}

	fcommerce := faker.Commerce()
	ftext := faker.Lorem()

	//Get slice of ids from oc_product table
	productIDslice := make([]int, 0)
	productIDs, err := db.Query("SELECT product_id FROM oc_product")
	if err != nil {
		log.Println(err.Error())
	}
	for productIDs.Next() {
		var id int
		err = productIDs.Scan(&id)
		if err != nil {
			log.Println(err.Error())
		}
		productIDslice = append(productIDslice, id)
	}

	for i := 0; i < 20000; i++ {
		productID := (productIDslice[i])
		languageID := 1
		name := fcommerce.ProductName()
		description := ftext.Sentence(30)
		tag := " "
		metaTitle := " "
		metaDescription := " "
		metaKeyword := " "
		metaH1 := " "

		resProductDescription, err := insForm.Exec(
			productID,
			languageID,
			name,
			description,
			tag,
			metaTitle,
			metaDescription,
			metaKeyword,
			metaH1)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("product description", resProductDescription)
	}
}

func insertCustomer() {

	db := dbConn()
	defer db.Close()
	rand.Seed(time.Now().UnixNano())

	insForm, err := db.Prepare("INSERT INTO oc_customer(customer_group_id, store_id, language_id, firstname, lastname, email, telephone, fax, password, salt, custom_field, ip, status, safe, token, code, date_added) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
	}

	fnumber := faker.Number()
	fdate := faker.Date()
	fname := faker.Name()
	fphone := faker.PhoneNumber()
	finternet := faker.Internet()

	//Get slice of ids from oc_store table
	storeIDslice := make([]int, 0)
	storeIDs, err := db.Query("SELECT store_id FROM oc_store")
	if err != nil {
		log.Println(err.Error())
	}
	for storeIDs.Next() {
		var id int
		err := storeIDs.Scan(&id)
		if err != nil {
			log.Println(err.Error())
		}
		storeIDslice = append(storeIDslice, id)
	}

	for i := 0; i < 1000000; i++ {

		storeID := (storeIDslice[rand.Intn(len(storeIDslice))])

		customerGroupID := fnumber.NumberInt(1)
		languageID := 1
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
		dateAdded := fdate.Birthday(3, 10)

		resCustomer, err := insForm.Exec(
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
		log.Println("customer", resCustomer)
	}
}

func insertOrder() {

	db := dbConn()
	defer db.Close()
	rand.Seed(time.Now().UnixNano())

	fnumber := faker.Number()
	fdate := faker.Date()
	finternet := faker.Internet()
	ftext := faker.Lorem()
	fcompany := faker.Company()
	faddress := faker.Address()
	shippingCompanyslice := make([]string, 0)
	shippingCompanyslice = append(shippingCompanyslice, "Nova Poshta", "UkrPoshta", "Some other poshta")

	var storeName string
	var storeURL string
	var customerGroupID int
	var firstName string
	var lastName string
	var email string
	var telephone string
	var fax string
	var customField string
	var zoneName string
	var countryID int
	var countryName string
	var languageID int
	var acceptLanguage string
	var ip string
	var currencyCode string

	insForm, err := db.Prepare("INSERT INTO oc_order(invoice_no, invoice_prefix, store_id, store_name, store_url, customer_id, customer_group_id, firstname, lastname, email, telephone, fax, custom_field, payment_firstname, payment_lastname, payment_company, payment_address_1, payment_address_2, payment_city, payment_postcode, payment_country, payment_country_id, payment_zone, payment_zone_id, payment_address_format, payment_custom_field, payment_method, payment_code, shipping_firstname, shipping_lastname, shipping_company, shipping_address_1, shipping_address_2, shipping_city, shipping_postcode, shipping_country, shipping_country_id, shipping_zone, shipping_zone_id, shipping_address_format, shipping_custom_field, shipping_method, shipping_code, comment, total, order_status_id, affiliate_id, commission, marketing_id, tracking, language_id, currency_id, currency_code, ip, forwarded_ip, user_agent, accept_language, date_added, date_modified) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
	}

	//Get slice of ids from oc_product table
	productIDslice := make([]int, 0)
	productIDs, err := db.Query("select product_id from oc_product")
	if err != nil {
		log.Println(err.Error())
	}
	for productIDs.Next() {
		var id int
		err = productIDs.Scan(&id)
		if err != nil {
			log.Println(err.Error())
		}
		productIDslice = append(productIDslice, id)
	}

	//Get slice of ids from oc_store table
	storeIDslice := make([]int, 0)
	storeIDs, err := db.Query("SELECT store_id FROM oc_store")
	if err != nil {
		log.Println(err.Error())
	}
	for storeIDs.Next() {
		var id int
		err = storeIDs.Scan(&id)
		if err != nil {
			log.Println(err.Error())
		}
		storeIDslice = append(storeIDslice, id)
	}

	//Get slice of ids from oc_customer table
	customerIDslice := make([]int, 0)
	customerIDs, err := db.Query("SELECT customer_id FROM oc_customer")
	if err != nil {
		log.Println(err.Error())
	}
	for customerIDs.Next() {
		var id int
		err = customerIDs.Scan(&id)
		if err != nil {
			log.Println(err.Error())
		}
		customerIDslice = append(customerIDslice, id)
	}

	//Get slice of ids from oc_zone table
	zoneIDslice := make([]int, 0)
	zoneIDs, err := db.Query("SELECT zone_id FROM oc_zone")
	if err != nil {
		log.Println(err.Error())
	}
	for zoneIDs.Next() {
		var id int
		err = zoneIDs.Scan(&id)
		if err != nil {
			log.Println(err.Error())
		}
		zoneIDslice = append(zoneIDslice, id)
	}

	//Get slice of ids from oc_order_status
	orderStatusIDslice := make([]int, 0)
	orderStatusIDs, err := db.Query("SELECT order_status_id FROM oc_order_status")
	for orderStatusIDs.Next() {
		var id int
		err = orderStatusIDs.Scan(&id)
		if err != nil {
			log.Println(err.Error())
		}
		orderStatusIDslice = append(orderStatusIDslice, id)
	}

	//Get slice of ids from oc_currency
	currencyIDslice := make([]int, 0)
	currencyIDs, err := db.Query("SELECT currency_id FROM oc_currency")
	if err != nil {
		log.Println(err.Error())
	}
	for currencyIDs.Next() {
		var id int
		err = currencyIDs.Scan(&id)
		if err != nil {
			log.Println(err.Error())
		}
		currencyIDslice = append(currencyIDslice, id)
	}

	for i := 0; i < 100000; i++ {

		storeID := (storeIDslice[rand.Intn(len(storeIDslice))])

		storeNameRow, err := db.Query("SELECT name FROM oc_store WHERE store_id=?", storeID)
		if err != nil {
			log.Println(err.Error())
		}
		for storeNameRow.Next() {
			err = storeNameRow.Scan(&storeName)
			if err != nil {
				log.Println(err.Error())
			}
		}

		storeURLRow, err := db.Query("SELECT url FROM oc_store WHERE store_id=?", storeID)
		if err != nil {
			log.Println(err.Error())
		}
		for storeURLRow.Next() {
			err = storeURLRow.Scan(&storeURL)
			if err != nil {
				log.Println(err.Error())
			}
		}

		customerID := (customerIDslice[rand.Intn(len(customerIDslice))])

		customerGroupIDRow, err := db.Query("SELECT customer_group_id FROM oc_customer WHERE customer_id=?", customerID)
		if err != nil {
			log.Println(err.Error())
		}
		for customerGroupIDRow.Next() {
			err = customerGroupIDRow.Scan(&customerGroupID)
			if err != nil {
				log.Println(err.Error())
			}
		}

		firstNameRow, err := db.Query("SELECT firstname FROM oc_customer WHERE customer_id=?", customerID)
		if err != nil {
			log.Println(err.Error())
		}
		for firstNameRow.Next() {
			err = firstNameRow.Scan(&firstName)
			if err != nil {
				log.Println(err.Error())
			}
		}

		lastNameRow, err := db.Query("SELECT lastname FROM oc_customer WHERE customer_id=?", customerID)
		if err != nil {
			log.Println(err.Error())
		}
		for lastNameRow.Next() {
			err = lastNameRow.Scan(&lastName)
			if err != nil {
				log.Println(err.Error())
			}
		}

		emailRow, err := db.Query("SELECT email FROM oc_customer WHERE customer_id=?", customerID)
		if err != nil {
			log.Println(err.Error())
		}
		for emailRow.Next() {
			err = emailRow.Scan(&email)
			if err != nil {
				log.Println(err.Error())
			}
		}

		telephoneRow, err := db.Query("SELECT telephone FROM oc_customer WHERE customer_id=?", customerID)
		if err != nil {
			log.Println(err.Error())
		}
		for telephoneRow.Next() {
			err = telephoneRow.Scan(&telephone)
			if err != nil {
				log.Println(err.Error())
			}
		}

		faxRow, err := db.Query("SELECT fax FROM oc_customer WHERE customer_id=?", customerID)
		if err != nil {
			log.Println(err.Error())
		}
		for faxRow.Next() {
			err = faxRow.Scan(&fax)
			if err != nil {
				log.Println(err.Error())
			}
		}

		ipRow, err := db.Query("SELECT ip FROM oc_customer WHERE customer_id=?", customerID)
		if err != nil {
			log.Println(err.Error())
		}
		for ipRow.Next() {
			err = ipRow.Scan(&ip)
			if err != nil {
				log.Println(err.Error())
			}
		}

		languageIDRow, err := db.Query("SELECT language_id FROM oc_customer WHERE customer_id=?", customerID)
		if err != nil {
			log.Println(err.Error())
		}
		for languageIDRow.Next() {
			err = languageIDRow.Scan(&languageID)
			if err != nil {
				log.Println(err.Error())
			}
		}

		acceptLanguageRow, err := db.Query("SELECT name FROM oc_language WHERE language_id=?", languageID)
		if err != nil {
			log.Println(err.Error())
		}
		for acceptLanguageRow.Next() {
			err = acceptLanguageRow.Scan(&acceptLanguage)
			if err != nil {
				log.Println(err.Error())
			}
		}

		zoneID := (zoneIDslice[rand.Intn(len(zoneIDslice))])

		zoneNameRow, err := db.Query("SELECT name FROM oc_zone WHERE zone_id=?", zoneID)
		if err != nil {
			log.Println(err.Error())
		}
		for zoneNameRow.Next() {
			err = zoneNameRow.Scan(&zoneName)
			if err != nil {
				log.Println(err.Error())
			}
		}

		countryIDRow, err := db.Query("SELECT country_id FROM oc_zone WHERE zone_id=?", zoneID)
		if err != nil {
			log.Println(err.Error())
		}
		for countryIDRow.Next() {
			err = countryIDRow.Scan(&countryID)
			if err != nil {
				log.Println(err.Error())
			}
		}

		countryNameRow, err := db.Query("SELECT name FROM oc_country WHERE country_id=?", countryID)
		if err != nil {
			log.Println(err.Error())
		}
		for countryNameRow.Next() {
			err = countryNameRow.Scan(&countryName)
			if err != nil {
				log.Println(err.Error())
			}
		}

		currencyID := (currencyIDslice[rand.Intn(len(currencyIDslice))])

		currencyCodeRow, err := db.Query("SELECT code from oc_currency where currency_id=?", currencyID)
		if err != nil {
			log.Println(err.Error())
		}
		for currencyCodeRow.Next() {
			err = currencyCodeRow.Scan(&currencyCode)
			if err != nil {
				log.Println(err.Error())
			}
		}

		orderStatusID := (orderStatusIDslice[rand.Intn(len(orderStatusIDslice))])

		addressFormat := (" ")
		invoiceNO := 1
		invoicePrefix := "invoice prefix"
		shippingCompany := (shippingCompanyslice[rand.Intn(len(shippingCompanyslice))])
		paymentCompany := fcompany.Name()
		address := faddress.StreetAddress()
		city := faddress.City()
		postcode := faddress.Postcode()
		comment := ftext.Sentence(15)
		paymentMethod := ("Payment method " + fnumber.Between(1, 10))
		paymentCode := finternet.Password(64, 128)
		shippingMethod := ("Shipping method" + fnumber.Between(1, 50))
		shippingCode := finternet.Password(50, 50)
		total := float64(rand.Intn(1000))
		dateAdded := fdate.Birthday(1, 3)
		dateModified := fdate.Birthday(0, 1)
		userAgent := "user agent"
		affiliateID := rand.Intn(10)
		tracking := shippingCode
		marketingID := 1
		commission := float64(rand.Intn(10))

		resOrder, err := insForm.Exec(
			invoiceNO,
			invoicePrefix,
			storeID,
			storeName,
			storeURL,
			customerID,
			customerGroupID,
			firstName,
			lastName,
			email,
			telephone,
			fax,
			customField,
			firstName,
			lastName,
			paymentCompany,
			address,
			address,
			city,
			postcode,
			countryName,
			countryID,
			zoneName,
			zoneID,
			addressFormat,
			customField,
			paymentMethod,
			paymentCode,
			firstName,
			lastName,
			shippingCompany,
			address,
			address,
			city,
			postcode,
			countryName,
			countryID,
			zoneName,
			zoneID,
			addressFormat,
			customField,
			shippingMethod,
			shippingCode,
			comment,
			total,
			orderStatusID,
			affiliateID,
			commission,
			marketingID,
			tracking,
			languageID,
			currencyID,
			currencyCode,
			ip,
			ip,
			userAgent,
			acceptLanguage,
			dateAdded,
			dateModified,
		)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("order", resOrder)
	}
}

func insertOrderProduct() {

	db := dbConn()
	defer db.Close()
	rand.Seed(time.Now().UnixNano())

	var name string
	var price float64
	var model string

	//Get slice of ids from oc_product table
	productIDslice := make([]int, 0)
	productIDs, err := db.Query("SELECT product_id FROM oc_product")
	if err != nil {
		log.Println(err.Error())
	}
	for productIDs.Next() {
		var id int
		err = productIDs.Scan(&id)
		if err != nil {
			log.Println(err.Error())
		}
		productIDslice = append(productIDslice, id)
	}
	//Get slice of ids from oc_order table
	orderIDslice := make([]int, 0)
	orderIDs, err := db.Query("SELECT order_id FROM oc_order")
	if err != nil {
		log.Println(err.Error())
	}
	for orderIDs.Next() {
		var id int
		err = orderIDs.Scan(&id)
		if err != nil {
			log.Println(err.Error())
		}
		orderIDslice = append(orderIDslice, id)
	}

	insForm, err := db.Prepare("INSERT INTO oc_order_product(order_id, product_id, name, model, quantity, price, total, tax, reward) VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
	}

	for i := 0; i < 100000; i++ {

		orderID := orderIDslice[i]

		productID := (productIDslice[rand.Intn(len(productIDslice))])

		nameRow, err := db.Query("SELECT name FROM oc_product_description WHERE product_id=?", productID)
		if err != nil {
			log.Println(err.Error())
		}
		for nameRow.Next() {
			err = nameRow.Scan(&name)
			if err != nil {
				log.Println(err.Error())
			}
		}

		priceRow, err := db.Query("SELECT price FROM oc_product WHERE product_id=?", productID)
		if err != nil {
			log.Println(err.Error())
		}
		for priceRow.Next() {
			err = priceRow.Scan(&price)
			if err != nil {
				log.Println(err.Error())
			}
		}

		modelRow, err := db.Query("SELECT model FROM oc_product WHERE product_id=?", productID)
		if err != nil {
			log.Println(err.Error())
		}
		for modelRow.Next() {
			err = modelRow.Scan(&model)
			if err != nil {
				log.Println(err.Error())
			}
		}

		quantity := float64(rand.Intn(9) + 1)
		total := (price * quantity)
		tax := (total / 5)
		reward := 1
		//		totalWithTax := (total + tax)
		//
		//		_, err = db.Query("UPDATE oc_order SET total=? where order_id=?", totalWithTax, orderID)
		//		if err != nil {
		//			log.Println(err.Error())
		//		}
		//
		resOrderProduct, err := insForm.Exec(
			orderID,
			i,
			name,
			model,
			quantity,
			price,
			total,
			tax,
			reward,
		)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("order product", i, resOrderProduct)

	}

}
