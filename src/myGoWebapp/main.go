package main

import (
   "net/http"
   "fmt"
   "time"
   "html/template"
   _ "github.com/go-sql-driver/mysql"
        "database/sql"

)
var uname = "root"   //Give the MySQL username
var pwd = "Siddhant.10" //Give the MySQL password
var url ="localhost:3306"  //Give the database URL
//Structure to store details of all users
type User struct {

  Uid, Username, Departname,Created string

}

//Structure to get Name from URL
type Welcome struct {
   Name string
   Time string
}

//Go application entrypoint
func main() {

welcome := Welcome{"Guest, to Web application in GO!", time.Now().Format(time.Stamp)}

http.HandleFunc("/create", createUser) //Function to handle create user request
http.HandleFunc("/update", updateUser) //Function to handle update user request
http.HandleFunc("/delete", deleteUser) //Function to handle delete user request
http.HandleFunc("/read", readUser) //Function to handle display user request

  // Provide which template to exxecute
  templates := template.Must(template.ParseFiles("template/welcome-template.html"))


  //All the static templates must reside in package /static in root directory
	http.Handle("/static/",
		http.StripPrefix("/static/",
		   http.FileServer(http.Dir("static"))))

      //Function to handle UI on certain URL
		  http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {
			// Takes name from URL and displays on UI
			if name := r.FormValue("name"); name != "" {
			   welcome.Name = name + ", to Web application in GO!";
			}else{
         name := "Guest"
         welcome.Name = name + ", to Web application in GO!";
      }

			if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			   http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		 })

		 fmt.Println("Listening");
		 fmt.Println(http.ListenAndServe(":8080", nil));
	  }

    func checkErr(err error) {
          if err != nil {
              panic(err)
          }
      }

    //Function to create user in database
    func createUser(w http.ResponseWriter, r *http.Request) {

      if r.Method == "GET" {
          t, _ := template.ParseFiles("login.gtpl")
          t.Execute(w, nil)
      } else {

          //Parse the form on UI to pick up entered values
          r.ParseForm()
          username :=  r.FormValue("username")
          department :=  r.FormValue("department")
          datecreated :=  r.FormValue("datecreated")

          upass := uname+":"+pwd+"@tcp("+url+")/godb?charset=utf8"
          db, err := sql.Open("mysql",upass)
          checkErr(err)

            //Prepare databse query
            stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
            checkErr(err)

            //Execute query
            stmt.Exec(username, department,datecreated)
            checkErr(err)

            if err==nil{
               http.Redirect(w, r, "/template/createSuccess.html", 301)
            }

      }
  }


  //Function to update user in database
  func updateUser(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        //Parse the form on UI to pick up entered values
        r.ParseForm()
        username :=  r.FormValue("username")
        department :=  r.FormValue("department")

        upass := uname+":"+pwd+"@tcp("+url+")/godb?charset=utf8"
        db, err := sql.Open("mysql",upass)
        checkErr(err)

          //Prepare databse query
          stmt, err := db.Prepare("UPDATE userinfo SET departname=? WHERE username=?")
          checkErr(err)

          //Execute query
          stmt.Exec(department,username)
          checkErr(err)

          if err==nil{
             http.Redirect(w, r, "/template/createSuccess.html", 301)
          }

    }
  }


  //Function to delete user in database
  func deleteUser(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
      //Parse the form on UI to pick up entered values
        r.ParseForm()
        username :=  r.FormValue("username")

        upass := uname+":"+pwd+"@tcp("+url+")/godb?charset=utf8"
        db, err := sql.Open("mysql",upass)
        checkErr(err)

          // insert
          stmt, err := db.Prepare("DELETE FROM userinfo WHERE username=?")
          checkErr(err)

          //Execute query
          stmt.Exec(username)
          checkErr(err)

          if err==nil{
             http.Redirect(w, r, "/template/createSuccess.html", 301)
          }
      }
  }

  //Function to read all users in database
  func readUser(w http.ResponseWriter, r *http.Request) {
    // fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        //Parse the form on UI to pick up entered values
        r.ParseForm()

        //Get database connection
        upass := uname+":"+pwd+"@tcp("+url+")/godb?charset=utf8"
        db, err := sql.Open("mysql",upass)
        checkErr(err)


          // Make and execute database query
          rows, err := db.Query("SELECT * FROM userinfo")
          checkErr(err)

            //Create a Structure object to store parsed information
            user := User{}
            //Create an array of Structure objects to store parsed information
            users := []User{}

            //Parse each row from database and store in a structure object stored in an array
            for rows.Next() {
              var Uid, Username, Departname, Created string
              err = rows.Scan(&Uid, &Username, &Departname, &Created)
              checkErr(err)
              user.Uid = Uid
              user.Username = Username
              user.Departname = Departname
              user.Created = Created
              users = append(users, user)
            }

        //Assign array of structure objects to a variable to be passed to frontend
        table := users

        //Provide template to be rendered for displaying the results
        var tmpl = template.Must(template.ParseFiles("template/layout.html"))
        //Render the template if no error
        err1 :=tmpl.ExecuteTemplate(w,"Index",table)
        if err1 != nil {
            panic(err)
        }

    }
  }
