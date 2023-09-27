# go-hotel-app

This is a personal Golang Project that creates a virtual hotel web app.
This project implements global app configuration utilizing repository function pattern
This project uses postgres as the database

- Built in Go version 1.19.1
- Uses [chi router](github.com/go-chi/chi) for routing and middleware purposes
- Uses [alex edwards scs session management](github.com/alexedwards/scs) for managing sessions purposes
- Uses [nosurf](github.com/justinas/nosurf) as a CSRFToken manager
- Uses [soda](https://github.com/gobuffalo/soda) as database migration tool
