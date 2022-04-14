package main 

import (
	"fmt"
	"os"
	"os/user"
	//"log"
	"github.com/hajimehoshi/oto"
	"github.com/hajimehoshi/go-mp3"
	"github.com/mbndr/figlet4go"
	"io"	
	// "strings"
)

//var mmtop string
var ascii = figlet4go.NewAsciiRender()
var renderStr, _ = ascii.Render("Hi " + usr.Username + " !")
var usr, err = user.Current()
var colorCyan = "\033[36m"

func main() {
	if len(os.Args) != 2 {
		fmt.Print(string(colorCyan), renderStr)
		fmt.Println("Enter path or file's name of .mp3")
		fmt.Println("Example: switchp <path> or <name>")
		return
	}

	arg := os.Args[1]
 
	// ascii := figlet4go.NewAsciiRender()
 
	// colorCyan := "\033[36m"
 
	// usr, err := user.Current()
	// if err != nil {
		// log.Fatal(err)
	// }

	//usr.HomeDir
	// renderStr, _ := ascii.Render("Hi " + usr.Username + " !")
	//fmt.Println("Hi " + usr.Username + " !")
	fmt.Print(string(colorCyan), renderStr)

	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	} 
		// fmt.Println("\nФайл был создан :)")
	defer f.Close()

	if arg == "help" {
		fmt.Println("\nИспользование: switchp <путь к файлу/название файла>")
	}  else {
	f, err = os.Open(arg + ".mp3")
	fmt.Println("I'm playing '" + arg + "' for you :)")
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		panic(err)
	}
	
	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
	
	}
}
	
	//os.Remove("test.txt")

	// var t string
	// t = "Glory to Russia!"

	// for i := 0; i < 9999; i++ {
		// f.WriteString(t)
	// }
	// fmt.Println("I wrote ->\t" + t)

	// fmt.Println(string(colorCyan), "\nGlory to Russia!")


	//fmt.Println("Enter the name of mp3: ")
	//fmt.Scanln(&arg)

	// if err != nil {
		// }
	// if err := run(); err != nil {
		// log.Fatal(err)
	// }	
}


