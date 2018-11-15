
-- +swan Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO system_settings (created_at,env, module, key, value, description)
VALUES (now(),'DEV', 'Books', 'maximum_rows_per_page', 10, 'A value of maximum books per page');

-- +swan Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM system_settings
WHERE module = 'Books' and key = 'maximum_rows_per_page';
