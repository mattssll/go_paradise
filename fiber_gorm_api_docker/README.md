# Fiber Gorm SQLite API
Very simple tutorial followed from a youtube video. <br>
I only added docker and docker-compose as an exercise. <br>
<br>
# How to run it?
Make sure you have docker running, then: 
```
docker-compose up -d myrestapi
```
My restapi is the service name defined in docker-compose.yml <br>
The container will be accessible in localhost:3000 <br>
API Routes: <br>
```
app.Get("api/v1/books", book.GetBooks)
app.Get("api/v1/book/:id", book.GetBook)
app.Post("api/v1/book", book.NewBook)
app.Delete("api/v1/book/:id", book.DeleteBook)
```

## Kubernetes deploy
The kubernetes deploy was made with kompose <br>
Kompose transforms docker-compose into kubernetes.yml files.
### How?
```
docker compose build myrestapi
docker push gcr.io/gcpproject/myrestapi 
kompose convert -f docker-compose.yml
bash k8deploy.sh
```

