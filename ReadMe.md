// About GO
1. statically typed language
2. strongly typed language
3. GO is compiled
4. Fast compile time
5. Built in concurrency [GO Routines]
6. Simplicity (automatic garbage collection)

// core concepts
1. Variables and basic data types
2. Functions and control structures
3. Arrays, slices, Maps, loops
4. Strings, Runes, Bytes
5. Structs, Interfaces

// main concepts
1. Pointers
2. Goroutines
3. Channels
4. Generics

// import "strings"

strings.ToUpper("go")         // "GO"
strings.ToLower("GO")         // "go"
strings.Contains("hello", "he") // true
strings.HasPrefix("hello", "he") // true
strings.HasSuffix("hello", "lo") // true
strings.Split("a,b,c", ",")   // []string{"a", "b", "c"}
strings.Join([]string{"a", "b"}, "-") // "a-b"
strings.ReplaceAll("go go go", "go", "run") // "run run run"
