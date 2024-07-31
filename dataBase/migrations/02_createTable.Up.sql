CREATE TABLE IF NOT EXISTS link(
    linkID SERIAL, longLink varchar(1000), shortLink varchar(255)
);

CREATE TABLE IF NOT EXISTS client(
    clientID SERIAL, email varchar(255), name varchar(255), password varchar(255)
);

CREATE TABLE IF NOT EXISTS clientlink (linkID int, clientID int);
