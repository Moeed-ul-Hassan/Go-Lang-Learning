package main

import (
	"fmt"
	"time"
)

//________        __________                     _ __
// /  _____/  ____  \______   \ ____  __ __  ____ |__|  |__   ____
// /   \  ___ /  _ \  |       _//  _ \|  |  \/    \|  |  |  \_/ __ \
// \    \_\  (  <_> ) |    |   (  <_> )  |  /   |  \  |  |  /\  ___/
//  \______  /\____/  |____|_  /\____/|____/|___|  /__|__|__/  \___  >
//         \/                \/                  \/                \/
//Every process which is running on the system is called a thread.
//Every process which is running concurrently is called a Routine. (GO ROUTINES)
//When GO keyword i added before a func call, it means that that func will run in background. (GO ROUTINE)
//Main Functions is also a GoRoutine. So when main is terminated, all routines are terminated.
//Lightweight.
//Provides Concurrency, means doing multiple things at the same time.

/* Real-Life Benefits (Where Does It Come In?)
1. Fast Response (API Response Time)
Suppose your app has a "Register User" button. When the user clicks it, they:

Make an entry in the database. (2 seconds)

Send a welcome email. (3 seconds)

Send a notification to the administrator. (1 second)

If you don't use Go, the user will see "Loading..." for 6 seconds. But if you put the email and notification in Go Routines, the user will see "Success" as soon as the entry is made to the database, and the rest of the work will continue in the background. The user will think the app is very fast!
*/

//Syntax to create a go routine
//go function_name()
//Routines are called by Go.

func firstprogram() {
	go say("Hello")
}

func say(word string) {
	fmt.Println(word)
}

//Channels in GOLang
//Channels are used to pass data between routines.
//Basically, Channels are like trasnmission lines between routines. That Tarnsfer Data from one routine to another.

// Syntax to create a channel
// channel_name := make(chan int)
func main() {
	//Sending Data to Channel
	channel_One := make(chan int, 1)
	fmt.Println("Channel Type :", channel_One)
	fmt.Println("Channel Value: ", channel_One)
	channel_One <- 1
	fmt.Println("Channel Value After updating: ", channel_One)
	//We have to close the channel after using it.
	close(channel_One)
	//If we dont do that it will cause a deadlock.
	//Deadlock is when a routine is waiting for data from another routine.
}

//If you started 100 background tasks and need to return their data (for example, price data from 100 websites), how do  you know when all the tasks are finished?

// Go uses two things for this:

// WaitGroups: This tells you to "Wait, don't close the program until all the helpers have finished."

// Channels: This is the pipe through which the background helper can tell you, "Hey, here's my result!"

//To understand Go Routines, always remember: "Fire and Forget." Initially, think of it as if you fired a shot and forgot. It will do its job, and you do yours. As long as you don't have to ask it for anything back, this "fire and forget" formula is best.

// Ye hamara function hai jo background mein chalega
func backgroundKaam(naam string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("👷 %s: Kaam ho raha hai... (Step %d)\n", naam, i)
		// We will add a delay to make sure output shows.
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// 1. First Go Routine
	go backgroundKaam("Helper-1")

	// 2. Second Go Routine
	go backgroundKaam("Helper-2")

	// 3. Program main func started.
	fmt.Println("🚀 Main Program: Maine dono helpers ko kaam pe laga diya hai!")
	fmt.Println("🚀 Main Program: Main ab apna baqi kaam kar raha hoon...")

	// ⚠️ ZARURI: Agar main ye Sleep na lagaoon, toh Main Program
	// foran khatam ho jayega aur Helpers ka kaam adhoora reh jayega.
	// (Kyunke Main 'Abbu' hai, wo ghar band karke chala jayega).
	time.Sleep(2 * time.Second)

	fmt.Println("🏁 Main Program: Saara kaam khatam, ab main ja raha hoon.")
}
