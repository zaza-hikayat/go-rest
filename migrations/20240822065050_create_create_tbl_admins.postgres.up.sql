CREATE TABLE "admin"
(
    id         bigserial NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL,
    xid        text NULL,
    "password" text      NOT NULL,
    email      text NULL,
    phone      text NULL,
    fullname   text NULL,
    "role"     text NULL,
    CONSTRAINT admin_pkey PRIMARY KEY (id),
    CONSTRAINT uni_admin_email UNIQUE (email),
    CONSTRAINT uni_admin_xid UNIQUE (xid)
);

CREATE INDEX idx_admin_deleted_at ON public.admin USING btree (deleted_at);