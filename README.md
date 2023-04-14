# jwt-token
A jwt-token generator and validator

How to use

1. Install [go from brew](https://formulae.brew.sh/formula/go) or based in [your Operating System.](https://go.dev/doc/install)
2. [Download this repository to your favourite destination and using your preferred method](https://docs.github.com/es/repositories/creating-and-managing-repositories/cloning-a-repository)
    
   Assuming you are using Mac or Linux
3. Go inside the repository folder ` % cd jwt-token-main `
    a. Enter the following command ` jwt-token-main % go mod init jwt-token.go `
    b. Then  ` jwt-token-main % go mod tidy `
    c. Build ` jwt-token-main % go build jwt-token.go `
4.An executable will be created, to run the jwt-token use (do not include the square braquets):
    a. create [arg] -->  [arg] can be a valid alphanumeric string OR use quotation marks to have any characters as "input data". When creating the token, will be shown in terminal
    b. validate [jwt-token] this is only valid for CastLabs example, copy the token once you created it and paste it in [jwt-token] 
    c. exit will terminate the app

Example, once created the command, run: ` $ ./jwt-token create holamundo `  (Remember you can use quotation marks to include special characters and blank spaces, like. "Hello World! You & I will email @ a special place {}"

this will create a jwt-token 'eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7ImlucHV0X2RhdGEiOiJob2xhbXVuZG8iLCJkYXRlIjoxNjgxNDM4OTk3fSwiaWF0IjoxNjgxNDM4OTk3LCJqdGkiOiJCOUM3NzhBOC0wODVGLTRBQzUtQTQyRS1GNzRDQUQxQjhFQzNcbiJ9.HXgpwyB0gDsTaEfThKmsMiTjlyP0QkMfkQrYqqINlXcvZJ5ekfDc766U6KNkMW00D5k0Hk_H49QI69inysvN-Q <nil>'
    
To validate, run: ` $ ./jwt-token validate eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7ImlucHV0X2RhdGEiOiJob2xhbXVuZG8iLCJkYXRlIjoxNjgxNDM4OTk3fSwiaWF0IjoxNjgxNDM4OTk3LCJqdGkiOiJCOUM3NzhBOC0wODVGLTRBQzUtQTQyRS1GNzRDQUQxQjhFQzNcbiJ9.HXgpwyB0gDsTaEfThKmsMiTjlyP0QkMfkQrYqqINlXcvZJ5ekfDc766U6KNkMW00D5k0Hk_H49QI69inysvN-Q ` (do not include <nil>)
this will throw:
` $ IS YOUR TOKEN VALID????
Data: holamundo + Date: 2023-04-13 20:23:17 -0600 CST + JTI: B9C778A8-085F-4AC5-A42E-F74CAD1B8EC3
 + Issued At: 2023-04-13 20:23:17 -0600 CST`
 
 As you can see, all the data and date is included. 
 
 If you don't follow the examples, then, an error will appear and the app won't run anymore.
 
 
