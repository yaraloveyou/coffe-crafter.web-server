----migrate -path migrations -database "postgres://localhost/coffe_crafter?sslmode=disable&user=postgres" down
DROP TRIGGER IF EXISTS set_updated_at ON users;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP TABLE IF EXISTS users;