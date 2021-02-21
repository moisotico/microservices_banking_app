## Write an API to create a new account for an existing customer

### Acceptance criteria

- [x] A new account ca only be opened with a min deposit of $5000.00 
- [x] Account can only be of saving or checking type. 
- [x] In case of an unexpected error, API should return status code 500 (Internal server error) along with error message.
- [x] The API should return the new account ID, when the new account is opened with the status code as 201 (CREATED).