CREATE TABLE  songs (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,
    song_title VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
    lyrics TEXT NOT NULL,
    video_link VARCHAR(255)
);
