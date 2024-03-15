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

INSERT INTO actor(name, male, birth_date) VALUES('Arnold Alois Schwarzenegger',true, '1947-07-30');
INSERT INTO film(name, description,release_date, rating) 
VALUES('Terminator','Terminator description',  '1999-01-01', 9);
INSERT INTO actor_film VALUES(1,1);