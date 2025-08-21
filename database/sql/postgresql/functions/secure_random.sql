CREATE OR REPLACE FUNCTION secure_random()
    RETURNS double precision
    AS $$
BEGIN
    RETURN('x' || ENCODE(gen_random_bytes(4), 'hex'))::bit(32)::bigint::float8 / 4294967295.0;
END;
$$
LANGUAGE plpgsql
VOLATILE PARALLEL SAFE;

