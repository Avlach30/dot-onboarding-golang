CREATE TABLE IF NOT EXISTS public.tickets (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    movie_schedule_id uuid NOT NULL,
    user_id uuid NOT NULL,
    selected_chairs jsonb NOT NULL,
    status varchar(255) NOT NULL DEFAULT 'confirmed',
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp NULL,
    CONSTRAINT idx_movie_schedule_id FOREIGN KEY (movie_schedule_id) REFERENCES public.movie_schedules (id) ON DELETE CASCADE,
    CONSTRAINT idx_user_id FOREIGN KEY (user_id) REFERENCES public.users (id) ON DELETE CASCADE
);
CREATE INDEX idx_tickets_movie_schedule_id ON public.tickets USING btree (movie_schedule_id);
CREATE INDEX idx_tickets_user_id ON public.tickets USING btree (user_id);