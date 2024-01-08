package mysqlaccesscontrol

import (
	"gameapp/entity"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	"gameapp/pkg/slice"
	"gameapp/repository/mysql"
	"strings"
	"time"
)

func (d *DB) GetUserPermissionsTitles(UserID uint, role entity.Role) ([]entity.PermissionTitle, error) {
	const op = "mysql.GetUserPermissionsTitles"

	roleACL := make([]entity.AccessControl, 0)

	rows, err := d.conn.Conn().Query(`select * from access_controles where actore_type= ? and actore_id=?`, entity.RoleActorType, role)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan()
		acl, err := scanAccessControl(rows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}
		roleACL = append(roleACL, acl)
	}
	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)

	}

	userACL := make([]entity.AccessControl, 0)
	userRows, err := d.conn.Conn().Query(`select * from access_controles where actore_type= ? and actore_id=?`, entity.UserActorType, UserID)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	defer userRows.Close()

	for userRows.Next() {
		userRows.Scan()
		acl, err := scanAccessControl(userRows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}
		userACL = append(userACL, acl)
	}
	if err := userRows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)

	}
	permissionIDs := make([]uint, 0)
	for _, r := range roleACL {
		if !slice.DoesExist(permissionIDs, r.PermissionID) {
			permissionIDs = append(permissionIDs, r.PermissionID)
		}
	}

	if len(permissionIDs) == 0 {
		return nil, nil
	}
	args := make([]any, len(permissionIDs))
	for i, id := range permissionIDs {
		args[i] = id
	}

	query := "select * from permissions where id in (?" +
		strings.Repeat(",?", len(permissionIDs)-1) +
		")"

	pRows, err := d.conn.Conn().Query(query,
		args...)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)

	}
	defer pRows.Close()
	permissionTitles := make([]entity.PermissionTitle, 0)
	for pRows.Next() {
		permission, err := scanPermission(rows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}
		permissionTitles = append(permissionTitles, permission.Title)
	}
	if err := pRows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)

	}
	return permissionTitles, nil
}
func scanAccessControl(scanner mysql.Scanner) (entity.AccessControl, error) {
	var createdAt time.Time
	var acl entity.AccessControl
	err := scanner.Scan(&acl.ID, &acl.ActorID, &acl.ActorType, &acl.PermissionID, &createdAt)
	return acl, err
}
