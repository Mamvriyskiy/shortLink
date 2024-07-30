const emailInput = document.getElementById("email");
const passwordInput = document.getElementById("password");
const loginInput = document.getElementById("login");
const loginButton = document.getElementById("login-button");
const errorMessage = document.getElementById("error-message");


loginButton.addEventListener("click", () => {
    const email = emailInput.value;
    const password = passwordInput.value;
    const login = loginInput.value;

    if (email === "" || password === "" || login == "") {
        errorMessage.textContent = "Пожалуйста, введите email и пароль.";
        return;
    }

    console.log(JSON.stringify({ login, email, password }))
        fetch('/auth/sing-up/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                //'Authorization': jwt
            },
            body: JSON.stringify({ login: login, password: password, email: email })
        })
        .then(response => response.json())
        .then(data => {
            //document.getElementById('shortlink').innerText = data.shortlink;
            alert('Регистрация прошла успешно');
            window.location.href = '/auth/sign-in';
        })
        .catch(error => {
            alert('Ошибка в регистрации' + error.message);
        });
        errorMessage.textContent = "";
    // window.location.href = "/welcome"; // Перенаправление
    });