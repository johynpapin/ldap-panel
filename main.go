package main

import (
	"net/http"
	"os"
	"text/template"

	"github.com/go-ldap/ldap/v3"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type SignInForm struct {
	Username string
	Password string
}

type ChangePasswordForm struct {
	OldPassword          string
	Password             string
	PasswordVerification string
}

func main() {
	// Templates

	signInViewTemplate, err := template.ParseFiles("templates/layouts/main.html", "templates/views/sign-in.html")
	if err != nil {
		panic(err)
	}

	dashboardViewTemplate, err := template.ParseFiles("templates/layouts/main.html", "templates/views/dashboard.html")
	if err != nil {
		panic(err)
	}

	// LDAP

	l, err := ldap.DialURL(os.Getenv("LDAP_PANEL_LDAP_URL"))
	if err != nil {
		panic(err)
	}

	err = l.Bind(os.Getenv("LDAP_PANEL_LDAP_USERNAME"), os.Getenv("LDAP_PANEL_LDAP_PASSWORD"))
	if err != nil {
		panic(err)
	}

	// Sessions

	store := sessions.NewCookieStore([]byte(os.Getenv("LDAP_PANEL_SESSION_KEY")))

	// Routes

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		username, ok := session.Values["username"].(string)
		if !ok || username == "" {
			signInViewTemplate.Execute(w, nil)
			return
		}

		dashboardViewTemplate.Execute(w, nil)
	}).Methods("GET")

	router.HandleFunc("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		signInForm := SignInForm{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		// TODO: authenticate user
		_ = signInForm

		session.Values["username"] = signInForm.Username
		session.Save(r, w)

		// Redirect to home
		http.Redirect(w, r, "/", http.StatusFound)
	}).Methods("POST")

	router.HandleFunc("/log-out", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		delete(session.Values, "username")
		session.Save(r, w)

		// Redirect to home
		http.Redirect(w, r, "/", http.StatusFound)
	}).Methods("GET")

	router.HandleFunc("/change-password", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		changePasswordForm := ChangePasswordForm{
			OldPassword:          r.FormValue("old-password"),
			Password:             r.FormValue("password"),
			PasswordVerification: r.FormValue("password-verification"),
		}

		username, _ := session.Values["username"].(string)

		// TODO: Use the real username
		passwordModifyRequest := ldap.NewPasswordModifyRequest(username, changePasswordForm.OldPassword, changePasswordForm.Password)
		if _, err = l.PasswordModify(passwordModifyRequest); err != nil {
			// TODO: Handle the error
		}

		// Redirect to home
		http.Redirect(w, r, "/", http.StatusFound)
	}).Methods("POST")

	http.ListenAndServe(":8080", router)

	// TODO: Close the LDAP connection
}
