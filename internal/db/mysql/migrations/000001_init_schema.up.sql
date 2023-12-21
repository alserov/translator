CREATE TABLE IF NOT EXISTS users
(
    Uuid             text NOT NULL,
    Username         text NOT NULL,
    Password         text NOT NULL,
    Email            text NOT NULL,
    Requests_Balance int CHECK ( Requests_Balance >= 0 ),
    Created_At       timestamp
)