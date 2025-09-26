CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON auth.users
    FOR EACH ROW
    EXECUTE PROCEDURE public.trigger_set_timestamp();

