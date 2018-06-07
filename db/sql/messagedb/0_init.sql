CREATE TABLE message (
    message_id UUID NOT NULL DEFAULT uuid_generate_v4(),
    message TEXT NOT NULL,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id)
);