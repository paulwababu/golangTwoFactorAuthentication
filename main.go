package main

import (
	//"database/sql"
	//"database/sql"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	"math/rand"
	"net/http"
	"os"

	"strconv"

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

var err error

//variables for phone number
var phone string
var usernam string
var email string
var password string
var random string
var secretNumber int

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

	min, max := 1, 10000
	rand.Seed(time.Now().UnixNano())
	secretNumber = rand.Intn(max-min) + min
	var random = strconv.Itoa(secretNumber)

	var phoneNumber = phone
	flag.Parse()
	parseFromEnv(apiKey, username)
	// apiKey and username are compulsory values
	if len(*apiKey) == 0 || len(*username) == 0 || len(phoneNumber) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	srv := sms.NewService(*apiKey, *username, *shortCode, *live)
	rep, err := srv.Send(random, []string{phoneNumber}, "")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Response: %v\n", rep)
	http.Redirect(res, req, "/2fa", 301)
}

func twoFactor(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.ServeFile(res, req, "2fa.html")
	}

	//form details
	var codeEntered = req.FormValue("pass")

	code, _ := strconv.Atoi(codeEntered)
	//var user string
	//var phonE = phone
	//var usernamE = usernam
	//var emaiL = email
	//var passworD = password
	// err = db.QueryRow("SELECT username FROM truth WHERE username=?", usernam).Scan(&user)

	//validate if the code entered is the one sent by system
	if code == secretNumber {
		fmt.Println(code)
		fmt.Println(secretNumber)
		fmt.Println("//////CLIENT USERNAME && PHONE && PASSWORD//////")
		fmt.Println(usernam)
		fmt.Println(phone)
		fmt.Println(password)
		http.Redirect(res, req, "/", 301)
		return
		// errr := db.QueryRow("SELECT username FROM truth WHERE username=?", usernam).Scan(&user)
	}
	//retry again
	http.Redirect(res, req, "/2fa", 301)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
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
	//db, err = sql.Open("mysql", "paulsaul:443556126216621@/users")
	//if err != nil {
	///	panic(err.Error())
	//	}
	//	defer db.Close()
	//
	//	err = db.Ping()
	//	if err != nil {
	//		panic(err.Error())
	//	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/2fa", twoFactor)
	http.HandleFunc("/signup", signupPage)
	log.Println("listening on Port 8080")
	http.ListenAndServe(":8080", nil)

}
