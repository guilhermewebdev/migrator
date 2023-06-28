DECLARE @IsLocked BOOLEAN = ?;
INSERT INTO migrations_lock (id, is_locked)   
VALUES (1, @IsLocked)  
ON DUPLICATE KEY UPDATE id = 1; 