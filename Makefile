.PHONY: download generate-yml

download:
	wget -O ./raw_data/rooms.csv 'https://docs.google.com/spreadsheets/d/1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8/export?format=csv&id=1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8&gid=0'
	wget -O ./raw_data/categories.csv 'https://docs.google.com/spreadsheets/d/1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8/export?format=csv&id=1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8&gid=1839493108'
	wget -O ./raw_data/schedule.csv 'https://docs.google.com/spreadsheets/d/1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8/export?format=csv&id=1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8&gid=1041297478'
	wget -O ./raw_data/slots.csv 'https://docs.google.com/spreadsheets/d/1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8/export?format=csv&id=1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8&gid=433889962'
	wget -O ./raw_data/partners.csv 'https://docs.google.com/spreadsheets/d/1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8/export?format=csv&id=1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8&gid=1712248286'
	wget -O ./raw_data/formats.csv 'https://docs.google.com/spreadsheets/d/1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8/export?format=csv&id=1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8&gid=563335428'
	wget -O ./raw_data/sessions.csv 'https://docs.google.com/spreadsheets/d/1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8/export?format=csv&id=1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8&gid=1529309388'
	wget -O ./raw_data/speakers.csv 'https://docs.google.com/spreadsheets/d/1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8/export?format=csv&id=1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8&gid=1791874699'

generate-yml:
	go run ./scripts/categories/main.go
	go run ./scripts/partners/main.go
	go run ./scripts/speakers/main.go
	go run ./scripts/slots/main.go
	go run ./scripts/formats/main.go
	go run ./scripts/rooms/main.go
	go run ./scripts/schedule/main.go
	go run ./scripts/sessions/main.go
