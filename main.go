package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func main() {
	DB, err = gorm.Open(sqlite.Open("apirest.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		return
	}

	DB.AutoMigrate(&User{})

	router := mux.NewRouter()

	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/login", login).Methods("POST") // Añadida la ruta /login

	log.Fatal(http.ListenAndServe(":8080", router))
}

var jwtKey = []byte("clave_secreta_del_jwt")

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

//funcion login para el acceso de ususarios

func login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	userID, err := Authenticate(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	token, err := GenerateToken(userID)
	if err != nil {
		http.Error(w, "Error al generar el token JWT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
}

func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func Authenticate(email, password string) (uint, error) {
	var user User
	DB.Where("email = ?", email).First(&user)
	if user.ID == 0 || !CheckPasswordHash(password, user.Password) {
		return 0, fmt.Errorf("credenciales inválidas")
	}
	return user.ID, nil
}

func ValidateToken(r *http.Request) (uint, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return 0, fmt.Errorf("token no proporcionado")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("token inválido")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, fmt.Errorf("error al obtener los claims del token")
	}

	return claims.UserID, nil
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error al hashear la contraseña", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	DB.Create(&user)

	token, err := GenerateToken(user.ID)
	if err != nil {
		http.Error(w, "Error al generar el token JWT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	_, err := ValidateToken(r)
	if err != nil {
		http.Error(w, "Acceso no autorizado", http.StatusUnauthorized)
		return
	}

	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	_, err := ValidateToken(r)
	if err != nil {
		http.Error(w, "Acceso no autorizado", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	_, err := ValidateToken(r)
	if err != nil {
		http.Error(w, "Acceso no autorizado", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	if user.Password != "" {
		hashedPassword, err := HashPassword(user.Password)
		if err != nil {
			http.Error(w, "Error al hashear la contraseña", http.StatusInternalServerError)
			return
		}
		user.Password = hashedPassword
	}
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	_, err := ValidateToken(r)
	if err != nil {
		http.Error(w, "Acceso no autorizado", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("Usuario eliminado")
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
