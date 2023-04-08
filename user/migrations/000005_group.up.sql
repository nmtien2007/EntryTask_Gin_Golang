create table "group"
(
	id bigserial
		constraint group_pk
			primary key,
	name varchar(100) not null
);