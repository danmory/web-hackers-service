# Solution of Test Task

## Usage

1. Clone the repository

2. Install Docker, Docker-Compose

3. Initialize .env file with the following configuration

    ```dotenv
        APP_ADDRESS=:8010
        REDIS_ADDRESS=db:6379
        REDIS_PASSWORD=yourpassword
    ```

4. Run the command

    `` $ docker-compose up -d --build ``

5. Populate Redis

    `` $ docker exec -it web-hackers-service_db_1 redis-cli -a yourpassword ``

    Then run in the opened command line the following:

    ```redis
        zadd hackers 1953 "Richard Stallman"  
        zadd hackers 1940 "Alan Kay"
        zadd hackers 1965 "Yukihiro Matsumoto"
        zadd hackers 1916 "Claude Shannon"
        zadd hackers 1969 "Linus Torvalds"
        zadd hackers 1912 "Alan Turing"
    ```

6. Exit redis-cli and run:

    `` $ curl -s http://localhost:8010/json/hackers |json_pp ``

## Author

Danila Moriakov(d.moriakov@gmail.com)
