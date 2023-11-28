package main

func main() {
	/*
		Write a program that plays the number guessing game from the previous exercise.

		1. Launch a process to execute the number guessing game from your program. Use
		   exec.Cmd, see https://pkg.go.dev/os/exec
		   Usually you set up the command to run with
		       cmd := exec.Command("programPath", "firstArgumentIfNeeded")
		   and then run the command in the background, so you can interact with it in
		   your code
		       cmd.Start()

		   Be careful to check for an error in case the start fails. If so, abort and
		   print the original error message.

		2. Read the prompt from the game and learn the maximum value.
		   Use Cmd.StdoutPipe() to read the game's process output.

		3. In a loop, keep reading messages from the game and send number guesses until
		   you receive the success message.
		   Use Cmd.StdinPipe() to send messages to the game's process.

		4. Print each guessed number to the terminal and also the number of guesses
		   it took to win.
	*/
}
