# go-user-notes-app
User auth and taking notes app with Golang, Fiber framework and Postgres as database.


**Docker command to run postgresql container:**
```bash
$ docker run -it -d -p 5432:5432 --name notes-postgre -e POSTGRES_PASSWORD=notes-pass -d postgres:latest
```

**Screenshots of the app:**


* **Sign up page:**
![Screenshot 1](./screenshots/signup.png)

* **Login page:**
![Screenshot 2](./screenshots/login.png)