
__Docker Commands__

```
//Rebuild and run all the containers
docker compose up -d --build

//Rebuild and run the specific container
docker-compose up --no-deps <<myservice>>

```
Mind that after each rebuild a copy of the old container persists, clean unused containers
Reference:https://betterstack.com/community/questions/how-to-restart-single-container-using-docker-compose/

__Leaflet Mandatory Requirements__
include the css from node_modules
```
import 'leaflet/dist/leaflet.css';
```

In index.css or App.css 
```
#root {
  height: 100vh;
  width: 100vw;
}

.leaflet-container {
  width: 100vw; 
  height: 100vh;
}

```
otherwise check official website for CDN header links

__vite.config.ts__ 

add this snippet for hot reload inside Docker container [check if it works]

```
  server: {
    host: true,
    port: 3000, // This is the port which we will use in docker
                // Thanks @sergiomoura for the window fix
                // add the next lines if you're using windows and hot reload doesn't work
     watch: {
       usePolling: true
     }
```

__Go pkgs__

```
go get github.com/gin-gonic/gin
```
An implementation of GraphQL in Go. Follows the official reference implementation graphql-js.
Supports: queries, mutations & subscriptions.
```
go get github.com/graphql-go/graphql
```
Golang HTTP.Handler for graphl-go
```
go get github.com/graphql-go/handler
```
pq - A pure Go postgres driver for Go's database/sql package
```
go get github.com/lib/pq
```