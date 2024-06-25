package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML content with embedded JavaScript
	fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html>
        <head>
            <title>adityanarayan.dev</title>
            <script>
                function getUserInfo() {
                    fetch('https://ipapi.co/json/')
                    .then(response => response.json())
                    .then(data => {
                        var city = data.city;
                        if (!city) {
                            city = "your city";
                        }
                        document.getElementById('message').innerHTML = "Hey, are you from " + city + "? Let's catch up! Ping me on <a href='https://x.com/svn_aditya' target='_blank'>Twitter</a>.";
                    })
                    .catch(error => console.log('Error:', error));
                }
            </script>
        </head>
        <body onload="getUserInfo()">
            <h1 id="message">Loading...</h1>
			<h2><a href='https://youtube.com/@aditya_narayana?si=C75E8vc6C6wp5JTw' target="_blank">Visit My YouTube Channel</a></h2>
		    <h2><a href='https://www.linkedin.com/in/aditya-narayana' target="_blank">Visit My Linkedin</a></h2>
        </body>
        </html>
    `)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
