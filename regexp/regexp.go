package main


    import "io/ioutil" 
import "regexp"
import "fmt"

func main() {

	re := regexp.MustCompile("(?sU)\n\t{.*\n\t}")


	f, _ := ioutil.ReadFile("../parser/go.y")

	fmt.Print(re.ReplaceAllLiteralString(string(f), "\n\t{\n\t}"))	
}
