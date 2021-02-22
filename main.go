package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

type ToJsObject map[string]interface{}

func (e *ToJsObject) Convert() (jsObject []byte) {
	element := reflect.ValueOf(e).Elem()
	//mapKeys := element.MapKeys()
	//typeField := element.Type()

	switch element.Kind() {
	case reflect.Map:
		fmt.Printf("string: %v\n", element.Type().String())
		fmt.Printf("name: %v\n", element.Type().Name())
		fmt.Printf("keys: %v\n", element.MapKeys())

		for _, i := range element.MapKeys() {
			field := i.Type().Key()
			typeField := i.Type().String()

			_ = field
			_ = typeField
		}

		//e := element.Elem()
		//for i := 0; i < e.NumField(); i += 1 {
		//	field := e.Field(i)
		//	typeField := e.Type().Field(i)
		//
		//	_ = field
		//	_ = typeField
		//}
	case reflect.Bool:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Int:
	case reflect.Slice:
	case reflect.String:
	case reflect.Struct:
	case reflect.Interface:
		fmt.Printf("name: %v\n", element.Type().Name())
		fmt.Printf("key: %v\n", element.Type().Key())
	}

	fmt.Printf("kind: %v\n", element.Kind())

	os.Exit(1)

	for i := 0; i < element.NumField(); i += 1 {
		field := element.Field(i)
		typeField := element.Type().Field(i)

		_ = field
		_ = typeField
	}

	return nil
}

func main() {
	var err error

	//var t ToJsObject = ToJsObject{
	//	"key": ToJsObject{
	//		"width": "100px",
	//	},
	//	"key_two": 1,
	//}
	//t.Convert()

	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})
	r.GET("/saveTimeLine", saveTimeLine)

	log.Println("Listening on :3000...")
	err = r.Run(":3000")
	if err != nil {
		log.Panic(err)
	}

}

func saveTimeLine(c *gin.Context) {
	var err error
	var body []byte

	body, err = ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", body)
}
