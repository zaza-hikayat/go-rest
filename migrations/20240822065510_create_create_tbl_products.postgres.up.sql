CREATE TABLE public."order"
(
    id               bigserial NOT NULL,
    created_at       timestamptz NULL,
    updated_at       timestamptz NULL,
    deleted_at       timestamptz NULL,
    order_num        text      NOT NULL,
    created_by       text      NOT NULL,
    recruiter_id     text NULL,
    member_snapshot  json NULL,
    product_snapshot json NULL,
    payment_method   text NULL,
    CONSTRAINT order_pkey PRIMARY KEY (id)
);

CREATE INDEX idx_order_deleted_at ON public."order" USING btree (deleted_at);