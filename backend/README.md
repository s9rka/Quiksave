## Authentication
- Users log in using their username and password.
- Upon successful login, users receive an `access_token` (expires in 15 minutes) and a `refresh_token` (expires in 1 day), both stored in cookies.
- To renew the `access_token`, a request is sent to the `/refresh` endpoint.
- If the `refresh_token` in the cookie is valid, a new `access_token` is issued.
- To authenticate the user and retrieve user details, call the `/me` endpoint.