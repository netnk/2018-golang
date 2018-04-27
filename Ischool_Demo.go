package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/xml"
	"bufio"
  "os"
)

var (
	message Message
	result3 Body

	key string
	dsnsname string
	id string

)

func main(){

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=====  Ischool Staff Test  =====")
	fmt.Println("Ex.")
	fmt.Println("Token Key: 12345")
	fmt.Println("Dsna Name: abcde")
	fmt.Println("Staff Id: 12345")
	fmt.Println("Result: Return result or Null")
	fmt.Println("")

	fmt.Print("Token Key : ")
	scanner.Scan()
    key = scanner.Text()

	str, err := get_url("https://ischool/?grant_type=refresh_token&client_id=12345&client_secret=12345&refresh_token=" + key)
	if err != nil {
		log.Fatal(err)
	}

	
	json.Unmarshal([]byte(str), &message)
//	fmt.Printf("%s", message.Access_token)  //get_access_token
	
    fmt.Print("Dsns Name : ")
    scanner.Scan()
    dsnsname = scanner.Text()

    fmt.Print("Staff Id : ")
    scanner.Scan()
    student_id = scanner.Text()



	str2, err := get_url("https://ischool" + dsnsname + "/a/get?stt=PassportAccessToken&AccessToken=" + message.Access_token + "&parser=spliter&content=id:" + id)
	if err != nil {
		log.Fatal(err)
	}

	
	xml.Unmarshal(str2, &result3)	
	fmt.Printf("%s", result3)

	fmt.Scanln()

}



func get_url(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

   
	return result, err
}



type Message struct{
	Access_token string  `json:"access_token"`
	Expires_in int  `json:"expires_in"`
	Token_type string  `json:"token_type"`
	Scope string  `json:"scope"`
	Refresh_token string  `json:"refresh_token"`
	Static bool  `json:"static"`
}

type Body struct {
	Staff []Staff `xml:staff>staff`
}

type Staff struct {	
	name string `xml:name`
}
