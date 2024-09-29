drop database if exists music;
create database music;
use music;

CREATE TABLE users (
    ID INT NOT NULL PRIMARY KEY AUTO_INCREMENT UNIQUE,
    login VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE artists (
    id INT AUTO_INCREMENT PRIMARY KEY,
    artist_name VARCHAR(255)
);

CREATE TABLE albums (
    id INT AUTO_INCREMENT PRIMARY KEY,
    album_name VARCHAR(255) NOT NULL,
    artist_id INT,
    FOREIGN KEY (artist_id)
        REFERENCES artists (id)
);

CREATE TABLE genres (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name_genre VARCHAR(100) NOT NULL
);

CREATE TABLE tracks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    duration INT NOT NULL,
    album_id INT,
    artist_id INT,
    genre_id INT,
    FOREIGN KEY (album_id)
        REFERENCES albums (id),
    FOREIGN KEY (artist_id)
        REFERENCES artists (id),
    FOREIGN KEY (genre_id)
        REFERENCES genres (id)
);

CREATE TABLE playlists (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name_playlist VARCHAR(100) NOT NULL,
    user_id INT,
    FOREIGN KEY (user_id)
        REFERENCES users (id)
);

CREATE TABLE playlist_tracks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    playlist_id INT,
    track_id INT,
    FOREIGN KEY (playlist_id)
        REFERENCES playlists (id),
    FOREIGN KEY (track_id)
        REFERENCES tracks (id)
);

CREATE TABLE listening_history (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    track_id INT,
    timestamp DATETIME,
    FOREIGN KEY (user_id)
        REFERENCES users (id),
    FOREIGN KEY (track_id)
        REFERENCES tracks (id)
);