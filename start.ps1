# Запуск локального сервера для проекта

Write-Host "Запуск локального сервера..."

# Если у тебя Node.js (npm run start):
# Start-Process -NoNewWindow npm run start

# Если Python:
# python -m http.server 8000

# Если нужно открыть сайт в браузере:
Start-Process "http://localhost:8000"

Write-Host "Сервер запущен! Нажми Ctrl+C для остановки."
