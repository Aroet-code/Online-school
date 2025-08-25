#!/bin/bash

echo "🚀 Запуск сервера школы программирования..."
echo "📦 Подготавливаем среду..."

# Создаем папку data если её нет
mkdir -p data

# Переходим в папку backend/go
cd src/backend/go

echo "🔧 Проверяем зависимости..."
go mod download 2>/dev/null || echo "⚠️ Предупреждение: не удалось скачать зависимости"

echo "✅ Готово к запуску"
echo "🎯 Запускаем сервер..."
echo ""
echo "📍 Сервер будет доступен по: http://localhost:8080"
echo "📊 API курсов: http://localhost:8080/api/courses"
echo "🩺 Проверка здоровья: http://localhost:8080/api/health"
echo ""
echo "📁 База данных создается в: ../../data/school.db"
echo "📁 Фронтенд обслуживается из: ../../frontend/"
echo ""
echo "⚡ Сервер запускается..."
echo "⏹️  Для остановки сервера нажмите Ctrl+C"
echo "----------------------------------------"

# Запускаем сервер напрямую
go run main.go db.go