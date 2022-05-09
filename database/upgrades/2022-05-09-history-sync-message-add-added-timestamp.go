package upgrades

import (
	"database/sql"
)

func init() {
	upgrades[44] = upgrade{"Add dispatch time to backfill queue", func(tx *sql.Tx, ctx context) error {
		// Add the inserted time TIMESTAMP column to history_sync_message
		_, err := tx.Exec(`
			ALTER TABLE history_sync_message
			ADD COLUMN inserted_time TIMESTAMP
		`)
		return err
	}}
}
