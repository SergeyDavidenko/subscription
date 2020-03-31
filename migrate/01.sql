CREATE TABLE IF NOT EXISTS subscriptions_user(
   id VARCHAR (55) NOT NULL,
   user_id VARCHAR (55) NOT NULL,
   start_subscription BIGINT,
   expipe_subscription BIGINT,
   activate BOOLEAN DEFAULT false
);