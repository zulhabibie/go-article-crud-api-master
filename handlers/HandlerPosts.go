package handlers

import (
	"encoding/json"
	"fmt"
	"go-crud-article/connection"
	"go-crud-article/repositories"
	"go-crud-article/structs"
	"go-crud-article/utils"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}


func CreateUser(w http.ResponseWriter, r *http.Request) {
    payloads, _ := ioutil.ReadAll(r.Body)

	var dataUser structs.User
	var risk_profile structs.Risk_profile
	json.Unmarshal(payloads, &dataUser)
	nameUser := dataUser.Name

	var tabel_user structs.User
	connection.DB.Where("name = ?", nameUser).First(&tabel_user)


	hash, errorHash := HashPassword(dataUser.Password)
	if errorHash != nil {
		http.Error(w, errorHash.Error(), http.StatusBadRequest)
	}
	dataUser.Password = hash
	connection.DB.Create(&dataUser)
	
	var age = dataUser.Age
	risk_profile.Userid = dataUser.Userid
	
	paramAge := (55 - age)
	// var risk_profile structs.Risk_profile
	if paramAge>=30 {
		risk_profile.STOCK = 72.5
		risk_profile.BOND = 21.5
		risk_profile.MM = 100-(risk_profile.STOCK+risk_profile.BOND)
	}
	if paramAge>=20 && paramAge<=29 {
		risk_profile.STOCK = 54.5
		risk_profile.BOND = 25.5
		risk_profile.MM = 100-(risk_profile.STOCK+risk_profile.BOND)
	}
	if paramAge<20  {
		risk_profile.STOCK = 34.5
		risk_profile.BOND = 45.5
		risk_profile.MM = 100-(risk_profile.STOCK+risk_profile.BOND)
	}
	json.Unmarshal(payloads, &risk_profile)
	connection.DB.Create(&risk_profile)

	res := structs.Result{Code: 200, Data: dataUser.Userid, Message: "Success create user"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func ListUser(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// customerID := vars["userid"]

	// var risk_profile structs.Risk_profile

	// connection.DB.Where("userid = ?", customerID).First(&risk_profile)
	// connection.DB.Model(&risk_profile).Select([]string{"id", "name", "age"}).Related(&risk_profile.User, "Userid")
	// // connection.DB.Model(&risk_profile).Association("MasterUser")
	// // connection.DB.First(&risk_profile, customerID)

	// res := structs.Result{Code: 200, Data: risk_profile, Message: "Success get customers"}
	// result, err := json.Marshal(res)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(result)

	// vars := mux.Vars(r)
	// limit := vars["limit"]
	// offset := vars["offset"]

	// articles := []structs.User{}

	// connection.DB.
	// 	Limit(limit).
	// 	Offset(offset).
		
	// 	Find(&articles)

	// res := structs.Result{Code: 200, Data: articles, Message: "Success get articles"}
	// results, err := json.Marshal(res)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(results)
	
	pagination := utils.GeneratePagination(r)
	var user structs.User
	userLists, error := repositories.GetAllUsers(&user, &pagination)

	if error != nil {
		fmt.Println(error.Error())
	}
	res := structs.Result{Code: 200, Data: userLists, Message: "Success get customers"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)

}

func DetailUser(w http.ResponseWriter, r *http.Request) {
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