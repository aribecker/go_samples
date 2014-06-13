// https://gist.github.com/ijt/950790
package main

import (
	"encoding/xml"
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    )

func main() {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}
	type Result2 struct {
		Customers  []string `xml:"CUSTOMER"`
	}
	v := Result2{}
		//response, _, err := http.Get("http://golang.org/")
    //response, err := http.Get("http://golang.org/")
    response, err := http.Get("http://www.thomas-bayer.com/sqlrest/CUSTOMER")
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
		fmt.Printf("Before unmarshal\n")
		err = xml.Unmarshal([]byte(contents), &v)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		fmt.Printf("Customers: %v\n", v.Customers)
		/*
		fmt.Printf("XMLName: %#v\n", v.XMLName)
		fmt.Printf("Name: %q\n", v.Name)
		fmt.Printf("Phone: %q\n", v.Phone)
		fmt.Printf("Email: %v\n", v.Email)
		fmt.Printf("Groups: %v\n", v.Groups)
		fmt.Printf("Address: %v\n", v.Address)
		*/
    }
}
