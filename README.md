1. To run all the services with an entire new environment run the following command
    rm -rf mysql/data && docker-compose rm -vf && docker-compose build --no-cache && docker-compose up

