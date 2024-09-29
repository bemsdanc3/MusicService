-- Вставка пользователей
INSERT INTO users (login, email, password) VALUES
('john_doe', 'john@example.com', 'hashedpassword123'),
('jane_smith', 'jane@example.com', 'hashedpassword456');

-- Вставка исполнителей
INSERT INTO artists (artist_name) VALUES
('Coldplay'),
('Taylor Swift'),
('Ed Sheeran');

-- Вставка альбомов
INSERT INTO albums (album_name, artist_id) VALUES
('A Rush of Blood to the Head', 1),  -- Coldplay
('1989', 2),                         -- Taylor Swift
('Divide', 3);                        -- Ed Sheeran

-- Вставка жанров
INSERT INTO genres (name_genre) VALUES
('Pop'),
('Rock'),
('Alternative Rock');

-- Вставка треков
INSERT INTO tracks (title, duration, album_id, artist_id, genre_id) VALUES
('The Scientist', 309, 1, 1, 3),    -- Coldplay, A Rush of Blood to the Head, Alternative Rock
('Blank Space', 231, 2, 2, 1),      -- Taylor Swift, 1989, Pop
('Shape of You', 233, 3, 3, 1);     -- Ed Sheeran, Divide, Pop

-- Вставка плейлистов
INSERT INTO playlists (name_playlist, user_id) VALUES
('Morning Vibes', 1),
('Workout Hits', 2);

-- Вставка треков в плейлисты
INSERT INTO playlist_tracks (playlist_id, track_id) VALUES
(1, 1),   -- The Scientist в плейлисте "Morning Vibes"
(1, 2),   -- Blank Space в плейлисте "Morning Vibes"
(2, 3);   -- Shape of You в плейлисте "Workout Hits"

-- Вставка истории прослушиваний
INSERT INTO listening_history (user_id, track_id, timestamp) VALUES
(1, 1, '2024-09-29 08:00:00'),   -- John Doe слушал The Scientist
(1, 2, '2024-09-29 09:15:00'),   -- John Doe слушал Blank Space
(2, 3, '2024-09-29 10:45:00');   -- Jane Smith слушала Shape of You