 ReST (Representational State Transfer)

CRUD => Create Read Update Destroy

HTTP Methods => {GET, PUT, POST, OPTIONS, PATCH, ...}

## Employee Management Server (JSON)

```
CRUD         Action             HTTP Method                   URI                   Request Body                     Response Body
-----------------------------------------------------------------------------------------------------------------------------------
Read         Index                GET                     /employees                    -                               [{...}, ...] -- Done
Read         Show                 GET                     /employees/{id}               -                               {...}        -- Done
Create       Create               POST                    /employees                  {...}                             {id: , ...}  -- Done
Update       Update               PUT                     /employees/{id}             {...}                             {...}
Update       Update               PATCH                   /employees/{id}             {selected attrs}                  {...}
Destroy      Destroy              DELETE                  /employess/{id}               -                                - / {...}
```


My SQL Server set-up in bash and windows -

install MySQL CLI- 
Visit the official MySQL website: https://dev.mysql.com/downloads/
Scroll down to the MySQL Community Downloads section.

go for install and follow as suggested.

Now open the MySQL command line client, and check using mysql command
then Set Envirrnent path to your pc, Add to path--> C:\Program Files\MySQL\MySQL Server 8.1\bin

Now mysql must be working for bash and windows cmd line
I'm using bash for further customization->>>>>>>>>

Sangwaniya ~/go/src
$ mysql --version
C:\Program Files\MySQL\MySQL Server 8.1\bin\mysql.exe  Ver 8.1.0 for Win64 on x86_64 (MySQL Community Server - GPL)

Sangwaniya ~/go/src
$ mysql -u root -p
Enter password: ****
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 11
Server version: 8.1.0 MySQL Community Server - GPL

Copyright (c) 2000, 2023, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
4 rows in set (0.01 sec)

mysql> create database employees;
Query OK, 1 row affected (0.00 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| employees          |
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.00 sec)


mysql> \q
Bye

Now we have to install mycli in pc, for that python must be installed, coz we using pip install, mac users can install using brew install mycli, for Windows users ->

now open cmd with administrative permissions and execute ->>>>>>>>
C:\Windows\System32>pip install mycli
Collecting mycli
  {Random log running...................}

Move back to git-bash and run following command->>>>>>>>
$ mycli mysql://root@localhost:3306/employees
Password:
MySQL
mycli 1.26.1
Home: http://mycli.net
Bug tracker: https://github.com/dbcli/mycli/issues
Thanks to the contributor - Kevin Schmeichel
MySQL root@localhost:employees> show tables;
+---------------------+
| Tables_in_employees |
+---------------------+
+---------------------+

0 rows in set
Time: 0.007s
MySQL root@localhost:employees>
