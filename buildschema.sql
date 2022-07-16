CREATE TABLE flags (
  id serial PRIMARY KEY,
  key varchar(40)--,
  -- created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  -- updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  -- deleted_at TIMESTAMPTZ
);

CREATE TABLE audiences (
  id serial PRIMARY KEY,
  key varchar(40),
  combine boolean NOT NULL DEFAULT false
);

CREATE TABLE attributes (
  id serial PRIMARY KEY,
  key varchar(40)
);

CREATE TABLE flagaudiences (
  id serial PRIMARY KEY,
  flag_id integer references flags(id) NOT NULL, --ON DELETE CASCADE
  audience_id integer references audiences(id) NOT NULL --ON DELETE CASCADE
);

CREATE TABLE conditions (
  id serial PRIMARY KEY,
  audience_id integer references audiences(id) NOT NULL, --ON DELETE CASCADE
  attr_id integer references attributes(id) NOT NULL, --ON DELETE CASCADE
  operator varchar(20) DEFAULT 'EQ',
  vals jsonb
);

CREATE TABLE projects (
  id serial PRIMARY KEY,
  key varchar(40)
);

CREATE TABLE flagaudiences (
  id serial PRIMARY KEY,
  flag_id integer references flags(id) NOT NULL, --ON DELETE CASCADE
  audience_id integer references audiences(id) NOT NULL --ON DELETE CASCADE
);
