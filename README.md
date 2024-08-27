# sd-leaderboard

## Sortedset redis
- https://www.youtube.com/watch?v=5jwuDM6Z3F8

```bash
ZADD contest:12345 10 user:0
ZADD contest:12345 30 user:1
ZADD contest:12345 100 user:2
ZADD contest:12345 50 user:3
ZADD contest:12345 70 user:4
ZADD contest:12345 90 user:5
ZADD contest:12345 20 user:6
ZADD contest:12345 80 user:7
ZADD contest:12345 40 user:8
ZADD contest:12345 60 user:9

# top 5 highest score
ZREVRANGE contest:12345 0 4 WITHSCORES

# top 5 lowest score
ZRANGE contest:12345 0 4 WITHSCORES

ZCOUNT contest:12345 50 100
```

- Estimation
  - Entry: user_id (8 bytes) + score (8 bytes) + overhead ~ 20 bytes
  - Entries per contest (4 mil users): 20 * 4 * 10^6 bytes = 80 MB
  - 500K contests: 80 * 5 * 10^5 MB = 4 * 10^7 MB = 4 * 10^4 GB = 40 TB 