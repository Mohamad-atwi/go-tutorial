### Interfaces

1. Create custom type (struct) which will come as a first argument in io.Copy function.
2. Make your type implement the required interface (https://pkg.go.dev/io#Copy). You can implement this function however
you want, but we suggest you simply print something (eg. `string(bs)` to the console - this might be a little hint ;) )
3. Run the code and check if it works!