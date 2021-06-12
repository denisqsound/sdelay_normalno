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

package main

import "fmt"

func main() {

	dirPath := "/Users/den/Desktop/Git/integration-tests/"
	//cmd := [6]string{"git", "-C", "{dirPath}", "log", "--pretty=format:==%an", "--numstat"}

	fmt.Println(dirPath)
	//fmt.Println(cmd)
}
