CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "countries" (
   "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   "name" varchar(255) NOT NULL UNIQUE
);

CREATE TABLE "cities" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" varchar(255) NOT NULL,
    "country_id" uuid NOT NULL REFERENCES countries(id) ON DELETE CASCADE
);

CREATE TABLE "districts" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" varchar(255) NOT NULL ,
    "city_id" uuid NOT NULL REFERENCES cities(id) ON DELETE CASCADE
);

CREATE TABLE "streets" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" varchar(255) NOT NULL,
    "district_id" uuid NOT NULL REFERENCES districts(id) ON DELETE CASCADE
);