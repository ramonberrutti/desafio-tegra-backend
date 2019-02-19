# desafio-tegra-backend

## TODO

- ~~Webserver~~
- Tests
- ~~Dockerfile~~
- ~~Remove logs~~
- Add Cache
- Add Database (dGraph?)

## Instalation
You need docker to run the test or have Go Lang installed on your machine.

- Clone the project:
`git clone git@github.com:ramonberrutti/desafio-tegra-backend.git`

- Compile the Dockerfile:
`docker build -t ramonberrutti/desafio-tegra-backend .`

- Run the Container:
```docker run -p 8080:8080 --rm -v `pwd`/data:/app/data ramonberrutti/desafio-tegra-backend```

If you have go installed in your machine you can run with:
`go run *.go`

## Endpoints

- Flights:
  * from: Airport from where the flight begin
  * to: Airport of destination
  * date: date of the flight

  Example: `http://localhost:8080/flights?from=VCP&to=BEL&date=2019-02-16`

- Airports:
  Example: `http://localhost:8080/airports`