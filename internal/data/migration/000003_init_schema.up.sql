CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "countries" (
   "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   "name" varchar(255)
);

CREATE TABLE "cities" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" varchar(255),
    "country_id" int
);

CREATE TABLE "districts" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" TEXT UNIQUE NOT NULL,
    "city_id" int
);

CREATE TABLE "streets" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" TEXT NOT NULL,
    "district_id" INT NOT NULL
);