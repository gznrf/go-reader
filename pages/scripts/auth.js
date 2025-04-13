// Здесь мы предполагаем, что конфиг загружается как строка JSON или YAML
const configYaml = `
api_base_url: "http://localhost:8080"
`;


// Теперь твой код
const formTitle = document.getElementById("form-title");
const nameInput = document.getElementById("name");
const toggleText = document.getElementById("toggle-text");
const toggleLink = document.getElementById("toggle-link");
const form = document.getElementById("auth-form");

let isLogin = true;

toggleLink.addEventListener("click", (e) => {
  e.preventDefault();
  isLogin = !isLogin;
  formTitle.textContent = isLogin ? "Вход" : "Регистрация";
  nameInput.classList.toggle("hidden", isLogin);
  toggleText.textContent = isLogin ? "Нет аккаунта?" : "Уже есть аккаунт?";
  toggleLink.textContent = isLogin ? "Зарегистрироваться" : "Войти";
});

form.addEventListener("submit", async (e) => {
  e.preventDefault();

  const name = nameInput.value.trim();
  const email = document.getElementById("email").value.trim();
  const password = document.getElementById("password").value.trim();

  // Подготовка данных для запроса
  const payload = isLogin
    ? { email, password }  // Для авторизации
    : { name, email, password };  // Для регистрации

  const url = isLogin
    ? `${config.api_base_url}/auth/authorization`  // Используем конфиг для авторизации
    : `${config.api_base_url}/auth/registration`;  // Используем конфиг для регистрации

  // Отправка данных на сервер
  try {
    const response = await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });

    const data = await response.json();

    if (response.ok) {
      alert(`${isLogin ? "Успешный вход" : "Регистрация успешна"}`);
      form.reset();
    } else {
      alert(data.message || "Ошибка сервера");
    }
  } catch (err) {
    console.error("Ошибка запроса:", err);
    alert("Сервер недоступен или ошибка сети.");
  }
});
