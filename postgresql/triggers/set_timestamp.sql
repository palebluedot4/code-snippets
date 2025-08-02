CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON example_table
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

