-- name: CreateUser :one
INSERT INTO users(name, age,password)
VALUES ($1, $2,$3)
RETURNING id, name, age;

-- name: GetUser :one
SELECT id, name, age
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, age
FROM users;

-- name: GetUserByName :one
SELECT id,name,age,password
FROM users
WHERE name = $1; 
