CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS postgis;

-- Таблица городов
CREATE TABLE cities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL UNIQUE,
    location GEOGRAPHY(POLYGON, 4326) NOT NULL -- Границы города
);

-- Таблица улиц
CREATE TABLE streets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    city_id UUID NOT NULL REFERENCES cities(id) ON DELETE CASCADE,
    location GEOGRAPHY(LINESTRING, 4326) NOT NULL -- Линия улицы
);
