CREATE TABLE public.log_integrations (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"url" text NOT NULL,
	request text NOT NULL,
	response text NOT NULL,
	"status" varchar(255) NOT NULL,
	"scheme" varchar(255) NOT NULL,
	deleted_at timestamptz NULL,
	updated_at timestamptz NULL,
	created_at timestamptz NULL,
	CONSTRAINT log_integrations_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_log_integrations_deleted_at ON public.log_integrations USING btree (deleted_at);
