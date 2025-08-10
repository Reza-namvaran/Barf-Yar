document.addEventListener('DOMContentLoaded', function() {
    const loginForm = document.getElementById('login-form');
    const errorMessage = document.getElementById('error-message');
    const successMessage = document.getElementById('success-message');
    const loading = document.getElementById('loading');
    const loginBtn = document.querySelector('.login-btn');

    loginForm.addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;
        
        errorMessage.style.display = 'none';
        successMessage.style.display = 'none';
        
        loading.style.display = 'block';
        loginBtn.disabled = true;
        
        try {
            const response = await fetch('/api/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    username: username,
                    password: password
                })
            });
            
            const data = await response.json();
            
            if (data.success) {
                successMessage.textContent = data.message;
                successMessage.style.display = 'block';
                
                setTimeout(() => {
                    window.location.href = '/dashboard';
                }, 1000);
            } else {
                errorMessage.textContent = data.message;
                errorMessage.style.display = 'block';
            }
        } catch (error) {
            console.error('Login error:', error);
            errorMessage.textContent = 'An error occurred. Please try again.';
            errorMessage.style.display = 'block';
        } finally {
            loading.style.display = 'none';
            loginBtn.disabled = false;
        }
    });
});
