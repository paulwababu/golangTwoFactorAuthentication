package main

import (
	"crypto/rand"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kingzbauer/africastalking-go/sms"
)

var (
	apiKey    = flag.String("k", "dcfb22e0bcc24fe4203bc7969d47ac67b0af58c94a721d145e9f69579c177364", "apiKey provided by AT")
	username  = flag.String("u", "PaulSaul", "username provided by AT")
	shortCode = flag.String("s", "", "Short code registered with your AT app")
	live      = flag.Bool("l", false, "Whether to make a live api call. Default is sandbox")
	// message   = flag.String("m", "hello", "Message to send")
	// number = flag.String("p", "+254797584194", "Phone number to receive the message")
)

//database and error variables
var db *sql.DB
var errr error

//variables for phone number
var phone string
var usernam string
var email string
var password string

//function for the sign up page
func signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.ServeFile(res, req, "signup.html")
		return
	}

	usernam = req.FormValue("username")
	phone = req.FormValue("phone")
	email = req.FormValue("email")
	password = req.FormValue("psw")

	http.Redirect(res, req, "/2fa", 301)
}

func twoFactor(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "2fa.html")
	}

	var RandomCrypto, _ = rand.Prime(rand.Reader, 12)
	var message = "Your four digit code is: " + RandomCrypto.String()
	var phoneNumber = phone
	flag.Parse()
	parseFromEnv(apiKey, username)
	// apiKey and username are compulsory values
	if len(*apiKey) == 0 || len(*username) == 0 || len(phoneNumber) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	srv := sms.NewService(*apiKey, *username, *shortCode, *live)
	rep, err := srv.Send(message, []string{phoneNumber}, "")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Response: %v\n", rep)
}

func parseFromEnv(apiKey, username *string) {
	if len(*apiKey) == 0 {
		*apiKey = os.Getenv("dcfb22e0bcc24fe4203bc7969d47ac67b0af58c94a721d145e9f69579c177364")
	}

	if len(*username) == 0 {
		*username = os.Getenv("PaulSaul")
	}
}

func main() {
	http.HandleFunc("/2fa", twoFactor)
	http.HandleFunc("/signup", signupPage)
	http.ListenAndServe(":8080", nil)

}
