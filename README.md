# Guess It 2

**Guess It 2** is an advanced command-line tool written in Go designed to read numerical data from standard input, compute statistical measures, and predict the range of the next number using linear regression and correlation analysis. 

This tool enhances accuracy by combining multiple statistical methods.

## Table of Contents

- [Installation](#installation)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Testing](#testing)
- [Contributing](#contributing)
- [Author](#author)

## Installation

To set up **Guess It 2**, ensure you have Go installed on your system:

1. Clone the repository:
   ```sh
   git clone https://learn.zone01kisumu.ke/git/shfana/guess-it-2.git
   cd guess-it-2
   ```
2. Build the project:

 ```sh
 go build -o guess-it-2 . 
 ```

 ## Project Structure

- **student/main.go**: The main entry point of the program.

- **student/readData/readdata.go**: Manages reading user input from standard input.

- **student/calculations/mean.go**: Computes the mean (average) of a slice of float64 numbers.

- **student/calculations/stddev.go:** Calculates the variance and standard deviation.

- **student/calculations/range.go**: Uses regression and correlation to calculate the prediction range.

- **student/calculations/predictNextValue.go**: Predicts the next value in the sequence using linear regression.

- **student/calculations/pearsonsCorrelation.go**: Computes the Pearson correlation coefficient.

- **student/calculations/linearRegression.go**: Performs linear regression to find the best-fit line and its parameters.

-  there are also unit tests that test the various functions of the program

-  to run the tests navigate to the calulations directory and run the following command

```bash
cd calculations
go test 
```

 output
 ```
 PASS
ok      guess-it-2/calculations 0.001s
```
- to check the test coverage run :

```
go test -cover
```

output 
```
PASS
coverage: 70.7% of statements
ok      guess-it-2/calculations 0.002s
```


## Usage

1.  Make the script executable:

```sh

chmod +x ./student/script.sh
```

2. Run the program using the provided bash script:

```sh

./student/script.sh
```

 Enter numerical values one by one and press Enter. The program will calculate and display the predicted range of the next number after each entry.

Example usage:

```go

5
8
7  10
```

3. To stop the program, press CTRL+SHIFT+C.

## Testing

To test the project, follow these steps:

1. Download the tester [zip file](https://assets.01-edu.org/guess-it/guess-it-dockerized.zip).
2. Extract the zip file into the root directory of the project.
3. Move the student/ folder into the extracted directory.
4. Follow the instructions in the tester's README file for running the tests.

## Contributions

Contributions are welcome! To contribute:

  1. Fork the repository.
  2. Create a new branch (git checkout -b feature-branch).
  3. Make your changes and commit (git commit -am 'Add new feature').
  4. Push to your branch (git push origin feature-branch).
  5. Create a Pull Request with details of your changes.

## Author

This project was built and maintaned by   [shfana](https://github.com/Wambita)