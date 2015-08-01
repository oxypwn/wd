@echo off
for /F "skip=2 tokens=2 delims=," %%a in ('wmic csproduct get uuid /format:csv') do if not defined uuid set uuid=%%a
wd.exe http://mbp.home.lan:8080/v1/%uuid%/postinstall postinstall



::PAUSE
