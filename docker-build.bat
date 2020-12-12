@ECHO OFF
IF "%1"=="NO-CACHE" docker build --no-cache --tag atlas-dis:latest .
IF NOT "%1"=="NO-CACHE" docker build --tag atlas-dis:latest .