#!/bin/sh
 
DATABASE=users.db
 
cat "db_init.sql" | sqlite3 "${DATABASE}"
