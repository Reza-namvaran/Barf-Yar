document.addEventListener('DOMContentLoaded', function() {
    const logoutLink = document.querySelector('.nav-link.logout');
    if (logoutLink) {
        logoutLink.addEventListener('click', async function(e) {
            e.preventDefault();
            
            try {
                const response = await fetch('/api/logout', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    }
                });
                
                if (response.ok) {
                    window.location.href = '/';
                }
            } catch (error) {
                console.error('Logout error:', error);
                window.location.href = '/';
            }
        });
    }

    const currentPath = window.location.pathname;
    const navLinks = document.querySelectorAll('.nav-link');
    
    navLinks.forEach(link => {
        if (link.getAttribute('href') === currentPath) {
            link.classList.add('active');
        }
    });
});
