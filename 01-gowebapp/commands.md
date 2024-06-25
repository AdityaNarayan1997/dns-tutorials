docker build -t gowebapp:0.0.1 .

docker run -d -p 80:8080 gowebapp:0.0.1