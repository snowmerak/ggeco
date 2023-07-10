docker run -e "ACCEPT_EULA=1" \
    -e "MSSQL_SA_PASSWORD=M@sterP@ssword1234!?" -e "MSSQL_PID=Developer" \
    -e "MSSQL_USER=SA" -p 1433:1433 --name=test_sql mcr.microsoft.com/azure-sql-edge