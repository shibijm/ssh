@echo off
setlocal
set GOARCH=amd64
set GOOS=windows
echo Building Windows binary...
go build -ldflags "-s -w" -trimpath -o out/
endlocal
if not %errorlevel%==0 (
	echo Build failed
	exit /b %errorlevel%
) else (
	echo Build successful
)
