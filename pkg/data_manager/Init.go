package data_manager

import "entry_task/database"

func Init(db *database.MyDB) error {
	err := startAttrManager(db)
	if err != nil {
		return err
	}

	err = startBrandManager(db)
	if err != nil {
		return err
	}

	err = StartCategoryManager(db)
	if err != nil {
		return err
	}
	return nil
}
