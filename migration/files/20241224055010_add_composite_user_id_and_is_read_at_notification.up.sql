-- Write your UP migration SQL here

CREATE INDEX IF NOT EXISTS idx_notifications_user_id_and_is_read ON public.notifications (user_id, is_read);
