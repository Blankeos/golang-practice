package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

func main() {
	// 1. Load env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load .env file: %s", err)
	}

	// 2. Connect to DB
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	fmt.Println("Connected to DB!")

	// 3. Query inside DB
	var title string
	var artist string

	err = conn.QueryRow(context.Background(), "SELECT title, artist FROM album WHERE id = $1", 2).Scan(&title, &artist)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(title, artist)

	// 4. Query another time (Using custom function Albums By Aritst)
	albums, err := albumsByArtist("Carlo Taleon", conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Albums by Artist failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Albums found: %v\n", albums)

	// 5. Query another time (Using custom function Album By ID)
	album, err := albumByID(1, conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Album by ID failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Abum found: %v\n", album)
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string, conn *pgx.Conn) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := conn.Query(context.Background(), "SELECT id, title, artist, price FROM album WHERE artist = $1", name)
	if (err != nil) {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist: %q: %v", name, err)
		}

		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64, conn *pgx.Conn) (Album, error) {
	
	var alb Album

	// Notice how I reverse the order of SELECT here     t👇  i👇
	row := conn.QueryRow(context.Background(), "SELECT title, id, artist, price FROM album WHERE id = $1", id)
	
	// 				      t👇       i👇 <- The rowscan is related to this order.
	if err := row.Scan(&alb.Title, &alb.ID, &alb.Artist, &alb.Price); err != nil {
		return alb, fmt.Errorf("albumByID: %q: %v", id, err)
	}

	return alb, nil
}