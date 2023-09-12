package postgres

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	errors "github.com/webitel/engine/model"
	"github.com/webitel/logger/model"
	"github.com/webitel/logger/storage"
	"strings"
	"time"
)

type Log struct {
	storage storage.Storage
}

func newLogStore(store storage.Storage) (storage.LogStore, errors.AppError) {
	if store == nil {
		return nil, errors.NewInternalError("postgres.log.new_log.check.bad_arguments", "error creating log interface to the log table, main store is nil")
	}
	return &Log{storage: store}, nil
}

func (c *Log) Get(ctx context.Context, opt *model.SearchOptions, filters ...model.Filter) ([]*model.Log, errors.AppError) {
	db, appErr := c.storage.Database()
	if appErr != nil {
		return nil, appErr
	}
	base := ApplyFiltersToBuilder(c.GetQueryBaseFromSearchOptions(opt), filters...)
	fmt.Println(base.ToSql())
	rows, err := base.RunWith(db).QueryContext(ctx)
	if err != nil {
		return nil, errors.NewInternalError("postgres.log.get_by_object_id.query_execute.fail", err.Error())
	}
	defer rows.Close()
	res, appErr := c.ScanRows(rows)
	if appErr != nil {
		return nil, appErr
	}
	return res, nil
}

func (c *Log) Insert(ctx context.Context, log *model.Log) errors.AppError {
	db, appErr := c.storage.Database()
	if appErr != nil {
		return appErr
	}
	_, err := db.ExecContext(ctx,
		`INSERT INTO
			logger.log(date, action, user_id, user_ip, new_state, record_id, config_id) 
		VALUES
			(
			$1, $2, $3, $4, $5, $6, $7
			)`,
		log.Date, log.Action, log.User.Id, log.UserIp, log.NewState, log.Record.Id, log.ConfigId,
	)
	if err != nil {
		return errors.NewInternalError("postgres.log.insert.scan.error", err.Error())
	}
	//defer rows.Close()
	//res, appErr := c.ScanRows(rows)
	//if appErr != nil {
	//	return nil, appErr
	//}
	//newModel = res[0]
	return nil
}

func (c *Log) InsertMany(ctx context.Context, log []*model.Log) errors.AppError {
	db, appErr := c.storage.Database()
	if appErr != nil {
		return appErr
	}
	base := sq.Insert("logger.log").Columns("date", "action", "user_id", "user_ip", "new_state", "record_id", "config_id").PlaceholderFormat(sq.Dollar)
	for _, v := range log {
		base = base.Values(v.Date, v.Action, v.User.Id, v.UserIp, v.NewState, v.Record.Id, v.ConfigId)
	}
	_, err := base.RunWith(db).ExecContext(ctx)
	if err != nil {
		return errors.NewInternalError("postgres.log.insert.query.error", err.Error())
	}
	return nil
}

func (c *Log) DeleteByLowerThanDate(ctx context.Context, date time.Time, configId int) (int, errors.AppError) {
	db, appErr := c.storage.Database()
	if appErr != nil {
		return 0, appErr
	}
	rows, err := db.ExecContext(ctx,
		`DELETE FROM logger.log WHERE log.date < $1 AND log.config_id = $2 `,
		date, configId,
	)
	if err != nil {
		return 0, errors.NewInternalError("postgres.log.delete_by_lowe_that_date.query.error", err.Error())
	}
	affected, err := rows.RowsAffected()
	if err != nil {
		return 0, errors.NewInternalError("postgres.log.delete_by_lowe_that_date.result.error", err.Error())
	}
	return int(affected), nil
}

func (c *Log) ScanRows(rows *sql.Rows) ([]*model.Log, errors.AppError) {
	if rows == nil {
		return nil, errors.NewInternalError("postgres.log.scan.check_args.rows_nil", "rows are nil")
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.NewInternalError("postgres.log.scan.get_columns.error", err.Error())
	}
	var logs []*model.Log

	for rows.Next() {
		var log model.Log
		binds := make([]func(dst *model.Log) interface{}, 0, 0)
		for _, v := range cols {

			switch v {
			case "id":
				binds = append(binds, func(dst *model.Log) interface{} { return &dst.Id })
			case "date":
				binds = append(binds, func(dst *model.Log) interface{} { return &dst.Date })
			case "user_id":
				binds = append(binds, func(dst *model.Log) interface{} { return &dst.User.Id })
			case "user_name":
				binds = append(binds, func(dst *model.Log) interface{} { return &dst.User.Name })
			case "user_ip":
				binds = append(binds, func(dst *model.Log) interface{} { return &dst.UserIp })
			case "new_state":
				binds = append(binds, func(dst *model.Log) interface{} { return &dst.NewState })
			case "record_id":
				binds = append(binds, func(dst *model.Log) interface{} { return &dst.Record.Id })
			case "action":
				binds = append(binds, func(dst *model.Log) interface{} { return &dst.Action })
			case "config_id":
				binds = append(binds, func(dst *model.Log) interface{} { return &dst.ConfigId })
			case "object_id":
				binds = append(binds, func(dst *model.Log) interface{} { return &dst.Object.Id })
			case "object_name":
				binds = append(binds, func(dst *model.Log) interface{} { return &dst.Object.Name })
			default:
				panic("postgres.log.scan.get_columns.error: columns gotten from sql don't respond to model columns. Unknown column: " + v)
			}

		}
		bindFunc := func(binds []func(*model.Log) any) []any {
			var fields []any
			for _, v := range binds {
				fields = append(fields, v(&log))
			}
			return fields
		}
		err = rows.Scan(bindFunc(binds)...)
		if err != nil {
			return nil, errors.NewInternalError("postgres.log.scan.scan.error", err.Error())
		}
		logs = append(logs, &log)
	}
	if len(logs) == 0 {
		return nil, errors.NewBadRequestError("postgres.log.scan.check_no_rows.error", sql.ErrNoRows.Error())
	}
	return logs, nil
}

func (c *Log) GetQueryBaseFromSearchOptions(opt *model.SearchOptions) sq.SelectBuilder {
	var fields []string
	for _, v := range opt.Fields {
		switch v {
		case "id":
			fields = append(fields, "log.id")
		case "date":
			fields = append(fields, "log.date")
		case "user_ip":
			fields = append(fields, "log.user_ip")
		case "new_state":
			fields = append(fields, "log.new_state")
		case "record":
			fields = append(fields, "log.record_id")
		case "action":
			fields = append(fields, "log.action")
		case "config_id":
			fields = append(fields, "log.config_id")
		case "user":
			fields = append(fields, "coalesce(wbt_user.name::varchar, wbt_user.username::varchar) as user_name", "log.user_id")
		case "object":
			fields = append(fields, "wbt_class.id as object_id", "wbt_class.name as object_name")
		}
	}
	if len(fields) == 0 {
		fields = append(fields,
			c.getFields()...)
	}
	base := c.GetQueryBase(fields)
	if opt.Search != "" {
		base = base.Where(sq.Like{"user_ip": opt.Search + "%"})
	}
	if opt.Sort != "" {
		splitted := strings.Split(opt.Sort, ":")
		if len(splitted) == 2 {
			order := splitted[0]
			column := splitted[1]
			if column == "user" {
				column = "user_name"
			}
			base = base.OrderBy(fmt.Sprintf("%s %s", column, order))
		}

	}
	offset := (opt.Page - 1) * opt.Size
	if offset < 0 {
		offset = 0
	}
	return base.Offset(uint64(offset)).Limit(uint64(opt.Size + 1))
}

func (c *Log) GetQueryBase(fields []string) sq.SelectBuilder {
	base := sq.Select(fields...).
		From("logger.log").
		JoinClause("LEFT JOIN directory.wbt_user ON wbt_user.id = log.user_id").
		JoinClause("LEFT JOIN logger.object_config ON object_config.id = log.config_id").
		JoinClause("LEFT JOIN directory.wbt_class ON wbt_class.id = object_config.object_id").
		PlaceholderFormat(sq.Dollar)

	return base
}

//func (c *Log) FillRecordLookup(ctx context.Context, models []*model.Log, domainId int64) errors.AppError {
//	var (
//		uniqueRecordIds map[string][]int64 // determining what records should be selected in the query
//		resultLookups   []*model.Lookup
//	)
//
//	// range through all records to know the unique record for each object
//	for _, record := range models {
//		if objectName := record.Object.Name.String(); objectName != "" { // check for valid object name
//			if record.Record.Id.Int() != 0 { // check for valid record id
//				if value, ok := uniqueRecordIds[objectName]; ok {
//					// ! append to the existing table
//					uniqueRecordIds[objectName] = append(value, record.Record.Id.Int64())
//				} else {
//					uniqueRecordIds[objectName] = []int64{record.Record.Id.Int64()}
//				}
//			}
//		}
//	}
//
//	db, appErr := c.storage.Database()
//	if appErr != nil {
//		return appErr
//	}
//	for key, values := range uniqueRecordIds {
//
//		base, appErr := c.GetRecordQuery(key, domainId, values)
//		if appErr != nil {
//			continue
//		}
//		query, args, _ := base.ToSql()
//		queryRes, err := db.QueryxContext(ctx, query, args)
//		//queryRes, err := base.RunWith(db).PlaceholderFormat(sq.Dollar).QueryContext(ctx)
//		if err != nil {
//			return errors.NewInternalError("postgres.log.fill_record_lookup.query.fail", err.Error())
//		}
//		for queryRes.Next() {
//			var res model.Lookup
//			err = queryRes.StructScan(&res)
//			if err != nil {
//				return errors.NewInternalError("postgres.log.fill_record_lookup.scan.fail", err.Error())
//			}
//			resultLookups = append(resultLookups, &res)
//		}
//		for _, log := range models {
//			for _, lookup := range resultLookups {
//				if log.Object.Name.String() == key && log.Record.Id.Int64() == lookup.Id.Int64() {
//					log.Record.Name = model.NewNullString(lookup.Name.String())
//				}
//			}
//		}
//	}
//	return nil
//}

//func (c *Log) GetRecordQuery(objectName string, domainId int64, recordsId []int64) (sq.SelectBuilder, errors.AppError) {
//	value, ok := storage.RecordTablesMapping[objectName]
//	if !ok {
//		errors.NewInternalError("postgres.log.get_record_query.search_for_table.fail", "there are record in logs that has unknown object")
//	}
//	base := sq.Select("id", value.ColumnName).From(value.GetFullPath()).Where(sq.Eq{value.ColumnDomain: domainId}).Where("id = any(?)", pq.Array(recordsId))
//
//	return base, nil
//}

func (c *Log) getFields() []string {
	return []string{
		"log.id",
		"log.date",
		"log.user_id",
		"coalesce(wbt_user.name::varchar, wbt_user.username::varchar) as user_name",
		"log.user_ip",
		"log.new_state",
		"log.record_id",
		"log.action",
		"log.config_id",
		"wbt_class.id as object_id",
		"wbt_class.name as object_name",
	}
}
