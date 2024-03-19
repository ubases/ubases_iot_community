
rem 生成整个库
..\\..\\tool\\mysql-to-model\\gentool.exe  -dsn "root:iLmz8sCXjkTYuh@tcp(192.168.5.56:3306)/iot_statistics?charset=utf8mb4&parseTime=True&loc=Local"  -outPath  orm -modelPkgName model

rem 生成单元测试加 -withUnitTest true

rem 生成某个表，后边加 -tables {table name}
