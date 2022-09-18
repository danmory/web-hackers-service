# Solution of Test Task

## Usage

1. Clone the repository

2. Install Docker, Docker-Compose

3. Initialize .env file with the following configuration

    ```dotenv
        APP_HOST=0.0.0.0
        APP_PORT=8010
        REDIS_HOST=db
        REDIS_PORT=6379
        REDIS_PASSWORD=yourpassword
    ```

4. Run the command

    `` $ docker-compose up -d --build ``

5. Populate Redis

    `` $ docker exec -i web-hackers-service_db_1 redis-cli -a youpassword < init.redis ``

6. Check result:

    `` $ curl -s http://localhost:8010/json/hackers |json_pp ``

## Author

Danila Moriakov(d.moriakov@gmail.com)
