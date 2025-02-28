# Executus-Server
MP3/Wav web server

To create the Docker image
docker build -t executus-server .

To run the Docker image
docker run -p 8080:8080 -e ServerPort=8080 executus-server