package main
import("fmt"
	   "bufio"
		"strconv"
		"os"
	)
type Person struct {
	lastName string
	firstName string
	iD int
}
func inPerson(p *Person , nextId int) (int, error){
	p.iD=nextId
	nextId++
	return nextId,nil
}
func printPerson(p Person){
	fmt.Println("First Name: "+p.firstName)
	fmt.Println("Last Name: "+p.lastName)
	fmt.Println("ID: "+strconv.Itoa(p.iD))
}
func main() {
	nextId := 101
	for {
		var (
			p Person
			err error
		)
		in:=bufio.NewReader(os.Stdin)
		fmt.Print("Enter First Name: ")
		firstName,_:=in.ReadString('\n')
		fmt.Print("Enter the Last Name: ")
		lastName,_:=in.ReadString('\n')
		p.firstName=firstName
		p.lastName=lastName
		nextId, err = inPerson(&p, nextId)
		if err != nil {
			fmt.Println("Invalid entry ... exiting")
			break
		}
		printPerson(p)
		}
	}
	