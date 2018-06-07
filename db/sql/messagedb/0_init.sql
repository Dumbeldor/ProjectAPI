CREATE TABLE message (
    message_id UUID NOT NULL DEFAULT uuid_generate_v4(),
    message TEXT NOT NULL,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    user_sender_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    user_receiver_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    PRIMARY KEY (message_id)
);