function navigateHome() {
    loadBooks(); 
}


function loadBooks() {
    const books = [
        { title: "Книга 1", cover: "https://via.placeholder.com/150x220?text=Book+1" },
        { title: "Книга 2", cover: "https://via.placeholder.com/150x220?text=Book+2" },
        { title: "Книга 3", cover: "https://via.placeholder.com/150x220?text=Book+3" },
        { title: "Книга 4", cover: "https://via.placeholder.com/150x220?text=Book+4" },
    ];

    const mainContent = document.getElementById('main-content');
    mainContent.innerHTML = '';

    books.forEach(book => {
        const bookCard = document.createElement('div');
        bookCard.className = 'book-card';
        bookCard.style.backgroundImage = `url('${book.cover}')`;
        bookCard.onclick = () => {window.location.href = "book_passport.html";}
        mainContent.appendChild(bookCard);
    });
}

// При загрузке страницы
document.addEventListener('DOMContentLoaded', loadBooks);
