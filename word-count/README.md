# Word count CLI  App

## About

This CLI app is a replika of *wc* in Unix OS 
## How to use
- Start by cloning this repository and navigate to the word-count app directory
```bash
  git clone <url> && cd word-count
```

- Run the application
```bash
  go run . [option] -file [filename]
```

#### Example:
```bash
  go run main.go -w test.txt
```
- The above command will read the **test.txt** file and count the words inside.include:


Other Options 
- `-c` to count bytes
- `-l` to count lines
- `-m` to count characters 

#### Run test
```bash
  go test
```

## Things to note
- You should have go installed


Enjoy <3
