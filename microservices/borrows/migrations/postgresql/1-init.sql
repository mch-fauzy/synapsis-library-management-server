-- +migrate Up
CREATE TABLE borrows (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    book_id INT NOT NULL,
    borrow_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    return_date TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(36) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(36) NOT NULL,
    deleted_at TIMESTAMP NULL,
    deleted_by VARCHAR(36) NULL
);

-- +migrate Down
DROP TABLE IF EXISTS borrows;
