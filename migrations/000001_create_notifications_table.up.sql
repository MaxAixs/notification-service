CREATE TABLE notifications (
     id UUID PRIMARY KEY,
     user_id UUID NOT NULL,
     email VARCHAR(255) NOT NULL,
    item_id INT NOT NULL ,
    topic VARCHAR(255),
     body TEXT,
     status VARCHAR(50) NOT NULL DEFAULT 'not_sent',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    sent_at TIMESTAMP
);