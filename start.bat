@echo off
chcp 65001 >nul
echo 🚀 Запуск сервера школы программирования...
echo 📦 Подготавливаем среду...

:: Создаем папку data если её нет
if not exist data mkdir data

:: Переходим в папку backend/go
cd src\backend\go

:: Проверяем, установлен ли Go
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo ❌ Go не найден. Установите Go с https://go.dev/dl/
    pause
    exit /b
)

echo 🔧 Проверяем зависимости...
go mod tidy

echo ✅ Готово к запуску
echo 🎯 Запускаем сервер...
echo.
echo 📍 Сервер будет доступен по: http://localhost:8080
echo 📊 API курсов: http://localhost:8080/api/courses
echo 🩺 Проверка здоровья: http://localhost:8080/api/health
echo 📁 База данных создается в: ..\..\data\school.db
echo 📁 Фронтенд обслуживается из: ..\..\frontend\
echo ----------------------------------------
echo ⏹️  Для остановки сервера закройте это окно
echo.

:: Запускаем сервер
go run main.go db.go
