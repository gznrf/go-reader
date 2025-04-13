// Валидация email
function validateEmail(email) {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]{2,}$/i;
    return re.test(email);
  }
  
  // Проверка на SQL инъекции
  function hasSQLInjection(input) {
    const sqlPattern = /('|--|;|\/\*|\*\/|xp_|drop|select|insert|update|delete)/i;
    return sqlPattern.test(input);
  }
  
  // Валидация данных
  function validateForm(name, email, password, isLogin) {
    if (!email || !password || (!isLogin && !name)) {
      return "Пожалуйста, заполните все поля.";
    }
  
    if (!validateEmail(email)) {
      return "Введите корректный email.";
    }
  
    if (password.length < 6) {
      return "Пароль должен быть не короче 6 символов.";
    }
  
    if (!isLogin && hasSQLInjection(name)) {
      return "Имя содержит подозрительные символы.";
    }
  
    return null; // Нет ошибок
  }
  