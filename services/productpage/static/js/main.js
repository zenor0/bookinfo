// 全局变量
let currentPage = 1;
let pageSize = 10;
let currentUser = '';

// 页面加载完成后执行
document.addEventListener('DOMContentLoaded', () => {
    loadBooks();
    // 从 localStorage 恢复用户选择
    const savedUser = localStorage.getItem('selectedUser');
    if (savedUser) {
        document.getElementById('user-selector').value = savedUser;
        currentUser = savedUser;
    }
});

// 加载图书列表
async function loadBooks(page = currentPage) {
    try {
        const response = await fetch(`/api/v1/books?page=${page}&page_size=${pageSize}`);
        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.error || '加载图书失败');
        }

        displayBooks(data.data);
        updatePagination(data.meta);
    } catch (error) {
        console.error('Error:', error);
        showError('加载图书列表失败');
    }
}

// 显示图书列表
function displayBooks(books) {
    const bookList = document.getElementById('book-list');
    bookList.innerHTML = '';

    books.forEach(book => {
        const bookCard = document.createElement('div');
        bookCard.className = 'book-card';
        bookCard.innerHTML = `
            <img src="${book.cover_image}" alt="${book.title}">
            <div class="content">
                <h3><a href="/books/${book.id}">${book.title}</a></h3>
                <p class="author">作者：${book.author}</p>
                <p class="price">¥${book.price}</p>
            </div>
        `;
        bookList.appendChild(bookCard);
    });
}

// 更新分页信息
function updatePagination(meta) {
    document.getElementById('page-info').textContent = `第 ${meta.page} 页`;
    document.getElementById('prev-page').disabled = meta.page <= 1;
    document.getElementById('next-page').disabled = meta.page * meta.page_size >= meta.total;
    currentPage = meta.page;
}

// 搜索图书
async function searchBooks() {
    const query = document.getElementById('search').value.trim();
    if (!query) return;

    try {
        const response = await fetch(`/api/v1/books/search?q=${encodeURIComponent(query)}`);
        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.error || '搜索失败');
        }

        displayBooks(data);
    } catch (error) {
        console.error('Error:', error);
        showError('搜索图书失败');
    }
}

// 更新用户选择
function updateUser() {
    const userSelector = document.getElementById('user-selector');
    currentUser = userSelector.value;
    localStorage.setItem('selectedUser', currentUser);
}

// 上一页
function prevPage() {
    if (currentPage > 1) {
        loadBooks(currentPage - 1);
    }
}

// 下一页
function nextPage() {
    loadBooks(currentPage + 1);
}

// 显示错误信息
function showError(message) {
    // TODO: 实现错误提示UI
    alert(message);
} 