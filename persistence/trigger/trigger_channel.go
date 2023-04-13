package trigger

const TriggerChannelInsertOrUpdate = "create_or_update_channel"

const SQLCreateInsertOrUpdateTrigger = `
CREATE OR REPLACE FUNCTION create_or_update_trigger_func()
RETURNS TRIGGER AS $$
BEGIN
	IF (TG_OP = 'INSERT') THEN
		PERFORM pg_notify('create_or_update_channel', TG_TABLE_NAME ||';'|| NEW.id::text ||';'|| 'inserted');
	ELSIF (TG_OP = 'UPDATE') THEN
		PERFORM pg_notify('create_or_update_channel', TG_TABLE_NAME ||';'|| NEW.id::text ||';'|| 'updated');
	ELSIF (TG_OP = 'DELETE') THEN
		PERFORM pg_notify('create_or_update_channel', TG_TABLE_NAME ||';'|| NEW.id::text ||';'|| 'deleted');
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;`

const (
	sqlCreateInsertUpdateTrigger = `
	CREATE TRIGGER create_or_update_%s_trigger
	AFTER INSERT OR UPDATE ON %s
	FOR EACH ROW EXECUTE FUNCTION create_or_update_trigger_func();
	`
)
