-- Write your UP migration SQL here

-- Create Notification Table as per the NotificationEntity struct
CREATE TABLE public.notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    content VARCHAR(255) NOT NULL,
    is_read BOOLEAN DEFAULT false,
    href VARCHAR(255) NOT NULL,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id UUID NOT NULL,

    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);