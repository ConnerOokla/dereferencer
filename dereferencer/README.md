# Dereferencer

This is a dereferencer utility that will take in a struct, remove the pointers (hence the dereferencing), and return a struct with the values.

Works for nested structs.

It uses reflect package
[Why should we not use reflection](https://stackoverflow.com/questions/34385735/should-the-usage-of-reflection-be-avoided-in-go)
