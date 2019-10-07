# build
docker image build -t example/cronjob:latest .

# run
docker container run -d --rm --name cronjob example/cronjob:latest

# view
docker container exec -it cronjob tail -f /var/log/cron.log