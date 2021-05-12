package main

import (
    "io"
    "fmt"
    "log"
    "time"
    "path"
    "net/http"
    "database/sql"
    "html/template"
    "encoding/json"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

type compromisedDetail struct {
    UserId      int    `json:"userid"`
    FirstName   string `json:"firstname"`
    LastName    string `json:"lastname"`
    Email       string `json:"email"`
}

const PORT = "9090"
var mysqlDbConn *sql.DB;

func MysqlDbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser   := "root"
    dbPass   := "Cloud#9"
    dbName   := "db1"

    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil { panic(err.Error()) }
    return db
}

func getPeopleJson(userid string) (string) {
    query := "SELECT * FROM compromised "
    if ( userid != "" ) {
        query += "WHERE userid = "+userid
    }

    cds := []compromisedDetail{}
    rows, err := mysqlDbConn.Query(query)
    defer rows.Close()
    if err != nil { panic(err.Error()) }

    for rows.Next() {
        cd := compromisedDetail{}
        err := rows.Scan(
            &cd.UserId,
            &cd.FirstName,
            &cd.LastName,
            &cd.Email,
        )
        if err != nil { panic(err.Error()) }

        cds = append(cds, cd)
    }

    outputBytes, err := json.Marshal(cds)
    return string(outputBytes)
}

func peopleHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    userid := vars["userid"]

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    io.WriteString(w, getPeopleJson(userid))
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
    filepath := path.Join("static", "index.html")
    tmpl, err := template.ParseFiles(filepath)
    if err != nil { panic(err.Error()) }

    var data = map[string]interface{}{
        "cds" : getPeopleJson(""),
    }
    err = tmpl.Execute(w, data)
    if err != nil { panic(err.Error()) }
}

func main() {
    // connect to mysql db
    mysqlDbConn = MysqlDbConn()
    defer mysqlDbConn.Close()

    // define router & service
    r := mux.NewRouter()
    // -- frontend
    r.HandleFunc("/", homeHandler)
    // -- backend
    r.HandleFunc("/people", peopleHandler)
    r.HandleFunc("/people/{userid}", peopleHandler)

    // run server
    server := &http.Server{
        Handler         : r,
        Addr            : ":"+PORT,
        WriteTimeout    : 300 * time.Second,
        ReadTimeout     : 300 * time.Second,
    }
    fmt.Println()
    log.Println("Running Server on localhost:"+PORT)
    fmt.Print("\n\n\n")
    log.Fatal(server.ListenAndServe())
}


