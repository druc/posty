-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  createdAt DATETIME NOT NULL
);

INSERT INTO posts (title, content, createdAt) VALUES ('The Joy of Coding', 'Coding is both an art and a science, blending creativity with logic.', '2024-08-01 10:15:00');
INSERT INTO posts (title, content, createdAt) VALUES ('Understanding Databases', 'Databases are essential for organizing, storing, and retrieving data efficiently.', '2024-08-02 14:30:00');
INSERT INTO posts (title, content, createdAt) VALUES ('JavaScript Tips', 'Here are some tips to improve your JavaScript coding skills.', '2024-08-03 09:45:00');
INSERT INTO posts (title, content, createdAt) VALUES ('Why Learn Python?', 'Python is versatile, beginner-friendly, and widely used in various fields.', '2024-08-04 11:00:00');
INSERT INTO posts (title, content, createdAt) VALUES ('Introduction to SQL', 'SQL is the language of databases, allowing you to interact with and manipulate data.', '2024-08-05 13:20:00');
INSERT INTO posts (title, content, createdAt) VALUES ('The Future of AI', 'Artificial intelligence is rapidly advancing, impacting many industries.', '2024-08-06 15:35:00');
INSERT INTO posts (title, content, createdAt) VALUES ('CSS Tricks', 'Master these CSS tricks to create visually appealing web pages.', '2024-08-07 08:10:00');
INSERT INTO posts (title, content, createdAt) VALUES ('Understanding APIs', 'APIs allow different software systems to communicate and share data.', '2024-08-08 16:45:00');
INSERT INTO posts (title, content, createdAt) VALUES ('Version Control with Git', 'Git is an essential tool for version control, enabling collaboration and tracking changes.', '2024-08-09 10:50:00');
INSERT INTO posts (title, content, createdAt) VALUES ('Building Responsive Websites', 'Responsive design ensures your website looks great on any device.', '2024-08-10 12:30:00');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
