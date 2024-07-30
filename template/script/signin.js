const emailInput = document.getElementById("email");
        const passwordInput = document.getElementById("password");
        const loginButton = document.getElementById("login-button");
        const errorMessage = document.getElementById("error-message");

        loginButton.addEventListener("click", () => {
            const email = emailInput.value;
            const password = passwordInput.value;

            if (email === "" || password === "") {
                errorMessage.textContent = "Пожалуйста, введите email и пароль.";
                return;
            }

            fetch('/auth/sing-in/', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ email, password })
                })
                .then(response => response.json())
                .then(data => {
                    if (data.token) {
                        localStorage.setItem('jwt', data.token);
                        console.log('Authorization token:', data.token);
                        window.location.href = '/app/create';
                    } else {
                        alert('Login failed');
                    }
                })
                .catch(error => console.error('Error:', error));
                errorMessage.textContent = "";
            // window.location.href = "/welcome"; // Перенаправление
            });