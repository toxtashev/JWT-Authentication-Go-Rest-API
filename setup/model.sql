create database login;

create table users(
	user_id serial not null primary key,
	username varchar not null,
	gmail varchar not null,
	password varchar not null,
	token varchar
);
