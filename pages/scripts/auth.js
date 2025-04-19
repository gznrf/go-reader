document.addEventListener("DOMContentLoaded", () => {
  const loginForm = document.querySelector("#loginForm form");
  const registerForm = document.querySelector("#registerForm form");

  loginForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const email = document.getElementById("loginEmail").value.trim();
    const password = document.getElementById("loginPassword").value.trim();
    await sendRequest("authorization", email, password);
  });

  registerForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const email = document.getElementById("registerEmail").value.trim();
    const password = document.getElementById("registerPassword").value.trim();
    await sendRequest("registration", email, password);
  });
});

async function sendRequest(endpoint, email, password) {
  if (!email || !password) {
    alert("Поля не должны быть пустыми");
    return;
  }

  const payload = { email, password };

  try {
    const response = await fetch(`${API_HOST}/auth/${endpoint}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload)
    });

    const text = await response.text();
    let data;
    try {
      data = JSON.parse(text);
    } catch (e) {
      console.warn("Ответ не JSON:", text);
    }

    if (response.ok) {
      alert(endpoint === "authorization" ? "Успешный вход" : "Регистрация успешна");
    } else {
      alert((data && data.message) || "Ошибка запроса");
    }

  } catch (error) {
    console.error("Ошибка:", error);
    alert("Не удалось подключиться к серверу");
  }
}

function switchTo(form) {
  const login = document.getElementById("loginForm");
  const register = document.getElementById("registerForm");

  if (form === "login") {
    login.classList.add("active");
    register.classList.remove("active");
  } else {
    login.classList.remove("active");
    register.classList.add("active");
  }
}
