package example

import "fmt"

// generate test data
func GenerateTestData() error {
	tableSql := `
CREATE TABLE pagination_users (
  id   INT          NOT NULL AUTO_INCREMENT
  COMMENT 'user table id',
  name VARCHAR(255) NOT NULL DEFAULT ''
  COMMENT 'user name',
  age  TINYINT(3)   NOT NULL DEFAULT '0'
  COMMENT 'user age',
  PRIMARY KEY (id)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT = 'user table';
`
	dataSql := `
INSERT INTO pagination_users (id, name, age)
VALUES
  (1, 'user_name_1', 1),
  (2, 'user_name_2', 2),
  (3, 'user_name_3', 3),
  (4, 'user_name_4', 4),
  (5, 'user_name_5', 5),
  (6, 'user_name_6', 6),
  (7, 'user_name_7', 7),
  (8, 'user_name_8', 8),
  (9, 'user_name_9', 9);
`
	db, err := newDbConnection()
	if err != nil {
		err = fmt.Errorf("newDbConnection error : %v", err)
		return err
	}
	defer db.Close()

	// create table
	if err := db.Exec(tableSql).Error; err != nil {
		err = fmt.Errorf("db.Exec(tableSql) error : %v", err)
		return err
	}

	// insert test data
	if err := db.Exec(dataSql).Error; err != nil {
		err = fmt.Errorf("db.Exec(dataSql) error : %v", err)
		return err
	}

	return nil
}
