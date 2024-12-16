
CREATE TABLE public.job_entities (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"task_name" varchar(255) NOT NULL,
	"payload" text NOT NULL,
	"booked" boolean NOT NULL DEFAULT false,
	deleted_at timestamptz NULL,
	updated_at timestamptz NULL,
	created_at timestamptz NULL,
	CONSTRAINT job_entities_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_job_entities_deleted_at ON public.job_entities USING btree (deleted_at);

CREATE TABLE public.job_failed_entities (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	job_id varchar(255) NOT NULL,
	error_message varchar(255) NOT NULL,
	deleted_at timestamptz NULL,
	updated_at timestamptz NULL,
	created_at timestamptz NULL,
	CONSTRAINT job_failed_entities_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_job_failed_entities_deleted_at ON public.job_failed_entities USING btree (deleted_at);
