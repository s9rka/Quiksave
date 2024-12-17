## Authentication
- User login with username and password.
- Upon successful login user receives `access-token` (exp 15min) in response body and `refresh-token` (exp 1 day) in cookie.
- To renew `access-token`, a request is sent to `/refresh`.
- User is issued a new `access-token` if the `refresh-token` in cookie is valid.

