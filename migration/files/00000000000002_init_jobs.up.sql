
CREATE TABLE IF NOT EXISTS public.jobs (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"task_name" varchar(255) NOT NULL,
	"payload" text NOT NULL,
	"booked" boolean NOT NULL DEFAULT false,
	deleted_at timestamptz NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT jobs_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_jobs_deleted_at ON public.jobs USING btree (deleted_at);

CREATE TABLE IF NOT EXISTS public.job_faileds (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	job_id varchar(255) NOT NULL,
	error_message varchar(255) NOT NULL,
	deleted_at timestamptz NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT job_faileds_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_job_faileds_deleted_at ON public.job_faileds USING btree (deleted_at);
