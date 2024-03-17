CREATE TABLE actor(
	id serial primary key,
	name varchar(255),
	male boolean,
	birth_date date
);
CREATE TABLE film(
	id serial primary key,
	name varchar(150) UNIQUE,
	description varchar(1000),
	release_date date,
	rating int
);
CREATE TABLE actor_film(
	actor_id int,
	film_id int,
	PRIMARY KEY(actor_id,film_id),
	FOREIGN KEY (actor_id) REFERENCES actor (id) ON DELETE CASCADE,
	FOREIGN KEY (film_id) REFERENCES film (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    role_id INTEGER,
    email VARCHAR(255) NULL,
    password VARCHAR(255),
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE 
);

INSERT INTO roles(name) VALUES('admin'),('user');

INSERT INTO users(role_id, email, password) VALUES(1,'admin@m.ru','passwordHash'), (2,'user@m.ru','passwordHash');

INSERT INTO actor(name, male, birth_date) VALUES ('Tom Hanks', true, '1956-07-09');
INSERT INTO actor(name, male, birth_date) VALUES ('Meryl Streep', false, '1949-06-22');
INSERT INTO actor(name, male, birth_date) VALUES ('Leonardo DiCaprio', true, '1974-11-11');
INSERT INTO actor(name, male, birth_date) VALUES ('Reese Witherspoon', false, '1976-03-22');
INSERT INTO actor(name, male, birth_date) VALUES ('Brad Pitt', true, '1963-12-18');
INSERT INTO actor(name, male, birth_date) VALUES ('Jennifer Lawrence', false, '1990-08-15');

INSERT INTO film(name, description, release_date, rating) VALUES ('Forrest Gump', 'A man with low IQ accomplishes great things in life.', '1994-07-06', 9);
INSERT INTO film(name, description, release_date, rating) VALUES ('The Devil Wears Prada', 'A young woman becomes an assistant to a powerful fashion magazine editor.', '2006-06-30', 8);
INSERT INTO film(name, description, release_date, rating) VALUES ('The Wolf of Wall Street', 'A stockbroker living the high life gets involved in crime and corruption.', '2013-12-25', 8);
INSERT INTO film(name, description, release_date, rating) VALUES ('Legally Blonde', 'A blonde sorority queen enrolls in Harvard Law School to win back her ex-boyfriend.', '2001-07-13', 7);
INSERT INTO film(name, description, release_date, rating) VALUES ('Fight Club', 'An insomniac office worker and a soap maker form an underground fight club.', '1999-10-15', 9);
INSERT INTO film(name, description, release_date, rating) VALUES ('Silver Linings Playbook', 'After a stint in a mental institution, a man tries to win back his estranged wife.', '2012-09-08', 8);

INSERT INTO actor_film(actor_id, film_id) VALUES (1, 1);
INSERT INTO actor_film(actor_id, film_id) VALUES (2, 2);
INSERT INTO actor_film(actor_id, film_id) VALUES (3, 3);
INSERT INTO actor_film(actor_id, film_id) VALUES (4, 4);
INSERT INTO actor_film(actor_id, film_id) VALUES (5, 5);
INSERT INTO actor_film(actor_id, film_id) VALUES (6, 6);