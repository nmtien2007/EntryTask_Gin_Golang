create table group_perm
(
	id bigserial
		constraint group_perm_pk
			primary key,
	group_id int not null
		constraint group_perm_group_id_fk
			references "group",
	perm_id int not null
		constraint group_perm_perm_id_fk
			references "perm",
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone not null
);

create unique index group_perm_group_id_perm_id_uindex
	on group_perm (group_id, perm_id);