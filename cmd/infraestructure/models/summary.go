package models

type Summary struct {
  ID string `json:"id"`
  UserId string `json:"user_id"`
  UserEmail string `json:"email"`
  Summary string `json:"summary"`
  ArtifactUrl string `json:"artifact_url"`
}