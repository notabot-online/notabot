-- Drop indexes
DROP INDEX IF EXISTS idx_transactions_tx_hash;
DROP INDEX IF EXISTS idx_transactions_verification_id;
DROP INDEX IF EXISTS idx_verifications_status;
DROP INDEX IF EXISTS idx_verifications_user_id;

-- Drop tables
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS verifications;


