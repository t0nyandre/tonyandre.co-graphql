CREATE TABLE public.posts
(
    id character varying(25) NOT NULL,
    title character varying(125) NOT NULL,
    excerpt text,
    text text NOT NULL,
    image text,
    slug character varying(125) NOT NULL,
    published boolean DEFAULT false,
    archived boolean DEFAULT false,
    created_at timestamp with time zone,
    published_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT posts_pkey PRIMARY KEY (id),
    CONSTRAINT posts_slug_key UNIQUE (slug)
)