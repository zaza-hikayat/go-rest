--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2
-- Dumped by pg_dump version 16.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: admin; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.admin (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    xid text,
    password text NOT NULL,
    email text,
    phone text,
    fullname text,
    role text
);


ALTER TABLE public.admin OWNER TO admin;

--
-- Name: admin_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.admin_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.admin_id_seq OWNER TO admin;

--
-- Name: admin_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.admin_id_seq OWNED BY public.admin.id;


--
-- Name: member; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.member (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    recruiter_id text,
    member_id text,
    full_name text,
    address text,
    password text,
    requiter_id text,
    email text,
    gender text,
    dob timestamp with time zone,
    phone text,
    bank text,
    rekening text
);


ALTER TABLE public.member OWNER TO admin;

--
-- Name: member_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.member_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.member_id_seq OWNER TO admin;

--
-- Name: member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.member_id_seq OWNED BY public.member.id;


--
-- Name: order; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public."order" (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    order_num text NOT NULL,
    created_by text NOT NULL,
    recruiter_id text,
    member_snapshot json,
    product_snapshot json,
    payment_method text
);


ALTER TABLE public."order" OWNER TO admin;

--
-- Name: order_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_id_seq OWNER TO admin;

--
-- Name: order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.order_id_seq OWNED BY public."order".id;


--
-- Name: product; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    code text,
    name text,
    qty bigint DEFAULT 0,
    description text,
    image text
);


ALTER TABLE public.product OWNER TO admin;

--
-- Name: product_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_id_seq OWNER TO admin;

--
-- Name: product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.product_id_seq OWNED BY public.product.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO admin;

--
-- Name: admin id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.admin ALTER COLUMN id SET DEFAULT nextval('public.admin_id_seq'::regclass);


--
-- Name: member id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.member ALTER COLUMN id SET DEFAULT nextval('public.member_id_seq'::regclass);


--
-- Name: order id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public."order" ALTER COLUMN id SET DEFAULT nextval('public.order_id_seq'::regclass);


--
-- Name: product id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product ALTER COLUMN id SET DEFAULT nextval('public.product_id_seq'::regclass);


--
-- Name: admin admin_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.admin
    ADD CONSTRAINT admin_pkey PRIMARY KEY (id);


--
-- Name: member member_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.member
    ADD CONSTRAINT member_pkey PRIMARY KEY (id);


--
-- Name: order order_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public."order"
    ADD CONSTRAINT order_pkey PRIMARY KEY (id);


--
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- Name: admin uni_admin_email; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.admin
    ADD CONSTRAINT uni_admin_email UNIQUE (email);


--
-- Name: admin uni_admin_xid; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.admin
    ADD CONSTRAINT uni_admin_xid UNIQUE (xid);


--
-- Name: member uni_member_member_id; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.member
    ADD CONSTRAINT uni_member_member_id UNIQUE (member_id);


--
-- Name: idx_admin_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_admin_deleted_at ON public.admin USING btree (deleted_at);


--
-- Name: idx_member_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_member_deleted_at ON public.member USING btree (deleted_at);


--
-- Name: idx_order_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_order_deleted_at ON public."order" USING btree (deleted_at);


--
-- Name: idx_product_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_product_deleted_at ON public.product USING btree (deleted_at);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

