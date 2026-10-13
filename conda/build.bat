@echo off
REM build.bat - For Windows builds

if not exist "%PREFIX%\bin" mkdir "%PREFIX%\bin"

copy bin\libpermissions-connector-s2.dll "%PREFIX%\bin\" || exit 1
copy permissions-connector-s2.pplugin "%PREFIX%\" || exit 1
copy permissions-connector-s2.yaml "%PREFIX%\" || exit 1

exit 0
