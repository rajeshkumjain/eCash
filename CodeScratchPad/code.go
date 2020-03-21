package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	ConstUserListTemplateFile = "pages/userlist.gohtml"
	ConstUserListingPage      = "userlist.gohtml"
	ConstUserListSQL          = `SELECT TOP(100) ru_id, ru_fname, ru_mname, ru_sname, ru_email_primary, ru_dob 
	FROM  registered_user 
	WHERE ru_cf_active_flag=1`
	ConstUserDetailsTemplateFile = "pages/userdetails.gohtml"
	ConstUserInsertTemplateFile  = "pages/userInsert.gohtml"
	ConstUserDetailsPage         = "userDetails.gohtml"
)

// database connection needs to come utils.go now

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "rajeshjain32", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "RJDELLLAPTOP", "the database server")
	user          = flag.String("user", "sa", "the database user")
	database      = flag.String("database", "sample", "The name of the database")
)

//* RegisteredUser : Registered User Structure
type RegisteredUser struct {
	Id, C_id, R_id, Rkc_id, Uc_id, Third_party_id              int
	Fname, Mname, Sname, Email_primary, Password               string
	Enable_flag, I_agree_t_n_c, Wants_newsletter               bool
	EmailURL, RU3rdparty_source, DOB, Rights                   string
	Cf_active_flag                                             bool
	Cf_date_created_on, Cf_date_last_updated_on                string
	Cf_active_date, Cf_action_status                           string
	Cf_archive_status, Cf_created_by_id, Cf_last_updated_by_id int
	Source                                                     string
}

//var tmplte = template.Must(template.ParseGlob("template/user*.gohtml"))

// dbConn : initial database connection
func dbConn() (db *sql.DB) {
	flag.Parse()
	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
		fmt.Printf(" database:%s\n", *database)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", *server, *user, *password, *database)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	condb, errdb := sql.Open("sqlserver", connString)
	if errdb != nil {
		log.Fatal("Open connection failed:", errdb.Error())
	}
	return condb
}

// end of dbConn

func Index(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl, err := template.ParseFiles("pages/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.ExecuteTemplate(w, "Index", nil)

}

// UserListing: List of Registered Users in the application
func UserListing(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl, err := template.ParseFiles(ConstUserListTemplateFile)
	if err != nil {
		log.Fatalln(err)
	}

	db := dbConn()
	qry, err := db.Query(ConstUserListSQL)
	if err != nil {
		log.Fatal("QUERY failed: ", err.Error())
	}
	regUser := RegisteredUser{}
	res := []RegisteredUser{}

	for qry.Next() {
		var id int
		var fname, mname, sname, email, dob string
		err = qry.Scan(&id, &fname, &mname, &sname, &email, &dob)
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
		}
		regUser.Id = id
		regUser.Fname = fname
		regUser.Mname = mname
		regUser.Sname = sname
		regUser.Email_primary = email
		regUser.DOB = dob

		res = append(res, regUser)
	}

	defer qry.Close()
	tmpl.ExecuteTemplate(w, ConstUserListingPage, res)
}

// ShowUserDetails: Registered User Details

func ShowUserDetails(w http.ResponseWriter, r *http.Request) {
	// In future comment bellow one and Use ParseGlob as problem top
	var tmpl *template.Template
	tmpl, err := template.ParseFiles(ConstUserDetailsTemplateFile)

	if err != nil {
		log.Fatalln(err)
	}

	nId := r.URL.Query().Get("id")
	/* Fields of interest are [ru_fname], [ru_mname], [ru_sname], [ru_email_primary],[ru_enable_flag],
	   [ru_i_agree_t_n_c], [ru_dob], [ru_rights], [c_id], [wants_newsletter] */

	selectSql := "SELECT ru_id, ru_fname, ru_mname, ru_sname, ru_email_primary, ru_dob, ru_enable_flag, ru_i_agree_t_n_c, ru_rights, c_id, wants_newsletter FROM registered_user WHERE ru_id =  " + nId
	fmt.Println(selectSql)
	db := dbConn()
	qry, err := db.Query(selectSql)
	/* see this in future - also refer to document / code at the of this file
	columnNames, err := qry.Columns()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(columnNames)
	*/
	if err != nil {
		log.Fatalln(err)

	}

	regUser := RegisteredUser{}
	for qry.Next() {

		var id int
		var c_id sql.NullInt32 // sql package for null value validation
		var fname, mname, sname, email, dob, ru_rights string
		var ru_enable_flag, ru_i_agree_t_n_c, wants_newsletter bool

		err = qry.Scan(&id, &fname, &mname, &sname, &email, &dob, &ru_enable_flag, &ru_i_agree_t_n_c, &ru_rights, &c_id, &wants_newsletter)
		qry.Columns()
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
		}
		fmt.Println(id, fname, mname, sname, email, dob, ru_enable_flag, ru_i_agree_t_n_c, ru_rights, c_id, wants_newsletter)
		regUser.Id = id
		regUser.Fname = fname
		regUser.Mname = mname
		regUser.Sname = sname
		regUser.Email_primary = email
		regUser.DOB = dob
		regUser.Enable_flag = ru_enable_flag
		regUser.I_agree_t_n_c = ru_i_agree_t_n_c
		regUser.Wants_newsletter = wants_newsletter
		regUser.Rights = ru_rights
	}

	defer qry.Close()
	fmt.Println(regUser)
	// tmpl.ExecuteTemplate(w, "userdetails.gohtml", regUser)
	tmpl.ExecuteTemplate(w, "show", regUser) // ConstUserDetailsPage

}

func New(w http.ResponseWriter, r *http.Request) {
	// In future comment bellow one and Use ParseGlob as problem top
	var tmpl *template.Template
	tmpl, err := template.ParseFiles(ConstUserInsertTemplateFile)
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		fname := r.FormValue("Fname")
		mname := r.FormValue("Mname")
		sname := r.FormValue("Sname")
		dob := r.FormValue("DOB")
		email_primary := r.FormValue("Email_primary")
		rights := r.FormValue("Rights")
		enable_flag := 0
		i_agree_t_n_c := 0
		wants_newsletter := 0
		if r.FormValue("Enable_flag") == "on" {
			enable_flag = 1
		}
		if r.FormValue("I_agree_t_n_c") == "on" {
			i_agree_t_n_c = 1
		}
		if r.FormValue("Wants_newsletter") == "on" {
			wants_newsletter = 1
		}

		// include default setting here - like created_on, last_update_on, source, ... etc.
		/*
			[ru_cf_active_flag] [tinyint] NULL,
			[ru_cf_date_created_on] [varchar](50) NULL,
			[ru_cf_date_last_updated_on] [varchar](50) NULL,
			[ru_cf_active_date] [int] NULL,
			[ru_cf_action_status] [varchar](100) NULL,
			[ru_cf_archive_status] [tinyint] NULL,
			[ru_cf_created_by_id] [int] NULL,
			[ru_cf_last_updated_by_id] [int] NULL,
			[ru_source] [varchar](100) NULL,
		*/
		tsql := fmt.Sprintf(`INSERT INTO registered_user
		 (ru_fname, ru_mname, ru_sname, ru_email_primary, ru_dob, ru_enable_flag, ru_i_agree_t_n_c, ru_rights, wants_newsletter) 
		 VALUES ('%s', '%s', '%s', '%s', '%s', '%d', '%d', '%s',  '%d') ;`, fname, mname, sname, email_primary, dob, enable_flag, i_agree_t_n_c, rights, wants_newsletter)
		fmt.Println(tsql)
		_, err := db.Exec(tsql)
		if err != nil {
			fmt.Println("insForm Error : ")
			log.Fatalln(err.Error())
		}
	}

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

// main: main calling program
func main() {
	log.Println("Server started on : http://localhost:8080...")
	http.HandleFunc("/", Index)
	http.HandleFunc("/UserListing", UserListing)
	http.HandleFunc("/show", ShowUserDetails)
	http.HandleFunc("/new", New)
	http.HandleFunc("/insert", Insert)
	http.ListenAndServe(":8080", nil)
}
