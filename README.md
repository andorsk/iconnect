Welcome to iConnect! 

###What is it? 
TODO

### To run (currently based on revel framework):

revel run myapp
http://localhost:9000/

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory
        proto/        Protobuf Definition

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites

``` 
    db/ 
       migrations/
```

```
    docs/
```