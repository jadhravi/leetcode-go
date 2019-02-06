[comment]: # (TODO: Add copyright)

# Readme
I am trying to learn [Go](https://golang.org/) by solving problems on [leetcode](https://leetcode.com/problemset/all/).

# Structure
1. Each problem has its own folder
2. Each folder contains two files, a.solution.go and b.solution_test.go
3. Run all the tests by using `go test` in each folder
4. I will also include any additional experimental solutions in separate files.
# Problems solved so far
1. [709-ToLowerCase](https://leetcode.com/problems/to-lower-case/) : Simple lowercase implementation.
2. [977-SortedSquares](https://leetcode.com/problems/squares-of-a-sorted-array/):  Simple implementation using built in sort. Also added a concurrent solution to play with and [go routines](https://tour.golang.com/concurrency/1) and [channels](https://tour.golang.com/concurrency/2)
3. [804-UniqueMorseRepresentations](https://leetcode.com/problems/unique-morse-code-words/) : Morse code implementation. This solution uses [strings.Builder](https://golang.org/pkg/strings/#Builder) and consumes more memory than submissions that include [slice append](https://tour.golang.org/moretypes/15) and [strings.Join](https://golang.org/pkg/strings/#Join) function. Although the submission passed, need to add more test cases.  
