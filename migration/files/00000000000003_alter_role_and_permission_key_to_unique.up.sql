-- Set role and permission key to unique
ALTER TABLE public.roles
  ADD CONSTRAINT roles_key_unique UNIQUE (key);

ALTER TABLE public.permissions
    ADD CONSTRAINT permissions_key_unique UNIQUE (key);