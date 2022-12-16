# Goal
**This all needs to sit on a kubecluster and have terraform**
 - Replication system that real times from pg || msSQL -> snowflake, bigquery, msSQL or pg
   * add type 1 & type 2 features
   * performance partitioning features
 - A rebuild of DBT in GoLang
 - An entity store that acts as the glue between pg and snowflake...home made
 - Web server that can access all this
    * can upload files from here
    * has a data lineage graph since we own the pipeline system
    * can schedule basic jobs like file uploads or exports from here
 - Light weight Scheduling Server
 - Snowflake Formatter/Linter...uses antlr4
 - Looker Formatter/Linter...hand rollded in go


# SQL Runner
- mimic basic features of dbt
- every model must have a schema.yml file to run in prod
  - there may be some nuance around this
- offer a build in CI runner
  - CI runner should also offer a faker dataset
- Can run the dag at any point in the project and offer runtime

# Scheduling server
- offer basic runtime over watcher

# Lineage tool
- this could plug right into the scheduler

# Exporter
- export basic lookml
- export a metrics layer
- export pipelines as dags into airflow
  - only supports airflow with task mapping

# Linter Engine
- Enforce basic rules


# Replication System
- replicate from postgres to snowflake as an mvp
- offer a type 1 & 2 option of a table

# Light weight programming language to reflect python
- For templating lets make a programming language the mimics python due to our end consumer


# Import of data

# Export of data