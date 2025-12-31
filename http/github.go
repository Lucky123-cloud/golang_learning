package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Given  github user login, return name and number of public repos

func main() {
	resp, err := http.Get("https://api.github.com/users/ardanlabs")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: bad status - %s\n", resp.Status)
		return
	}

	ctype := resp.Header.Get("Content-Type")
	fmt.Println("content-type:", ctype)

	// io.Copy(os.Stdout, resp.Body)

	var reply struct {
		Name         string
		Public_Repos int
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&reply); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(reply.Name, reply.Public_Repos)

}

/* JSON <-> GO
Types
string <-> string
true/false -> Bool
number -> flaot64, int,
array -> []T []any
object ->map[string]any, struct

encoding/json API
json -> []bytes -> Go: Unmarshal
Go -> []byte -> JSON: Marshal
JSON -> io.Reader -> Go: Decoder
Go -> io.Writer -> JSON: Encoder

*/
