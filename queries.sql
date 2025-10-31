-- Таблица режиссёров
CREATE TABLE directors (
    id SERIAL PRIMARY KEY,
    last_name TEXT NOT NULL,
    first_name TEXT NOT NULL
);

-- Таблица фильмов
CREATE TABLE films (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    release_date DATE NOT NULL,
    director_id INTEGER NOT NULL REFERENCES directors(id) ON DELETE RESTRICT,
    uuid UUID UNIQUE NOT NULL,
    rating DECIMAL(3,1) CHECK (rating >= 0 AND rating <= 10),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


-- Вставляем режиссёров
INSERT INTO directors (last_name, first_name) VALUES
    ('Нолан', 'Кристофер'),
    ('Кэмерон', 'Джеймс'),
    ('Тарантино', 'Квентин'),
    ('Скорсезе', 'Мартин'),
    ('Бертон', 'Тим'),
    ('Спилберг', 'Стивен'),
    ('Бондарчук', 'Фёдор'),
    ('Михалков', 'Никита');

-- Вставляем фильмы
INSERT INTO films (title, release_date, director_id, uuid, rating) VALUES
    ('Начало', '2010-07-16', 1, 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 8.8),
    ('Интерстеллар', '2014-11-06', 1, 'b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 8.6),
    ('Титаник', '1997-12-19', 2, 'c2eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 7.9),
    ('Аватар', '2009-12-18', 2, 'd3eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 7.8),
    ('Криминальное чтиво', '1994-10-14', 3, 'e4eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 8.9),
    ('Однажды в Голливуде', '2019-07-26', 3, 'f5eebc99-9c0b-4ef8-bb6d-6bb9bd380a16', 7.6),
    ('Отступники', '2006-09-26', 4, 'c7633f68-cba4-4314-b3c1-21f32864eb1c', 8.5),
    ('Волк с Уолл-стрит', '2013-12-25', 4, '823d41bc-2f8b-497a-8165-8e66a49ab047', 8.2);


-- режиссёры и количество их фильмов
SELECT
    d.id as director_id,
    d.first_name || ' ' || d.last_name as director_full_name,
    COUNT(f.id) as films_count
FROM directors d
    JOIN films f ON d.id = f.director_id
WHERE f.release_date >= '2000-01-01'
GROUP BY d.id, d.first_name, d.last_name
HAVING COUNT(f.id) >= 1
ORDER BY films_count DESC, director_full_name ASC;

CREATE INDEX idx_film_release_date ON films(release_date);


DROP TABLE films;
DROP TABLE directors;

DROP INDEX idx_film_release_date;
