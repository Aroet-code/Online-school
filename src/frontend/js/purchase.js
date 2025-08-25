// purchase.js - Обработчик страницы покупки курса

document.addEventListener('DOMContentLoaded', function() {
    const urlParams = new URLSearchParams(window.location.search);
    const courseId = urlParams.get('course_id');
    
    if (courseId) {
        loadCourseInfo(courseId);
    } else {
        showError('Не указан ID курса');
    }
    
    document.getElementById('purchaseForm').addEventListener('submit', async function(e) {
        e.preventDefault();
        await processPurchase(courseId);
    });
});

async function loadCourseInfo(courseId) {
    try {
        const response = await fetch('http://localhost:8080/api/courses');
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

function displayCourseInfo(course) {
    const purchaseInfo = document.getElementById('purchaseInfo');
    
    purchaseInfo.innerHTML = `
        <div class="course-purchase-info">
            <h3>${course.title}</h3>
            <div class="course-details">
                <p><strong>Язык:</strong> ${course.language.toUpperCase()}</p>
                <p><strong>Цена:</strong> ${course.price} руб.</p>
            </div>
        </div>
    `;
}

async function processPurchase(courseId) {
    const email = document.getElementById('email').value;
    
    if (!email) {
        showError('Пожалуйста, введите email для отправки чека');
        return;
    }
    
    try {
        showMessage('🔄 Обрабатываем покупку...', 'info');
        
        // Совершаем покупку
        const purchaseResponse = await fetch('http://localhost:8080/api/purchase', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                email: email,
                course_id: parseInt(courseId)
            })
        });
        
        const result = await purchaseResponse.json();
        
        if (purchaseResponse.ok) {
            showMessage('✅ ' + result.message + 
                       '\n💰 Новый баланс: ' + result.user.wallet + ' руб.', 'success');
            
            setTimeout(() => {
                window.location.href = 'index.html';
            }, 3000);
        } else {
            throw new Error(result.error || 'Ошибка при покупке');
        }
        
    } catch (error) {
        console.error('Ошибка покупки:', error);
        showError(error.message);
    }
}

function showMessage(text, type) {
    const messageDiv = document.getElementById('message');
    messageDiv.textContent = text;
    messageDiv.className = `message ${type}`;
    messageDiv.style.display = 'block';
}

function showError(text) {
    showMessage('❌ ' + text, 'error');
}