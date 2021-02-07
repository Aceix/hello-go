/*
 * Note that the global scope exectuctes
 * fist, then init() is called, before main()
 *
 */

package main

import "fmt"

func answerToLife() int {
	return 42
}

func init() {
	whatIsThe = 0
}

func main() {
	if whatIsThe == 0 {
		fmt.Println("It's all a lie.")
	}
}

var whatIsThe = answerToLife()
