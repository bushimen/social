CREATE TABLE instagram (
  shortcode VARCHAR(20) PRIMARY KEY,
  caption TEXT NOT NULL,
  tags TEXT[] NOT NULL,
  likes INTEGER NOT NULL,
  comments INTEGER NOT NULL,
  type VARCHAR(10) NOT NULL,
  thumbnail TEXT NOT NULL,
  image TEXT NOT NULL,
  timestamp TIMESTAMP WITH TIME ZONE
);

CREATE INDEX instagram_likes_idx on instagram (likes);
CREATE INDEX instagram_comments_idx on instagram (comments);
CREATE INDEX instagram_timestamp_idx on instagram (timestamp);
