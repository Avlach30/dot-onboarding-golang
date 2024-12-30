-- Write your UP migration SQL here

CREATE INDEX IF NOT EXISTS idx_composite_user_role ON public.user_roles (user_id, role_id);
