# Getting Started
`docker run miguelpragier/simpleservice:latest -p 3001:80`
<br>
`curl http://127.0.0.1:3001/simple/ping/`

# Simple Service
Containerized service for learn and test docker basics

Here's the currently available methods in this webservice:

- GET /simple/uptime/
Returns the time since the service booted up as text/plain

- GET /simple/ping/
Returns "pong" as text/plain

- GET /simple/capital/{country}/
Returns the respective country capital as application/json

- GET /simple/capital/
Returns the list of all available countries
