/*
Copyright (c) 2021, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package sql

import (
	"context"
	"database/sql"
	"qz/cas"
	"qz/util"
	"reflect"
	"time"
)

// DbCasHandler implements qz/commands.HelperFactory
type DbCasHandler struct {
}

// Name implements qz/commands.HelperFactory
func (ch *DbCasHandler) Name() string {
	return reflect.TypeOf(ch).Elem().String()
}

// Help implements qz/commands.HelperFactory
func (ch *DbCasHandler) Help() string {
	return `
	"component": "qz/helpers.DbCasHandler",
	"param": {
		"path": "HandlerHelperFactCtxName"
	}
	`
}

// ComponentType implements qz/commands.HelperFactory
func (ch *DbCasHandler) ComponentType() reflect.Type {
	return reflect.TypeOf(ch)
}

// DbHelper creates a db operator, implements the CAS Handler interfaces
type dbHelper struct {
	db *sql.DB
}

// DbConfig DB Config
type DbConfig struct {
	DriverName      string `json:"driverName"`
	DbURL           string `json:"dbURL"`
	ConnMaxLifetime string `json:"connMaxLifetime"` // 0
	MaxIdleConns    int    `json:"maxIdleConns"`    // 50
	MaxOpenConns    int    `json:"maxOpenConns"`    // 50
}

func newDbHelper(ctx context.Context, cfg *DbConfig) (dbh *dbHelper, err error) {
	d, err := time.ParseDuration(cfg.ConnMaxLifetime)

	dbh.db, err = sql.Open(cfg.DriverName, cfg.DbURL)
	if err != nil {
		return
	}

	dbh.db.SetConnMaxLifetime(d)
	dbh.db.SetMaxIdleConns(cfg.MaxIdleConns)
	dbh.db.SetMaxOpenConns(cfg.MaxOpenConns)

	go func() {
		for {
			select {
			case <-ctx.Done():
				if dbh.db != nil {
					dbh.db.Close()
					dbh.db = nil
				}
				util.DebugInfo(ctx, "db closed with ctx Done")
				return
			}
		}
	}()
	return
}

func (dh *dbHelper) LoadJSON(namespace, key string) (rec *cas.CasJSON, err error) {

	return
}

// SaveJSON saves record, updates hash and timestamp
func (dh *dbHelper) SaveJSON(rec *cas.CasJSON) (err error) {

	return
}

/*
DbSchema DB Schema  backup for reference
*/
const DbSchema = `
--
-- PostgreSQL database dump
--

-- Dumped from database version 12.5 (Ubuntu 12.5-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.5 (Ubuntu 12.5-0ubuntu0.20.04.1)

-- Started on 2021-02-03 21:06:34 IST
-- AB: removed config table

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

DROP DATABASE traceit;
--
-- TOC entry 2989 (class 1262 OID 24220)
-- Name: traceit; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE traceit WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_IN' LC_CTYPE = 'en_IN';


ALTER DATABASE traceit OWNER TO postgres;

-- \connect traceit

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
-- TOC entry 203 (class 1259 OID 24228)
-- Name: nskey; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.nskey (
    key character varying(256) NOT NULL,
    meta json,
    tmstamp timestamp with time zone,
    namespace character varying(256) DEFAULT 'default'::character varying NOT NULL,
    hash bytea NOT NULL
);


ALTER TABLE public.nskey OWNER TO postgres;

--
-- TOC entry 2990 (class 0 OID 0)
-- Dependencies: 203
-- Name: TABLE nskey; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.nskey IS 'Namespace key parent to json_data, bin_data';


--
-- TOC entry 205 (class 1259 OID 24247)
-- Name: bin_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bin_data (
    val bytea,
    mimetype character varying(256) DEFAULT 'text'::character varying
)
INHERITS (public.nskey);


ALTER TABLE public.bin_data OWNER TO postgres;



--
-- TOC entry 204 (class 1259 OID 24241)
-- Name: json_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.json_data (
    val json
)
INHERITS (public.nskey);


ALTER TABLE public.json_data OWNER TO postgres;

--
-- TOC entry 2853 (class 2604 OID 24256)
-- Name: bin_data namespace; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bin_data ALTER COLUMN namespace SET DEFAULT 'default'::character varying;


--
-- TOC entry 2851 (class 2604 OID 24255)
-- Name: json_data namespace; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.json_data ALTER COLUMN namespace SET DEFAULT 'default'::character varying;


--
-- TOC entry 2857 (class 2606 OID 24258)
-- Name: nskey nskey_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.nskey
    ADD CONSTRAINT nskey_pkey PRIMARY KEY (hash);


--
-- TOC entry 2855 (class 1259 OID 24259)
-- Name: key_namespace; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX key_namespace ON public.nskey USING btree (key, namespace);


--
-- TOC entry 2854 (class 1259 OID 24227)
-- Name: url; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX url ON public.config USING btree (url);


-- Completed on 2021-02-03 21:06:35 IST

--
-- PostgreSQL database dump complete
--



`
