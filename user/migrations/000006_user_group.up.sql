create table user_group
(
	id serial
		constraint user_group_pk
			primary key,
	user_id int not null
		constraint user_group_user_id_fk
			references "user",
	group_id int not null
		constraint user_group_group_id_fk
			references "group",
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone not null
);

create unique index user_group_user_id_group_id_uindex
	on user_group (user_id, group_id);