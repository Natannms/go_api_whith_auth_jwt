package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Natannms/go_api/database"
	"github.com/Natannms/go_api/httpResponses"
	"github.com/Natannms/go_api/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key") // Altere essa chave secreta para uma de sua preferência

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	passwordBody := user.Password
	// Busca o usuário no banco de dados
	result := database.Conn().Where("email = ?", user.Email).First(&user)
	if result.Error != nil {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordBody))
	if err != nil {
		http.Error(w, "Senha incorreta", http.StatusUnauthorized)
		return
	}

	// // Gera um token JWT
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Não foi possível gerar o token JWT", http.StatusInternalServerError)
		return
	}

	// Retorna o token JWT para o cliente
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(httpResponses.LoginResponse{
		User: user,
		HttpResponseMessage: httpResponses.HttpResponseMessage{
			Message: "Logged",
		},
		Cookie: http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		},
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Remove o cookie do token de autenticação do cliente
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	// Responde com uma mensagem de sucesso em formato JSON
	response := map[string]string{"message": "Logout efetuado com sucesso"}
	json.NewEncoder(w).Encode(response)
}
