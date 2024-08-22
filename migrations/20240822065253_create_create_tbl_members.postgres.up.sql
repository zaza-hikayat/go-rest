CREATE TABLE public."member"
(
    id           bigserial NOT NULL,
    created_at   timestamptz NULL,
    updated_at   timestamptz NULL,
    deleted_at   timestamptz NULL,
    recruiter_id text NULL,
    member_id    text NULL,
    full_name    text NULL,
    address      text NULL,
    "password"   text NULL,
    requiter_id  text NULL,
    email        text NULL,
    gender       text NULL,
    dob          timestamptz NULL,
    phone        text NULL,
    bank         text NULL,
    rekening     text NULL,
    CONSTRAINT member_pkey PRIMARY KEY (id),
    CONSTRAINT uni_member_member_id UNIQUE (member_id)
);
CREATE INDEX idx_member_deleted_at ON public.member USING btree (deleted_at);