package main

import (
	"database/sql"
	"fmt"
	ray "github.com/gen2brain/raylib-go/raylib"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	textbox "raytodo"
)

type Note struct {
	Content string
	ID      int
}

var db *sql.DB

func initDB() {
	if _, err := os.Stat("notes.db"); os.IsNotExist(err) {
		file, err := os.Create("notes.db")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		fmt.Println("notes.db file created")
	}
	var err error
	db, err = sql.Open("sqlite3", "./notes.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS notes (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"content" TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	fmt.Println("Connected to sql database and ensured notes table exists")
}
func insertNote(stuff string) {
	insertNoteSQL := `INSERT INTO notes (content) VALUES (?)`
	_, err := db.Exec(insertNoteSQL, stuff)
	if err != nil {
		fmt.Println("Error inserting note: ", err)
	}
}
func removeNote(id int) {
	removeNoteSQL := `DELETE FROM notes WHERE id = ?`
	_, err := db.Exec(removeNoteSQL, id)
	if err != nil {
		fmt.Println("Error deleting note: ", err)
	}
}
func fetchNotes() []Note {
	rows, err := db.Query("SELECT id, content FROM notes")
	if err != nil {
		fmt.Println("Error Fetching note: ", err)
		return nil
	}
	defer rows.Close()
	var notes []Note
	for rows.Next() {
		var note Note
		err = rows.Scan(&note.ID, &note.Content)
		if err != nil {
			fmt.Println("Error scanning row : ", err)
			continue
		}
		notes = append(notes, note)
	}
	return notes

}

func main() {
	initDB()
	defer db.Close()

	textbox1 := textbox.NewTextbox(100, 100, 600, 50, 50)
	notes := fetchNotes()

	ray.InitWindow(840, 600, "Quick-Task")
	defer ray.CloseWindow()
	ray.SetTargetFPS(60)
	for !ray.WindowShouldClose() {
		ray.BeginDrawing()
		ray.ClearBackground(ray.Black)
		textbox1.HandleInput()
		textbox1.Render()
		var but ray.Rectangle
		but = ray.NewRectangle(725, 100, 80, 45)
		ray.DrawRectangleRec(but, ray.LightGray)
		ray.DrawText("Add", 745, 110, 20, ray.Black)
		if ray.IsMouseButtonPressed(ray.MouseLeftButton) {
			if ray.CheckCollisionPointRec(ray.GetMousePosition(), but) {
				if textbox1.Text != "" && textbox1.Text != " " {
					insertNote(textbox1.Text)
          var note Note
          note.Content = textbox1.Text
					notes = append(notes, note)
					textbox1.Text = ""
				}
			}
		}
		ray.DrawText("Quick Task List", 100, 50, 20, ray.LightGray)
		if textbox1.Text != "" && textbox1.Text != " " {
			if ray.IsKeyPressed(ray.KeyEnter) {
				if textbox1.Text != " " {
					insertNote(textbox1.Text)
          var note Note
          note.Content = textbox1.Text
					notes = append(notes, note)
					textbox1.Text = ""
				}
			}
		}
		yPosition := 200
		for _, note := range notes {
			ray.DrawText(note.Content, 100, int32(yPosition), 20, ray.Green)
			rbut := ray.NewRectangle(600, float32(yPosition), 100, 40)
			ray.DrawRectangleRec(rbut, ray.LightGray)
			ray.DrawText("Remove", 610, int32(yPosition), 20, ray.Black)
			if ray.IsMouseButtonPressed(ray.MouseLeftButton) {
				if ray.CheckCollisionPointRec(ray.GetMousePosition(), rbut) {
					removeNote(note.ID)
					notes = fetchNotes()
					textbox1.Text = ""
				}
			}
			yPosition += 50
		}
		ray.EndDrawing()
	}

}
