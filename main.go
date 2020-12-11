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

type results struct {
	Addition       float64
	Subtraction    float64
	Multiplication float64
	Division       float64
}

type simpleinterest struct {
	Amount    float64
	Principal float64
	Rate      float64
	Time      float64
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := gin.Default()
	r.GET("/", LoadCalculator)
	r.POST("/Mathematical", LoadMathematics)
	r.POST("/Simpleinterest", LoadInterest)
	r.POST("/Calculatemaths", Mathematics)
	r.POST("/Calculateinterest", Interest)
	r.POST("/Reset", Reset)

	r.Run(":" + port)
	// r.Run(":8080")

}

func LoadCalculator(c *gin.Context) {
	t, err := template.ParseFiles("view/main.html")
	t.Execute(c.Writer, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadMathematics(c *gin.Context) {
	t, err := template.ParseFiles("view/index.html")
	t.Execute(c.Writer, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadInterest(c *gin.Context) {
	t, err := template.ParseFiles("view/interest.html")
	t.Execute(c.Writer, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Mathematics(c *gin.Context) {

	a := c.PostForm("Number")

	x, err := strconv.ParseFloat(a, 8)
	if err != nil {
		log.Fatal(err)
	}

	//b = append(b, i)
	d := []float64{x}
	b = append(b, d...)

	if len(b) < 2 {
		t, err := template.ParseFiles("view/indexcopy.html")
		t.Execute(c.Writer, nil)
		if err != nil {
			log.Fatal(err)
		}
	} else {

		r := results{
			Addition:       Add(),
			Subtraction:    Sub(),
			Multiplication: Mul(),
			Division:       Div(),
		}

		fmt.Println(r)

		t, err := template.ParseFiles("view/index.html")
		t.Execute(c.Writer, r)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func Add() float64 {
	var d float64
	d = 0
	for x, _ := range b {
		d = d + b[x]
	}
	return d
}

func Sub() float64 {
	var d float64
	d = 0

	for i := 0; i < (len(b) - 1); i++ {
		d = b[i] - b[i+1]
	}
	return d
}

func Mul() float64 {

	var d float64
	d = 1
	for x, _ := range b {
		d = d * b[x]
	}

	return d

}

func Div() float64 {
	var d float64
	d = 0

	for i := 0; i < (len(b) - 1); i++ {
		d = b[i] / b[i+1]
	}
	return d
}

func Reset(c *gin.Context) {
	b = nil
	t, err := template.ParseFiles("view/index.html")
	t.Execute(c.Writer, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Interest(c *gin.Context) {
	var a, p, r1, t float64

	n1 := c.PostForm("Amount")
	n2 := c.PostForm("Principal")
	n3 := c.PostForm("Rate")
	n4 := c.PostForm("Time")
	fmt.Println(n1, n2, n3, n4)

	// var np1 *string
	// np1 = &n1

	// var np2 *string
	// np2 = &n2

	// var np3 *string
	// np3 = &n3

	// var np4 *string
	// np4 = &n4

	a, err := strconv.ParseFloat(n1, 8)
	if err != nil {
		a = 0
	}

	fmt.Println(a)

	p, err1 := strconv.ParseFloat(n2, 8)
	if err1 != nil {
		p = 0
	}

	fmt.Println(p)

	r1, err2 := strconv.ParseFloat(n3, 8)
	if err2 != nil {
		r1 = 0
	}

	fmt.Println(r1)

	t, err3 := strconv.ParseFloat(n4, 8)
	if err3 != nil {
		t = 0
	}

	fmt.Println(t)

	if a == 0 {
		a = p * (1 + ((r1 / 100) * t))
	}

	if p == 0 {
		p = a / (1 + ((r1 / 100) * t))
	}

	if r1 == 0 {
		r1 = ((1 / t) * ((a / p) - 1)) * 100
	}

	if t == 0 {
		t = (1 / (r1 / 100)) * ((a / p) - 1)
	}

	roi := simpleinterest{
		Amount:    a,
		Principal: p,
		Rate:      r1,
		Time:      t,
	}

	fmt.Println(roi)

	t1, err5 := template.ParseFiles("view/interestanswers.html")
	t1.Execute(c.Writer, roi)
	if err5 != nil {
		log.Fatal(err)
	}

}
