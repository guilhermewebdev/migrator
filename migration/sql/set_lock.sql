DECLARE @IsLocked BOOLEAN;
SET @IsLocked = ?;

MERGE INTO migrations_lock AS target
USING (SELECT 1 AS id) AS source
ON (target.id = source.id)
WHEN MATCHED THEN
    UPDATE SET target.is_locked = @IsLocked
WHEN NOT MATCHED THEN
    INSERT (id, is_locked)
    VALUES (1, @IsLocked);