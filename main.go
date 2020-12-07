package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var b []float64

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := gin.Default()
	r.GET("/", LoadCalculator)
	r.POST("/receivenumber", InputNumber)
	r.POST("/addition", Add)
	r.POST("/subtraction", Sub)
	r.POST("/multiplication", Mul)
	r.POST("/division", Div)
	r.POST("/reset", Reset)
	r.Run(":" + port)
}

func LoadCalculator(c *gin.Context) {
	t, err := template.ParseFiles("view/index.html")
	t.Execute(c.Writer, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func InputNumber(c *gin.Context) {
	a := c.PostForm("Number")

	x, err := strconv.ParseFloat(a, 8)
	if err != nil {
		log.Fatal(err)
	}

	//b = append(b, i)
	d := []float64{x}
	b = append(b, d...)

	fmt.Println(b)

	t, err := template.ParseFiles("view/indexcopy.html")
	t.Execute(c.Writer, b)
	if err != nil {
		log.Fatal(err)
	}
}

func Add(c *gin.Context) {
	var d float64
	d = 0
	for x, _ := range b {
		d = d + b[x]
	}

	t, err := template.ParseFiles("view/index.html")
	t.Execute(c.Writer, d)
	if err != nil {
		log.Fatal(err)
	}
}

func Sub(c *gin.Context) {
	var d float64
	d = 0
	if len(b) > 1 {

		for i := 0; i < (len(b) - 1); i++ {
			d = b[i] - b[i+1]
		}

		t, err := template.ParseFiles("view/index.html")
		t.Execute(c.Writer, d)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		t, err := template.ParseFiles("view/index.html")
		t.Execute(c.Writer, "Please enter more than one number")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Mul(c *gin.Context) {
	var d float64
	d = 1
	for x, _ := range b {
		d = d * b[x]
	}

	t, err := template.ParseFiles("view/index.html")
	t.Execute(c.Writer, d)
	if err != nil {
		log.Fatal(err)
	}

}

func Div(c *gin.Context) {
	var d float64
	d = 0
	if len(b) > 1 {

		for i := 0; i < (len(b) - 1); i++ {
			d = b[i] / b[i+1]
		}

		t, err := template.ParseFiles("view/index.html")
		t.Execute(c.Writer, d)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		e := "Please enter more than one number"
		t, err := template.ParseFiles("view/index.html")
		t.Execute(c.Writer, e)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Reset(c *gin.Context) {
	b = nil
	t, err := template.ParseFiles("view/index.html")
	t.Execute(c.Writer, nil)
	if err != nil {
		log.Fatal(err)
	}
}
