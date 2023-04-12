package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {

	// Creating the JSON Payload structure
	type Payload struct {
		Data string `json:"input_data"`
		Date int    `json:"date"`
	}

	// Creating a token using a custom claims type. RegisteredClaims is embedded in the custom type to allow for easy encoding, pargins and validation of registered claims
	type MyCustomClaims struct {
		Payload Payload
		jwt.RegisteredClaims
	}

	// The HEX Secret
	hmacSampleSecret := []byte("a9ddbcaba8c0ac1a0a812dc0c2f08514b23f2db0a68343cb8199ebb38a6d91e4ebfb378e22ad39c2d01d0b4ec9c34aa91056862ddace3fbbd6852ee60c36acbf")

	// Receiving "data" as parameter

	if len(os.Args) < 2 {
		fmt.Println("Expected 'create [input_data]', 'validate [jwt-token]', or 'exit' commands")
		os.Exit(1)
	}

	commands := os.Args[1]

	switch commands {
	case "create":
		input_data := os.Args[2]

		// Create the claims
		claims := MyCustomClaims{
			Payload{
				Data: input_data,
				Date: int(time.Now().Unix()),
			},
			jwt.RegisteredClaims{
				// A usual scenario is to set the expiration time relative to the current time
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    "test",
				Subject:   "somebody",
				ID:        "1",
				Audience:  []string{"somebody_else"},
			},
		}

		// Using UUID as JTI
		newJti, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Fatal(err)
		}

		payload := Payload{
			Data: input_data,
			Date: int(time.Now().Unix()), //change to date when entering data
		}

		// Create claims while leaving out some of the optional fields
		claims = MyCustomClaims{
			payload,
			jwt.RegisteredClaims{
				// Also fixed dates can be used for the NumericDate
				IssuedAt: jwt.NewNumericDate(time.Now()),
				ID:       string(newJti),
			},
		}

		//--------------------------

		// Create a new token object, specifying signing method and the claims
		// you would like it to contain.
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)

		fmt.Println(tokenString, err)
	case "validate":
		tokenString := os.Args[2]
		fmt.Println("IS YOUR TOKEN VALID????")
		//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiYXVkIjoic2luZ2xlIn0.QAWg1vGvnqRuCFTMcPkjZljXHh8U3L_qUjszOtQbeaA"

		tokenValid, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return hmacSampleSecret, nil
		})

		if claims, ok := tokenValid.Claims.(*MyCustomClaims); ok && tokenValid.Valid {
			fmt.Printf("Data: %v + Date: %v + JTI: %v + Issued At: %v", claims.Payload.Data, time.Unix(int64(claims.Payload.Date), 0), claims.RegisteredClaims.ID, claims.RegisteredClaims.IssuedAt)
		} else {
			fmt.Println(err)
		}

	case "exit":
		fmt.Println("Thanks for using me! Good bye!")
		os.Exit(1)

	}

}
