register:
	curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Dachi","email":"dachimari9@gmail.com","password":"secret123","age":25}'
token:
	curl "http://localhost:8080/api/v1/verify?token=PASTE_TOKEN_FROM_EMAIL_HERE"
login:
	curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"email":"dachimari9@gmail.com","password":"secret123"}'
users:
	curl http://localhost:8080/api/v1/users \
	-H "Authorization: Bearer YOUR_TOKEN"	
singleusers:
	curl http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
createusers:
	curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{"name":"Bob","email":"bob@example.com","age":25}'