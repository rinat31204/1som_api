-- 1. Создание ENUM-типов
CREATE TYPE currency_type AS ENUM ('KGS', 'KZT', 'USD');
CREATE TYPE transaction_type AS ENUM ('credit', 'debit','refund');

-- 2. Таблица пользователей (ID — внешний ключ от Firebase/Auth)
CREATE TABLE users (
                       id         TEXT PRIMARY KEY,
                       email      TEXT NOT NULL,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 3. Таблица кошельков (у пользователя может быть несколько кошельков)
CREATE TABLE wallets (
                         id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         user_id    TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                         created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 4. Таблица счетов (у каждого кошелька несколько счетов в разных валютах)
CREATE TABLE accounts (
                          id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          wallet_id      UUID NOT NULL REFERENCES wallets(id) ON DELETE CASCADE,
                          account_number TEXT NOT NULL UNIQUE,
                          balance        BIGINT NOT NULL DEFAULT 0 CHECK (balance >= 0),
                          currency       currency_type NOT NULL,
                          created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                          updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 5. Таблица транзакций
CREATE TABLE transactions (
                              id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              account_number TEXT NOT NULL REFERENCES accounts(account_number) ON DELETE RESTRICT,
                              amount         BIGINT NOT NULL CHECK (amount >= 0),
                              currency       currency_type NOT NULL,
                              description    TEXT,
                              type           transaction_type NOT NULL,
                              created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 6. Индексы для ускорения запросов
CREATE INDEX idx_wallets_user_id ON wallets(user_id);
CREATE INDEX idx_accounts_wallet_id ON accounts(wallet_id);
CREATE INDEX idx_accounts_account_number ON accounts(account_number);
CREATE INDEX idx_transactions_account_number ON transactions(account_number);
CREATE INDEX idx_transactions_created_at ON transactions(created_at);

-- 7. Триггер для автоматического обновления updated_at в accounts
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_accounts_updated_at
    BEFORE UPDATE ON accounts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();