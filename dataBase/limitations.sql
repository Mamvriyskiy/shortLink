ALTER TABLE client
ALTER COLUMN password
SET NOT NULL,
ALTER COLUMN name
SET NOT NULL,
ALTER COLUMN email
SET NOT NULL,
ADD check (password != ''),
ADD check (name != ''),
ADD check (email != ''),
ADD primary key (clientID);

ALTER TABLE link
ALTER COLUMN longlink
SET NOT NULL,
ALTER COLUMN shortlink
SET NOT NULL,
ADD check (longlink != ''),
ADD check (shortlink != ''),
ADD primary key (linkID);

ALTER TABLE clientlink
ALTER COLUMN clientID
SET NOT NULL,
ALTER COLUMN linkID
SET NOT NULL,
ADD FOREIGN KEY (clientID) REFERENCES client (clientID) ON DELETE CASCADE,
ADD FOREIGN KEY (linkID) REFERENCES link (linkID) ON DELETE CASCADE;
