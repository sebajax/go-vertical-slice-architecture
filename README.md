## go-architecture-angrycoders

#### Seba Ituarte & Luis Fernando Miranda - This structure will help in the project building using golang and fiber to isolate the dependencies and have a cleaner code

![alt text](./go_vertical-slicing-architecture.drawio.png)

### This app uses conventional commits

[Conventional commits url](https://www.conventionalcommits.org/en/v1.0.0/)

### structure

    cmd
        contains the main.go file that is our starting point to execute
    migrations
        contains all the database configuration for the api (if needed)
    internal
        contains all the api logic

### Docker usage

    Build server
        docker-compose -p go-architecture-angrycoders build

    Start server
        docker-compose up -d

    Stop server
        docker-compose down

### Standalone usage

    uvicorn app.main:app --reload

### Testing

    To run unit testing

    To run unit testing coverage

### Database migration script

    To run the script

### Environment variables

To modify/add configuration via environment variables, use the `.env` file, which contains basic app configuration.
