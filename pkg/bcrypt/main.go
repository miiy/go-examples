package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main()  {
	password := "123456"
	// generate hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	hashedPasswordStr := string(hashedPassword)
	log.Println(hashedPasswordStr)

	// validate
	err = bcrypt.CompareHashAndPassword([]byte(hashedPasswordStr), []byte(password))
	if err != nil {
		log.Println(err)
	}
	log.Println("mached")

	// error password
	password += "111"
	err = bcrypt.CompareHashAndPassword([]byte(hashedPasswordStr), []byte(password))
	if err != nil {
		log.Println(err)
	}
}
