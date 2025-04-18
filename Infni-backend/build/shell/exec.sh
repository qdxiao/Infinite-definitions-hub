goctl api go -api apps/user.api -dir apps/api -style gozero
goctl model mysql ddl --src sql/user.sql --dir apps/api/internal/model