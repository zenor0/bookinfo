// 全局变量
let currentUser = '';
let bookId = '';

// 页面加载完成后执行
document.addEventListener('DOMContentLoaded', () => {
    // 从 URL 获取图书 ID
    const pathParts = window.location.pathname.split('/');
    bookId = pathParts[pathParts.length - 1];

    // 从 localStorage 恢复用户选择
    const savedUser = localStorage.getItem('selectedUser');
    if (savedUser) {
        document.getElementById('user-selector').value = savedUser;
        currentUser = savedUser;
    }

    // 加载图书详情
    loadBookDetails();
});

// 加载图书详情
async function loadBookDetails() {
    try {
        const headers = {};
        if (currentUser) {
            headers['end-user'] = currentUser;
        }

        const response = await fetch(`/api/v1/books/${bookId}`, {
            headers: headers
        });
        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.error || '加载图书详情失败');
        }

        displayBookDetails(data);
    } catch (error) {
        console.error('Error:', error);
        showError('加载图书详情失败');
    }
}

// 显示图书详情
function displayBookDetails(data) {
    // 显示评分
    const ratingDiv = document.getElementById('rating');
    if (data.rating) {
        const stars = '★'.repeat(Math.round(data.rating.average_rating));
        ratingDiv.innerHTML = `
            <span class="stars" style="color: ${data.reviews[0]?.color || '#f39c12'}">
                ${stars}
            </span>
            <span class="rating-value">(${data.rating.average_rating.toFixed(1)})</span>
        `;
    }

    // 显示评论
    const reviewsDiv = document.getElementById('reviews');
    reviewsDiv.innerHTML = '';

    if (data.reviews && data.reviews.length > 0) {
        data.reviews.forEach(review => {
            const reviewElement = document.createElement('div');
            reviewElement.className = 'review';
            reviewElement.innerHTML = `
                <div class="user">${review.username}</div>
                <div class="rating" style="color: ${review.color || '#f39c12'}">
                    ${review.rating_html || ''}
                </div>
                <div class="content">${review.content}</div>
                <div class="date">${new Date(review.created_at).toLocaleDateString()}</div>
            `;
            reviewsDiv.appendChild(reviewElement);
        });
    } else {
        reviewsDiv.innerHTML = '<p>暂无评论</p>';
    }
}

// 提交评论
async function submitReview(event) {
    event.preventDefault();

    if (!currentUser) {
        showError('请先选择用户');
        return;
    }

    const form = event.target;
    const content = form.content.value;
    const rating = parseInt(form.rating.value);

    try {
        const response = await fetch('/api/v1/reviews', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'end-user': currentUser
            },
            body: JSON.stringify({
                book_id: parseInt(bookId),
                user_id: 1, // TODO: 使用真实的用户ID
                username: currentUser,
                content: content,
                rating: rating
            })
        });

        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.error || '提交评论失败');
        }

        // 重新加载图书详情
        loadBookDetails();
        form.reset();
    } catch (error) {
        console.error('Error:', error);
        showError('提交评论失败');
    }
}

// 更新用户选择
function updateUser() {
    const userSelector = document.getElementById('user-selector');
    currentUser = userSelector.value;
    localStorage.setItem('selectedUser', currentUser);
    // 重新加载图书详情以获取新的评分展示
    loadBookDetails();
}

// 显示错误信息
function showError(message) {
    // TODO: 实现错误提示UI
    alert(message);
} 