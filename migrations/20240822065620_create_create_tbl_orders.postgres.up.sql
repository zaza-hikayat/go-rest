CREATE TABLE "product"
(
    id          bigserial NOT NULL,
    created_at  timestamptz NULL,
    updated_at  timestamptz NULL,
    deleted_at  timestamptz NULL,
    code        text NULL,
    name        text NULL,
    qty         int8 DEFAULT 0 NULL,
    description text NULL,
    image       text NULL,
    CONSTRAINT product_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_product_deleted_at ON public.product USING btree (deleted_at);