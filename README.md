# unicorn-factory

## Project structure

```bash
/unicorn-factory
    /cmd
        /server     # Main application entry point
    /pkg            # Reusable packages (if any)
    /internal       # Application-specific packages
    /api            # API handlers and route definitions
    /config         # Configuration files
    /db             # Database-related code
    /middleware     # Middleware functions
    /model          # Data models
    /router         # Router setup and middleware registration
    /util           # Utility functions
```

## To request unicorn production

```bash
POST http://localhost:8081/v1/unicorn
```

Body:

```json
{
  "amount": 100
}
```

The reponse will be a Id that can be used to get the unicorn production is going

## To know how the production is going

```bash
GET http://localhost:8081/v1/unicorn/id
```

```json
[
  {
    "Name": "dim-javier",
    "Capabilities": ["lazy", "cry", "code"],
    "ProducedAt": "2023-09-01T15:03:56.023361Z"
  },
  {
    "Name": "responsible-selma",
    "Capabilities": ["fullfill wishes", "change color", "swim"],
    "ProducedAt": "2023-09-01T15:03:56.988422Z"
  }
]
```

if passing the header "Response-Type = minimal" you wil get the amount of the unicorn produced until now, so you will know when the prodcution is done when you get the same amount sent in the request.
