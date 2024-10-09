# User Management Service

This module is responsible for user management, including registration and authentication of user.

## Features

- User Registration: Allows new users to create accounts.
- Authentication: Verifies user credentials to grant access.

## Requirements

- Flask
- SQLAlchemy
- PostgreSQL
- Docker and Docker Compose (if used in a container)
   
## API

### User Registration

- Endpoint: /users/register
- Method: POST
- Request Body: {
    "username": "example"
  }
  
### Authentication

- Endpoint: /users/login
- Method: GET
- Request Query Param: ?username="some username"
  

  