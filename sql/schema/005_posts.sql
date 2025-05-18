-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(150) NOT NULL,        
    url VARCHAR(255) UNIQUE NOT NULL,    
    description VARCHAR(1000) NOT NULL, 
    published_at TIMESTAMP,    
    feed_id UUID,             
    FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;