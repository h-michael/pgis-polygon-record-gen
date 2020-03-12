# pgis-polygon-record-gen

```sh
  -t string
        required, target table name
  -c string
        required, target column name
  -s int
        optional, default 1000, record size (default 1000)
```

```sh
$ pgis-polygon-record-gen -t test_table -c test_polygon -s 2 > seed.sql
$ cat ./seed.sql
INSERT INTO "test_table" ("test_column") VALUES
	(ST_PolygonFromText('POLYGON((-10 15, -53 98, 70 115, 88 6, -10 15))', 4326)),
	(ST_PolygonFromText('POLYGON((46 -142, -77 13, -32 -40, 72 114, 46 -142))', 4326));
```
