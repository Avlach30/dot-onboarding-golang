CREATE TABLE IF NOT EXISTS public.movies (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    title varchar(255) NOT NULL,
    genre varchar(255) NOT NULL,
    poster_url varchar(255) NOT NULL,
    duration_in_minutes int NOT NULL,
    desciption varchar(255) NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamptz NULL,
    CONSTRAINT movies_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS movies_title_idx ON public.movies USING btree (created_at);