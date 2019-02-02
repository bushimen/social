CREATE TABLE bilibili (
  aid INTEGER PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  favourites INTEGER NOT NULL,
  coins INTEGER NOT NULL,
  likes INTEGER NOT NULL,
  dislikes INTEGER NOT NULL,
  danmaku INTEGER NOT NULL,
  comments INTEGER NOT NULL,
  shares INTEGER NOT NULL,
  views INTEGER NOT NULL,
  duration INTEGER NOT NULL,
  thumbnail TEXT NOT NULL,
  image TEXT NOT NULL,
  timestamp TIMESTAMP WITH TIME ZONE
);

CREATE INDEX bilibili_duration_idx on bilibili (duration);
CREATE INDEX bilibili_views_idx on bilibili (views);
CREATE INDEX bilibili_likes_idx on bilibili (likes);
CREATE INDEX bilibili_shares_idx on bilibili (shares);
CREATE INDEX bilibili_danmaku_idx on bilibili (danmaku);
CREATE INDEX bilibili_timestamp_idx on bilibili (timestamp);
