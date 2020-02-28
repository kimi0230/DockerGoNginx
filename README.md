#  Kimi's Gin Framework
## Project structure
    app
        controllers/
            api/
                yourfolder/
                    xxx.go
        
        middleware/
                yourfolder/
                    xxx.go

        models/		---  db model
            xxx.go

        services/
            awsservices/
                awsservices.go
            commomservices/
                commomservices.go --- useful function
            emailservices/
                emailservices.go --- send email
            firebaseservices/
                firebaseservices.go --- firebase connection
            ginservices/
                ginservices.go --- for gin request, response function
            modelservices/
                modelservices.go --- model useful funcion
            redisservices/
                redisservices.go --- redis connection
            cinfig.go 
            constant.go

    config/
        databases/
            mysql/
                mysql.go --- db connection
        errorCode/
            errorCode.go --- API return error code
        constant.go --- constant variable
        validateParam.go --- verify struct

    database/
        migrate.go --- execute migrate databases
            ex.  go run migrate -m=[auto/drop/refresh/create]
            auto: Automatically migrate your schema, to keep your schema update to date.
            drop: Drop table.
            create: Create Table.
            refresh: Execute drop and auto.

    logs/
        YYYYMMDD_stdout.log

    routes/
        routes.go --- url routes
    test/ --- unit test

    .env		--- Production Environment
    .env.dev 	--- RD Test Environment
    .env.qa 	--- QA Test Environment
    .env.example	--- Environment Example
    main.go		--- start project
        ex. go run main.go [app/dev/qa]
    pm2app.json	--- start production environment by pm2
    pm2dev.json	--- start RD test environment by pm2
    pm2qa.json	--- start qa test environment by pm2
    
## Init vendor, will create vendor.go
    govendor init

## Copy used package which locate at $GOPAH to vendor folder
    govendor add +external

## Install package 
    govendor fetch PACKAGE_NAME

## Install package from vendor/vendor.go
    govendor sync

## Create ".env" ".env.dev"  and ".env.qa" files
    cp .env.example .env
    cp .env.example .env.dev
    cp .env.example .env.qa

## Migrate Database, -m=auto/drop/create/refresh
    go run database/migrate.go -m=auto

## fresh server
    fresh -c fresh.conf

## Run will run defaut env (.env.dev)
    go run main.go

## Run Production  (.env)
    go run main.go app

## Run Development (.env.dev)
    go run main.go dev

## Run QA (.env.qa)
    go run main.go qa