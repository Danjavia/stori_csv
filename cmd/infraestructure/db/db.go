package database

import (
    "context"
    "database/sql"

    "github.com/danjavia/stori_csv/cmd/infraestructure/models"
    _ "github.com/lib/pq" // Import the Postgres driver
)

func CreateSummary(ctx context.Context, db *sql.DB, summary *models.Summary) error {
    query := `INSERT INTO summary (id, user_id, user_email, summary, artifact_url)
               VALUES ($1, $2, $3, $4, $5)`
    stmt, err := db.PrepareContext(ctx, query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.ExecContext(ctx, summary.ID, summary.UserId, summary.UserEmail, summary.Summary, summary.ArtifactUrl)
    return err
}

func GetSummaries(ctx context.Context, db *sql.DB, userId string) ([]*models.Summary, error) {
    query := `SELECT id, user_id, user_email, summary, artifact_url FROM summary WHERE user_id = $1`
    rows, err := db.QueryContext(ctx, query, userId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    summaries := []*models.Summary{}
    for rows.Next() {
        summary := &models.Summary{}
        err := rows.Scan(&summary.ID, &summary.UserId, &summary.UserEmail, &summary.Summary, &summary.ArtifactUrl)
        if err != nil {
            return nil, err
        }
        summaries = append(summaries, summary)
    }

    return summaries, nil
}