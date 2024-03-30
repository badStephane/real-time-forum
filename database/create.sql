CREATE TABLE user (
 id INTEGER NOT NULL PRIMARY KEY,
 nickname VARCHAR(30) NOT NULL,
 passwrd VARCHAR(100) NOT NULL,
 email VARCHAR(30) NOT NULL,
 fname VARCHAR(30) NOT NULL,
 lname VARCHAR(30) NOT NULL,
 age INTEGER NOT NULL,
 gender VARCHAR(10) NOT NULL,
 created_at DATETIME NOT NULL
);

CREATE TABLE category (
 id INTEGER NOT NULL PRIMARY KEY,
 category_name VARCHAR(30) NOT NULL,
 descript VARCHAR(100),
 created_at DATETIME NOT NULL
);


CREATE TABLE post (
 id INTEGER NOT NULL PRIMARY KEY,
 user_id INTEGER NOT NULL,
 title VARCHAR(30) NOT NULL,
 content TEXT NOT NULL,
 created_at DATETIME NOT NULL,
 updated_at DATETIME NOT NULL,
 FOREIGN KEY(user_id) REFERENCES user(id)
);

CREATE TABLE comment (
 id INTEGER NOT NULL PRIMARY KEY,
 user_id INTEGER NOT NULL,
 post_id INTEGER NOT NULL,
 content TEXT NOT NULL,
 created_at DATETIME NOT NULL,
 updated_at DATETIME NOT NULL,
 FOREIGN KEY(user_id) REFERENCES user(id),
 FOREIGN KEY(post_id) REFERENCES post(id)
);

CREATE TABLE category_relation (
 id INTEGER NOT NULL PRIMARY KEY,
 category_id INTEGER NOT NULL,
 post_id INTEGER NOT NULL,
 FOREIGN KEY(category_id) REFERENCES category(id),
 FOREIGN KEY(post_id) REFERENCES post(id)
);

CREATE TABLE message (
 id INTEGER NOT NULL PRIMARY KEY,
 from_user INTEGER NOT NULL,
 to_user INTEGER NOT NULL,
 is_read TINYINT(1) NOT NULL,
 message TEXT NOT NULL,
 created_at DATETIME NOT NULL,
 FOREIGN KEY(from_user) REFERENCES user(id),
 FOREIGN KEY(to_user) REFERENCES user(id)
);

CREATE TABLE sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nickname TEXT,
    user_id INTEGER,
    cookie TEXT,
    expired_time TEXT,
    web_socket_conn TEXT
);




INSERT INTO category (id,category_name,descript,created_at)
VALUES
    (1,'Food & Beverage','',DateTime('now','localtime')),
    (2,'Technology & Gadgets','',DateTime('now','localtime')),
    (3,'Travel & Tourism','',DateTime('now','localtime')),
    (4,'Health & Wellness','',DateTime('now')),
    (5,'Fashion & Style','',DateTime('now')),
    (6,'Arts & Entertainment','',DateTime('now')),
    (7,'Sports & Fitness','',DateTime('now')),
    (8,'Education & Learning','',DateTime('now')),
    (9,'Home & Garden','',DateTime('now')),
    (10,'Business & Finance','',DateTime('now')),
    (11,'Cuisines','Recommendation regarding food in Mariehamn',DateTime('now','localtime')),
    (12,'Places','Places worth a visit in Mariehamn',DateTime('now','localtime')),
    (13,'Activities','Interesting events happening in Mariehamn',DateTime('now','localtime'));

