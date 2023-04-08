create table perm
(
	id bigserial
		constraint perm_pk
			primary key,
	name varchar(100) not null,
	code varchar(50) not null,
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone not null
);

create unique index perm_code_uindex
	on permission (code);