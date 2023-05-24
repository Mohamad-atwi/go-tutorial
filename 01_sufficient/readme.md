### pass by value vs by reference

1. Add a new type called ContactInfo with fields email and address (both of type string) and add it to person struct.
2. Create a new function called updateAddress with person as receiver (so that you can call yourPerson.updateAddress("new address")).
3. Adjust initialization of jim in the main function.
4. Change jim's address using your new function. Did the address change? Why/why not?
5. Finally, change the print function so that the field names are added (https://pkg.go.dev/fmt)