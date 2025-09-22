1. Libriaries: 
- github.com/gin-gonic/gin
- github.com/lib/pq
- github.com/dgrijalva/jwt-go
- golang.org/x/crypto/bcrypt

2. Server PostgreSQL
  Via terminal:
  2.1 Install PostgreSQL: brew install postgresql
  2.2 Check status PostgreSQL: brew services list
  2.3 If status false, launch PostgreSQL: brew services start postgresql
  2.4 Connect to PostgreSQL: psql -U postgres

3. Launch server: go run main.go
