package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getFrontendPath() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return "../../frontend"
	}
	return filepath.Join(currentDir, "..", "..", "frontend")
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(data)
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/health" {
			healthHandler(w, r)
			return
		}
		if r.URL.Path == "/api/courses" && r.Method == "GET" {
			coursesHandler(w, r)
			return
		}
		if r.URL.Path == "/api/register" && r.Method == "POST" {
			registerHandler(w, r)
			return
		}
		if r.URL.Path == "/api/login" && r.Method == "POST" {
			loginHandler(w, r)
			return
		}

		frontendPath := getFrontendPath()
		http.FileServer(http.Dir(frontendPath)).ServeHTTP(w, r)
	})(w, r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":   "healthy",
		"database": "connected",
	}
	sendJSONResponse(w, response)
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	courses, err := GetAllCourses()
	if err != nil {
		sendErrorResponse(w, "Failed to load courses", http.StatusInternalServerError)
		return
	}

	var coursesResponse []map[string]interface{}
	for _, course := range courses {
		coursesResponse = append(coursesResponse, map[string]interface{}{
			"id":        course.ID,
			"title":     course.Title,
			"language":  course.Language,
			"price":     course.Price,
			"page_path": course.PagePath,
		})
	}

	sendJSONResponse(w, coursesResponse)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user, err := RegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"success":  true,
		"message":  "User registered successfully",
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
	}
	sendJSONResponse(w, response)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user, err := LoginUser(req.Email, req.Password)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Login successful",
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"wallet":   user.WalletBalance,
		},
	}
	sendJSONResponse(w, response)
}

func main() {
	fmt.Println("🚀 Запуск сервера школы программирования...")
	InitDatabase()

	if err := CheckTables(); err != nil {
		log.Printf("⚠️ Ошибка таблиц: %v", err)
	}

	courses, err := GetAllCourses()
	if err != nil {
		log.Printf("⚠️ Ошибка получения курсов: %v", err)
	} else {
		fmt.Printf("✅ Загружено %d курсов\n", len(courses))
	}

	http.HandleFunc("/", mainHandler)

	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	fmt.Println("📊 Доступные эндпоинты:")
	fmt.Println("   GET  /api/health    - Проверка здоровья сервера")
	fmt.Println("   GET  /api/courses   - Получить все курсы")
	fmt.Println("   POST /api/register  - Регистрация пользователя")
	fmt.Println("   POST /api/login     - Вход пользователя")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
