package collector

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"time"
)

// Container solution
// The better solution for this is to update a mounted volume from inside the non-privileged
// namespaces container, and have another listening process run as root outside the container
// , to do the actual shutdown work at the host machine. This provides a secure interface between
// the container and the host machine.
func terminate(target string) {
	fmt.Printf("target: %s\n", target)
	if target == TargetServer || target == "localhost" {
		// run locally
		fmt.Println("Hey I am going to turn you off! Server.")
		fmt.Printf("runtime.GOOS: %v\n", runtime.GOOS)
		b, err := ioutil.ReadFile("/var/run/shutdown_signal") // linux

		if err != nil {
			log.Println("File doesn't exist, create it to write.")
		}
		fmt.Printf("shutdown_signal: %s\n", string(b))

		signal := []byte("true")
		err = ioutil.WriteFile("/var/run/shutdown_signal", signal, 0644)

		if err != nil {
			panic(err)
		}
		fmt.Printf("Successfully write signal to %s\n", string(signal))
	} else {
		// TODO: turn off remote server
		// cmd := exec.Command("ssh", "-t", "-t", "-p", "{{port}}", "{{hostip}}", "init 6")
		// out, err := cmd.CombinedOutput()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Printf("cmd Output:%v", string(out))
	}
}

// Binary solution
func shutdownCommand() {

	fmt.Println("The computer is going to be shutdown in 5 seconds .....")
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("Bye")
	// shutdown locally
	fmt.Println(runtime.GOOS)
	cmd := exec.Command("shutdown", "-h", "now") // linux/darwin
	// if runtime.GOOS == "windows" {
	// 	cmd = exec.Command("shutdown", "/s")
	// }

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("cmd Output:%v", string(out))
}
