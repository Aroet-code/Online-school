// Проверка авторизации при загрузке
function checkAuth() {
    const user = JSON.parse(localStorage.getItem('user'));
    const authButtons = document.getElementById('authButtons');
    const userInfo = document.getElementById('userInfo');
    
    if (user) {
        authButtons.style.display = 'none';
        userInfo.style.display = 'flex';
        document.getElementById('userName').textContent = user.username;
        document.getElementById('userBalance').textContent = `${user.wallet} ₽`;
    }
}

// Выход из системы
function logout() {
    localStorage.removeItem('user');
    location.reload();
}

// Загрузка курсов
async function loadCourses() {
    const coursesContainer = document.getElementById('coursesContainer');
    
    try {
        coursesContainer.innerHTML = '<div class="loading">🔄 Загрузка курсов...</div>';
        
        const response = await fetch('http://localhost:8080/api/courses');
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const courses = await response.json();
        
        if (courses.length === 0) {
            coursesContainer.innerHTML = '<div class="loading">😕 Курсы не найдены</div>';
            return;
        }
        
        displayCourses(courses);
        
    } catch (error) {
        console.error('❌ Ошибка загрузки курсов:', error);
        coursesContainer.innerHTML = `
            <div class="error">
                <p>⚠️ Не удалось загрузить курсы</p>
                <p><strong>Ошибка:</strong> ${error.message}</p>
                <p>Проверьте: </p>
                <ul>
                    <li>Запущен ли сервер на localhost:8080</li>
                    <li>Доступен ли API: <a href="http://localhost:8080/api/courses" target="_blank">/api/courses</a></li>
                </ul>
                <button onclick="location.reload()">🔄 Обновить страницу</button>
            </div>
        `;
    }
}

// Отображение курсов
function displayCourses(courses) {
    const coursesContainer = document.getElementById('coursesContainer');
    
    coursesContainer.innerHTML = courses.map(course => `
        <div class="language-card ${course.language}-card">
            <div class="language-header">
                <div style="display: flex; align-items: center; flex: 1;">
                    <div class="language-icon">${getLanguageEmoji(course.language)}</div>
                    <div class="language-title">
                        <h2>${course.title}</h2>
                    </div>
                </div>
                <span class="difficulty ${course.language}">${getDifficulty(course.language)}</span>
            </div>
            
            <div class="language-content">
                <p class="purpose">${getLanguageDescription(course.language)}</p>
                
                <div class="price-container">
                    <div class="price-amount">${course.price}</div>
                    <div class="price-currency">руб.</div>
                </div>
            </div>
            
            <div class="course-actions">
                <a href="purchase.html?course_id=${course.id}" class="btn-buy">
                    💰 Купить
                </a>
                <a href="trial.html?course_id=${course.id}" class="btn-trial">
                    👀 Триал
                </a>
            </div>
        </div>
    `).join('');
}

// Emoji для языков
function getLanguageEmoji(language) {
    const emojis = {
        'go': '🐹',
        'python': '🐍', 
        'csharp': '🔷'
    };
    return emojis[language] || '🎓';
}

// Сложность курса
function getDifficulty(language) {
    const difficulties = {
        'go': 'Средний',
        'python': 'Начальный',
        'csharp': 'Средний'
    };
    return difficulties[language] || 'Начальный';
}

// Описание языка
function getLanguageDescription(language) {
    const descriptions = {
        'go': 'Go (Golang) - компилируемый язык от Google с упором на простоту и производительность. Идеален для сетевых приложений и микросервисов.',
        'python': 'Python - интерпретируемый язык с чистым синтаксисом. Популярен в data science, веб-разработке и искусственном интеллекте.',
        'csharp': 'C# - современный объектно-ориентированный язык от Microsoft. Применяется для разработки desktop, web, mobile и игр на Unity.'
    };
    return descriptions[language] || 'Современный язык программирования для профессиональной разработки.';
}

// Инициализация при загрузке страницы
document.addEventListener('DOMContentLoaded', function() {
    checkAuth();
    loadCourses();
});

// Глобальные функции
window.logout = logout;
window.loadCourses = loadCourses;