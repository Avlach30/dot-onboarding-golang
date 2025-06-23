CREATE TABLE IF NOT EXISTS public.movie_studios (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	name varchar(255) NOT NULL,
	chair_capacity int NOT NULL,
	additional_capacities jsonb,
	deleted_at timestamptz NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT movie_studios_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_movie_studios_deleted_at ON public.movie_studios USING btree (deleted_at);