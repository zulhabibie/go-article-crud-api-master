package handlers

import (
	"go-crud-article/connection"

	"encoding/json"
	"go-crud-article/structs"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Login(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	//baca input user dari body
	var dataUser structs.User
	json.Unmarshal(payloads, &dataUser)
	name_data := dataUser.Name
	password_data := dataUser.Password
	//init data untuk pengecekan
	var tbUser structs.User
	connection.DB.Where("name = ?", name_data).First(&tbUser)
	name_tb := tbUser.Name
	customer_password_hash := tbUser.Password

	//validasi inputan
	if name_data != "" && password_data != "" {
		//cek username
		if name_data != name_tb {
			http.Error(w, "name not available", http.StatusBadRequest)
		} else {
			//cek password
			match := CheckPasswordHash(password_data, customer_password_hash)
			if !match {
				http.Error(w, "password not match", http.StatusBadRequest)
			} else {
				res := structs.Result{Code: 200, Data: name_tb, Message: "Success login"}
				result, err := json.Marshal(res)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(result)
			}
		}
	} else {
		http.Error(w, "invalid input", http.StatusBadRequest)
	}

}
