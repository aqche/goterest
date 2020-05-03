# goterest

A simplified Pinterest clone built with Go.

## Features

- Secure registration and authentication.
- Create and delete pins.
- View all pins.
- View a specific user's pins.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development. See deployment for notes on how to deploy the project.

### Prerequisites

To run this application you need [Go](https://golang.org/) and [PostgreSQL](https://www.postgresql.org/).

### Local Setup

Clone the project.

```
git clone https://github.com/aqche/goterest.git
```

Setup a PostgreSQL database.

```
sudo -u postgres createdb goterest
sudo -u postgres psql -d goterest -f setup.sql
```

Create a role for working with the new database.

```
sudo -u postgres psql -d goterest
postgres=#CREATE USER <user> WITH ENCRYPTED PASSWORD '<password>';
postgres=#GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO <user>;
postgres=#GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO <user>;
```

Start up the Go application.

```
go run ./cmd/web/
```

## Testing

TBD

## Deployment

TBD

## Built With

- [Gorilla Web Toolkit](https://www.gorillatoolkit.org/) - Go Libraries for Routing, Session Management, and CSRF Protection
- [Bulma](https://bulma.io/) - CSS Framework

## Contributing

Feel free to submit a pull request!

## Authors

- **aqche** - _Author_ - [aqche](https://github.com/aqche)

See also the list of [contributors](https://github.com/aqche/goterest/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for more details.

## Acknowledgments

- [Pinterest](https://www.pinterest.com/) - The inspiration for this site.
- [LogoHub](https://logohub.io/) - For the neat logo.
- [favicon.io](https://favicon.io/) - For the matching favicon.
