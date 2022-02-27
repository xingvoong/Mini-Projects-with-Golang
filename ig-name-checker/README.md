# Instagram username checker

a scipt to check whether an IG username exist

## Author
- [@xingvoong](https://github.com/xingvoong)

## Scope
- To focus on the Go language itself, concurrency, channels, and for the scope and simplicity of this project, a username is invalid when there is at least one space
in the given string
- To enhance the project, one could use IG API to check for a username.

## Run Locally
Check required tech is installed (see below).

Clone the project
```bash
https://github.com/xingvoong/Mini-Projects-with-Golang
```
Go to the project directory
```bash
cd ig-name-checker
```

Run the modules
```bash
go run main.go
```

## Expected Output
```
No concurrency
[200] IG username barackobama exists 
[200] IG username MichelleObama exists 
[200] IG username elonmusk exists 
IG username 12fdfvscc. does not exist 
IG username xinvoongacxczxccxxc  12e3e does not exist 
3.838212838s

Wait Group
IG username xinvoongacxczxccxxc  12e3e does not exist 
IG username 12fdfvscc. does not exist 
[200] IG username barackobama exists 
[200] IG username elonmusk exists 
[200] IG username MichelleObama exists 
1.210572362s

Channel
IG username xinvoongacxczxccxxc  12e3e does not exist 

[200] IG username elonmusk exists
[200] IG username MichelleObama exists
[200] IG username barackobama exists
[200] IG username 12fdfvscc. exists
473.985011ms
```
## Requirements
- Golang installed

## Acknowledgement
- https://medium.com/wesionary-team/practical-example-of-concurrency-on-golang-fc4609ea8ed1
