# Golang-practice

I've returned to picking up Golang, in a more serious way. 

In this repo, I'm uploading various small programs, utilities, challenges that I've written, and/or currently expanding. 

### Why
I'm focused on writing Go code in a way that is idiomatic, performant, and uses the language features in the best way possible.   
I think that commiting the code can help me look back, identify my own pitfalls, allow more exprienced Gophers to offer feedback and document my own 'Gotchas' and 'aha!' moments.

Many codefiles contain a `// Possible Improvements :` comment, on things that should be improved in a future commit, or if the code was to be deployed in a mature environment.

## What code 

## Show me the goodies!

- [Goof](https://github.com/tpaschalis/goof)   
Got my first [Github stars](https://tpaschalis.github.io/a-dozen-github-stars/) with this small project, that me and some colleagues have adopted on our day-to-day workflow.    
*Goof* (Go Offer File) is a small utility to easily serve a single file or directory over a network.  
All you need is Goof, and a command line, or a web browserA small utility to 

- [Timezones overlap](timezones-overlap)  
A quick tool to compare timezones between each other. You can use this to coordinate the 9-to-5 working hours between timezones, or when to watch the San Antonio Spurs from Greece's GMT+2.

- [Producer-Consumer concurrency](producer-consumer)   
A quick and simplistic starting point on how to implement your own [producer-consumer pattern](https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem) for some concurrent efficiency.


### Solutions to [Advent of Code 2018](https://adventofcode.com/)   
Preface : Ipsum Lorem    
All solutions are located [here](AdventOfCode2018).   

### Solutions to [Harvard's CS50](https://online-learning.harvard.edu/course/cs50-introduction-computer-science)

- [Luhn's Algorithm](credit-luhn)   
An implementation of [Luhn's algorihm](https://en.wikipedia.org/wiki/Luhn_algorithm) for AMEX, Visa and Mastercards credit card numbers.

- [Crack DES crypt hash](crack-des-hash)   
Crack some passwords that were hashed and salted with C’s DES-based (not MD5-based) `crypt` function. Correctly uses concurrency to speed up the brute force approach

- [Recover JPEG Images](recover-images)   
Identify JPEG signature from `.raw` file and recover images. Learned how to play with raw bytes in Go, and remembered how 'blocks' and 'file signatures' work.

- [Caesar Cipher](caesar-cipher)   
The classic `rot X` cipher from the Roman times :P learned that ORing an ASCII chat with `0x20` converts uppercase to lowercase but leaves numbers and symbols untouched.

- [Vigenere Cipher](vigenere-cipher)   
An extension of Caesar's Cipher, where you'd `rot` the alphabet through a different letter, based on a `key`.

- [Whodunit??](whodunit)  
Image manipulation to solve a riddle *a-la* Cluedo, by mimicking the iconic red lens.

- [Mario Pyramid](mario-pyramid)   
Prints a double half-pyramid *a-la* 2D-platformer Super Mario obstacles.

- [FizzBuzz](fizzbuzz)   
Because reasons
