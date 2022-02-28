package handlers

import (
	"encoding/json"
	"fmt"
	"go-crud-article/connection"
	"go-crud-article/structs"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
func Login(w http.ResponseWriter, r *http.Request)  {
	// name     := r.URL.Query().Get("name")
	// password := r.URL.Query().Get("name")


	res := structs.Result{Code: 200, Message: "Success get articles"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	keyPass := make(map[string]string)
	keyVal := make(map[string]int)
    payloads, _ := ioutil.ReadAll(r.Body)

	name	 := 	keyVal["age"]
	
	var user structs.User
	
	var risk_profile structs.Risk_profile
	var param int = 55
	if param-name>=30 {
		risk_profile.STOCK = 72.5
		risk_profile.BOND = 21.5
		risk_profile.MM = 100-(risk_profile.STOCK+risk_profile.BOND)
	}
	if param-name>=20 && param-name<=29 {
		risk_profile.STOCK = 54.5
		risk_profile.BOND = 25.5
		risk_profile.MM = 100-(risk_profile.STOCK+risk_profile.BOND)
	}
	if param-name<20  {
		risk_profile.STOCK = 34.5
		risk_profile.BOND = 45.5
		risk_profile.MM = 100-(risk_profile.STOCK+risk_profile.BOND)
	}
	hash, _ := HashPassword(keyPass["password"]) 
	json.Unmarshal(payloads, &user)
	json.Unmarshal(payloads, &risk_profile)


	connection.DB.Create(&user)
	risk_profile.Userid = user.Userid
	user.Password = hash
	connection.DB.Create(&risk_profile)
	res := structs.Result{Code: 200, Data: user, Message: "Success create user"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func ListUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset := vars["offset"]

	articles := []structs.User{}

	connection.DB.
		Limit(limit).
		Offset(offset).
		
		Find(&articles)

	res := structs.Result{Code: 200, Data: articles, Message: "Success get articles"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetDetailUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]

	var article structs.Risk_profile
	connection.DB.Preload("User").First(&article, articleID)

	res := structs.Result{Code: 200, Data: article, Message: "Success get detail user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var articleUpdates structs.User
	
	json.Unmarshal(payloads, &articleUpdates)

	var article structs.User
	connection.DB.First(&article, articleID)
	connection.DB.Model(&article).Updates(&articleUpdates)

	res := structs.Result{Code: 200, Data: article, Message: "Success update article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]

	var article structs.User
	connection.DB.First(&article, articleID)
	connection.DB.Delete(&article)

	res := structs.Result{Code: 200, Message: "Success delete article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}