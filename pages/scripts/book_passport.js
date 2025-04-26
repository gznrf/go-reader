const books = [
    {
      id: 1,
      title: "Книга Первая",
      cover: "https://via.placeholder.com/300x450?text=Book+1",
      description: "Описание первой книги. Очень интересное и захватывающее!",
      tags: ["Фэнтези", "Приключения", "Бестселлер"]
    },
    {
      id: 2,
      title: "Книга Вторая",
      cover: "https://via.placeholder.com/300x450?text=Book+2",
      description: "Описание второй книги. Познавательная история.",
      tags: ["Научная фантастика", "Драма"]
    },
    {
      id: 3,
      title: "Книга Третья",
      cover: "https://via.placeholder.com/300x450?text=Book+3",
      description: "Описание третьей книги. Романтические приключения.",
      tags: ["Роман", "Приключения"]
    },
  ];
  
  function getBookById(id) {
    return books.find(book => book.id === id);
  }
  
  function loadBook() {
    const params = new URLSearchParams(window.location.search);
    const bookId = 1 // parseInt(params.get('id'));
  
    const book = getBookById(bookId);
  
    if (!book) {
      document.getElementById('book-details').innerHTML = "<h2>Книга не найдена</h2>";
      return;
    }
  
    const bookDetails = `
      <div class="book-page">
        <div class="book-cover-large" style="background-image: url('${book.cover}');"></div>
  
        <div class="book-info">
          <div class="book-title">${book.title}</div>
  
          <div class="book-description">${book.description}</div>
  
          <div class="book-tags">
            ${book.tags.map(tag => `<span class="tag">${tag}</span>`).join('')}
          </div>
  
          <div class="book-buttons">
            <button class="button">⭐ В избранное</button>
            <button class="button">📖 Читать</button>
          </div>
        </div>
      </div>
    `;
  
    document.getElementById('book-details').innerHTML = bookDetails;
  }
  
  document.addEventListener('DOMContentLoaded', loadBook);
  