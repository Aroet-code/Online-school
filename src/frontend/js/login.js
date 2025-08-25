document.getElementById('loginForm').addEventListener('submit', async function(e) {
    e.preventDefault();
    
    const formData = {
        email: document.getElementById('email').value,
        password: document.getElementById('password').value
    };
    
    try {
        showMessage('🔄 Вход...', 'info');
        
        const response = await fetch('http://localhost:8080/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData)
        });
        
        const result = await response.json();
        
        if (response.ok) {
            showMessage('✅ Вход выполнен успешно!', 'success');
            localStorage.setItem('user', JSON.stringify(result.user));
            setTimeout(() => {
                window.location.href = 'index.html';
            }, 1000);
        } else {
            showMessage('❌ Ошибка: ' + result.error, 'error');
        }
    } catch (error) {
        console.error('Ошибка сети:', error);
        showMessage('❌ Ошибка сети. Проверьте запущен ли сервер', 'error');
    }
});

function showMessage(text, type) {
    const messageDiv = document.getElementById('message');
    messageDiv.textContent = text;
    messageDiv.className = `message ${type}`;
}