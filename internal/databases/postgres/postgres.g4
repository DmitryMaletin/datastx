grammar PostgreSQL;

// root
//   : statement+
//   ;

// statement
//   : create_table
//   | select
//   | insert
//   | update
//   | delete
//   | create_index
//   | drop_table
//   | alter_table
//   | truncate
//   | grant
//   | revoke
//   ;

// create_table
//   : 'CREATE TABLE' table_name '(' column_definition (',' column_definition)* ')'
//   ;

// column_definition
//   : column_name data_type (NOT NULL)?
//   ;

// data_type
//   : INTEGER
//   | BIGINT
//   | SMALLINT
//   | REAL
//   | DOUBLE PRECISION
//   | NUMERIC
//   | CHARACTER
//   | CHARACTER VARYING
//   | DATE
//   | TIME
//   | TIMESTAMP
//   | INTERVAL
//   | BOOLEAN
//   | ENUM
//   | JSON
//   | JSONB
//   ;

// select
//   : 'SELECT' select_clause 'FROM' table_name (where_clause)? (order_by_clause)? (limit_clause)?
//   ;

// select_clause
//   : '*'
//   | select_item (',' select_item)*
//   ;

// select_item
//   : column_name
//   | function_call
//   | '*'
//   ;

// function_call
//   : function_name '(' (expression (',' expression)*)? ')'
//   ;

// where_clause
//   : 'WHERE' expression
//   ;

// order_by_clause
//   : 'ORDER BY' order_by_item (',' order_by_item)*
//   ;

// order_by_item
//   : expression (ASC | DESC)?
//   ;

// limit_clause
//   : 'LIMIT' INTEGER
//   ;

// insert
//   : 'INSERT INTO' table_name '(' column_name (',' column_name)* ')' 'VALUES' '(' value (',' value)* ')'
//   ;

// update
//   : 'UPDATE' table_name 'SET' assignment (',' assignment)* (where_clause)?
//   ;

// assignment
//   : column_name '=' expression
//   ;

// delete
//   : 'DELETE FROM' table_name (where_clause)?
//   ;

// create_index
//   : 'CREATE INDEX' index_name 'ON' table_name '(' column_name ')'
//   ;

// drop_table
//   : 'DROP TABLE' table_name
//   ;

// alter_table
//   : 'ALTER TABLE' table_name 'ADD COLUMN' column_definition
//   | 'ALTER TABLE' table_name 'DROP COLUMN' column_name
//   | 'ALTER TABLE' table_name 'RENAME TO' table_name
//   ;

// truncate
//   : 'TRUNCATE TABLE' table_name
//   ;

// grant
//   : 'GRANT' privilege (',' privilege)* 'ON TABLE' table_name 'TO' role_name
