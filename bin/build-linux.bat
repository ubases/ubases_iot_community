SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

#go mod tidy
#go mod vendor

for /d %%i in ("../iot_*_service") do (
  echo %%i
  go build -tags="ce" ../%%i
)
pause
