#!/usr/bin/env bash
MYSQL_USERNAME="root"
MYSQL_PASSWORD="lark2022"
MYSQL_HOST="lark.com"
MYSQL_PORT=13306
MYSQL_DB="lark_user"

folder="mysql/lark_user/sqls"

for file in ${folder}/*
do
  mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} -D${MYSQL_DB} < ${file}
done

# 测试数据
#for i in {1..10};
#do
#  INSERT="INSERT INTO users ( uid, lark_id, mobile ) VALUES ( ${i}, ${i}, ${i} );"
#  mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} -D${MYSQL_DB} -e "$INSERT"
#done
#
## 测试数据
#for i in {1..10};
#do
#  INSERT="INSERT INTO chat_members
#          ( chat_id,chat_type, uid, alias, member_avatar,server_id)
#          VALUES
#          ( 3333336666669999990,2, ${i},CONCAT('name:',${i}),CONCAT('avatar',${i}),10000);"
#  mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} -D${MYSQL_DB} -e "$INSERT"
#done

<<xxxx

DROP_PROCEDURE="DROP PROCEDURE IF EXISTS insert_users;"
CREATE_PROCEDURE="CREATE PROCEDURE insert_users () BEGIN
                  	DECLARE
                  		n INT DEFAULT 1;
                  	WHILE
                  			n < 10001 DO
                  			INSERT INTO users ( uid, lark_id )
                  		VALUES
                  			( n, n );

                  		SET n = n + 1;

                  	END WHILE;

                  END;"
CALL_PROCEDURE="call insert_users();"
mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} -D${MYSQL_DB} -e "$DROP_PROCEDURE $CREATE_PROCEDURE $CALL_PROCEDURE"

xxxx

<<xxxx

DROP PROCEDURE IF EXISTS insert_users;
CREATE PROCEDURE insert_users () BEGIN
	DECLARE
		n INT DEFAULT 1;
	WHILE
			n < 10001 DO
			INSERT INTO users ( uid, lark_id )
		VALUES
			( n, n );

		SET n = n + 1;

	END WHILE;

END;

call insert_users();

xxxx

<<xxxx

mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} -D${MYSQL_DB} < ${file}

mysql -h127.0.0.1 -P3306 -uroot -p123456 -Dsdb
#参数
-h:host主机
-P:port端口
-u:user用户名
-p:password密码
-D:database数据库

xxxx
