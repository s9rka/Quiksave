## Set Up Database
Ensure Docker and Docker Compose are installed and running.
Navigate to the build folder.

Start the database for the first time:
```
docker compose up -d
```

This will start the PostgreSQL container (nota-bene-db).
The database will automatically initialize using the init.sql script located in the database/ folder.

## Start Database
To start the database (if already set up):

Run:
```
docker compose up
```

This will start the nota-bene-db container if it’s stopped.
Verify it’s running:
`docker ps`

## Remove and Reinitialize Database
If you’ve made changes to the database schema or init.sql:

Stop and remove the container:
```
docker compose down -v
```

The -v flag ensures the data volume is also removed, resetting the database completely.
Reinitialize the database:
`docker compose up -d`

The init.sql script will run again, applying the new schema and setup.
