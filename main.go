package main

import (
	"encoding/json"
	"fmt"
	"gameapp/repository/mysql"
	"gameapp/service/userservice"
	"io"
	"log"
	"net/http"
)

const (
	JwtSignKey = "hgfhhkgghf"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", healthCheckHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)
	mux.HandleFunc("/users/login", userLoginHandler)
	mux.HandleFunc("/users/profile", userProfileHandler)
	log.Println("server is listening on port 8787....")
	http.ListenAndServe(":8787", mux)
}
func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error":"invalid method"}`)
		return
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":%s}`, err.Error())))
		return
	}
	var uReq userservice.RegisterRequest
	err = json.Unmarshal(data, &uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":%s}`, err.Error())))
		return
	}
	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo, JwtSignKey)
	_, err = userSvc.Register(uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":%s}`, err.Error())))
		return
	}
	writer.Write([]byte(`{"message":"user create"}`))
}
func healthCheckHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, `{"message":"everything is goood!"}`)

}
func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error":"invalid method"}`)
		return
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":%s}`, err.Error())))
		return
	}
	var lReq userservice.LoginRequest
	err = json.Unmarshal(data, &lReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":%s}`, err.Error())))
		return
	}
	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo, JwtSignKey)
	resp, err := userSvc.Login(lReq)
	data, err = json.Marshal(resp)

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":%s}`, err.Error())))
		return
	}
	writer.Write(data)
}
func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprintf(writer, `{"error":"invalid method"}`)
		return
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":%s}`, err.Error())))
		return
	}
	//jwtToken := req.Header.Get("Authorization")

	pReq := userservice.ProfileRequest{UserID: 0}
	err = json.Unmarshal(data, &pReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":%s}`, err.Error())))
		return
	}
	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo, JwtSignKey)
	resp, err := userSvc.Profile(pReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":%s}`, err.Error())))
		return
	}
	data, err = json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":%s}`, err.Error())))
		return
	}
	writer.Write(data)

}
