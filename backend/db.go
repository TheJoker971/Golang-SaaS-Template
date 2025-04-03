import (
	"os"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	pseudo string
	email string
	password string
	admin bool
}

func ConnectDb() *sql.DB{
	if not != nil {
		log.Fatal("L'url de la database est inexistante !")
	}
	db, err := sql.Open("mysql","./saas.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS Users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		pseudo VARCHAR,
		email VARCHAR,
		password VARCHAR,
		admin BOOLEAN
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("❌ Erreur création de la table: %s", err)
	}

	fmt.Println("✅ Table prête.")

	return db
}

func Register(db *sql.DB,user User) {
	query := `
	INSERT INTO Users(
		pseudo VARCHAR,
		email VARCHAR,
		password VARCHAR,
		admin BOOLEAN
	)
	VALUES(?, ?, ?, ?);`
	_,err := db.Exec(query,user.pseudo,user.email,user.password,user.admin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ User insert in base")
}

