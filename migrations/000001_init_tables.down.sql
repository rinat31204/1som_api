-- Удаление триггера и функции
DROP TRIGGER IF EXISTS trigger_accounts_updated_at ON accounts;
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Удаление таблиц (порядок важен из-за зависимостей)
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS users;

-- Удаление ENUM-типов
DROP TYPE IF EXISTS transaction_type;
DROP TYPE IF EXISTS currency_type;