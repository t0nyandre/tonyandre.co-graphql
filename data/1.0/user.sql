CREATE TABLE public.users
(
    id VARCHAR(30) NOT NULL,
    first_name text NOT NULL,
    last_name text,
    email text UNIQUE NOT NULL,
    password text NOT NULL,
    confirmed boolean DEFAULT false,
    created_at text NOT NULL,
    updated_at text,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)