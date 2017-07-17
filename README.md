## findup
findup searches for duplicate files in given directories. It scans all the directories given
as command line arguments and prints list of duplicate files on standard out.

findup is written in Go.

### Installation
Install and update with go: 
`go get -u github.com/jasilven/findup`

### Usage
To search and get report of duplicate files in directories `dir1, dir2` and `dir3`
run: `$GOPATH/bin/findup dir1 dir2 dir3`

### Dependencies
No dependencies.
