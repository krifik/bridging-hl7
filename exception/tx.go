package exception

import "gorm.io/gorm"

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfNeeded(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfNeeded(errorCommit)
	}
}
