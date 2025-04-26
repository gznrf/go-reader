document.addEventListener('DOMContentLoaded', () => {
    const menuItems = document.querySelectorAll('.sidebar li');
  
    menuItems.forEach(item => {
      item.addEventListener('click', () => {
        const page = item.getAttribute('data-page');
        if (page) {
          window.location.href = page;
        }
      });
    });
  
    const uploadForm = document.getElementById('uploadForm');
  
    if (uploadForm) {
      uploadForm.addEventListener('submit', async (e) => {
        e.preventDefault();
  
        const fileInput = document.getElementById('file');
        const titleInput = document.getElementById('title');
        const tagsInput = document.getElementById('tags');
        const descriptionInput = document.getElementById('description');
  
        const formData = new FormData();
        formData.append('file', fileInput.files[0]);
        formData.append('title', titleInput.value);
        formData.append('tags', tagsInput.value);
        formData.append('description', descriptionInput.value);
  
        try {
          const response = await fetch('/api/upload', {
            method: 'POST',
            body: formData
          });
  
          if (response.ok) {
            alert('Книга успешно загружена!');
            uploadForm.reset();
          } else {
            alert('Ошибка загрузки книги!');
          }
        } catch (error) {
          console.error(error);
          alert('Ошибка сервера!');
        }
      });
    }
  });
  