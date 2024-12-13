CREATE TABLE public.permission_entities (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"name" varchar(255) NOT NULL,
	"key" varchar(255) NOT NULL,
	deleted_at timestamptz NULL,
	updated_at timestamptz NULL,
	created_at timestamptz NULL,
	CONSTRAINT permission_entities_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_permission_entities_deleted_at ON public.permission_entities USING btree (deleted_at);


CREATE TABLE public.role_entities (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"name" varchar(255) NOT NULL,
	"key" varchar(255) NOT NULL,
	deleted_at timestamptz NULL,
	updated_at timestamptz NULL,
	created_at timestamptz NULL,
	CONSTRAINT role_entities_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_role_entities_deleted_at ON public.role_entities USING btree (deleted_at);


CREATE TABLE public.user_entities (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"name" varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	deleted_at timestamptz NULL,
	updated_at timestamptz NULL,
	created_at timestamptz NULL,
	CONSTRAINT user_entities_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_user_entities_deleted_at ON public.user_entities USING btree (deleted_at);

CREATE TABLE public.role_permission_entities (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	role_id uuid NULL,
	permission_id uuid NULL,
	deleted_at timestamptz NULL,
	some_data timestamptz NULL,
	created_at timestamptz NULL,
	CONSTRAINT role_permission_entities_pkey PRIMARY KEY (id),
	CONSTRAINT fk_role_permission_entities_permisison FOREIGN KEY (permission_id) REFERENCES public.permission_entities(id) ON DELETE CASCADE,
	CONSTRAINT fk_role_permission_entities_role FOREIGN KEY (role_id) REFERENCES public.role_entities(id) ON DELETE CASCADE
);
CREATE INDEX idx_composite_role_permission ON public.role_permission_entities USING btree (role_id, permission_id);
CREATE INDEX idx_role_permission_entities_created_at ON public.role_permission_entities USING btree (created_at);
CREATE INDEX idx_role_permission_entities_deleted_at ON public.role_permission_entities USING btree (deleted_at);
CREATE INDEX idx_role_permission_entities_updated_at ON public.role_permission_entities USING btree (some_data);