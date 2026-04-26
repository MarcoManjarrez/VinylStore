# Vinyl Store API
A Go API for managing a vinyl record registry, with JWT authentication and Gin framework.

## Instalation
### Prerequisites
> Go version 1.20 or higher
> curl 8.18.0 or higher

### Setup
On the terminal, run the following commands:
```
go mod download
go run .\main.go
```

## Endpoint table
| Endpoint | Method | Description|
|----------|--------|------------|
| /createAccount| POST | Allows for the creation of an account in the server |
| /login | POST | Logs user into the system. Provides a token which must be used for other endpoints |
| /logout | POST | Logs user out of the system. Invalidates tokens generated in login |
| /albums | GET | Returns the entire list of albums recorded in the registry. Locked behind auth token. |
| /albums/:id | GET | Returns the listing of a single vinyl by id. Locked behind auth token. |
| /post-album | POST | Adds a new album to the registry. Locked behind auth token. |
| /createAlbum | POST | Adds a new album to the registry. Locked behind auth token. |
| /status | GET | Gets system and user status. Locked behind auth token. |

## Usage examples
All of the following examples will be demonstrations for curl requests. 

### /createAccount
```
curl -X POST http://localhost:8080/createAccount \
     -H "Content-Type: application/json" \
     -d "{\"username\": \"<New username>\", \"password\": \"<New password>\"}"
```

### /login
```
 curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d "{\"username\": \"<Username>\", \"password\":\"<Password>\"}
```
This endpoint will return a message similar to the following
```
{
"message": "Hi username, welcome to the Store System",
"token": <Your token>
}
```
>The token will be necessary to use the other protected endpoints.

### /logout
```
curl -H "Authorization: Bearer <Your token>" http://localhost:8080/logout
```

### /albums
```
curl -H "Authorization: Bearer <Your token>" http://localhost:8080/albums
```

### /albums/:id
```
curl -H "Authorization: Bearer <Your token>" http://localhost:8080/albums/<Existing album ID>
```

### /post-album
```
curl http://localhost:8080/post-album
--include
--header "Authorization: Bearer <Your token>"
--header "Content-Type: application/json"
--request "POST"
--data '{"id": "4","title": "Benny Golson's New York Scene","artist": "Benny
Golson","price": 49.99}'
```

### /status 
```
curl -H "Authorization: Bearer <Your token>" http://localhost:8080/status
```


## Technical details
The system was developed with the following technologies:
- Framework: The system utilizes Gin Gonic for routing management
- Security: bcrypt hashing was used for secure password storage  
- Authentication: JWT, or JSON web tokens were utilized for stateless session management 