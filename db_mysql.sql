DROP TABLE IF EXISTS votes;
DROP TABLE IF EXISTS category_nominations;
DROP TABLE IF EXISTS nominations;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS files;
DROP TABLE IF EXISTS voters;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id INT(11) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255),
    password VARCHAR(255),
    token VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW()
);

CREATE TABLE voters(
    id INT(11) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    regis_num VARCHAR(255),
    full_name VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW()
);

CREATE TABLE files(
    id INT(11) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    ori_name VARCHAR(255),
    hash_name VARCHAR(255),
    path VARCHAR(255),
    mime_type VARCHAR(255),
    size DOUBLE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW()
);

CREATE TABLE categories(
    id INT(11) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    file_id INT(11) UNSIGNED NOT NULL,
    name VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),

    FOREIGN KEY(file_id)
    REFERENCES files(id)
    ON DELETE CASCADE
);

CREATE TABLE nominations(
    id INT(11) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    file_id INT(11) UNSIGNED NOT NULL,
    full_name VARCHAR(255),
    dob DATE,
    poo VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),

    FOREIGN KEY(file_id)
    REFERENCES files(id)
    ON DELETE CASCADE
);

CREATE TABLE category_nominations(
	id INT(11) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    category_id INT(11) UNSIGNED NOT NULL,
    nomination_id INT(11) UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),

    FOREIGN KEY(category_id)
    REFERENCES categories(id)
    ON DELETE CASCADE,

    FOREIGN KEY(nomination_id)
    REFERENCES nominations(id)
    ON DELETE CASCADE
);

CREATE TABLE votes(
	id INT(11) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    voter_id INT(11) UNSIGNED NOT NULL,
    category_nomination_id INT(11) UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),

    FOREIGN KEY(category_nomination_id)
    REFERENCES category_nominations(id)
    ON DELETE CASCADE,

    FOREIGN KEY(voter_id)
    REFERENCES voters(id)
    ON DELETE CASCADE
);