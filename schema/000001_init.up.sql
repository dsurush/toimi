create table adverts(
    id serial not null unique,
    title text not null,
    description text not null,
    photos text[] not null, 
    price integer not null,
	create_date timestamp not null
);