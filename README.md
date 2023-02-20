# Go-Cat
Go-Cat is a command-line utility that reads one or more files and prints them to the standard output. It is similar to the Unix cat command, but with additional options for displaying non-printing characters and controlling the formatting of the output.

## License
Go-Cat is licensed under the Apache License, Version 2.0.

## Copyright
© 2023

## Usage
`go-cat [OPTION]... [FILE]...`
Go-Cat takes zero or more options, followed by zero or more file names. If no file names are provided, it reads from the standard input. If multiple files are provided, their contents are concatenated in the order specified.

The available options are:

`-A` : equivalent to `-vET`

`-E` : display $ at the end of each line

`-T` : display TAB characters as `^I`

`-s` : squeeze multiple blank lines into one

# Examples
To display the contents of a file named example.txt, simply run:

`go-cat example.txt`

To display the contents of a file with non-printing characters, use the -A options:

`go-cat -A example.txt`

To display the contents of multiple files, run:

`go-cat file1.txt file2.txt file3.txt`

To display the contents of a file and number each line, use the cat command in combination with the nl command:

`go-cat example.txt | nl`

# Implementation
Go-Cat reads the contents of each file specified on the command line, concatenates them, and prints them to the standard output. It reads the files in chunks of 4KB, and can handle large files without running out of memory.

Go-Cat also provides additional functionality for displaying non-printing characters and controlling the formatting of the output, including the ability to display TAB characters and end each line with a $ symbol.