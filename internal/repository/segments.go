package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log/slog"
	"segments-api/internal/logger/sl"
	"segments-api/internal/model/segment"
	"segments-api/pkg/database"
	"strings"
	"time"
)

type SegmentRepositoryImpl struct {
	client database.Client
	logger *slog.Logger
}

func New(client database.Client, logger *slog.Logger) *SegmentRepositoryImpl {
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
	//TODO implement me
	panic("implement me")
}

type idScanner struct {
	id int
}

func (r SegmentRepositoryImpl) AddUser(add []string, remove []string, userId int) error {
	r.logger.Info("executing sql...")

	addSql := `
		INSERT INTO USER_SEGMENTS(USER_ID, SEGMENT_ID, CREATED, EXPIRED)
		VALUES ($1, (SELECT ID FROM SEGMENTS WHERE NAME = $2), $3, $4)
	`

	deleteSql := `
		DELETE
		FROM USER_SEGMENTS
		WHERE USER_ID = $1 AND SEGMENT_ID IN (select id from segments where name = $2)
	`

	var batch = &pgx.Batch{}
	for _, elem := range add {
		batch.Queue(addSql, userId, elem, time.Now(), time.Now())
	}
	for _, elem := range remove {
		batch.Queue(deleteSql, userId, elem)
	}

	br, err := r.client.SendBatch(context.Background(), batch).Exec()
	if err != nil {
		r.logger.Error("batch operation failed", sl.Err(err))
	}

	fmt.Printf("rows affected => %d\n", br.RowsAffected())

	return err
}

func (r SegmentRepositoryImpl) GetAllByUser(userID int) ([]segment.Segment, error) {
	//TODO implement me
	panic("implement me")
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}
