-- Create table user_roles for many to many relation between users and roles

CREATE TABLE public.user_roles (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    user_id uuid NULL,
    role_id uuid NULL,
    deleted_at timestamptz NULL,
    created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT user_roles_pkey PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (role_id) REFERENCES roles (id)
);