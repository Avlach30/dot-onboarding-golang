-- Set role and permission key to unique
ALTER TABLE public.role_entities
  ADD CONSTRAINT roles_key_unique UNIQUE (key);

ALTER TABLE public.permission_entities
    ADD CONSTRAINT permissions_key_unique UNIQUE (key);