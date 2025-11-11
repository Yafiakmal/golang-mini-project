CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
CREATE TABLE posts (
    id BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    category_id BIGINT REFERENCES categories(id) NOT NULL,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    content TEXT NOT NULL,
    published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);
CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    post_id BIGINT REFERENCES posts(id) NOT NULL,
    user_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE tags (
    id BIGINT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
CREATE TABLE post_tags (
    post_id BIGINT REFERENCES posts(id),
    tag_id BIGINT REFERENCES tags(id),
    PRIMARY KEY (post_id, tag_id)
);