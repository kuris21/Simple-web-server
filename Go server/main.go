// This project creates a go Server with three routes, route, / hello and / forum. 
// The root opens index. html
// the /hello calls the function hello func
// The /forum calls forum func which will open form.html

// defines the package name for the go file 
package main

// import imports the three packages below 
import(
	"fmt"	// For formatting and printing output
	"log"	// For logging messages (helps in debugging and logging server optional errors)
	"net/http"	// Provides HTTP client and server implimentations
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err; %v", )
	}
	fmt.Fprintf(w,"POST request sucessful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n ", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
if r.URL.Path != "/hello"{
	http.Error(w, "404 not found", http.StatusNotFound)
	return
}
if r.Method != "GET"{
	http.Error(w, "method is not supported", http.StatusNotFound)
	return
}
fmt.Fprintf(w, "hello")
}

// defines main function which is the entry point for any executable go program
func main(){
	fileServer := http.FileServer(http.Dir("./static"))	// sets a file server to serve static files
	http.Handle("/", fileServer)	// registers file server as the handler for all requests to the root url
	http.HandleFunc("/form", formHandler)	// resisters the function form handler to handle Http requests at the /form url path
	http.HandleFunc("/hello", helloHandler)	// resisters the function hello handler to handle Http requests at the /hello url path

	fmt.Print("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080",nil); err !=nil {
		log.Fatal(err)
	}
}