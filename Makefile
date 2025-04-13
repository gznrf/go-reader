m_up:
	migrate -path migrations -database "postgresql://postgres:123456@localhost:5432/database_name?sslmode=disable" -verbose up
m_down:
	migrate -path migration -database "postgresql://postgres:123456@localhost:5432/database_name?sslmode=disable" -verbose down

