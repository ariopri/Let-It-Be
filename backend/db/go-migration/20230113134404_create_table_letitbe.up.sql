CREATE TABLE kursus
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    users_id INT NOT NULL,
    deskripsi VARCHAR(255) NOT NULL,
    modul VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (users_id) REFERENCES users(id)
);

CREATE TABLE komentar
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    rate INT NOT NULL,
    content VARCHAR(255) NOT NULL,
    users_id INT NOT NULL,
    kursus_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (users_id) REFERENCES users(id),
    FOREIGN KEY (kursus_id) REFERENCES kursus(id)
);

CREATE TABLE subjects
(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE classes
(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    subject_id INT NOT NULL, 
    teacher VARCHAR(255) NOT NULL,
    rating INT NOT NULL,
    FOREIGN KEY (subject_id) REFERENCES subjects(id)
);