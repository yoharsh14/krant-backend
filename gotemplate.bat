@echo off
REM this ensure the script run without echoing every line to the console

REM -- 1. Creating the cmd folder --
REM MD stand

MD cmd
REM creating a file
COPY NUL "cmd/api.go" > NUL
COPY NUL "cmd/main.go" > NUL


MD internal
MD "internal/adapters"
MD "internal/env"
MD "internal/json"
MD "internal/business"
MD "internal/routers"
MD "internal/models"
@echo off
set ENV_FILE=internal\env\env.go

(
echo package env
echo.
echo import "os"
echo.
echo func GetString(key, fallback string^) string ^{
echo     if val ^:= os.Getenv(key^); val ^!= "" ^{
echo         return val
echo     ^}
echo     return fallback
echo ^}
) > "%ENV_FILE%"



@echo off
setlocal DisableDelayedExpansion

set FILE=internal\json\json.go


(
echo package json
echo.
echo import ^(
echo     "encoding/json"
echo     "net/http"
echo ^)
echo.
echo func Write(w http.ResponseWriter, status int, data any^) ^{
echo     w.Header^(^).Set^("Content-Type", "application/json"^)
echo     w.WriteHeader^(http.StatusOK^)
echo     json.NewEncoder^(w^).Encode^(data^)
echo ^}
echo.
echo func Read(r *http.Request, data any^) error ^{
echo     decoder ^:= json.NewDecoder^(r.Body^)
echo     decoder.DisallowUnknownFields^(^)
echo     return decoder.Decode^(data^)
echo ^}
) > "%FILE%"

echo Go file generated at %JSON%


echo Script finished. Press any key to exit.
pause > NUL