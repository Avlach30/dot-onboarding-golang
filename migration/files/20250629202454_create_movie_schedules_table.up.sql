CREATE TABLE IF NOT EXISTS public.movie_schedules
(
    id uuid NOT NULL,
    movie_id uuid NOT NULL,
    movie_studio_id uuid NOT NULL,
    show_datetime timestamp with time zone NOT NULL,
    price numeric(10, 2) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NULL,
    CONSTRAINT movie_schedules_pkey PRIMARY KEY (id),
    CONSTRAINT movie_schedules_movie_id_fkey FOREIGN KEY (movie_id)
        REFERENCES public.movies (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT movie_schedules_movie_studio_id_fkey FOREIGN KEY (movie_studio_id)
        REFERENCES public.movie_studios (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
CREATE INDEX movie_schedules_movie_id_idx ON public.movie_schedules USING btree (movie_id);
CREATE INDEX movie_schedules_movie_studio_id_idx ON public.movie_schedules USING btree (movie_studio_id);
