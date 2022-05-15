// package crndm (create randoms) is a rdrand instruction package.
//
// For more information about rdrand, see: https://en.wikipedia.org/wiki/RDRAND
//
// The intended goal is not to generate a fully featured  random package.
// I wanted to create a Golang package that uses CPU instructions to generate random
// numbers and wraps it for simplicity. Rdrand generates random numbers using some
// noise from the CPU (eg silicon noise).
//
// For example, if you want to generate numbers in a range `(Range(0, 100)`, this
// package does not support it. You can implement any random number generation algorithm
// using this package.
//
// `Rdseed` functions may not return true sometimes. These instructions may not work
// under certain conditions as it depends on some hardware features / statuses. Do not
// forget to check the values returned from the functions or the values you pass as a reference.
//
// These intrinsics are already slow in terms of what they do. Additionally, Golang
// **cannot inline assembly**. Do not expect good performance.
// For Golang assembly: https://dave.cheney.net/2019/08/20/go-compiler-intrinsics
package crndm
