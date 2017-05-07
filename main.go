package main

import (
"fmt"
"net/http"
"strings"
"log"
"os"
"net"
)

var class string = "default"
var cc string = "MM"
var course_code string = "XX"

/*
Takes in the request and appends to a csv
 */
func submitForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Parse arguments
	r.ParseForm()
	firstname := strings.Join(r.Form["firstname"], "")
	lastname := strings.Join(r.Form["lastname"], "")
	gender := strings.Join(r.Form["gender"], "")
	email := strings.Join(r.Form["email"], "")
	password := strings.Join(r.Form["password"], "")
	parentfirstname := strings.Join(r.Form["parentfirstname"], "")
	parentlastname := strings.Join(r.Form["parentlastname"], "")
	parentemail := strings.Join(r.Form["parentemail"], "")
	photoconsent := strings.Join(r.Form["photoconsent"], "")
	parentsignature := strings.Join(r.Form["parentsignature"], "")

	text := class + cc + course_code + "," + cc + "," + email + "," + password + "," + firstname + "," + lastname + "," + "," +
		"," + "," + "," + gender + "," + "," + "Markham/Toronto" + "," + "," + parentfirstname + "," +
		parentlastname + ", 5555555555 ," + parentemail + "," + photoconsent + "," + parentsignature + "\n"

	fmt.Print(text)

	// Write to the file
	f, err := os.OpenFile("tim_studentdata_golangGenerated.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		// If the file does not exist, create it first
		f, err = os.OpenFile("tim_studentdata_golangGenerated.csv", os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString("Course Code, Community Centre, Student E-mail, Password, First Name, " +
			"Last Name, Birth Year,Birth Month, Birth Day, Home Phone, Gender, Address, City, Postal Code, " +
			"Parent First Name, Parent Last Name, Cell Phone, Parent E-mail, Photo-Consent, " +
			"ParentSignature\n"); err != nil {
			panic(err)
		}

	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}

	// Respond back to the user
	w.Write([]byte("Success! Press back to add another student"))
}

/*
Allows the instructor to change the course and location.
All future requests coming in will use this
 */
func changeAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Parse arguments
	r.ParseForm()
	course_code = strings.Join(r.Form["course-code"], "")
	class = strings.Join(r.Form["course"], "")
	cc = strings.Join(r.Form["cc"], "")

	fmt.Println("\n\nUpdated course code: " + course_code)
	fmt.Println("Updated class: " + class)
	fmt.Println("Updated CC: " + cc + "\n")
	fmt.Println("Final course code being used: " + class + cc + course_code + "\n\n")

	w.Write([]byte("Success!"))
}

func main() {
	http.HandleFunc("/submit", submitForm)
	http.HandleFunc("/changeAdmin", changeAdmin)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Find IPs
	fmt.Println("Possible IPs of this server")
	ifaces, err := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err == nil{
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				fmt.Print(ip)
				fmt.Print("\n")
			}
		}

	}

	fmt.Println("\n\n====Remember to head to localhost:1231/admin.html to set course and CC!====")

	// Start Server
	err = http.ListenAndServe(":1231", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
