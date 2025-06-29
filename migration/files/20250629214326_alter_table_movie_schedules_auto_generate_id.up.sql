ALTER TABLE public.movie_schedules
    ALTER COLUMN id
    SET DEFAULT uuid_generate_v4();