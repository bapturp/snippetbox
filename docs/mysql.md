## MySQL

```sh
# Start MySQL daemon
brew services start mysql

# Open MySQL Command-Line Client
mysql -u root

# Stop MySQL daemon
brew services stop mysql
```

```sql
-- use a database
USE snippetbox;

-- show tables
show tables;

-- select some data
SELECT id, title, expires FROM snippets WHERE id = 4;

-- exit
exit;
```
