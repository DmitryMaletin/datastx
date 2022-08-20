# Goal
**This all needs to sit on a kubecluster and have terraform**
 - Replication system that real times from pg || msSQL -> snowflake, bigquery, msSQL or pg
 - A rebuild of DBT in GoLang
 - An entity store that acts as the glue between pg and snowflake...home made
 - Web server that can access all this
    * can upload files from here
    * has a data lineage graph since we own the pipeline system
    * can schedule basic jobs like file uploads or exports from here
 - Light weight Scheduling Server
 - Snowflake Formatter/Linter...uses antlr4
 - Looker Formatter/Linter...hand rollded in go