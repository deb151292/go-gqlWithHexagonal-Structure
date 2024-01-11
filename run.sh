
echo "Building App"

docker-compose -f docker-compose-app.yml build

echo "Starting App"

docker-compose -f docker-compose-app.yml up
