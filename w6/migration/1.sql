CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE recipe (
	recipe_id uuid NOT NULL DEFAULT uuid_generate_v4(),
	name varchar(60) NOT NULL,
	url varchar(255) NOT NULL,
	rating float8 NOT NULL,
	CONSTRAINT recipe_pkey PRIMARY KEY (recipe_id)
);

CREATE TABLE menu (
	menu_id uuid NOT NULL DEFAULT uuid_generate_v4(),
	recipe_id uuid NOT NULL
);
