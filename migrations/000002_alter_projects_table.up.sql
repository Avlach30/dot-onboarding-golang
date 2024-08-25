-- Check if the column 'owner' exists in the 'projects' table
SET @column_exists = (
    SELECT COUNT(*)
    FROM information_schema.columns
    WHERE table_name = 'projects'
    AND column_name = 'owner'
    AND table_schema = DATABASE()
    );

-- Conditionally alter the table to rename 'owner' to 'user_id' and change its type
SET @alter_column_sql = IF(
    @column_exists > 0,
    'ALTER TABLE projects CHANGE COLUMN owner user_id INT(11);',
    'SELECT "Column does not exist";'
    );

-- Execute the SQL for altering the column
PREPARE stmt FROM @alter_column_sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;


-- Check if the foreign key constraint for 'user_id' already exists
SET @fk_exists = (
    SELECT COUNT(*)
    FROM information_schema.key_column_usage
    WHERE table_name = 'projects'
    AND column_name = 'user_id'
    AND constraint_schema = DATABASE()
    AND referenced_table_name IS NOT NULL
    );

-- Conditionally add the foreign key constraint
SET @add_fk_sql = IF(
    @fk_exists = 0,
    'ALTER TABLE projects ADD CONSTRAINT fk_projects_user_id FOREIGN KEY (user_id) REFERENCES users(id);',
    'SELECT "Foreign key constraint already exists";'
    );

-- Execute the SQL for adding the foreign key constraint
PREPARE stmt2 FROM @add_fk_sql;
EXECUTE stmt2;
DEALLOCATE PREPARE stmt2;