CREATE TABLE If NOT EXISTS comments(
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL NOT NULL,
  post_id BIGSERIAL NOT NULL,
  content TEXT NOT NULL,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
