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

## Load testing with Yandex Tank

Load testing was carried out using Yandex Tank:

The service can withstand about 20,000 RPS
on Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz and 16 GB RAM.

Try it out(config file is *load.yaml*):

```bash
    docker run \
        -v $(pwd):/var/loadtest \
        -v $SSH_AUTH_SOCK:/ssh-agent -e SSH_AUTH_SOCK=/ssh-agent \
        --net host \
        -it direvius/yandex-tank
```

## Author

Danila Moriakov(d.moriakov@gmail.com)
