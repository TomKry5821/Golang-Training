package main

type jsonGame struct {
	Id    int
	Name  string
	Genre string
	Price int
}

type item struct {
	id    int
	name  string
	price int
}
type game struct {
	item
	genre string
}

/*
const (
	WON_TEXTS  = 2
	LOSE_TEXTS = 2
)
*/
//const size = 1e7

func main() {
	// ---------------------------------------------------------
	// EXERCISE: Random Messages
	//
	//  Display a few different won and lost messages "randomly".
	//
	// HINTS
	//  1. You can use a switch statement to do that.
	//  2. Set its condition to the random number generator.
	//  3. I would use a short switch.
	//
	// EXAMPLES
	//  The Player wins: then randomly PrintGames one of these:
	//
	//  go run main.go 5
	//    YOU WON
	//  go run main.go 5
	//    YOU'RE AWESOME
	//
	//  The Player loses: then randomly PrintGames one of these:
	//
	//  go run main.go 5
	//    LOSER!
	//  go run main.go 5
	//    YOU LOST. TRY AGAIN?
	// ---------------------------------------------------------
	//wm := ['You won!', 'You are awesome!']
	//lm := ['You lost!', 'You lost. Try again?']
	/*isCorrect, isVerbose := checkCommandLineArgs()
	if !isCorrect {
		return
	}
	argsIndex := setArgsIndex(isVerbose)
	fn, fError, sn, sError := setPickedNumbersAndSetMaxRounds(argsIndex)
	if checkErrors(fError, sError) {
		return
	}
	maxRounds := setMaxRounds(fn, sn)
	drawNumbersAndShowResult(fn, sn, isVerbose, maxRounds)*/

	// ---------------------------------------------------------
	// EXERCISE: Case Insensitive Search
	//
	//  Allow for case-insensitive searching
	//
	// EXAMPLE
	//  Let's say that the user runs the program like this:
	//    go run main.go LAZY
	//
	//  Or like this: go run main.go lAzY
	//  Or like this: go run main.go lazy
	//
	//  For all cases above, the program should find
	//  the "lazy" keyword.
	// ---------------------------------------------------------
	/*const (
		phrase = "some words like that lazy srazy"
	)
	if len(os.Args) < 2 {
		fmt.Println("Pick a word to find!")
		return
	}
	w := strings.ToLower(os.Args[1])
	wordsArr := strings.Split(phrase, " ")

	for i := 0; i < len(wordsArr); i++ {
		if (w == wordsArr[i]) {
			fmt.Println("Word found!")
			return
		}
	}
	fmt.Println("Word not found!")*/

	// ---------------------------------------------------------
	// EXERCISE: Crunch the primes
	//
	//  1. Get numbers from the command-line.
	//  2. `for range` over them.
	//  4. Convert each one to an int.
	//  5. If one of the numbers isn't an int, just skip it.
	//  6. Print the ones that are only the prime numbers.
	//
	// RESTRICTION
	//  The user can run the program with any number of
	//  arguments.
	//
	// HINT
	//  Find whether a number is prime using this algorithm:
	//  https://stackoverflow.com/a/1801446/115363
	//
	// EXPECTED OUTPUT
	//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
	//    2 7 11 13
	//
	//  go run main.go 1 2 3 5 7 A B C
	//    2 3 5 7
	// ---------------------------------------------------------
	/*numbers := os.Args[1:]
	for i := 0; i < len(numbers); i++ {
		n, error := strconv.Atoi(numbers[i])
		if error != nil {
			continue
		}
		if isPrimeNumber(n) {
			fmt.Println(n)
		}
	}*/

	// ---------------------------------------------------------
	// EXERCISE: Declare empty arrays
	//
	//  1. Declare and PrintGames the following arrays with their types:
	//
	//    1. The names of your three best friends (names array)
	//
	//    2. The distances to five different locations (distances array)
	//
	//    3. A data buffer with five bytes of capacity (data array)
	//
	//    4. Currency exchange ratios only for a single currency (ratios array)
	//
	//    5. Up/Down status of four different web servers (alives array)
	//
	//    6. A byte array that doesn't occupy memory space (zero array)
	//
	//  2. Print only the types of the same arrays.
	//
	//  3. Print only the elements of the same arrays.
	//
	// HINT
	//   When printing the elements of an array, you can use the usual Printf verbs.
	//
	//   For example:
	//     When printing a string array, you can use "%q" verb as usual.
	//
	// EXPECTED OUTPUT
	//  names    : [3]string{"", "", ""}
	//  distances: [5]int{0, 0, 0, 0, 0}
	//  data     : [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
	//  ratios   : [1]float64{0}
	//  alives   : [4]bool{false, false, false, false}
	//  zero     : [0]uint8{}
	//
	//  names    : [3]string
	//  distances: [5]int
	//  data     : [5]uint8
	//  ratios   : [1]float64
	//  alives   : [4]bool
	//  zero     : [0]uint8
	//
	//  names    : ["" "" ""]
	//  distances: [0 0 0 0 0]
	//  data     : [0 0 0 0 0]
	//  ratios   : [0.00]
	//  alives   : [false false false false]
	//  zero     : []
	// ---------------------------------------------------------
	/*
		names := [...]string{"test1", "test2", "test3"}
		distances := [...]int{0, 1, 12, 23, 34}
		data := [...]uint8{0xAD, 0xC, 0xD, 0xE, 0xF}
		ratios := [...]float64{2.23}
		alives := [...]bool{false, true, false, false}
		zero := [...]uint8{}

		for i := range names {
			fmt.Printf("names[%d]:  %q", i, names[i])
			fmt.Println()
		}
		fmt.Println()
		for i := range distances {
			fmt.Printf("distances[%d]:  %d", i, distances[i])
			fmt.Println()
		}
		fmt.Println()
		for i := range data {
			fmt.Printf("data[%d]:  %d", i, data[i])
			fmt.Println()
		}
		for i := range ratios {
			fmt.Printf("ratios[%d]:  %.2f", i, ratios[i])
			fmt.Println()
		}
		fmt.Println()
		for i := range alives {
			fmt.Printf("alives[%d]:  %t", i, alives[i])
			fmt.Println()
		}
		fmt.Println()
		for i := range zero {
			fmt.Printf("zero[%d]:  %d", i, zero[i])
			fmt.Println()
		}
	*/

	// ---------------------------------------------------------
	// EXERCISE: Fix
	//
	//  1. Uncomment the code
	//
	//  2. And fix the problems
	//
	//  3. BONUS: Simplify the code
	// ---------------------------------------------------------
	/*
		names := [3]string{
			"Einstein",
			"Shepard", "Tesla"}

		books := [...]string{
			"Kafka's Revenge",
			"Stay Golden",
			"",
			"",
			""}

		fmt.Printf("%q\n", names)
		fmt.Printf("%q\n", books)
	*/

	// ---------------------------------------------------------
	// EXERCISE: Compare the Arrays
	//
	//  1. Uncomment the code
	//
	//  2. Fix the problems so that arrays become comparable
	//
	// EXPECTED OUTPUT
	//  true
	//  true
	//  false
	// ---------------------------------------------------------
	/*
		week := [...]string{"Monday", "Tuesday"}
		wend := [2]string{"Saturday", "Sunday"}

		fmt.Println(week != wend)

		evens := [...]int32{2, 4, 6, 8, 10}
		evens2 := [...]int32{2, 4, 6, 8, 10}

		fmt.Println(evens == evens2)

		// Use     : uint8 for one of the arrays instead of byte here.
		// Remember: Aliased types are the same types.
		image := [5]byte{'h', 'i'}
		var data [5]byte

		fmt.Println(data == image)
	*/
	/*
		books := [...]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

		upper, lower := books, books

		for i := range books {
			upper[i] = strings.ToUpper(upper[i])
			lower[i] = strings.ToLower(lower[i])
		}

		fmt.Printf("books: %q\n", books)
		fmt.Printf("upper: %q\n", upper)
		fmt.Printf("lower: %q\n", lower)
	*/

	/*
		fmt.Printf("names:   %#v\n", names)
		fmt.Printf("distances:   %#v\n", distances)
		fmt.Printf("data:   %#v\n", data)
		fmt.Printf("ratios:   %#v\n", ratios)
		fmt.Printf("alives:   %#v\n", alives)
		fmt.Printf("0:   %#v\n", zero)

		fmt.Printf("names:   %T\n", names)
		fmt.Printf("distances:   %T\n", distances)
		fmt.Printf("data:   %T\n", data)
		fmt.Printf("ratios:   %T\n", ratios)
		fmt.Printf("alives:   %T\n", alives)
		fmt.Printf("0:   %T\n", zero)

		fmt.Printf("names    : %q\n", names)
		fmt.Printf("distances: %d\n", distances)
		fmt.Printf("data     : %d\n", data)
		fmt.Printf("ratios   : %.2f\n", ratios)
		fmt.Printf("alives   : %t\n", alives)
		fmt.Printf("zero     : %d\n", zero)
	*/
	// ---------------------------------------------------------
	// EXERCISE: Wizard Printer
	//
	//   In this exercise, your goal is to display a few famous scientists
	//   in a pretty table.
	//
	//   1. Create a multi-dimensional array
	//   2. In each inner array, store the scientist's name, lastname and his/her
	//      nickname
	//   3. Print their information in a pretty table using a loop.
	//
	// EXPECTED OUTPUT
	//   First Name      Last Name       Nickname
	//   ==================================================
	//   Albert          Einstein        time
	//   Isaac           Newton          apple
	//   Stephen         Hawking         blackhole
	//   Marie           Curie           radium
	//   Charles         Darwin          fittest
	// ---------------------------------------------------------
	/*
		scientist := [5][3]string{{"Albert", "Einstein", "time"},
			{"Isaac", "Newton", "apple"},
			{"Stephen", "Hawking", "blackhole"},
			{"Marie", "Curie", "radium"},
			{"Charles", "Darwin", "fittest"}}

		fmt.Println(`First Name    Last Name    Nickname`)
		for i := range scientist {
			fmt.Println()
			for j := range scientist[i] {
				fmt.Printf("%-15s    ", scientist[i][j])
			}
		}
	*/

	// ---------------------------------------------------------
	// EXERCISE: Currency Converter
	//
	//   In this exercise, you're going to display currency exchange ratios
	//   against USD.
	//
	//   1. Declare a few constants with iota. They're going to be the keys
	//      of the array.
	//
	//   2. Create an array that contains the conversion ratios.
	//
	//      You should use keyed elements and the contants you've declared before.
	//
	//   3. Get the USD amount to be converted from the command line.
	//
	//   4. Handle the error cases for missing or invalid input.
	//
	//   5. Print the exchange ratios.
	//
	// EXPECTED OUTPUT
	//   go run main.go
	//     Please provide the amount to be converted.
	//
	//   go run main.go invalid
	//     Invalid amount. It should be a number.
	//
	//   go run main.go 10.5
	//     10.50 USD is 9.24 EUR
	//     10.50 USD is 8.19 GBP
	//     10.50 USD is 1186.71 JPY
	//
	//   go run main.go 1
	//     1.00 USD is 0.88 EUR
	//     1.00 USD is 0.78 GBP
	//     1.00 USD is 113.02 JPY
	// ---------------------------------------------------------
	/*
		const (
			EUR = iota
			GBP
			JPY
		)
		currencies := [...]float64{
			EUR: 0.88,
			GBP: 0.78,
			JPY: 113.02}

		names := [...]string{
			"EUR",
			"GBP",
			"JPY"}

		arg := os.Args[1]
		dolars, error := strconv.ParseFloat(arg, 64)
		if error != nil {
			fmt.Println("Wrong dolars amount!")
			return
		}

		for i := range currencies {
			fmt.Printf("%.2f dolars is %.2f %s\n", dolars, (dolars * currencies[i]), names[i])
		}
	*/

	// ---------------------------------------------------------
	// EXERCISE: Hipster's Love Bookstore Search Engine
	//
	//  Your goal is to allow people to search for books.
	//
	//  1. Create an array with the following book titles:
	//      Kafka's Revenge
	//      Stay Golden
	//      Everythingship
	//      Kafka's Revenge 2nd Edition
	//
	//  2. Get the search query from the command-line argument
	//
	//  3. Search for the books in the books array
	//
	//  4. When the program finds the book, PrintGames it.
	//  5. Otherwise, PrintGames that the book doesn't exist.
	//
	//  6. Handle the errors.
	//
	// RESTRICTION:
	//   + The search should be case insensitive.
	//
	// EXPECTED OUTPUT
	//   go run main.go
	//     Tell me a book title
	//
	//   go run main.go STAY
	//     Search Results:
	//     + Stay Golden
	//
	//   go run main.go sTaY
	//     Search Results:
	//     + Stay Golden
	//
	//   go run main.go "Kafka's Revenge"
	//     Search Results:
	//     + Kafka's Revenge
	//     + Kafka's Revenge 2nd Edition
	//
	//   go run main.go void
	//     Search Results:
	//     We don't have the book: "void"
	//
	// HINTS:
	//   + To find out whether a string contains another string value, you can use the strings.Contains function.
	//   + To convert a string value to lowercase, you can use the strings.ToLower function.
	//   + Check out the strings package for more information.
	// ---------------------------------------------------------
	/*
		books := [...]string{
			"Kafka's Revenge",
			"Stay Golden",
			"Everythingship",
			"Kafka's Revenge 2nd Edition"}

		if len(os.Args) < 2 {
			fmt.Println("Pass book title!")
			return
		}
		query := strings.ToUpper(os.Args[1])
		isFound := false
		for i := range books {
			if strings.Contains(strings.ToUpper(books[i]), query) {
				fmt.Printf("%s\n", books[i])
				isFound = true
			}
		}
		if isFound == false {
			fmt.Println("We cannot find any book :(")
	*/

	// ---------------------------------------------------------
	// EXERCISE: Find the Average
	//
	//  Your goal is to fill an array with numbers and find the average.
	//
	//  1. Get the numbers from the command-line.
	//
	//  2. Create an array and assign the given numbers to that array.
	//
	//  3. Print the given numbers and their average.
	//
	// RESTRICTION
	//   + Maximum 5 numbers can be provided
	//   + If one of the arguments are not a valid number, skip it
	//
	// EXPECTED OUTPUT
	//   go run main.go
	//     Please tell me numbers (maximum 5 numbers).
	//
	//   go run main.go 1 2 3 4 5 6
	//     Please tell me numbers (maximum 5 numbers).
	//
	//   go run main.go 1 2 3 4 5
	//     Your numbers: [1 2 3 4 5]
	//     Average: 3
	//
	//   go run main.go 1 a 2 b 3
	//     Your numbers: [1 0 2 0 3]
	//     Average: 2
	// ---------------------------------------------------------
	/*
		args := os.Args[1:]
		correctedNumbers := correctNumbers(args)
		sum := 0
		count := len(correctedNumbers)
		for i := range correctedNumbers {
			tmp, _ := strconv.Atoi(correctedNumbers[i])
			sum += tmp
		}
		avg := float64(sum) / float64(count)
		fmt.Printf("Average - %.2f", avg)
	*/

	// ---------------------------------------------------------
	// EXERCISE: Number Sorter
	//
	//  Your goal is to sort the given numbers from the command-line.
	//
	//  1. Get the numbers from the command-line.
	//
	//  2. Create an array and assign the given numbers to that array.
	//
	//  3. Sort the given numbers and PrintGames them.
	//
	// RESTRICTION
	//   + Maximum 5 numbers can be provided
	//   + If one of the arguments is not a valid number, skip it
	//
	// HINTS
	//  + You can use the bubble-sort algorithm to sort the numbers.
	//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
	//
	//  + When swapping the elements, do not check for the last element.
	//
	//    Or, you will receive this error:
	//    "panic: runtime error: index out of range"
	//
	// EXPECTED OUTPUT
	//   go run main.go
	//     Please give me up to 5 numbers.
	//
	//   go run main.go 6 5 4 3 2 1
	//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
	//
	//   go run main.go 5 4 3 2 1
	//     [1 2 3 4 5]
	//
	//   go run main.go 5 4 a c 1
	//     [0 0 1 4 5]
	// ---------------------------------------------------------
	/*
		args := os.Args[1:]
		numbers := correctNumbers(args)
		numbersSorted := [5]int{}
		if len(numbers) > 5 {
			fmt.Println("We can sort At most 5 numbers")
			return
		}
		for i := range numbers {
			tmp, _ := strconv.Atoi(numbers[i])
			numbersSorted[i] = tmp
		}
		fmt.Println(sort(numbersSorted))
	*/

	// ---------------------------------------------------------
	// EXERCISE: Word Finder
	//
	//   Your goal is to search for the words inside the corpus.
	//
	//   Note: This exercise is similar to the previous word finder program:
	//   https://github.com/inancgumus/learngo/tree/master/13-loops/10-word-finder-labeled-switch
	//
	//   1. Get the search query from the command-line (it can be multiple words)
	//
	//   2. Filter these words, do not search for them:
	//      and, or, was, the, since, very
	//
	//      To do this, use an array for the filtered words.
	//
	//   3. Print the words found.
	//
	// RESTRICTION
	//   + The search and the filtering should be case insensitive
	//
	// HINT
	//   + strings.Fields function converts a given string to a slice.
	//
	//     You can find its example in the word finder program that I've mentioned
	//     above.
	//
	// EXPECTED OUTPUT
	//   go run main.go
	//     Please give me a word to search.
	//
	//   go run main.go and was
	//
	//   go run main.go AND WAS
	//
	//   go run main.go cat beginning
	//     #2 : "cat"
	//     #11: "beginning"
	//
	//   go run main.go Cat Beginning
	//     #2 : "cat"
	//     #11: "beginning"
	// ---------------------------------------------------------
	/*
		const corpus = "lazy cat jumps again and again and again since the beginning this was very important"
		if len(os.Args) < 2 {
			fmt.Println("Pass word to find!")
			return
		}
		query := os.Args[1]
		wordsArray := strings.Split(corpus, " ")

		for i := range wordsArray {
			if query == wordsArray[i] {
				fmt.Printf("\\#%-2d : %q\n", i+1, wordsArray[i])
			}
		}
	*/

	// ---------------------------------------------------------
	// EXERCISE: Refactor
	//
	//  Goal is refactoring the clock project by moving some of its parts to
	//  a new file in the same package.
	//
	//  1. Create a new file: placeholders.go
	//  2. Move the placeholder type to placeholder.go
	//  3. Move all the placeholders (zero to nine and the colon) to placeholder.go
	//  4. Move the digits array to placeholders.go
	//
	// HINT
	//  + placeholders.go file should belong to main package as well
	//
	//    To remember how to do so: Rewatch the "PART I — Packages, Scopes and Importing"
	//    section.
	//
	//  + Short declaration won't work in the package scope.
	//    Remember why by watching: "Short Declaration: Package Scope" lecture
	//
	//  + If you receive the following error and you don't know what to do watch:
	//    "Packages - Learn how to run multiple files" lecture.
	//
	//    # command-line-arguments
	//    undefined: placeholder
	//    undefined: colon
	//
	// EXPECTED OUTPUT
	//  The same output before the refactoring.
	// ---------------------------------------------------------
	/*
		screen.Clear()
		alarmCount := 0
		for {
			screen.MoveTopLeft()

			now := time.Now()
			hour, min, sec, milisec := now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1e8

			clock := [...]placeholder{
				digits[hour/10], digits[hour%10],
				colon,
				digits[min/10], digits[min%10],
				colon,
				digits[sec/10], digits[sec%10],
				dot,
				digits[milisec%10],
			}
			if alarmCount == 99 {
				screen.Clear()
				fmt.Println(alarm)
				time.Sleep(time.Second)
				screen.Clear()
				alarmCount = 0
				continue
			}

			for line := range clock[0] {
				for index, digit := range clock {
					// colon blink
					next := clock[index][line]
					if digit == colon && sec%2 == 0 {
						next = "   "
					}
					fmt.Print(next, "  ")
				}
				fmt.Println()
			}
			alarmCount += 1
			time.Sleep(time.Second / 10)
		}

		return

	*/

	/*screen.Clear()
	shift := 0
	//raising := true

	screen.MoveTopLeft()

	for {
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}
		screen.Clear()
		screen.MoveTopLeft()
		for i := range clock {
			if i < shift {
				i = shift
			}
			clock[i-shift] = clock[i]
		}
		for i := 0; i < shift; i++ {
			clock[len(clock)-1-i] = placeholder{}
		}
		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}

				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		shift += 1
		if shift > 7 {
			shift = 0
		}
		time.Sleep(time.Second)
	}
	*/

	// ---------------------------------------------------------
	// EXERCISE: Declare nil slices
	//
	//  1. Declare the following slices as nil slices:
	//
	//    1. The names of your friends (names slice)
	//
	//    2. The distances to locations (distances slice)
	//
	//    3. A data buffer (data slice)
	//
	//    4. Currency exchange ratios (ratios slice)
	//
	//    5. Up/Down status of web servers (alives slice)
	//
	//
	//  2. Print their type, length and whether they're equal to nil value or not.
	//
	//
	// EXPECTED OUTPUT
	//  names    : []string 0 true
	//  distances: []int 0 true
	//  data     : []uint8 0 true
	//  ratios   : []float64 0 true
	//  alives   : []bool 0 true
	// ---------------------------------------------------------
	/*
		names := []string{"diho raz", "sentinel", "sentino"}
		var distances = []int{10, 23, 32}
		var data = []uint8{0xA, 0b011, 0x23}
		var ratios = []float64{1.23, 2.34}
		var alives = []bool{true, false}

		fmt.Printf("names - type: %T, length: %d, is nil: %t\n", names, len(names), names == nil)
		fmt.Printf("distances - type: %T, length: %d, is nil: %t\n", distances, len(distances), distances == nil)
		fmt.Printf("data  - type: %T, length: %d, is nil: %t\n", data, len(data), data == nil)
		fmt.Printf("ratios - type: %T, length: %d, is nil: %t\n", ratios, len(ratios), ratios == nil)
		fmt.Printf("alives - type: %T, length: %d, is nil: %t\n", alives, len(alives), alives == nil)
	*/

	// ---------------------------------------------------------
	// EXERCISE: Declare the arrays as slices
	//
	//   1. First run the following program as it is
	//
	//   2. Then change the array declarations to slice declarations
	//
	//   3. Observe whether anything changes or not (on the surface :)).
	//
	// EXPECTED OUTPUT
	//  names    : []string ["Einstein" "Tesla" "Shepard"]
	//  distances: []int [50 40 75 30 125]
	//  data     : []uint8 [72 69 76 76 79]
	//  ratios   : []float64 [3.14]
	//  alives   : []bool [true false true false]
	//  zero     : []uint8 []
	// ---------------------------------------------------------
	/*
		names := []string{"Einstein", "Tesla", "Shepard"}
		distances := []int{50, 40, 75, 30, 125}
		data := []byte{'H', 'E', 'L', 'L', 'O'}
		ratios := []float64{3.14145}
		alives := []bool{true, false, true, false}
		zero := []byte{}

		fmt.Printf("names    : %[1]T %[1]q\n", names)
		fmt.Printf("distances: %[1]T %[1]d\n", distances)
		fmt.Printf("data     : %[1]T %[1]d\n", data)
		fmt.Printf("ratios   : %[1]T %.2[1]f\n", ratios)
		fmt.Printf("alives   : %[1]T %[1]t\n", alives)
		fmt.Printf("zero     : %[1]T %[1]d\n", zero)
	*/

	// ---------------------------------------------------------
	// EXERCISE: Fix the problems
	//
	//  1. Uncomment the code
	//
	//  2. Fix the problems
	//
	//  3. BONUS: Simplify the code
	//
	//
	// EXPECTED OUTPUT
	//   "Einstein and Shepard and Tesla"
	//   ["Fire" "Kafka's Revenge" "Stay Golden"]
	//   [1 2 3 5 6 7 8 9]
	// ---------------------------------------------------------
	/*
		names := []string{
			"Einstein",
			"Shepard",
			"Tesla"}

		// -----------------------------------
		books := []string{
			"Stay Golden",
			"Fire",
			"Kafka's Revenge"}

		sort.Strings(books)

		//-----------------------------------
		// this time, do not change the nums array to a slice
		nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}

		// use the slicing expression to change the nums array to a slice below
		sort.Ints(nums[0:])

		// -----------------------------------
		// Here: Use the strings.Join function to join the names
		// (see the expected output)
		fmt.Printf("%s\n", strings.Join(names, " and "))
		fmt.Printf("%q\n", books)
		fmt.Printf("%d\n", nums)
	*/

	// ---------------------------------------------------------
	// EXERCISE: Compare the slices
	//
	//  1. Split the namesA string and get a slice
	//
	//  2. Sort all the slices
	//
	//  3. Compare whether the slices are equal or not
	//
	//
	// EXPECTED OUTPUT
	//
	//   They are equal.
	//
	//
	// HINTS
	//
	//   1. strings.Split function splits a string and
	//      returns a string slice
	//
	//   2. Comparing slices: First check whether their length
	//      are the same or not; only then compare them.
	//
	// ---------------------------------------------------------
	/*
		namesA := []string{"Da Vinci, Wozniak, Carmack"}
		newNamesA := strings.Split(namesA[0], ", ")
		sort.Strings(newNamesA)
		namesB := []string{"Wozniak", "Da Vinci", "Carmack"}
		sort.Strings(namesB)
		msg := ""
		for i := range newNamesA {
			if namesB[i] != newNamesA[i] {
				msg = "not"
				break
			}
		}
		fmt.Printf("Slices  are %s equal", msg)
	*/

	// ---------------------------------------------------------
	// EXERCISE: Append
	//
	//  Please follow the instructions within the code below.
	//
	// EXPECTED OUTPUT
	//  They are equal.
	//
	// HINTS
	//  bytes.Equal function allows you to compare two byte
	//  slices easily. Check its documentation: go doc bytes.Equal
	// ---------------------------------------------------------
	/*
		// 1. uncomment the code below
		png, header := []byte{'P', 'N', 'G'}, []byte{}
		header = append(header, png...)
		if bytes.Equal(png, header) {
			fmt.Println("Slices are equal")
		}

		// 2. append elements to header to make it equal with the png slice

		// 3. compare the slices using the bytes.Equal function

		// 4. PrintGames whether they're equal or not
	*/

	// ---------------------------------------------------------
	// EXERCISE: Append #2
	//
	//  1. Create the following nil slices:
	//     + Pizza toppings
	//     + Departure times
	//     + Student graduation years
	//     + On/off states of lights in a room
	//
	//  2. Append them some elements (use your creativity!)
	//
	//  3. Print all the slices
	//
	//
	// EXPECTED OUTPUT
	// (Your output may change, mine is like so:)
	//
	//  pizza       : [pepperoni onions extra cheese]
	//
	//  graduations : [1998 2005 2018]
	//
	//  departures  : [2019-01-28 15:09:31.294594 +0300 +03 m=+0.000325020
	//  2019-01-29 15:09:31.294594 +0300 +03 m=+86400.000325020
	//  2019-01-30 15:09:31.294594 +0300 +03 m=+172800.000325020]
	//
	//  lights      : [true false true]
	//
	//
	// HINTS
	//  + For departure times, use the time.Time type. Check its documentation.
	//
	//      now := time.Now()     -> Gives you the current time
	//      now.Add(time.Hour*24) -> Gives you a time.Time 24 hours after `now`
	//
	//  + For graduation years, you can use the int type
	// ---------------------------------------------------------
	/*
		pizza := []string{}
		graduations := []int{}
		departures := []time.Time{}
		lights := []bool{}

		pizza = append(pizza, "Extra cheese")
		graduations = append(graduations, 2019)
		departures = append(departures,
			time.Now(),
			time.Now().Add(time.Hour*24),
			time.Now().Add(time.Hour*48))
		lights = append(lights, true)

		fmt.Printf("pizza       : %s\n", pizza)
		fmt.Printf("\ngraduations : %d\n", graduations)
		fmt.Printf("\ndepartures  : %s\n", departures)
		fmt.Printf("\nlights      : %t\n", lights)
	*/

	// ---------------------------------------------------------
	// EXERCISE: Append #3 — Fix the problems
	//
	//  Fix the problems in the code below.
	//
	// BONUS
	//
	//  Simplify the code.
	//
	// EXPECTED OUTPUT
	//
	//  toppings: [black olives green peppers onions extra cheese]
	//
	// ---------------------------------------------------------
	/*
		toppings := []string{"black olives", "green peppers"}
		var pizza []string
		pizza = append(pizza, toppings...)
		pizza = append(toppings, "onions")
		toppings = append(pizza, "extra cheese")
		fmt.Printf("pizza       : %s\n", pizza)
	*/

	// ---------------------------------------------------------
	// EXERCISE: Append and Sort Numbers
	//
	//  We'll have a []string that should contain numbers.
	//
	//  Your task is to convert the []string to an int slice.
	//
	//  1. Get the numbers from the command-line
	//
	//  2. Append them to an []int
	//
	//  3. Sort the numbers
	//
	//  4. Print them
	//
	//  5. Handle the error cases
	//
	//
	// EXPECTED OUTPUT
	//
	//  go run main.go
	//    provide a few numbers
	//
	//  go run main.go 4 6 1 3 0 9 2
	//    [0 1 2 3 4 6 9]
	//
	//  go run main.go a b c
	//    []
	//
	//  go run main.go 4 a 1 c d 9
	//    [1 4 9]
	//
	// ---------------------------------------------------------
	/*
		args := os.Args[1:]
		var numbers []int
		for i := range args {
			tmp, err := strconv.Atoi(args[i])
			if err != nil {
				continue
			}
			numbers = append(numbers, tmp)
		}
		sort.Ints(numbers)
		fmt.Println(numbers)
	*/

	// ---------------------------------------------------------
	// EXERCISE: Housing Prices
	//
	//  We have received housing prices. Your task is to load the data into
	//  appropriately typed slices then PrintGames them.
	//
	//  1. Check out the expected output
	//
	//
	//  2. Check out the code below
	//
	//     1. header   : stores the column headers
	//     2. data     : stores the real data related to each column
	//     3. separator: you will use it to separate the data
	//
	//
	//  3. Parse the data
	//
	//     1. Separate it into rows by using the newline character ("\n")
	//
	//     2. For each row, separate it by using the separator (",")
	//
	//
	//  4. Load the data into distinct slices
	//
	//     1. Load the locations into a []string
	//     2. Load the sizes into []int
	//     3. Load the beds into []int
	//     4. Load the baths into []int
	//     5. Load the prices into []int
	//
	//
	//  5. Print the header
	//
	//     1. Separate it by using the separator
	//
	//     2. Print each column 15 character wide ("%-15s")
	//
	//
	//  6. Print the rows from the slices that you've created, line by line
	//
	//
	// EXPECTED OUTPUT
	//
	//  Location       Size           Beds           Baths          Price
	//  ===========================================================================
	//  New York       100            2              1              100000
	//  New York       150            3              2              200000
	//  Paris          200            4              3              400000
	//  Istanbul       500            10             5              1000000
	//
	//
	// HINTS
	//
	//  + strings.Split function can separate a string into a []string
	//
	// ---------------------------------------------------------
	/*
			const (
				header = "Location,Size,Beds,Baths,Price"
				data   = `New York,100,2,1,100000
		New York,150,3,2,200000
		Paris,200,4,3,400000
		Istanbul,500,10,5,1000000`

				separator = ","
			)
			var location []string
			var size []int
			var beds []int
			var baths []int
			var price []int
			rows := strings.Split(data, "\n")
			for _, el := range rows {
				tmp := strings.Split(el, ",")

				location = append(location, tmp[0])

				s, _ := strconv.Atoi(tmp[1])
				size = append(size, s)

				bds, _ := strconv.Atoi(tmp[2])
				beds = append(beds, bds)

				bths, _ := strconv.Atoi(tmp[3])
				baths = append(baths, bths)

				p, _ := strconv.Atoi(tmp[4])
				price = append(price, p)
			}

			fmt.Println("Location\tSize\tBeds\tBaths\tPrice")
			fmt.Println("=============================================")
			for i := range rows {
				fmt.Printf("%-15s%d\t%d\t%d\t%d\t\n", location[i], size[i], beds[i], baths[i], price[i])
			}
			fmt.Println("=============================================")
			fmt.Printf("\t%.2f\t%.2f\t%.2f\t%.2f", average(size), average(beds), average(baths), average(price))

		}
	*/
	/*
		const (
			header = "Location,Size,Beds,Baths,Price"
			data   = `New York,100,2,1,100000
			New York,150,3,2,200000
			Paris,200,4,3,400000
			Istanbul,500,10,5,1000000`

			separator = ","
		)
		var location []string
		var size []int
		var beds []int
		var baths []int
		var price []int

		rows := strings.Split(data, "\n")
		for _, el := range rows {
			tmp := strings.Split(el, ",")

			location = append(location, tmp[0])

			s, _ := strconv.Atoi(tmp[1])
			size = append(size, s)

			bds, _ := strconv.Atoi(tmp[2])
			beds = append(beds, bds)

			bths, _ := strconv.Atoi(tmp[3])
			baths = append(baths, bths)

			p, _ := strconv.Atoi(tmp[4])
			price = append(price, p)
		}

		query := os.Args[1:]
		names := strings.Split(header, ",")
		toUpper(query)
		toUpper(names)
		var from int
		elF := strings.ToUpper(os.Args[1])
		switch elF {
		case "LOCATION":
			from = 0
		case "PRICE":
			from = 4
		case "BEDS":
			from = 2
		case "BATHS":
			from = 3
		case "SIZE":
			from = 1
		}
		if len(os.Args) == 2 {
			fmt.Printf("%-15s", names[from:])
		} else {
			var to int
			elS := strings.ToUpper(os.Args[2])
			switch elS {
			case "LOCATION":
				to = 0
			case "PRICE":
				to = 4
			case "BEDS":
				to = 2
			case "BATHS":
				to = 3
			case "SIZE":
				to = 1
			}
			fmt.Printf("%-15s", names[from:to])
		}
	*/

	// ---------------------------------------------------------
	// EXERCISE: Slice the numbers
	//
	//   We've a string that contains even and odd numbers.
	//
	//   1. Convert the string to an []int
	//
	//   2. Print the slice
	//
	//   3. Slice it for the even numbers and PrintGames it (assign it to a new slice variable)
	//
	//   4. Slice it for the odd numbers and PrintGames it (assign it to a new slice variable)
	//
	//   5. Slice it for the two numbers at the middle
	//
	//   6. Slice it for the first two numbers
	//
	//   7. Slice it for the last two numbers (use the len function)
	//
	//   8. Slice the evens slice for the last one number
	//
	//   9. Slice the odds slice for the last two numbers
	//
	//
	// EXPECTED OUTPUT
	//   go run main.go
	//
	//     nums        : [2 4 6 1 3 5]
	//     evens       : [2 4 6]
	//     odds        : [1 3 5]
	//     middle      : [6 1]
	//     first 2     : [2 4]
	//     last 2      : [3 5]
	//     evens last 1: [6]
	//     odds last 2 : [3 5]
	//
	//
	// NOTE
	//
	//  You can also use my prettyslice package for printing the slices.
	//
	//
	// HINT
	//
	//  Find a function in the strings package for splitting the string into []string
	//
	// ---------------------------------------------------------
	/*
		data := "2 4 6 1 3 5"
		var numbers []int
		splitedData := strings.Split(data, " ")

		for _, el := range splitedData {
			tmp, _ := strconv.Atoi(el)
			numbers = append(numbers, tmp)
		}
		fmt.Println("nums: ", numbers)
		fmt.Println("evens: ", findEven(numbers))
		fmt.Println("odds: ", findOdds(numbers))
		fmt.Println("first 2: ", numbers[:2])
		fmt.Println("last 2: ", numbers[len(numbers)-2:])
		fmt.Println("middle 2: ", numbers[(len(numbers)/2)-1:(len(numbers)/2)+1])
	*/

	// ---------------------------------------------------------
	// EXERCISE: Slicing by arguments
	//
	//   We've a []string, you will get arguments from the command-line,
	//   then you will PrintGames only the elements that are requested.
	//
	//   1. Print the []string (it's in the code below)
	//
	//   2. Get the starting and stopping positions from the command-line
	//
	//   3. Print the []string slice by slicing it using the starting and stopping
	//      positions
	//
	//   4. Handle the error cases (see the expected output below)
	//
	//   5. Add new elements to the []string slice literal.
	//      Your program should work without changing the rest of the code.
	//
	//   6. Now, play with your program, get a deeper sense of how the slicing
	//      works.
	//
	//
	// EXPECTED OUTPUT
	//
	//  go run main.go
	//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
	//
	//    Provide only the [starting] and [stopping] positions
	//
	//
	//  (error because: we expect only two arguments)
	//
	//  go run main.go 1 2 4
	//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
	//
	//    Provide only the [starting] and [stopping] positions
	//
	//
	//  (error because: starting index >= 0 && stopping pos <= len(slice) )
	//
	//  go run main.go -1 5
	//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
	//
	//    Wrong positions
	//
	//
	//  (error because: starting <= stopping)
	//
	//  go run main.go 3 2
	//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
	//
	//    Wrong positions
	//
	//
	//  go run main.go 0
	//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
	//
	//    [Normandy Verrikan Nexus Warsaw]
	//
	//
	//  go run main.go 1
	//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
	//
	//    [Verrikan Nexus Warsaw]
	//
	//
	//  go run main.go 2
	//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
	//
	//    [Nexus Warsaw]
	//
	//
	//  go run main.go 3
	//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
	//
	//    [Warsaw]
	//
	//
	//  go run main.go 4
	//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
	//
	//    []
	//
	//
	//  go run main.go 0 3
	//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
	//
	//    [Normandy Verrikan Nexus]
	//
	//
	//  go run main.go 1 3
	//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
	//
	//    [Verrikan Nexus]
	//
	//
	//  go run main.go 1 2
	//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
	//
	//    [Verrikan]
	//
	// ---------------------------------------------------------
	/*
		ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

		args := os.Args[1:]
		msg, err := catchErrors(args)
		if err == true {
			fmt.Println(msg)
			return
		}
		f, _ := strconv.Atoi(args[0])
		if len(args) == 1 {
			fmt.Println(ships[f:])
			return
		}
		s, _ := strconv.Atoi(args[1])
		fmt.Println(ships[f:s])
	*/

	// ---------------------------------------------------------
	// EXERCISE: Fix the backing array problem
	//
	//  Ensure that changing the elements of the `mine` slice
	//  does not change the elements of the `nums` slice.
	//
	//
	// CURRENT OUTPUT (INCORRECT)
	//
	//  Mine         : [-50 -100 -150 25 30 50]
	//  Original nums: [-50 -100 -150]
	//
	//
	// EXPECTED OUTPUT
	//
	//  Mine         : [-50 -100 -150]
	//  Original nums: [56 89 15]
	//
	// ---------------------------------------------------------
	/*
		// DON'T TOUCH THE FOLLOWING CODE
		nums := []int{56, 89, 15, 25, 30, 50}

		// ----------------------------------------
		// ONLY ADD YOUR CODE HERE
		//
		// Ensure that nums slice never changes even though
		// the mine slice changes.
		var mine []int
		mine = append(mine, nums[0:3]...)
		// ----------------------------------------

		// DON'T TOUCH THE FOLLOWING CODE
		//
		// This code changes the elements of the nums
		// slice.
		//
		mine[0], mine[1], mine[2] = -50, -100, -150

		fmt.Println("Mine         :", mine)
		fmt.Println("Original nums:", nums[:3])
	*/

	// ---------------------------------------------------------
	// EXERCISE: Sort the backing array
	//
	//  1. Sort only the middle 3 items.
	//
	//  2. All the slices should see your changes.
	//
	//
	// RESTRICTION
	//
	//  Do not sort manually. Sort by using the sort package.
	//
	//
	// EXPECTED OUTPUT
	//
	//  Original: [pacman mario tetris doom galaga frogger asteroids simcity metroid defender rayman tempest ultima]
	//
	//  Sorted  : [pacman mario tetris doom galaga asteroids frogger simcity metroid defender rayman tempest ultima]
	//
	//
	// HINT:
	//
	//   Middle items are         : [frogger asteroids simcity]
	//
	//   After sorting they become: [asteroids frogger simcity]
	//
	// ---------------------------------------------------------
	/*
		items := []string{
			"pacman", "mario", "tetris", "doom", "galaga", "frogger",
			"asteroids", "simcity", "metroid", "defender", "rayman",
			"tempest", "ultima",
		}
		mt := items[5:8]
		sort.Strings(mt)

		fmt.Println("Original:", items)
		// ADD YOUR CODE HERE
		fmt.Println()
		fmt.Println("Sorted  :", items)
	*/

	/*
		// DO NOT CHANGE ANYTHING IN THIS CODE.

		// get the first three elements from api.temps
		received := api.Read(0, 3)

		// append changes the api package's temps slice's
		// backing array as well.
		received = append(received, []int{1, 3}...)

		fmt.Println("api.temps     :", api.All())
		fmt.Println("main.received :", received)
	*/

	// ---------------------------------------------------------
	// EXERCISE: Fix the memory leak
	//
	//  WARNING
	//
	//    This is a very difficult exercise. You need to
	//    do some research on your own to solve it. Please don't
	//    get discouraged if you can't solve it yet.
	//
	//
	//  GOAL
	//
	//    In this exercise, your goal is to reduce the memory
	//    usage. To do that, you need to find and fix the memory
	//    leak within `main()`.
	//
	//
	//  PROBLEM
	//
	//    `main()` calls `api.Report()` that reports the current
	//    memory usage.
	//
	//    After that, `main()` calls `api.Read()` that returns
	//    a slice with 10 millions of elements. But you only need
	//    the last 10 elements of the returned slice.
	//
	//
	//  WHAT YOU NEED TO DO
	//
	//    You only need to change the code in `main()`. Please
	//    do not touch the code in `api/api.go`.
	//
	//
	//  CURRENT OUTPUT
	//
	//    > Memory Usage: 113 KB
	//
	//    Last 10 elements: [...]
	//
	//    > Memory Usage: 65651 KB
	//
	//      + Before `api.Read()` call: It uses 113 KB of memory.
	//
	//      + After `api.Read()` call : It uses  65 MB of memory.
	//
	//      + This means that, `main()` never releases the memory.
	//        This is the leak.
	//
	//      + Your goal is to release the unused memory. Remember,
	//        you only need 10 elements but in the current code
	//        below you have a slice with 10 millions of elements.
	//
	//
	//  EXPECTED OUTPUT
	//
	//    > Memory Usage: 116 KB
	//
	//    Last 10 elements: [...]
	//
	//    > Memory Usage: 118 KB
	//
	//      + In the expected output, `main()` releases the memory.
	//
	//        It no longer uses 65 MB of memory. Instead, it only
	//        uses 118 KB of memory. That's why the second
	//        `api.Report()` call reports 118 KB.
	//
	//
	//  ADDITIONAL NOTE
	//
	//    Memory leak means: Your program is using unnecessary
	//    computer memory. It doesn't release memory that is
	//    no longer needed.
	//
	//    See this for more information:
	//    https://en.wikipedia.org/wiki/Memory_leak
	//
	//
	//  HINTS
	//
	//    Check out `hints.md` file if you get stuck.
	//
	// ---------------------------------------------------------
	// reports the initial memory usage
	/*
		api.Report()

		// returns a slice with 10 million elements.
		// it allocates 65 MB of memory space.
		millions := api.Read()

		// -----------------------------------------------------
		// ✪ ONLY CHANGE THE CODE IN THIS AREA ✪
		millions = append([]int(nil), millions[len(millions)-10:]...)
		last10 := millions[len(millions)-10:]

		fmt.Printf("\nLast 10 elements: %d\n\n", last10)

		// ✪ ONLY CHANGE THE CODE IN THIS AREA ✪
		// -----------------------------------------------------
		//millions = []int(nil)
		api.Report()

		// don't worry about this code.
		fmt.Fprintln(ioutil.Discard, millions[0])
	*/

	// ---------------------------------------------------------
	// EXERCISE: Print daily requests
	//
	//  You've got request logs of a web server. The log data
	//  contains 8-hourly totals per each day. It is stored
	//  in the `reqs` slice.
	//
	//  Find and PrintGames the total requests per day, as well as
	//  the grand total.
	//
	//  See the `reqs` slice and the steps in the code below.
	//
	//
	// RESTRICTIONS
	//
	//   1. You need to produce the daily slice, don't just loop
	//      and PrintGames the element totals directly. The goal is
	//      gaining more experience in slice operations.
	//
	//   2. Your code should work even if you add to or remove the
	//      existing elements from the `reqs` slice.
	//
	//      For example, after solving the exercise, try it with
	//      this new data:
	//
	//      reqs := []int{
	// 	      500, 600, 250,
	// 	      200, 400, 50,
	// 	      900, 800, 600,
	// 	      750, 250, 100,
	// 	      150, 654, 235,
	// 	      320, 534, 765,
	// 	      121, 876, 285,
	// 	      543, 642,
	// 	      // the last element is missing (your code should be able to handle this)
	// 	      // that is why you shouldn't calculate the `size` below manually.
	//      }
	//
	//      The grand total of the new data should be 10525.
	//
	//
	// EXPECTED OUTPUT
	//
	//   Please run `solution/main.go` to see the expected
	//   output.
	//
	//   go run solution/main.go
	//
	// ---------------------------------------------------------
	// There are 3 request totals per day (8-hourly)
	/*
		const N = 3

		// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
		reqs := []int{
			500, 600, 250,
			200, 400, 50,
			900, 800, 600,
			750, 250, 100,
			150, 654, 235,
			320, 534, 765,
			121, 876, 285,
			543, 642,
			// 	      // the last element is missing (your code should be able to handle this)
			// 	      // that is why you shouldn't calculate the `size` below manually.
		}

		// ================================================
		// #1: Make a new slice with the exact size needed.

		//_ = reqs // remove this when you start
		var days int
		if len(reqs)%3 == 0 {
			days = len(reqs) / 3
		} else {
			days = (len(reqs) / 3) + 1
		}
		fmt.Println(days)
		daily := make([][]int, days)
		for i := 0; i < days; i++ {
			daily[i] = make([]int, 0, N)
		}
		// ================================================

		// ================================================
		// #2: Group the `reqs` per day into the slice: `daily`.
		daysc := 0
		sum := 0
		for i, el := range reqs {
			daily[daysc] = append(daily[daysc], el)
			sum += el
			if (i+1)%3 == 0 {
				daysc += 1
			}
		}
		//
		// So the daily will be:
		// [
		//  [500, 600, 250]
		//  [200, 400, 50]
		//  [900, 800, 600]
		//  [750, 250, 100]
		// ]

		//_ = daily // remove this when you start

		// ================================================
		// #3: Print the results

		// Print a header
		fmt.Printf("%-10s%-10s\n", "Day", "Requests")
		fmt.Println(strings.Repeat("=", 20))
		for i, el := range daily {
			fmt.Printf("%d day: %v\n", i+1, el)
		}
		fmt.Printf("Total: %d", sum)

		// Loop over the daily slice and its inner slices to find
		// the daily totals and the grand total.
		// ...

		// ================================================
	*/

	// ---------------------------------------------------------
	// EXERCISE: Sort and write items to a file
	//
	//  1. Get arguments from command-line
	//
	//  2. Sort them
	//
	//  3. Write the sorted slice to a file
	//
	//
	// EXPECTED OUTPUT
	//
	//   go run main.go
	//     Send me some items and I will sort them
	//
	//   go run main.go orange banana apple
	//
	//   cat sorted.txt
	//     apple
	//     banana
	//     orange
	//
	//
	// HINTS
	//
	//   + REMEMBER: os.Args is a []string
	//
	//   + String slices are sortable using `sort.Strings`
	//
	//   + Use ioutil.WriteFile to write to a file.
	//
	//   + But you need to convert []string to []byte to be able to
	//     write it to a file using the ioutil.WriteFile.
	//
	//   + To do that, create a new []byte and append the elements of your
	//     []string.
	//
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Sort and write items to a file with their ordinals
	//
	//  Use the previous exercise.
	//
	//  This time, PrintGames the sorted items with ordinals
	//  (see the expected output)
	//
	//
	// EXPECTED OUTPUT
	//
	//   go run main.go
	//     Send me some items and I will sort them
	//
	//   go run main.go orange banana apple
	//
	//   cat sorted.txt
	//     1. apple
	//     2. banana
	//     3. orange
	//
	//
	// HINTS
	//
	//   ONLY READ THIS IF YOU GET STUCK
	//
	//   + You can use strconv.AppendInt function to append an int
	//     to a byte slice. strconv contains a lot of functions for appending
	//     other basic types to []byte slices as well.
	//
	//   + You can append individual characters to a byte slice using
	//     rune literals (because: rune literal are typeless numerics):
	//
	//     var slice []byte
	//     slice = append(slice, 'h', 'i', ' ', '!')
	//     fmt.Printf("%s\n", slice)
	//
	//     Above code prints: hi !
	// ---------------------------------------------------------
	/*
		args := os.Args[1:]
		if len(args) < 1 {
			fmt.Println("Pass some words to sort and write into file")
			return
		}
		sort.Strings(args)
		var out []byte
		var ord string
		for i := range args {

			out = strconv.AppendInt(out, int64(i+1), 10)
			out = append(out, ". "...)
			out = append(out, ord...)
			out = append(out, args[i]...)
			out = append(out, "\n"...)
		}
		ioutil.WriteFile("Test.txt", out, 0644)
	*/

	// ---------------------------------------------------------
	// EXERCISE: Find and write the names of subdirectories to a file
	//
	//  Create a program that can get multiple directory paths from
	//  the command-line, and prints only their subdirectories into a
	//  file named: dirs.txt
	//
	//
	//  1. Get the directory paths from command-line
	//
	//  2. Append the names of subdirectories inside each directory
	//     to a byte slice
	//
	//  3. Write that byte slice to dirs.txt file
	//
	//
	// EXPECTED OUTPUT
	//
	//   go run main.go
	//     Please provide directory paths
	//
	//   go run main.go dir/ dir2/
	//
	//   cat dirs.txt
	//
	//     dir/
	//             subdir1/
	//             subdir2/
	//
	//     dir2/
	//             subdir1/
	//             subdir2/
	//             subdir3/
	//
	//
	// HINTS
	//
	//   ONLY READ THIS IF YOU GET STUCK
	//
	//   + Get all the files in a directory using ioutil.ReadDir
	//     (A directory is also a file)
	//
	//   + You can use IsDir method of a FileInfo value to detect
	//     whether a file is a directory or not.
	//
	//     Check out its documentation:
	//
	//     go doc os.FileInfo.IsDir
	//
	//     # or using godocc
	//     godocc os.FileInfo.IsDir
	//
	//   + You can use '\t' escape sequence for indenting the subdirs.
	//
	//   + You can find a sample directory structure under:
	//     solution/ directory
	//
	// ---------------------------------------------------------
	/*
		args := os.Args[1:]
		if len(args) < 1 {
			fmt.Println("Pass some directories")
			return
		}

		for ind, i := range args {
			dir, err := ioutil.ReadDir(i)
			if err != nil {
				println("Bad directory!")
				continue
			}
			fmt.Println(args[ind])
			for _, file := range dir {
				if file.IsDir() == true {
					fmt.Println("\t", file.Name())
				}
			}
		}
	*/

	// ---------------------------------------------------------
	// EXERCISE: Find the Bug
	//
	//  As I've annotated in the lectures, there is a bug
	//  in this code. Please find the bug and fix it.
	//
	//
	// HINT #1
	//
	//  💀 Read this only if you get stuck.
	//
	//  Print the width*height and the capacity of the drawing buffer
	//  after a single drawing loop ends. You might be surprised.
	//
	//
	// HINT #2
	//
	//  💀 Read this only if you get stuck.
	//
	//  The bug is in the drawing buffer. It doesn't include the
	//  newline and space characters when creating the buffer. So
	//  the buffer is not large enough to hold all the characters.
	//  So new backing arrays are getting allocated.
	//
	// ---------------------------------------------------------
	/*
		const (
			cellEmpty = ' '
			cellBall  = '⚾'

			maxFrames = 1200
			speed     = time.Second / 20
		)

		var (
			px, py int    // ball position
			vx, vy = 1, 1 // velocities
			//	ppx, ppy      int
			cell          rune // current cell (for caching)
			width, height = screen.Size()
		)

		// create the board
		//board := make([]bool, width*height)
		//for column := range board {
		//	board[column] = make([]bool, height)
		//}

		// create a drawing buffer
		buf := make([]rune, 0, (width*2+1)*height)

		// clear the screen once
		screen.Clear()

		for i := 0; i < maxFrames; i++ {
			// calculate the next ball position
			px += vx
			py += vy

			// when the ball hits a border reverse its direction
			if px <= 0 || px >= width-1 {
				vx *= -1
			}
			if py <= 0 || py >= height-1 {
				vy *= -1
			}

			// remove the previous ball
			//for y := range board[0] {
			//	for x := range board {
			//buf[(ppx*height)+ppy] = cellEmpty
			//	}
			//}

			// put the new all and save last position
			//buf[(px*height)+py] = cellBall
			//ppx = px
			//ppy = py

			// rewind the buffer (allow appending from the beginning)
			buf = buf[:0]

			// draw the board into the buffer
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					buf = append(buf, cellEmpty)
					if x == px && y == py {
						buf = append(buf, cellBall)
					}
					buf = append(buf, cell, ' ')
				}
				buf = append(buf, '\n')
			}

			// PrintGames the buffer
			screen.MoveTopLeft()
			fmt.Print(string(buf))

			// slow down the animation
			time.Sleep(speed)
		}
	*/

	// ---------------------------------------------------------
	// EXERCISE: Convert the strings
	//
	//   1. Loop over the words slice
	//
	//   2. In the loop:
	//      1. Convert each string value to a byte slice
	//      2. Print the byte slice
	//      3. Append the byte slice to the `bwords`
	//
	//   3. Print the words using the `bwords`
	//
	// EXPECTED OUTPUT
	//  [103 111 112 104 101 114]
	//  [112 114 111 103 114 97 109 109 101 114]
	//  [103 111 32 108 97 110 103 117 97 103 101]
	//  [103 111 32 115 116 97 110 100 97 114 100 32 108 105 98 114 97 114 121]
	//  gopher
	//  programmer
	//  go language
	//  go standard library
	// ---------------------------------------------------------
	/*
		words := []string{
			"gopher",
			"programmer",
			"go language",
			"go standard library",
		}

		var bwords [][]byte

		for _, el := range words {
			var bs []byte
			bs = append(bs, el...)
			fmt.Println(bs)
			bwords = append(bwords, bs)
		}
		for _, el := range bwords {
			fmt.Println(string(el))
		}
	*/

	// ---------------------------------------------------------
	// EXERCISE: Print the runes
	//
	//  1. Loop over the "console" word and PrintGames its runes one by one,
	//     in decimals, hexadecimals and binary.
	//
	//  2. Manually put the runes of the "console" word to a byte slice, one by one.
	//
	//     As the elements of the byte slice use only the rune literals.
	//
	//     Print the byte slice.
	//
	//  3. Repeat the step 2 but this time, as the elements of the byte slice,
	//     use only decimal numbers.
	//
	//  4. Repeat the step 2 but this time, as the elements of the byte slice,
	//     use only hexadecimal numbers.
	//
	//
	// EXPECTED OUTPUT
	//   Run the solution to see the expected output.
	// ---------------------------------------------------------
	/*
		const word = "console"

		for _, w := range word {
			fmt.Printf("%c\n", w)
			fmt.Printf("\tdecimal: %[1]d\n", w)
			fmt.Printf("\thex    : 0x%[1]x\n", w)
			fmt.Printf("\tbinary : 0b%08[1]b\n", w)
		}

		// PrintGames the word manually using runes
		fmt.Printf("with runes       : %s\n",
			string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}))

		// PrintGames the word manually using decimals
		fmt.Printf("with decimals    : %s\n",
			string([]byte{99, 111, 110, 115, 111, 108, 101}))

		// PrintGames the word manually using hexadecimals
		fmt.Printf("with hexadecimals: %s\n",
			string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}))
	*/

	/*
		words := []string{
			"cool",
			"güzel",
			"jīntiān",
			"今天",
			"read 🤓",
		}

		for _, s := range words {
			fmt.Printf("%q\n", s)

			// Print the byte and rune length of the strings
			// Hint: Use len and utf8.RuneCountInString
			fmt.Printf("\thas %d bytes and %d runes\n",
				len(s), utf8.RuneCountInString(s))

			// Print the bytes of the strings in hexadecimal
			// Hint: Use % x verb
			fmt.Printf("\tbytes   : % x\n", s)

			// Print the runes of the strings in hexadecimal
			// Hint: Use % x verb
			fmt.Print("\trunes   :")
			for _, r := range s {
				fmt.Printf("% x", r)
			}
			fmt.Println()

			// Print the runes of the strings as rune literals
			// Hint: Use for range
			fmt.Print("\trunes   :")
			for _, r := range s {
				fmt.Printf("%q", r)
			}
			fmt.Println()

			// Print the first rune and its byte size of the strings
			// Hint: Use utf8.DecodeRuneInString
			r, size := utf8.DecodeRuneInString(s)
			fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)

			// Print the last rune of the strings
			// Hint: Use utf8.DecodeLastRuneInString
			r, size = utf8.DecodeLastRuneInString(s)
			fmt.Printf("\tlast    : %q (%d bytes)\n", r, size)

			// Slice and PrintGames the first two runes of the strings
			_, first := utf8.DecodeRuneInString(s)
			_, second := utf8.DecodeRuneInString(s[first:])
			fmt.Printf("\tfirst 2 : %q\n", s[:first+second])

			// Slice and PrintGames the last two runes of the strings
			_, last1 := utf8.DecodeLastRuneInString(s)
			_, last2 := utf8.DecodeLastRuneInString(s[:len(s)-last1])
			fmt.Printf("\tlast 2  : %q\n", s[len(s)-last2-last1:])

			// Convert the string to []rune
			// Print the first and last two runes
			rs := []rune(s)
			fmt.Printf("\tfirst 2 : %q\n", string(rs[:2]))
			fmt.Printf("\tlast 2  : %q\n", string(rs[len(rs)-2:]))
		}
	*/

	/*
		var http []byte
		http = append(http, 104, 116, 116, 112, 58, 47, 47)
		if len(os.Args) < 2 {
			fmt.Println("Pass correct page url!")
			return
		}
		arg := os.Args[1]
		burl := make([]byte, 0, len(arg)+5)
		buffer := make([]byte, 0, len(arg)+5)
		burl = append(burl, arg...)
		canMask := false
		fmt.Println(burl)

		for _, el := range burl {
			if !canMask && bytes.Contains(buffer, http) {
				canMask = true
				http[1] = 0
			}
			if canMask {
				if el == ' ' {
					canMask = false
				} else {
					buffer = append(buffer, '*')
				}
			} else {
				buffer = append(buffer, el)
			}
		}
		buffer = AddRest(buffer, os.Args[2:])
		fmt.Println(string(buffer))
	*/

	/*
		if len(os.Args) < 2 {
			fmt.Println("Pass a quote to wrap!")
			return
		}
		var q string
		q = os.Args[1]
		cc := 0
		var buffer []byte
		for ind, _ := range q {
			if cc >= 40 && q[ind] == ' ' {
				buffer = append(buffer, "\n"...)
				cc = 0
			} else {
				buffer = append(buffer, q[ind])
			}
			cc += 1
		}
		fmt.Println(string(buffer))
	*/

	// ---------------------------------------------------------
	// EXERCISE: Warm-up
	//
	//  Create and PrintGames the following maps.
	//
	//  1. Phone numbers by last name
	//  2. Product availability by Product ID
	//  3. Multiple phone numbers by last name
	//  4. Shopping basket by Customer ID
	//
	//     Each item in the shopping basket has a Product ID and
	//     quantity. Through the map, you can tell:
	//     "Mr. X has bought Y bananas"
	//
	// ---------------------------------------------------------
	/*
		phoneBook := map[string]string{
			"bowen": "202-555-0179",
			"dulin": "03.37.77.63.06",
			"greco": "03489940240",
		}
		magazine := map[int]bool{
			617841573: true,
			879401371: false,
			576872813: true,
		}
		phonesBook := map[string][]string{
			"bowen": {"202-555-0179"},
			"dulin": {"03.37.77.63.06 03.37.70.50.05 02.20.40.10.04"},
			"greco": {"03489940240", "03489900120"},
		}
		shoppingBasket := map[int]map[int]int{
			100: {617841573: 4, 576872813: 2},
			101: {576872813: 5, 657473833: 20},
			102: {},
		}

		fmt.Printf("phones     : %#v\n", phoneBook)
		fmt.Printf("products   : %#v\n", magazine)
		fmt.Printf("multiPhones: %#v\n", phonesBook)
		fmt.Printf("basket     : %#v\n", shoppingBasket)

		who, phone := "dulin", "N/A"
		if v, ok := phoneBook[who]; ok {
			phone = v
		}
		fmt.Printf("%s's phone number: %s\n", who, phone)

		// Is Product ID 879401371 available?
		id, status := 879401371, "available"
		if !magazine[id] {
			status = "not " + status
		}
		fmt.Printf("Product ID #%d is %s\n", id, status)

		// What is Greco's second phone number?
		who, phone = "greco", "N/A"
		if phones := phonesBook[who]; len(phones) >= 2 {
			phone = phones[1]
		}
		fmt.Printf("%s's 2nd phone number: %s\n", who, phone)

		// How many of 576872813 the customer 101 is going to buy?
		cid, pid := 101, 576872813
		fmt.Printf("Customer #%d is going to buy %d from Product ID #%d.\n", cid, shoppingBasket[cid][pid], pid)
	*/

	// ---------------------------------------------------------
	// EXERCISE: Students
	//
	//  Create a program that returns the students by the given
	//  Hogwarts house name (see the data below).
	//
	//  Print the students sorted by name.
	//
	//  "bobo" doesn't belong to Hogwarts, remove it by using
	//  the delete function.
	//
	//
	// RESTRICTIONS
	//
	//  + Add the following data to your map as is.
	//    Do not sort it manually and do not modify it.
	//
	//  + Slices in the map shouldn't be sorted (changed).
	//    HINT: Copy them.
	//
	//
	// EXPECTED OUTPUT
	//
	//  go run main.go
	//
	//  Please type a Hogwarts house name.
	//
	//
	//  go run main.go bobo
	//
	//  Sorry. I don't know anything about "bobo".
	//
	//
	//  go run main.go hufflepuf
	//
	//  ~~~ hufflepuf students ~~~
	//
	//        + diggory
	//        + helga
	//        + scamander
	//        + wenlock
	//
	// ---------------------------------------------------------
	/*
		houses := map[string][]string{
			"gryffindor": {
				"weasley",
				"hagrid",
				"dumbledore",
				"lupin",
			},
			"hufflepuf": {
				"wenlock",
				"scamander",
				"helga",
				"diggory",
			},
			"ravenclaw": {
				"flitwick",
				"bagnold",
				"wildsmith",
				"montmorency",
			},
			"slytherin": {
				"horace",
				"nigellus",
				"higgs",
				"scorpius",
			},
		}

		args := os.Args[1:]
		if len(args) < 1 {
			fmt.Println("Pass house name!")
			return
		}
		for i := 0; i < len(args); i++ {
			fmt.Println(strings.Join(houses[args[i]], "\n"))
		}
	*/

	// ---------------------------------------------------------
	// EXERCISE: Warm Up
	//
	//  Starting with this exercise, you'll build a command-line
	//  game store.
	//
	//  1. Declare the following structs:
	//
	//     + item: id (int), name (string), price (int)
	//
	//     + game: embed the item, genre (string)
	//
	//
	//  2. Create a game slice using the following data:
	//
	//     id  name          price    genre
	//
	//     1   god of war    50       action adventure
	//     2   x-com 2       30       strategy
	//     3   minecraft     20       sandbox
	//
	//
	//  3. Print all the games.
	//
	// EXPECTED OUTPUT
	//  Please run the solution to see the output.
	// ---------------------------------------------------------
	/*
	   	const data = `
	   [
	           {
	                   "id": 1,
	                   "name": "god of war",
	                   "genre": "action adventure",
	                   "price": 50
	           },
	           {
	                   "id": 2,
	                   "name": "x-com 2",
	                   "genre": "strategy",
	                   "price": 40
	           },
	           {
	                   "id": 3,
	                   "name": "minecraft",
	                   "genre": "sandbox",
	                   "price": 20
	           }
	   ]`

	   	games := []game{
	   		{
	   			item:  item{id: 1, name: "god of war", price: 50},
	   			genre: "action adventure",
	   		},
	   		{
	   			item:  item{id: 2, name: "x-com 2", price: 20},
	   			genre: "strategy",
	   		},
	   		{
	   			item:  item{id: 3, name: "minecraft", price: 25},
	   			genre: "sandbox",
	   		},
	   	}
	   	gamesMap := fillMap(games)

	   	in := bufio.NewScanner(os.Stdin)
	   	for {
	   		screen.Clear()
	   		screen.MoveTopLeft()
	   		fmt.Println("Welcome in game store\nWhat you wanna do?\n1. Print games\n2. Find game by id\n3. Save\n4. Load\n5. Quit")
	   		if in.Scan() {
	   			cmd := in.Text()
	   			switch cmd {
	   			case "1":
	   				PrintGames(games)
	   			case "2":
	   				screen.Clear()
	   				screen.MoveTopLeft()
	   				fmt.Println("Pass id of the game")
	   				if in.Scan() {
	   					arg := in.Text()
	   					if id := PassAndCheckId(arg); id == -1 {
	   						fmt.Println("Wrong id")
	   						time.Sleep(2 * time.Second)
	   						break
	   					} else {
	   						g, found := FindGameById(gamesMap, id)
	   						if !found {
	   							fmt.Println("We can not find the game")
	   						} else {
	   							fmt.Printf("#%d %-15s %-15d %-15s\n", g.id, g.name, g.price, g.genre)
	   							time.Sleep(2 * time.Second)
	   						}
	   					}
	   				}
	   			case "3":
	   				SaveGamesToJson(games)
	   			case "4":
	   				gs := LoadGamesFromJson(data)
	   				PrintGames(gs)
	   			case "5":
	   				break
	   			default:
	   				fmt.Println("Unknown command")
	   				time.Sleep(2 * time.Second)
	   			}
	   		}
	   	}
	*/

	/*
		// create a nil pointer with the type of pointer to a computer
		var cptr *computer
		cptr = nil
		// compare the pointer variable to a nil value, and say it's nil
		if cptr == nil {
			fmt.Println("cptr ptr is nil ")
		}
		// create an apple computer by putting its address to a pointer variable
		apple := &computer{"apple"}
		// put the apple into a new pointer variable
		newApple := apple
		// compare the apples: if they are equal say so and print their addresses
		if apple == newApple {
			fmt.Println("Pointers are equal")
			fmt.Printf("%p - first address, %p - second address\n", apple, newApple)
		}
		// create a sony computer by putting its address to a new pointer variable
		sony := &computer{"sony"}
		// compare apple to sony, if they are not equal say so and print their
		// addresses
		if sony != apple {
			fmt.Println("Sony is not equal to apple")
			fmt.Printf("%p - sony address, %p - apple address\n", sony, apple)
		}
		// put apple's value into a new ordinary variable
		ordApple := *apple
		// print apple pointer variable's address, and the address it points to
		// and, print the new variable's addresses as well
		// compare the value that is pointed by the apple and the new variable
		// if they are equal say so
		if ordApple == *apple {
			fmt.Println("apple and ordApple values are equal")
			fmt.Printf("%p - ordApple address, %p - apple address\n", &ordApple, apple)
		}
		// print the values:
		// the value that is pointed by the apple and the new variable
		fmt.Printf("%s - ordApple value, %s - apple value\n", ordApple.brand, (*apple).brand)
		// create a new function: change
		// the func can change the given computer's brand to another brand
		change(apple)
		// change sony's brand to hp using the func — print sony's brand
		fmt.Printf("%s - new brand\n", (*apple).brand)
		// write a func that returns the value that is pointed by the given *computer
		// print the returned value
		fmt.Printf("%v- printed value\n", val(apple))
		// write a new constructor func that returns a pointer to a computer
		// and call the func 3 times and print the returned values' addresses
	*/

	// ---------------------------------------------------------
	// EXERCISE: Swap
	//
	//  Using funcs, swap values through pointers, and swap
	//  the addresses of the pointers.
	//
	//
	//  1- Swap the values through a func
	//
	//     a- Declare two float variables
	//
	//     b- Declare a func that can swap the variables' values
	//        through their memory addresses
	//
	//     c- Pass the variables' addresses to the func
	//
	//     d- Print the variables
	//
	//
	//  2- Swap the addresses of the float pointers through a func
	//
	//     a- Declare two float pointer variables and,
	//        assign them the addresses of float variables
	//
	//     b- Declare a func that can swap the addresses
	//        of two float pointers
	//
	//     c- Pass the pointer variables to the func
	//
	//     d- Print the addresses, and values that are
	//        pointed by the pointer variables
	//
	// ---------------------------------------------------------
	/*
		var a, b *float64
		va := 2.67
		vb := 3.09
		a = &va
		b = &vb
		fmt.Printf("a - address: %p, value %.2f\n", &va, va)
		fmt.Printf("b - address: %p, value %.2f\n", &vb, vb)
		swapValues(&va, &vb)
		fmt.Printf("a - address: %p, value %.2f\n", &va, va)
		fmt.Printf("b - address: %p, value %.2f\n", &vb, vb)
		swapAddresses(&a, &b)
		fmt.Printf("a - address: %p, value %.2f\n", a, va)
		fmt.Printf("b - address: %p, value %.2f\n", b, vb)
	*/

	/*
		var c *computer = &computer{brand: nil}
		change(c, "apple")
		fmt.Printf("brand: %s\n", c.brand)
	*/
	/*
		var schools []map[int]string

		schools = append(schools, make(map[int]string))
		load(schools[0], []string{"batman", "superman"})

		schools = append(schools, make(map[int]string))
		load(schools[1], []string{"spiderman", "wonder woman"})

		fmt.Println(schools[0])
		fmt.Println(schools[1])
	*/
}

/*
func load(m map[int]string, students []string) {
	for i, name := range students {
		m[i+1] = name
	}
}
*/
/*
func change(c *computer, brand string) {
	c.brand = &brand
}

type computer struct {
	brand *string
}
*/
/*
func swapValues(a, b *float64) {
	//tmp := *a
	//*a = *b
	//*b = tmp
	*a, *b = *b, *a
	//return *a, *b
}
func swapAddresses(a, b **float64) {
	*a, *b = *b, *a

}
*/
/*
type computer struct {
	brand string
}

func val(p *computer) computer {
	return *p
}
func change(p *computer) {
	p.brand = "Another Brand"
	return
}
*/
/*
func PrintGames(games []game) {
	fmt.Printf("id %-15s %-15s %-15s\n", "title", "price", "genre")
	for _, g := range games {
		fmt.Printf("#%d %-15s %-15d %-15s\n", g.id, g.name, g.price, g.genre)
	}
	fmt.Println("Backing to menu...")
	time.Sleep(5 * time.Second)
}

func fillMap(games []game) map[int]game {

	gamesMap := map[int]game{}
	for _, el := range games {
		gamesMap[el.id] = el
	}
	return gamesMap
}

func PassAndCheckId(arg string) int {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return -1
	}
	return id
}
func FindGameById(gamesMap map[int]game, id int) (game, bool) {
	g, ok := gamesMap[id]
	return g, ok
}
func SaveGamesToJson(games []game) {
	for _, g := range games {
		jg := jsonGame{Name: g.name, Id: g.id, Price: g.price, Genre: g.genre}
		eg, err := json.MarshalIndent(jg, "", "\t")
		if err != nil {
			break
		}
		fmt.Println(string(eg))
	}
}
func LoadGamesFromJson(data string) []game {
	var bd []byte
	bd = append(bd, data...)
	var decodedGames []jsonGame
	json.Unmarshal(bd, &decodedGames)
	games := make([]game, 0, len(decodedGames))
	for _, el := range decodedGames {
		tmp := game{item: item{id: el.Id, name: el.Name, price: el.Price}, genre: el.Genre}
		games = append(games, tmp)
	}
	return games
}
*/
/*
func AddRest(buffer []byte, args []string) []byte {
	for _, el := range args {
		buffer = append(buffer, ' ')
		buffer = append(buffer, el...)
	}
	return buffer
}
*/
/*
// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
*/
/*
func toUpper(arr []string) {
	for _, el := range arr {
		strings.ToUpper(el)
	}
	return
}
*/
/*
func catchErrors(args []string) (string, bool) {
	if len(args) < 1 {
		return "Provide start and/or stop position!", true
	} else if len(args) > 2 {
		return "Provide only start and/or stop position", true
	} else if len(args) == 1 {
		if _, err := strconv.Atoi(args[0]); err == nil {
			return "", false
		} else {
			return "You have provided bad start!", true
		}
	}
	f, err := strconv.Atoi(args[0])
	if err != nil || f < 0 {
		return "You have provided bad start!", true
	}
	if s, errs := strconv.Atoi(args[1]); errs != nil || s < 0 || s < f {
		return "You have provided bad stop!", true
	}
	return "", false
}
*/
/*
func findOdds(arr []int) []int {
	var result []int
	for _, el := range arr {
		if el%2 == 1 {
			result = append(result, el)
		}
	}
	return result
}
*/
/*
func findEven(arr []int) []int {
	var result []int
	for _, el := range arr {
		if el%2 == 0 {
			result = append(result, el)
		}
	}
	return result
}
*/
/*
func printPart(array string, count int) {
	if len(array) == 1 {
		fmt.Println(array)
		time.Sleep(100)
		return
	}
	fmt.Println(array)
	time.Sleep(100)
	arr := array[count:]
	count += 1
	printPart(arr, count)
}
*/
/*
func sort(array [5]int) [5]int {
	for i := range array {
		for j := range array {
			if array[i] <= array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}
*/
/*
func correctNumbers(array []string) []string {
	for i := range array {

		if _, error := strconv.Atoi(array[i]); error != nil {
			array[i] = "0"
		}
	}
	return array
}
*/
/*
func setArgsIndex(isVerbose bool) int {
	if isVerbose {
		return 2
	} else {
		return 1
	}
}
func checkCommandLineArgs() (bool, bool) {
	isVerbose := false
	isCorrect := false
	switch len(os.Args) {
	case 1, 2:
		fmt.Println("Pick Two numbers!")
	case 3:
		if os.Args[1] == "-v" {
			fmt.Println("Pick Two numbers!")
		} else {
			isCorrect = true
		}
	case 4:
		if os.Args[1] == "-v" {
			isCorrect = true
			isVerbose = true
		}
	default:
	}
	return isCorrect, isVerbose
}

func setPickedNumbersAndSetMaxRounds(argsIndex int) (int, error, int, error) {
	fpn := os.Args[argsIndex]
	fn, fError := strconv.Atoi(fpn)
	spn := os.Args[argsIndex+1]
	sn, sError := strconv.Atoi(spn)
	return fn, fError, sn, sError
}

func drawNumbersAndShowResult(fn int, sn int, isVerbose bool, maxRounds int) {
	rand.Seed(time.Now().UnixNano())
	r := pickBiggerNumber(fn, sn)
	if isSmallNumber(r) {
		r = 10
	}
	for i := 1; i <= maxRounds; i++ {
		cm := rand.Intn(r)
		if isVerbose {
			fmt.Printf("Round %d. Picked number - %d\n", i, cm)
		}
		if cm == fn || cm == sn {
			winnerInfo()
			return
		}
	}
	loserInfo()
}

func checkErrors(fError error, sError error) bool {
	if fError != nil {
		fmt.Println("Błąd!", fError)
		return true
	}
	if sError != nil {
		fmt.Println("Błąd!", sError)
		return true
	}
	return false
}

func winnerInfo() {
	switch rand.Intn(WON_TEXTS) {
	case 0:
		fmt.Println("You won!")
	case 1:
		fmt.Println("You are awesome!")
	default:
		fmt.Println("Oops, we hava a problem")
	}
	return
}

func loserInfo() {
	switch rand.Intn(LOSE_TEXTS) {
	case 0:
		fmt.Println("You lost!")
	case 1:
		fmt.Println("You lost. Try again?")
	default:
		fmt.Println("Oops, we hava a problem")
	}
	return
}

func pickBiggerNumber(n1 int, n2 int) int {
	if n1 >= n2 {
		return n1
	} else {
		return n2
	}
}

func isSmallNumber(n int) bool {
	return n < 10
}

func setMaxRounds(fn int, sn int) int {
	maxGuess := pickBiggerNumber(fn, sn)
	if maxGuess <= 10 {
		return 5
	} else if maxGuess > 10 && maxGuess <= 30 {
		return 10
	} else if maxGuess > 30 && maxGuess <= 50 {
		return 15
	} else {
		return 20
	}
}

func isPrimeNumber(number int) bool {
	for i := 2; i < (number/2)+1; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

*/
/*
func average(arr []int) float64 {
	sum := 0
	count := len(arr)
	for _, el := range arr {
		sum += el
	}
	return float64(sum) / float64(count)
}
*/
