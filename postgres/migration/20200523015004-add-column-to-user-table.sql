
-- +migrate Up
CREATE TABLE test_table (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name text NOT NULL
);

INSERT INTO 
    test_table(name)
VALUES
    ('hello world');

-- +migrate Down
drop table test_table;