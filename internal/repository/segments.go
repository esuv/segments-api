package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"segments-api/internal/logger/sl"
	"segments-api/internal/model/segment"
	"strings"
	"time"
)

type SegmentRepositoryImpl struct {
	client *pgxpool.Pool
	logger *slog.Logger
}

func New(client *pgxpool.Pool, logger *slog.Logger) *SegmentRepositoryImpl {
	return &SegmentRepositoryImpl{client: client, logger: logger}
}

func (r SegmentRepositoryImpl) Create(name string) (segment.Segment, error) {
	var sgm = segment.Segment{Name: name}
	sql := `
		INSERT INTO SEGMENTS(NAME)
		VALUES ($1)
		RETURNING ID
	`

	r.logger.Debug(fmt.Sprintf("executing sql: %s", formatQuery(sql)))

	if err := r.client.QueryRow(context.Background(), sql, name).Scan(&sgm.ID); err != nil {
		return segment.Segment{}, err
	}

	return sgm, nil
}

func (r SegmentRepositoryImpl) Delete(name string) error {
	deleteUS := `
		DELETE
		FROM USER_SEGMENTS
		WHERE SEGMENT_ID = (SELECT ID FROM SEGMENTS WHERE NAME = $1);
	`

	deleteSegment := `
		DELETE FROM SEGMENTS WHERE NAME = $1;
	`

	var batch = &pgx.Batch{}

	batch.Queue(deleteUS, name)
	batch.Queue(deleteSegment, name)

	return r.executeBatch(batch)
}

func (r SegmentRepositoryImpl) AddUser(add []string, remove []string, userId int) error {
	r.logger.Info("executing sql...")

	batch := prepareBatch(add, remove, userId)

	return r.executeBatch(batch)
}

func prepareBatch(add []string, remove []string, userId int) *pgx.Batch {
	addSql := `
		INSERT INTO USER_SEGMENTS(USER_ID, SEGMENT_ID, CREATED, EXPIRED)
		VALUES ($1, (SELECT ID FROM SEGMENTS WHERE NAME = $2), $3, $4)
		ON CONFLICT DO NOTHING;
	`

	deleteSql := `
		DELETE
		FROM USER_SEGMENTS
		WHERE USER_ID = $1 AND SEGMENT_ID IN (select id from segments where name = $2);
	`

	var batch = &pgx.Batch{}
	defaultTTL := 30 * time.Hour * 24

	for _, elem := range add {
		batch.Queue(addSql, userId, elem, time.Now(), time.Now().Add(defaultTTL))
	}
	for _, elem := range remove {
		batch.Queue(deleteSql, userId, elem)
	}

	return batch
}

func (r SegmentRepositoryImpl) GetAllByUser(userID int) ([]segment.Segment, error) {
	sql := `
		SELECT DISTINCT S.ID, S.NAME
		FROM SEGMENTS S
			LEFT JOIN PUBLIC.USER_SEGMENTS US ON S.ID = US.SEGMENT_ID
		WHERE US.USER_ID = $1
	`

	r.logger.Debug(fmt.Sprintf("executing sql: %s", formatQuery(sql)))

	rows, err := r.client.Query(context.Background(), sql, userID)
	if err != nil {
		return nil, err
	}

	segments, err := pgx.CollectRows(rows, pgx.RowToStructByName[segment.Segment])
	if err != nil {
		r.logger.Error(fmt.Sprintf("CollectRows error: %v", err), sl.Err(err))
		return nil, err
	}

	return segments, nil
}

func (r SegmentRepositoryImpl) executeBatch(batch *pgx.Batch) error {
	br, err := r.client.SendBatch(context.Background(), batch).Exec()
	if err != nil {
		r.logger.Error("batch operation failed", sl.Err(err))
	}

	r.logger.Info(fmt.Sprintf("rows affected => %d\n", br.RowsAffected()))
	return err
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}
