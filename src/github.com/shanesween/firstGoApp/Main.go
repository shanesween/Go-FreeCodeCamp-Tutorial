package main

// func main() {
//Arrays - Require set length. must input integer or .../
// var students [4]string
// fmt.Printf("studetns: %v", students)
// students[0] = "Lisa"
// students[1] = "Bart"
// students[2] = "Maggie"
// fmt.Printf("Students: %v", students)

// var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
// fmt.Println(identityMatrix)

// var identityMatrix [3][3]int
// identityMatrix[0] = [3]int{1, 0, 0}
// identityMatrix[1] = [3]int{0, 1, 0}
// identityMatrix[2] = [3]int{0, 0, 1}
// fmt.Println(identityMatrix)

//arrays are literal copy. not pointing to same data. must re-assign entire length of array

//use pointers
// a := [...]int{1, 2, 3}
// b := &a
// //b will point to same data when using & so both will change
// b[1] = 5
// fmt.Println(a)
// fmt.Println(b)

// //Slices
// a := []int{1, 2, 3, 4, 5, 6, 7}
// b := a[:]   //slice of all elements
// c := a[3:]  //4th element to end
// d := a[:6]  //first 6
// e := a[3:6] //4-6 elements
// a[5] = 42
// //slices are reference types, refer to same data
// fmt.Println(a)
// fmt.Println(b)
// fmt.Println(c)
// fmt.Println(d)
// fmt.Println(e)

// fmt.Printf("Length: %v\n", len(a))
// //capacity = slice-only; # of elements in slice do not have to match slice is only projection of that underlying array.
// fmt.Printf("Capacity: %v\n", cap(a))

// a := make([]int, 3)
// fmt.Println(a)
// fmt.Printf("Length: %v\n", len(a))
// fmt.Printf("Capacity: %v\n", cap(a))

// //Appending slices
// a := []int{}
// fmt.Println(a)
// fmt.Printf("Length: %v\n", len(a))
// fmt.Printf("Capacity: %v\n", cap(a))
// a = append(a, 1)
// fmt.Println(a)
// fmt.Printf("Length: %v\n", len(a))
// fmt.Printf("Capacity: %v\n", cap(a))
// a = append(a, []int{2, 3, 4, 5}...)
// //to concat two slices, place spread operator after the appending element
// fmt.Println(a)
// fmt.Printf("Length: %v\n", len(a))
// fmt.Printf("Capacity: %v\n", cap(a))

// //Pop Operation
// a := []int{1, 2, 3, 4, 5}
// b := a[1:]
// fmt.Println(b)
// c := a[len(a)-1]
// fmt.Println(c)

// //Remove element from slice
// a := []int{1, 2, 3, 4, 5}
// fmt.Println(a) // [1,2,3,4,5]
// b := append(a[:2], a[3:]...)
// fmt.Println(b) // [1,2,4,5]
// fmt.Println(a) // [1,2,4,5,5]
//if you remove an item from a slice, you may get unexpected behavior if you have references to the first array

// // Summary
// Arrays:
// 	fixed size
// 	declaration styles:
// 		a := [3]int{1,2,3}
// 		a := [...]int{1,2,3}
// 		var a [3]int
// 	copies refer to different underlying data
// 	len function

// Slices:
// 	Backed by an array
// 	creation styles:
// 		slice existing array or slice
// 		literl stylesvia make function
// 		a := make([]int,10) //create slice with capacity and length == 10
// 		a := make([]int,10) //create slice with length == 10 and capacity == 100
// 	len, cap, append functions
// 		append may cause odd operations if underlying array is too small
// 	copies refer to same underlying array

// Maps
// statePopulations := make(map[string]int) //good setup for using map with loop
// statePopulations = map[string]int{
// 	"California": 1000000,
// 	"Texas":      5550000,
// 	"Florida":    23412351,
// }
// fmt.Println(statePopulations)
// statePopulations["Georgia"] = 1244666
// fmt.Println(statePopulations["Georgia"]) //prints population #
// fmt.Println(statePopulations)            //adds Georgia to population array
// //Return order is not guaranteed in a map
// delete(statePopulations, "Georgia")    //Deletes Georgia from map,
// fmt.Println(statePopulations["Texas"]) //but if you print Georgia, it will show a 0
// fmt.Println(statePopulations)

// //Use OK syntax to confirm if key is in map or not
// // pop, ok := statePopulations["Texas"]
// // fmt.Println(pop, ok)

// _, ok := statePopulations["Texas"]
// fmt.Println(ok)

// //Refererence to statePopulations still affects it
// sp := statePopulations
// delete(sp, "Texas")
// fmt.Println(sp)
// fmt.Println(statePopulations)
// }

// type Doctor struct {
// 	// Structs - like a class?

// 	Number     int
// 	ActorName  string
// 	Companions []string
// }

// func main() {
//Most practical/Verbose version of struct
// aDoctor := Doctor{
// 	Number:    3,
// 	ActorName: "Jon Pertwee",
// 	Companions: []string{
// 		"Liz Shaw",
// 		"Jo Grant",
// 		"Jane Smith",
// 	},
// }
// fmt.Println(aDoctor)
// fmt.Println(aDoctor.ActorName)
// fmt.Println(aDoctor.Companions[1])

// //Anonymous version of struct - wont be available for export
// aDoctor := struct{ name string }{name: "John Pertwee"}
// fmt.Println(aDoctor)

//structs are value types
// aDoctor := struct{ name string }{name: "John Pertwee"}
// anotherDoctor := aDoctor
// anotherDoctor.name = "Shane Sweeney"
// fmt.Println(aDoctor)       //John..if used and '&' symbol on line 166 as a pointer, it would change both to Shane
// fmt.Println(anotherDoctor) //Shane
// }

//GO DOESN'T HAVE INHERITENCE; DOES NOT SUPPORT OBJECT ORIENTED PRINCIPLES
//USES COMPOSITION
// DIFFERENCE?
//	Inheritence = a cardinal is a bird is an animal
//	Compososition (through Imbedded) = a bird has animal-like characteristics by imbedding the animal struct

// type Animal struct {
// 	Name   string
// 	Origin string
// }
// type Bird struct {
// 	Animal   //Imbed the animal struct right into the bird struct
// 	SpeedMPH float32
// 	CanFly   bool
// }

// func main() {
// b := Bird{}
// b.Name = "Emu"
// b.Origin = "Australia"
// b.SpeedMPH = 48
// b.CanFly = false
// fmt.Println(b) //{{Emu Australia} 48 false}...bird HAS A name/origin

// // Literal syntax
// b := Bird{
// 	Animal:   Animal{Name: "Emu", Origin: "Australia"},
// 	SpeedMPH: 48,
// 	CanFly:   false,
// }
// fmt.Println(b.Name)

// }

//Add Tags!

// type Animal struct {
// 	Name   string `required`
// 	Origin string `max: "100"`
// }

// func main() {
// 	t := reflect.TypeOf(Animal{})
// 	field, _ := t.FieldByName("Name")
// 	fmt.Println(field.Tag) //"required"
// }

//Summary
//Maps:
// Collections of value types that are accessed via keys
//created via literals or via make fn
//Members accessed via [key] syntax
//myMap["key"] = "value"
//check for presence with "value, ok" form of result
//Structs:
//Collections of disparate data types that describe a single concept
//Keyed by named fields
//Normally created as types, but anonymous structs are allowed
//structs are value types
//No inheritence but can use composition via embedding
//Tags can be added to struct fields to describe field

//CONTROL FLOW - IF and SWITCH
// func main() {
// if false {
// 	fmt.Println("The Test is true")
// }

// if pop, ok := statePopulations["Florida"]; ok {
// 	fmt.Println(pop)
// }

//Comparing #'s with decimals
//Below method will round for you
// myNum := 0.123
// if math.Abs(myNum / math.Pow(math.Sqrt(myNum), 2) - 1) < 0.001 {
// 	fmt.Println("These are the same")
// }

//Switch Statements
//Example 1
// switch 2 {
// case 1:
// 	fmt.Println("one")
// case 2:
// 	fmt.Println("two")
// default:
// 	fmt.Println("not one or two")
// } // this switch case will print 'two' because of the # on switch line

// //Example 2
// switch 5 {
// case 1, 5, 10:
// 	fmt.Println("one, five, or ten")
// case 2, 4, 6:
// 	fmt.Println("two, four, six")
// default:
// 	fmt.Println("not 1, 2, 4, 5, 6, or 10")
// } // this switch case will print 'one, five, or ten' because of the # on switch line

// //Example 3
// switch i := 2 + 3; i {
// case 1, 5, 10:
// 	fmt.Println("one, five, or ten")
// case 2, 4, 6:
// 	fmt.Println("two, four, six")
// default:
// 	fmt.Println("not 1, 2, 4, 5, 6, or 10")
// } // this switch case will print 'one, five, or ten' because of i = 2+3

//Example 4 - Tagless Syntax
// i := 10
// switch {
// case i <= 10:
// 	fmt.Println("10 or less")
// 	fallthrough
// case i <= 20:
// default:
// 	fmt.Println("greater than 20")
// } // this switch case will print '10 or less' because of i = 2+3
//No case numbers, just allowing them to be..Even though the answer is both first and second case, it will execute first one and choose that
//break keyword is implied betwen cases!!!
//what if you want it to continue to the 20 case as well as 10 case? use 'fallthrough'

// 	//Example 5 - Type Switching
// 	var i interface{} = [3]int{}
// 	switch i.(type) {
// 	case int:
// 		fmt.Println("i is an int")
// 	case float64:
// 		fmt.Println("i is a float64")
// 	case string:
// 		fmt.Println("i is a string")
// 	case [3]int:
// 		fmt.Println("i is [3]int")
// 	default:
// 		fmt.Println("i is another type")
// 	}
// }

// func main() {
// for i := 0; i < 5; i++ {
// 	fmt.Println(i)
// }
// for i := 0; i < 5; i = i + 2 {
// 	fmt.Println(i)
// }
// for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
// 	fmt.Println(i, j)
// }
// i := 0 // scoped to main function instead of the for loop
// for i < 5 {
// 	fmt.Println(i)
// 	i++ //can solve the do/while problems moving it down here
// 	if i == 5 { //
// 		break
// 	}
// }
// for i := 0; i < 10; i++ {
// 	if i%2 == 0 {
// 		continue //will skip this index element and continue to next - here runs only odds
// 	}
// 	fmt.Println(i)
// }
// Nested loops
// for i := 1; i <= 3; i++ {
// 	for j := 1; j <= 3; j++ {
// 		fmt.Println(i * j)
// 	}
// }
//Example 2 - Labels
// Loop:
// 	for i := 1; i <= 3; i++ {
// 		for j := 1; j <= 3; j++ {
// 			fmt.Println(i * j)
// 			if i*j >= 3 {
// 				break Loop
// 			}
// 		}
// 	}
// s := []int{1, 2, 3}
// fmt.Println(s) // [1 2 3]

// //for range loop
// s := []int{1, 2, 3}
// for k, v := range s {
// 	fmt.Println(k, v)
// }

// statePopulations := map[string]int{
// 	"Cali":     1111,
// 	"Texas":    2222,
// 	"New York": 3333,
// 	"Penn":     4444,
// }
// for k, v := range statePopulations {
// 	fmt.Println(k, v)
// }

// s := "Hello Go!"
// for k, v := range s {
// 	// fmt.Println(k, v) //will print position and numeric unicode of each letter
// 	fmt.Println(k, string(v)) //will print a string letter
// }

// statePopulations := map[string]int{
// 	"Cali":     1111,
// 	"Texas":    2222,
// 	"New York": 3333,
// 	"Penn":     4444,
// }
// for k := range statePopulations {
// 	fmt.Println(k) //prints only keys (states)
// }

// //What about only values?
// statePopulations := map[string]int{
// 	"Cali":     1111,
// 	"Texas":    2222,
// 	"New York": 3333,
// 	"Penn":     4444,
// }
// for _, v := range statePopulations { // utilize the underscore/write-only variable
// 	fmt.Println(v) //prints only values (#s)
// }

//FOR LOOP Summary
//	simple loops:
// for initializer; test; incrementer {}
// for test {}
// for {}
// Exiting early
//break - breaks out of loop
//continue - doesnt break out of loop, just current iteration
//labels - set up label, with break label. Good for nested loops
//Looping over collections
// for k, v := range collection {}

// //Defer
// fmt.Println("start")
// defer fmt.Println("middle")
// fmt.Println("end")

// //Defer runs in LIFO
// defer fmt.Println("start")
// defer fmt.Println("middle")
// defer fmt.Println("end") // end middle start

// res, err := http.Get("http://www.google.com/robots.txt")
// if err != nil {
// 	log.Fatal(err)
// }
// robots, err := ioutil.ReadAll(res.Body)
// res.Body.Close()
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Printf("%s", robots)

// a := "start"
// defer fmt.Println(a) //takes in accepted data at time of defer - will hold.
// a = "end"

//Panic
// a, b := 1, 0
// ans := a / b
// fmt.Println(ans) // cannot run this operation in Go; run time will show a panic.

// //Built in Panic
// fmt.Println("start")
// panic("something bad happened")
// fmt.Println("end")

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello Go"))
// })
// err := http.ListenAndServe(":8080", nil)
// if err != nil {
// 	panic(err.Error())
// }

// fmt.Println("start")
// fmt.Println("start")
// panicker()
// fmt.Println("end")
// }

// func panicker() {
// 	fmt.Println("about to panic")
// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Println("Error:", err)
// 			panic(err)
// 		}
// 	}()
// 	panic("something bad happened")
// 	fmt.Println("done panicking")
// }

// SUMMARY
// Defer
// 	useful to delay executuion of a statement until function exists
// 	useful to group "open" and "close" functions together
// 		be careful in loops
// 	Run in LIFO (last in first out) order
// 	Arguments evaluated at time defer is executed, not at time of called function execution
// Panic
// 	Occur when program cannot continue at all
// 		Dont use when file cant be opened, unless it is critical
// 		use for unrecoverable events - cannot obtain TCP port for web server
// 	Function will stop executing
// 		 Deferred functions will still fire even after panic
// Recover
// 	Used to recover from panics
// 	Only used in deferred function that will look for panic situation; recover will decide how to proceed.
// 	Current function will not attempt to continue, but higher functions in call stack will.
//

// Pointers!
// func main() {

// a := 42
// b := a
// fmt.Println(a, b) // 42 42
// a = 27
// fmt.Println(a, b) // 27 42

// var a int = 42
// var b *int = &a // 'b' is pointing to an integer, and that integer is 'a'. * declares the pointer
// // fmt.Println(a, b)  // 42 0xc000016070...'b' isnt holding the value 42, it is holding the memory location that is holding 'a's data.
// // fmt.Println(&a, b) // 0xc000016070 0xc000016070
// fmt.Println(a, *b) // 0xc000016070 0xc000016070. Finds memory location and pulls out that data
// a = 27
// fmt.Println(a, *b) //27 27
// *b = 14
// fmt.Println(a, *b) //14 14

// a := [3]int{1, 2, 3}              // array of 1, 2, 3
// b := &a[0]                        // pointer to memory location of index 0
// c := &a[1]                        // pointer to memory location of index
// fmt.Printf("%v %p %p\n", a, b, c) //[1 2 3] 0xc0000b6000 0xc0000b6008...cannot use arithmetic like subtract 8 from 'c' to get 'b's memory location

// var ms myStruct
// ms = myStruct{foo: 42}
// fmt.Println(ms) // will print {42}
// }

// 	var ms *myStruct
// 	ms = &myStruct{foo: 42}
// 	fmt.Println(ms) // will print &{42}
// }

// var ms *myStruct
// ms = new(myStruct)
// (*ms).foo = 42         //de-referencing operator has a lower precedence than dot operator; need params
// fmt.Println((*ms).foo) // will print out 42

// //but that ^ is messy!! Try this:
// var ms *myStruct
// ms = new(myStruct)
// ms.foo = 42
// fmt.Println(ms.foo) // will print out 42!
// // pointer ms doesnt really have field foo on it; pointer ms is pointing to a structure that has foo
// a := [3]int{1, 2, 3}
// b := a
// fmt.Println(a, b) //[1 2 3] [1 2 3]
// a[1] = 42
// fmt.Println(a, b) //[1 42 3] [1 2 3]

// //but if you remove the index and turn the array into a slice; a and b will both change
// a := []int{1, 2, 3}
// b := a
// fmt.Println(a, b) //[1 2 3] [1 2 3]
// a[1] = 42
// fmt.Println(a, b) //[1 42 3] [1 42 3]
// // ..in an array, its values are considered intrinsic to the variable. 'a' is holding variables and size.
// // However a slice doesnt contain the data itself, a slice contains a pointer to the first element that the slice projects of the underlying array

// //Maps and pointers
// a := map[string]string{"foo": "bar", "baz": "buz"}
// b := a
// fmt.Println(a, b) //map[baz:buz foo:bar] map[baz:buz foo:bar]
// a["foo"] = "qux"
// fmt.Println(a, b) //map[baz:buz foo:qux] map[baz:buz foo:qux]
// //maps have pointer to underlying data, not data themselves
// //when working with slices and maps in Go, keep in mind who has access to that underlying data
// //because passing maps and slices around can get you into situations where data may change
// }

// type myStruct struct {
// 	foo int
// }
//Pointers Summary
//Creating Pointers
//	Pointer types use an asterik(*) as a prefix to type pointed to
//			i.e.: *int - a pointer to an integer
// 		Use the addressof operator (&) to get address of variable in memory location.
//	De-reference pointers
//		Dereference a pointer by preceding with an asterik(*)
//		Complex types (i.e. structs) are automatically dereferenced.
//	Create pointers to objects
//		Can use the addressof operator (&) if value type already exists
//		ms := my struct{foo:42}
//		p := &ms
//	Use addressof operator before initializer
//		&mystruct{foo:42}
//	Use the new keyword
//		Can't initialize fields at the same time
//	Types with internal pointers
//		All assignment operations in Go are copy operations.
//			if we create a and set b to a, all data in a is going to be copied to create b
//		Slices and maps contain internal pointers, so copies point to same underlying data

//FUNCTIONS!!!!!
// Example 1
// func main() {
// 	sayMessage("HEllo GO!!")
// }

// func sayMessage(msg string) {
// 	fmt.Println(msg)
// }

// Example 2
// func main() {
// 	for i := 0; i < 5; i++ {
// 		sayMessage("Hello Go!", i)
// 	}
// }

// func sayMessage(msg string, idx int) {
// 	fmt.Println(msg)
// 	fmt.Println("The value of the index is", idx)
// }

// Example 3
// func main() {
// 	for i := 0; i < 5; i++ {
// 		sayMessage("Hello Go!", i)
// 	}
// }

// func sayMessage(msg string, idx int) {
// 	fmt.Println(msg)
// 	fmt.Println("The value of the index is", idx)
// }

// Example 4
// func main() {
// 	sayGreeting("Hello", "Riley")
// }

// func sayGreeting(greeting, name string) { // when data type is the same, we can write it at the end, compiler will infer all the same
// 	fmt.Println(greeting, name)
// }

// // Example 5
// func main() {
// 	greeting := "hello"
// 	name := "Riley"
// 	sayGreeting(greeting, name)
// 	println(name) // Riley - data not change even when name is passed into another function
// }

// func sayGreeting(greeting, name string) { // when data type is the same, we can write it at the end, compiler will infer all the same
// 	fmt.Println(greeting, name) // hello Riley
// 	name = "Shane"
// 	fmt.Println(greeting, name) // hello Shane

// }

// // Example 6
// func main() {
// 	greeting := "hello"
// 	name := "Riley"
// 	sayGreeting(&greeting, &name)
// 	println(name) //shane - accepts the pointer so now instead of copying, we are pointing to main variable. changed variable in scope of main and sayGreeting
// }

// func sayGreeting(greeting, name *string) {
// 	fmt.Println(*greeting, *name) //hello Riley
// 	*name = "Shane"
// 	fmt.Println(*name) // shane
// 	//Why? a lot of times our functions need to act on the paramters that are passed into them,
// 	//pointers only way to do that. and much more efficient to pass in pointers than value of data
// }

// //Example 7
// func main() {
// 	sum(1, 2, 3, 4, 5)
// }

// func sum(values ...int) {
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	fmt.Println("The sum is ", result)
// } 	//Can also be written like this:

// func main() {
// 	sum("The sum is", 1, 2, 3, 4, 5)
// }
// func sum(msg string, values ...int) {
// 	fmt.Println(values) // [1 2 3 4 5]
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	fmt.Println(msg, result) // The sum is 15
// }

// func main() {

// 	s := sum(1, 2, 3, 4, 5)
// 	fmt.Println("The sum is", s)
// }
// func sum(values ...int) int { //expect to return an integer, so write the type here
// 	fmt.Println(values) //
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	return result
// }

// func main() {

// 	s := sum(1, 2, 3, 4, 5)
// 	fmt.Println("The sum is", *s) //add pointer; still prints The sum is 15
// }
// func sum(values ...int) *int { //expect to return an integer, so write the type here
// 	fmt.Println(values) // still prints [1 2 3 4 5]
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	return &result //add address of result
// 	//In GO, when it recognizes that you are returning a value that's generated on the local stack,
// 	//its automatically gonna promote this variable for you to be on the shared memory of the computer aka the heath memory.
// }

// //Named Values - can get confusing to read - since return variables named at the top, while return action is below.
// func main() {

// 	s := sum(1, 2, 3, 4, 5)
// 	fmt.Println("The sum is", s) //add pointer; still prints The sum is 15
// }
// func sum(values ...int) (result int) { //declaring a 'result' variable which will be implicitly returned.
// 	fmt.Println(values)
// 	for _, v := range values {
// 		result += v
// 	}
// 	return
// }

// func main() {
// 	d := divide(5.0, 3.0)
// 	fmt.Println(d) // 1.666666
// }

// func divide(a, b float64) float64 {
// 	return a / b
// }

// func main() {
// 	d, err := divide(5.0, 2.0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(d)
// }

// func divide(a, b float64) (float64, error) { // can add second return variable. Can return as many as you want in a function
// 	if b == 0.0 {
// 		return 0.0, fmt.Errorf("Cannot divide by 0")
// 	}
// 	return a / b, nil
// }

// //A very simple Anonymous function
// func main() {
// 	func() {
// 		fmt.Println("Hello GO")
// 	}() //makes it an immediately defined function
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		func(i int) {
// 			fmt.Println(i)
// 		}(i)
// 	}
// }

// func main() {
// 	f := func() {
// 		fmt.Println("Hello")
// 	}
// 	f()
// }

// //divide function called within main fn instead of globally
// func main() {
// 	var divide func(float64, float64) (float64, error)
// 	divide = func(a, b float64) (float64, error) {
// 		if b == 0.0 {
// 			return 0.0, fmt.Errorf("Cannot divide by zero")
// 		} else {
// 			return a / b, nil
// 		}
// 	}
// 	d, err := divide(5.0, 0.0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(d)
// }

// //A Method!! provides context that that function is executing in
// func main() {
// 	g := greeter{
// 		greeting: "Hello",
// 		name:     "Go",
// 	}
// 	g.greet() // method invocation
// 	fmt.Println("The new name is", g.name)
// }

// type greeter struct {
// 	greeting string
// 	name     string
// }

// func (g greeter) greet() { //specifying greeter as a value type: this is known as a Value Receiever
// 	fmt.Println(g.greeting, g.name)
// 	g.name = "" // operating on the copy of the greeter object not the greeter itself
// }

// //pointer receiver
// func (g *greeter) greet() { //specifying greeter as a value type: this is known as a Value Receiever
// 	fmt.Println(g.greeting, g.name)
// 	g.name = "" // operating on the copy of the greeter object not the greeter itself
// }

//Functions Summary
//comma delimited list of variables and types
//	func foo(bar string, baz int)
//parameters of same type list type once
//	func foo(bar, baz int)
//when pointers are passed in, the function can change the value in the caller
//	this is always true for data of slices and maps
//use variadic parameters to send list of same types in
//	must be last parameter
//	received as a slice
//	func foo (bar string, baz ...int)
//return values
// 	single return values just list ype
//		func foo() int
// 	multiple return value list types surrounded by parenthesis
// 		func foo() (int, error)
// 		the (result tupe, error) paradigm is a very common idiom
// 	can use named return values
// 		initializes returned variables
// 	can return addresses of local variables
// 		automatically promoted from local memory (stack) to shared memory
//
// Anonymous Functions
// 	dont have name if they are immediately invoked
// 	assigned to a variable or passed as an argument to a function
//
// Functions as types
// 	can assign functions to variables or use as arguments and return values in functions
// 	type sugnatures is like function signature, with no parameter names
// 		var f func(string, string, int) (int, error)
//
// Methods
// 	Function that executes in context of a type
//  Format:
// 		func (g greeter) greet() {
// 			...
// 		}
// Reveiver can be value or pointer
// 	Value receiver gets copy of type
// 	Pointer receiver gets pointer to type
//
//
//
//
//
// Interfaces - interfaces usually take on '-er' after the method (write vs writer)
//
// Example 1
// func main() {
// 	var w Writer = ConsoleWriter{}
// 	w.Write([]byte("Hello GOGO"))
// }
//
// type Writer interface {
// 	Write([]byte) (int, error)
// }
// type ConsoleWriter struct{}

// func (cw ConsoleWriter) Write(data []byte) (int, error) {
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }
//
// // Example 2
// func main() {
// 	myInt := IntCounter(0)
// 	var inc Incrementer = &myInt
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(inc.Increment())
// 	}
// }

// type Incrementer interface {
// 	Increment() int
// }

// type IntCounter int

// func (ic *IntCounter) Increment() int {
// 	*ic++
// 	return int(*ic)
// }
//
// // Example 3
// func main() {
// 	var wc WriterCloser = NewBufferedWriterCloser()
// 	wc.Write([]byte("Hello Youtube Listeners, this is a test"))
// 	wc.Close()

// 	// bwc := wc.(*BufferedWriterCloser)
// 	// fmt.Println(bwc)
// 	r, ok := wc.(*BufferedWriterCloser)
// 	if ok {
// 		fmt.Println(r)
// 	} else {
// 		fmt.Println("Conversion failed")
// 	}
// }

// // Example 4
// func main() {
// 	var myObj interface{} = NewBufferedWriterCloser()
// 	if wc, ok := myObj.(WriterCloser); ok {
// 		wc.Write([]byte("Hello Youtube Listeners, this is a test"))
// 		wc.Close()
// 	}

// 	r, ok := myObj.(io.Reader)
// 	if ok {
// 		fmt.Println(r)
// 	} else {
// 		fmt.Println("Conversion failed")
// 	}
// }

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type Closer interface {
// 	Close() error
// }

// type WriterCloser interface {
// 	Writer
// 	Closer
// }

// type BufferedWriterCloser struct {
// 	buffer *bytes.Buffer
// }

// func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
// 	n, err := bwc.buffer.Write(data)
// 	if err != nil {
// 		return 0, err
// 	}
// 	v := make([]byte, 8)
// 	for bwc.buffer.Len() > 8 {
// 		_, err := bwc.buffer.Read(v)
// 		if err != nil {
// 			return 0, err
// 		}
// 		_, err = fmt.Println(string(v))
// 		if err != nil {
// 			return 0, err
// 		}
// 	}
// 	return n, nil
// }

// func (bwc *BufferedWriterCloser) Close() error {
// 	for bwc.buffer.Len() > 0 {
// 		data := bwc.buffer.Next(8)
// 		_, err := fmt.Println(string(data))
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func NewBufferedWriterCloser() *BufferedWriterCloser {
// 	return &BufferedWriterCloser{
// 		buffer: bytes.NewBuffer([]byte{}),
// 	}
// }

// func main() {
// 	var i interface{} = "0"
// 	switch i.(type) {
// 	case int:
// 		fmt.Println("i is an integer")
// 	case string:
// 		fmt.Println("i is a string")
// 	default:
// 		fmt.Println("I don't know what i is!!!")
// 	}
// }
//
// func main() {
// 	var wc WriterCloser = myWriterCloser{}
// 	fmt.Println(wc)
// }

// type Writer interface {
// 	Write([]byte) (int, error)
// }
// type Closer interface {
// 	Close() error
// }

// type WriterCloser interface {
// 	Writer
// 	Closer
// }
// type myWriterCloser struct{}

// func (mwc myWriterCloser) Write(data []byte) (int, error) { //Cannot make myWriterCloser a pointer...Write method has a pointer receiver; because you are using myWriterCloser directly; could add & to line 1001
// 	return 0, nil
// }
// func (mwc myWriterCloser) Close() error {
// 	return nil
// }
//
//
// Interface Best Practices
// Use many, small interfaces - preferable in long run than a few monolithic ones
// 	Single method interfaces are some of the most powerful and flexible
// 		io.Writer, io.Reader, interface {} (all have 1 or 0 methods on them)
// Dont export interfaces for types that will be consumed.
// DO export interfaces for types that will be used by package
//
// Interfaces Summary
// 	Basics
// 		type Writer interface {
// 			Write([]byte)(int,error)
// 		}
//
// 		type ConsoleWriter struct {}
//
// 		func (cw ConsoleWriter) Write(data []byte) (int, error) {
// 			n, err := fmt.Println(string(data))
// 			return n, err
// 		}
// 	Composing Interfaces
// 		type Writer interface {
// 			Write([]byte)(int,error)
// 		}
//
// 		type Closer interface {
// 			Close() error
// 		}
//
// 		type WriterCloser interface {
// 			Writer
// 			Closer
// 		}
// 	Type Conversion
// 		var wc WriterCloser = NewBufferedWriterCloser()
// 		bwc := wc.(*BufferedWriterCloser)
// 	Empty Interface and Type Switches
// 		var i interface{} = 0
// 		switch i.(type) {
// 			case int:
// 				fmt.Println("i is an integer")
// 			case string:
// 				fmt.Println("i is a string")
// 			default:
// 				fmt.Println("I don't know what i is!!!")
// 		}
// 	Implementing with values and pointers
// 		Method set of value is all methods with value receivers
// 			If you are going to try and impolement an interface with a value type, then all of the methods of that interface have to have receivers
// 		Method set of pointer is all methods, regardless of receiver type
// 			Gives access to underlying data of that type.
