# mysql-client

基于gorm的实现

```bash
go get github.com/icowan/mysql-client
```

## Demo

```golang
dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local&timeout=20m&collation=utf8mb4_unicode_ci",
		"root",
		"password",
		"127.0.0.1",
		3306,
		"dbname")

db, err := NewMysql(dbUrl, true)
if err != nil {
    log.Fatal(err)
}
```