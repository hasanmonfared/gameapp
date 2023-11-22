package mysql

import (
	"database/sql"
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/richerror"
)

func (d MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number= ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("can't scan query result :%w", err)
	}
	return false, nil
}
func (d MySQLDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name,phone_number,password) values(?,?,?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command:%w", err)
	}
	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil
}
func (d MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number= ?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}
		return entity.User{}, false, richerror.New(err, "mysql.GetUserByPhoneNumber", "can't scan query result", richerror.KindUnexpected, nil)
	}
	return user, true, nil
}
func (d MySQLDB) GetUserByID(userID uint) (entity.User, error) {

	row := d.db.QueryRow(`select * from users where id= ?`, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(err, "mysql.GetUserByID", "record not found.", richerror.KindNotFound, nil)
		}
		return entity.User{}, richerror.New(err, "mysql.GetUserByID", "can't scan query result", richerror.KindUnexpected, nil)
	}
	return user, nil
}
func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt []uint8
	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password)
	return user, err
}
