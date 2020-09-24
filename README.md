# Golang Two Factor authentication implementation using AfricasTalking SMS Gateway

The web program is a fully functional registration page that has a sms two factor authentication using Africa's talking unofficial sdk
This is still under active development

# Requirements
go >= 1.13

# How To Run

Walk first:) Create a new database with a users table

mysql> CREATE TABLE truth(

    -> id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    
    -> username VARCHAR(50),
    
    -> phone VARCHAR(30),
    
    -> email VARCHAR(255),
    
    -> password VARCHAR(120)
    
    -> );
    
Go get both the required packages below

# Installations

go get github.com/kingzbauer/africastalking-go

_ "github.com/go-sql-driver/mysql"

go get golang.org/x/crypto/bcrypt

"github.com/JesusIslam/goinblue"

"database/sql"

