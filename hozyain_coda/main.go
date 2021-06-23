//package main

//import subprocess
//
//dirPath = "/Users/den/Desktop/Git/integration-tests/"
//cmd = ["git", "-C", f"{dirPath}", "log", "--pretty=format:==%an", "--numstat"]
//
//all_code = subprocess.run(cmd, capture_output=True)
//all_code_strip = all_code.stdout.decode("utf-8").rstrip()
//
//name = None
//stats = {}
//aliases = {"Denis Mitin": ["denis mitin", "denis", "Денис", "Denis", "Mitin Denis"],
//"Hayk Kirakosyan": ["Haik Kirakosyan", "Hayk Kirakosyan"],
//"Eduard Makedonskii": ["Eduard Makedonskii", "emakedonskii", "e.makedonskii"]}
//
//for line in all_code_strip.split("\n"):
//if line.startswith("=="):
//name = line[2:]
//for k, v in aliases.items():
//if name in v:
//name = k
//
//if name not in stats:
//stats[name] = 0
//
//elif len(line) > 0 and line[0].isdigit():
//elements = line.split()
//stats[name] += int(elements[0]) + int(elements[1])
//
//score = {k: v for k, v in reversed(sorted(stats.items(), key=lambda item: item[1]))}
//
//for k, v in score.items():
//print(f"{k} = {v}")
//

//import (
//	"flag"
//	"fmt"
//	"github.com/spf13/pflag"
//)

//func main() {
//
//	dirPath := "/Users/den/Desktop/Git/integration-tests/"
//	//cmd := [6]string{"git", "-C", "{dirPath}", "log", "--pretty=format:==%an", "--numstat"}
//	pflag.Args("git", "-C", "dirPath", "log", "--pretty=format:==%an", "--numstat")
//
//
//	git :=flag.Args("git")
//
//
//	//pflag.BoolSlice("git", "-C", "dirPath", "log", "--pretty=format:==%an", "--numstat")
//	//pflag.lag.Int("flagname", 1234, "help message for flagname")
//
//	fmt.Println(dirPath)
//	//fmt.Println(cmd)

//import (
//"flag"
//"github.com/spf13/pflag"
//	"github.com/spf13/viper"
//)
//
//func main() {
//
//	// using standard library "flag" package
//	flag.Int("flagname", 1234, "help message for flagname")
//
//	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
//	pflag.Parse()
//	viper.BindPFlags(pflag.CommandLine)
//
//	i := viper.GetInt("flagname") // retrieve value from viper
//
//}

//package main
//
//import "fmt"
//
//func main() {
//	//a := fmt.Scan(&a)
//	var a string
//	_, err := fmt.Scan(&a)
//	if err != nil {
//		return
//	}
//	fmt.Println(a)
//
//}

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	//fmt.Printf("%v\n", p)

	//fmt.Printf("%+v\n", p)

	//fmt.Printf("%#v\n", p)

	//fmt.Printf("%T\n", p)

	//fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14)

	fmt.Printf("%c\n", 33)

	fmt.Printf("%x\n", 456)

	fmt.Printf("%f\n", 78.9)

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")

	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
