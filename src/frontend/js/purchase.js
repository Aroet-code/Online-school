// purchase.js - Обработчик страницы покупки курса

document.addEventListener('DOMContentLoaded', function() {
    // Получаем ID курса из URL параметров
    const urlParams = new URLSearchParams(window.location.search);
    const courseId = urlParams.get('course_id');
    
    if (courseId) {
        loadCourseInfo(courseId);
    } else {
        showError('Не указан ID курса');
    }
    
    // Обработчик формы покупки
    document.getElementById('purchaseForm').addEventListener('submit', async function(e) {
        e.preventDefault();
        await processPurchase(courseId);
    });
});

// Загрузка информации о курсе
async function loadCourseInfo(courseId) {
    try {
        const response = await fetch('/api/courses');
        const courses = await response.json();
        
        const course = courses.find(c => c.id == courseId);
        if (course) {
            displayCourseInfo(course);
        } else {
            showError('Курс не найден');
        }
    } catch (error) {
        console.error('Ошибка загрузки курса:', error);
        showError('Не удалось загрузить информацию о курсе');
    }
}

// Отображение информации о курсе
function displayCourseInfo(course) {
    const purchaseInfo = document.getElementById('purchaseInfo');
    
    purchaseInfo.innerHTML = `
        <div class="course-purchase-info">
            <h3>${course.title}</h3>
            <div class="course-details">
                <p><strong>Язык:</strong> ${course.language.toUpperCase()}</p>
                <p><strong>Уровень:</strong> ${course.level}</p>
                <p><strong>Длительность:</strong> ${course.duration_days} дней</p>
                <p class="course-price"><strong>Цена:</strong> ${course.price} руб.</p>
            </div>
        </div>
    `;
}

// Обработка покупки
async function processPurchase(courseId) {
    const email = document.getElementById('email').value;
    
    if (!email) {
        showError('Пожалуйста, введите email');
        return;
    }
    
    try {
        showMessage('🔄 Обрабатываем покупку...', 'info');
        
        // Сначала получаем информацию о пользователе
        const userResponse = await fetch('/api/user', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email: email })
        });
        
        if (!userResponse.ok) {
            const errorData = await userResponse.json();
            throw new Error(errorData.error || 'Пользователь не найден');
        }
        
        const userData = await userResponse.json();
        
        // Затем совершаем покупку
        const purchaseResponse = await fetch('/api/purchase', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                user_id: userData.id,
                course_id: courseId
            })
        });
        
        const result = await purchaseResponse.json();
        
        if (purchaseResponse.ok) {
            showMessage('✅ Покупка успешно завершена!', 'success');
            
            // Перенаправляем на главную страницу через 2 секунды
            setTimeout(() => {
                window.location.href = 'index.html';
            }, 2000);
        } else {
            throw new Error(result.error || 'Ошибка при покупке');
        }
        
    } catch (error) {
        console.error('Ошибка покупки:', error);
        showError(error.message);
    }
}

// Показать сообщение
function showMessage(text, type) {
    const messageDiv = document.getElementById('message');
    messageDiv.textContent = text;
    messageDiv.className = `message ${type}`;
    messageDiv.style.display = 'block';
}

// Показать ошибку
function showError(text) {
    showMessage('❌ ' + text, 'error');
}

// Показать информационное сообщение
function showInfo(text) {
    showMessage('ℹ️ ' + text, 'info');
}