CREATE TABLE public.users
(
    id character varying(25) NOT NULL,
    username character varying(30) NOT NULL,
    email character varying(200) NOT NULL,
    password text NOT NULL,
    confirmed boolean DEFAULT false,
    disabled boolean DEFAULT false,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_username_key UNIQUE (username)
)