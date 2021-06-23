# forum

This project consists in creating a web forum that allows:
* communication between users.
* associating categories to posts.
* liking and disliking posts and comments.
* filtering posts.


### To run:

```go
go run cmd/app/main.go
OR
make && ./forum
``` 
**To run in Docker:**

```
1.docker build -t <custom name of the image> .

2.docker container run -p :17555 --name <custom name of the container> <name of the image>
``` 


