-- 创建数据库
SELECT 'CREATE DATABASE bookinfo'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'bookinfo')\gexec

-- 切换到 bookinfo 数据库
\c bookinfo;

-- 创建图书表
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    isbn VARCHAR(13) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    publisher VARCHAR(255) NOT NULL,
    language VARCHAR(50) NOT NULL,
    pages INTEGER NOT NULL,
    year INTEGER NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    description TEXT,
    cover_image VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 创建评论表
CREATE TABLE IF NOT EXISTS reviews (
    id SERIAL PRIMARY KEY,
    book_id INTEGER NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL,
    username VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 创建评分表
CREATE TABLE IF NOT EXISTS ratings (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
    reviewer VARCHAR(255) NOT NULL,
    comment TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 插入示例图书数据
INSERT INTO books (isbn, title, author, publisher, language, pages, year, price, description, cover_image) VALUES
('9787111213826', '深入理解计算机系统', 'Randal E. Bryant', '机械工业出版社', '中文', 756, 2016, 139.00, 
'本书从程序员的视角详细阐述计算机系统的本质概念，并展示这些概念如何实实在在地影响应用程序的正确性、性能和实用性。', 
'/static/img/csapp.jpg'),

('9787115546081', 'Go语言编程', '许式伟', '人民邮电出版社', '中文', 380, 2020, 89.00,
'本书由Go语言之父郝林（Rob Pike）作序推荐，系统讲解Go语言的特性，并通过实例讲解Go语言的使用。',
'/static/img/golang.jpg'),

('9787111641247', 'Kubernetes权威指南', '龚正', '机械工业出版社', '中文', 621, 2021, 149.00,
'本书详细介绍了Kubernetes的架构原理、核心概念和基本使用方法，是一本Kubernetes入门与实践的优秀教程。',
'/static/img/k8s.jpg');

-- 插入示例评论数据
INSERT INTO reviews (book_id, user_id, username, content) VALUES
(1, 1, 'alice', '这本书非常适合想要深入理解计算机系统的读者，内容详实，讲解清晰。'),
(1, 2, 'bob', '虽然内容很好，但是对于初学者来说可能有点难度。'),
(2, 1, 'alice', 'Go语言的入门书籍，对新手很友好。'),
(3, 3, 'charlie', '全面介绍了Kubernetes的各个方面，很实用的一本书。');

-- 插入示例评分数据
INSERT INTO ratings (product_id, rating, reviewer, comment) VALUES
(1, 5, 'alice', '非常好的一本书'),
(1, 4, 'bob', '内容详实'),
(1, 5, 'charlie', '经典著作'),
(2, 4, 'alice', '入门必读'),
(2, 5, 'david', '讲解清晰'),
(3, 5, 'charlie', '很实用'),
(3, 4, 'eve', '内容全面'); 