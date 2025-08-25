package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID            uint    `gorm:"primaryKey"`
	Username      string  `gorm:"uniqueIndex;not null;size:50"`
	Email         string  `gorm:"uniqueIndex;not null;size:100"`
	Password      string  `gorm:"not null"` // Временно без хеширования
	WalletBalance float64 `gorm:"default:10000.0"`
}

type Course struct {
	ID       uint    `gorm:"primaryKey"`
	Title    string  `gorm:"not null;size:200"`
	Language string  `gorm:"not null;size:20"`
	Price    float64 `gorm:"not null"`
	PagePath string  `gorm:"not null"`
}

type Purchase struct {
	ID             uint    `gorm:"primaryKey"`
	UserID         uint    `gorm:"not null"`
	CourseID       uint    `gorm:"not null"`
	PurchaseAmount float64 `gorm:"not null"`
	Status         string  `gorm:"default:'active';size:20"`
}

func getDBPath() string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("⚠️ Ошибка получения текущей директории: %v", err)
		return "school.db"
	}
	dataDir := filepath.Join(currentDir, "..", "..", "data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Printf("⚠️ Ошибка создания папки data: %v", err)
		return "school.db"
	}
	return filepath.Join(dataDir, "school.db")
}

func InitDatabase() {
	dbPath := getDBPath()
	fmt.Printf("📊 База данных: %s\n", dbPath)

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Ошибка подключения к базе:", err)
	}

	// Автомиграция
	err = DB.AutoMigrate(&User{}, &Course{}, &Purchase{})
	if err != nil {
		log.Fatal("❌ Ошибка миграции:", err)
	}
	fmt.Println("✅ Таблицы созданы успешно")

	// Добавляем тестовые курсы
	initCourses()
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
			result := DB.Create(&course)
			if result.Error != nil {
				fmt.Printf("⚠️ Ошибка добавления курса %s: %v\n", course.Title, result.Error)
			} else {
				fmt.Printf("✅ Добавлен курс: %s\n", course.Title)
			}
		}
		fmt.Printf("✅ Всего добавлено %d курсов\n", len(courses))
	}
}

func GetAllCourses() ([]Course, error) {
	var courses []Course
	result := DB.Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func RegisterUser(username, email, password string) (*User, error) {
	user := &User{
		Username:      username,
		Email:         email,
		Password:      password,
		WalletBalance: 10000.0,
	}
	result := DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func LoginUser(email, password string) (*User, error) {
	var user User
	result := DB.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("пользователь не найден или неверный пароль")
	}
	return &user, nil
}

func CheckTables() error {
	// Проверяем существование таблиц
	if !DB.Migrator().HasTable(&User{}) {
		return fmt.Errorf("таблица users не существует")
	}
	if !DB.Migrator().HasTable(&Course{}) {
		return fmt.Errorf("таблица courses не существует")
	}
	if !DB.Migrator().HasTable(&Purchase{}) {
		return fmt.Errorf("таблица purchases не существует")
	}
	fmt.Println("✅ Все таблицы существуют")
	return nil
}
