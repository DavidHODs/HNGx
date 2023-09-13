-- Create a table named person
CREATE TABLE IF NOT EXISTS person (
    id serial PRIMARY KEY,
    fullName VARCHAR NOT NULL,
    isDeleted  BOOLEAN,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP,
    deletedAt TIMESTAMP
);