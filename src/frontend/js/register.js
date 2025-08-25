document.getElementById('registerForm').addEventListener('submit', async function(e) {
    e.preventDefault();
    
    const formData = {
        username: document.getElementById('username').value,
        email: document.getElementById('email').value,
        password: document.getElementById('password').value
    };
    
    if (!formData.username || !formData.email || !formData.password) {
        showMessage('❌ Все поля обязательны для заполнения', 'error');
        return;
    }
    
    try {
        showMessage('🔄 Регистрация...', 'info');
        
        const response = await fetch('http://localhost:8080/api/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData)
        });
        
        const result = await response.json();
        
        if (response.ok) {
            showMessage('✅ Регистрация успешна!', 'success');
            setTimeout(() => {
                window.location.href = 'login.html';
            }, 2000);
        } else {
            showMessage('❌ Ошибка: ' + result.error, 'error');
        }
    } catch (error) {
        console.error('Ошибка сети:', error);
        showMessage('❌ Ошибка сети. Проверьте запущен ли сервер на localhost:8080', 'error');
    }
});

function showMessage(text, type) {
    const messageDiv = document.getElementById('message');
    messageDiv.textContent = text;
    messageDiv.className = `message ${type}`;
}