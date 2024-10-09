package main

import (
    "html/template"
    "net/http"
    "strconv"
)

type Student struct {
    FirstName string
    LastName  string
    Age       int
    Gender    string
}

type Class struct {
    Name        string
    Department  string
    Level       string
    StudentList []Student
}

type User struct {
    FirstName string
    LastName  string
    BirthDate string
    Gender    string
}

var viewCount int

func promoHandler(w http.ResponseWriter, r *http.Request) {
    students := []Student{
        {FirstName: "John", LastName: "Doe", Age: 20, Gender: "masculin"},
        {FirstName: "Jane", LastName: "Smith", Age: 19, Gender: "féminin"},
    }

    class := Class{
        Name:        "B1 Informatique",
        Department:  "Informatique",
        Level:       "Bachelor 1",
        StudentList: students,
    }

    tmpl, err := template.ParseFiles("templates/promo.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, class)
}

func changeHandler(w http.ResponseWriter, r *http.Request) {
    viewCount++
    var message string
    if viewCount%2 == 0 {
        message = "Le nombre de vues est pair : " + strconv.Itoa(viewCount)
    } else {
        message = "Le nombre de vues est impair : " + strconv.Itoa(viewCount)
    }

    tmpl, err := template.ParseFiles("templates/change.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, message)
}

func userFormHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/form.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}

func userTreatmentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    firstName := r.FormValue("firstName")
    lastName := r.FormValue("lastName")
    birthDate := r.FormValue("birthDate")
    gender := r.FormValue("gender")

    if len(firstName) == 0 || len(lastName) == 0 || len(firstName) > 32 || len(lastName) > 32 {
        http.Error(w, "Erreur : Nom ou prénom invalide", http.StatusBadRequest)
        return
    }

    user := User{
        FirstName: firstName,
        LastName:  lastName,
        BirthDate: birthDate,
        Gender:    gender,
    }

    tmpl, err := template.ParseFiles("templates/display.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, user)
}

func main() {
    http.HandleFunc("/promo", promoHandler)
    http.HandleFunc("/change", changeHandler)
    http.HandleFunc("/user/form", userFormHandler)
    http.HandleFunc("/user/treatment", userTreatmentHandler)

	http.ListenAndServe("localhost:8080", nil)
}
