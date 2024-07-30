const longUrlInput = document.getElementById("long-url");
const shortenButton = document.getElementById("shorten-button");
const shortLinkContainer = document.getElementById("short-link-container");

shortenButton.addEventListener("click", () => {
    const longlink = longUrlInput.value;

    jwt = localStorage.getItem('jwt');
    if (jwt) {
        console.log("Token:" + jwt)
    } else {
        console.log("Token:" + jwt)
        alert('Пройдите регистрацию или войдите в аккаунт')
        window.location.href = '/auth/sign-up';
        return;
    }

    if (longlink === "") {
        alert("Введите ссылку!");
        return;
    }

    //'Authorization': 'Bearer ' + jwt
    console.log(JSON.stringify({ longlink }))
    fetch('/api/create/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': jwt
        },
        body: JSON.stringify({ longlink: longlink })
    })
    .then(response => response.json())
    .then(data => {
        const shortURL = data.shortlink
        //alert(data.shortlink);
        //document.getElementById('shortlink').innerText = data.shortlink;
        const shortLinkElement = document.createElement("a");
        shortLinkElement.href = shortURL;
        shortLinkElement.textContent = shortURL;
        shortLinkElement.classList.add("short-link");
        shortLinkContainer.innerHTML = ""; // Очистка предыдущего результата
        shortLinkContainer.appendChild(shortLinkElement);
    })
    .catch(error => {
        alert('Ошибка создания короткой ссылки: ' + error.message);
    });

    //const shortUrl = "https://your-short-link.com/short-link"; // Замените на реальную короткую ссылку
});
