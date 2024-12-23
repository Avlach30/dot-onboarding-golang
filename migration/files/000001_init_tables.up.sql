CREATE TABLE IF NOT EXISTS public.permissions (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"name" varchar(255) NOT NULL,
	"key" varchar(255) NOT NULL,
	deleted_at timestamptz NULL,
	updated_at timestamptz NULL,
	created_at timestamptz NULL,
	CONSTRAINT permissions_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_permissions_deleted_at ON public.permissions USING btree (deleted_at);


CREATE TABLE IF NOT EXISTS public.roles (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"name" varchar(255) NOT NULL,
	"key" varchar(255) NOT NULL,
	deleted_at timestamptz NULL,
	updated_at timestamptz NULL,
	created_at timestamptz NULL,
	CONSTRAINT roles_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_roles_deleted_at ON public.roles USING btree (deleted_at);


CREATE TABLE IF NOT EXISTS public.users (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"name" varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	deleted_at timestamptz NULL,
	updated_at timestamptz NULL,
	created_at timestamptz NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON public.users USING btree (deleted_at);

CREATE TABLE IF NOT EXISTS public.role_permissions (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	role_id uuid NULL,
	permission_id uuid NULL,
	deleted_at timestamptz NULL,
	some_data timestamptz NULL,
	created_at timestamptz NULL,
	CONSTRAINT role_permissions_pkey PRIMARY KEY (id),
	CONSTRAINT fk_role_permissions_permisison FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE,
	CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS idx_composite_role_permission ON public.role_permissions USING btree (role_id, permission_id);
CREATE INDEX IF NOT EXISTS idx_role_permissions_created_at ON public.role_permissions USING btree (created_at);
CREATE INDEX IF NOT EXISTS idx_role_permissions_deleted_at ON public.role_permissions USING btree (deleted_at);
CREATE INDEX IF NOT EXISTS idx_role_permissions_updated_at ON public.role_permissions USING btree (some_data);