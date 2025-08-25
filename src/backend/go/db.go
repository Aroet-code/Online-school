package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID            uint      `gorm:"primaryKey"`
	Username      string    `gorm:"uniqueIndex;not null;size:50"`
	Email         string    `gorm:"uniqueIndex;not null;size:100"`
	PasswordHash  string    `gorm:"not null" json:"-"`
	WalletBalance float64   `gorm:"default:0.0"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

type Course struct {
	ID       uint    `gorm:"primaryKey"`
	Title    string  `gorm:"not null;size:200"`
	Language string  `gorm:"not null;size:20"`
	Price    float64 `gorm:"not null"`
	PagePath string  `gorm:"not null"`
}

type Purchase struct {
	ID             uint      `gorm:"primaryKey"`
	UserID         uint      `gorm:"not null"`
	CourseID       uint      `gorm:"not null"`
	PurchaseAmount float64   `gorm:"not null"`
	PurchaseDate   time.Time `gorm:"autoCreateTime"`
	Status         string    `gorm:"default:'active';size:20"`
}

// Функция для получения DSN (Data Source Name)
func getDSN() string {
	// Получаем параметры подключения из переменных окружения
	// или используем значения по умолчанию
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "school_system"
	}

	// Формат: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func InitDatabase() {
	dsn := getDSN()
	fmt.Printf("📊 Подключение к MySQL: %s\n", dsn)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("❌ Ошибка подключения к базе:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("❌ Ошибка получения DB объекта:", err)
	}

	// Настройки пула соединений для MySQL
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("❌ Ошибка ping базы данных:", err)
	}

	fmt.Println("✅ Соединение с MySQL установлено")

	// Создаем базу данных если она не существует
	err = DB.Exec("CREATE DATABASE IF NOT EXISTS school_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci").Error
	if err != nil {
		log.Fatal("❌ Ошибка создания базы данных:", err)
	}

	// Миграция таблиц
	err = DB.AutoMigrate(&User{}, &Course{}, &Purchase{})
	if err != nil {
		log.Fatal("❌ Ошибка миграции:", err)
	}

	fmt.Println("✅ Таблицы созданы успешно")
	initCourses()
	createAdminUser()
}

func initCourses() {
	var count int64
	DB.Model(&Course{}).Count(&count)

	if count == 0 {
		courses := []Course{
			{Title: "Go для начинающих", Language: "go", Price: 5000.0, PagePath: "program_study/Go.html"},
			{Title: "Python основы", Language: "python", Price: 6000.0, PagePath: "program_study/Python.html"},
			{Title: "C# программирование", Language: "csharp", Price: 4500.0, PagePath: "program_study/Csharp.html"},
		}

		for _, course := range courses {
			DB.Create(&course)
		}
		fmt.Printf("✅ Добавлено %d курсов\n", len(courses))
	}
}

func createAdminUser() {
	var adminCount int64
	DB.Model(&User{}).Where("username = ?", "admin").Count(&adminCount)

	if adminCount == 0 {
		hashedPassword, err := HashPassword("admin123")
		if err != nil {
			fmt.Printf("⚠️ Ошибка хеширования пароля админа: %v\n", err)
			return
		}

		admin := &User{
			Username:      "admin",
			Email:         "admin@school.ru",
			PasswordHash:  hashedPassword,
			WalletBalance: 200000000.0,
		}

		result := DB.Create(admin)
		if result.Error != nil {
			fmt.Printf("⚠️ Ошибка создания админа: %v\n", result.Error)
		} else {
			fmt.Printf("✅ Администратор создан: %s\n", admin.Username)
		}
	}
}

// Остальные функции остаются без изменений...
func GetAllCourses() ([]Course, error) {
	var courses []Course
	result := DB.Find(&courses)
	return courses, result.Error
}

func RegisterUser(username, email, password string) (*User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("ошибка хеширования пароля: %v", err)
	}

	user := &User{
		Username:     username,
		Email:        email,
		PasswordHash: hashedPassword,
	}

	result := DB.Create(user)
	if result.Error != nil {
		return nil, fmt.Errorf("ошибка создания пользователя: %v", result.Error)
	}

	return user, nil
}

func LoginUser(email, password string) (*User, error) {
	var user User
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("пользователь не найден")
	}

	if !CheckPasswordHash(password, user.PasswordHash) {
		return nil, fmt.Errorf("неверный пароль")
	}

	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("пользователь не найден")
	}
	return &user, nil
}

func GetUserByID(id uint) (*User, error) {
	var user User
	result := DB.First(&user, id)
	if result.Error != nil {
		return nil, fmt.Errorf("пользователь не найден")
	}
	return &user, nil
}

func GetCourseByID(id uint) (*Course, error) {
	var course Course
	result := DB.First(&course, id)
	if result.Error != nil {
		return nil, fmt.Errorf("курс не найден")
	}
	return &course, nil
}

func CreatePurchase(userID, courseID uint, amount float64) (*Purchase, error) {
	purchase := &Purchase{
		UserID:         userID,
		CourseID:       courseID,
		PurchaseAmount: amount,
		Status:         "active",
	}

	result := DB.Create(purchase)
	if result.Error != nil {
		return nil, fmt.Errorf("ошибка создания покупки: %v", result.Error)
	}

	return purchase, nil
}

func CheckTables() error {
	// Проверяем существование таблиц
	tables := []string{"users", "courses", "purchases"}
	for _, table := range tables {
		result := DB.Exec("SELECT 1 FROM " + table + " LIMIT 1")
		if result.Error != nil {
			return fmt.Errorf("таблица %s не существует или недоступна", table)
		}
	}
	return nil
}
