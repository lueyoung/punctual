## 0 Paper

## 1 Configure Cassandra
Before you execute the program, Launch `cqlsh` and execute:
```
CREATE KEYSPACE punctual WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
CREATE TABLE punctual.history(key text, value text, timeline UUID, PRIMARY KEY(key));
CREATE INDEX on punctual.history(timeline);
```

## 2 Run
run:
```
make
```
