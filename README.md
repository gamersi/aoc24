# Advent of Code Template in Go

This is a template to put advent of code problems into a single executable. It
allows you to measure the execution time of each part of each problem, and to
ignore some of the boilerplate. Not sure if anyone else will be interested, but
I thought I would give it a try because I have seen a lot of templates like
this.

# Instructions

To start a new day, generate code from a template using the command:

```
make start
```

It will default to adding code for the current day. This will create a
directory called `dayXXp1` where `XX` is replaced by the day number. Inside
will be a file called `solution.go` with a `Solve` function in which to put
your solution, and a `solution_test.go` file to write your tests. The `Solve`
function takes an `io.Reader` argument which is the input and returns a
solution which can be any type.

If you wish to start a problem for a specific day, say the 21st, you can create
the desired directory from the template by using the make command to create the
part 1 directory for that day using the command below.

```
make day21p1
```

To run the last code you worked on use the command:

```
make run
```

This will generate a `run.go` file and run the most recently modified code. You
can run all the days with the command:

```
make runall
```

Finally, you can do tests with the command:

```
make test
```
