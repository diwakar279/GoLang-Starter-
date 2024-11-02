# Steps for installation

1. [Install go](go.dev)

2. Install a extension if you are using VS Code -> GO extension

3. Run "go install -v golang.org/x/tools/cmd/goimports@latest" to install.

# Variables

1. **Using var keyword**

    *var variable_name type = value*
```
   var name String ="bye"
```

You can also do this 
```
   var name = "bye"
```
but there is a catch You can not do declaration like this
```
   var name
```
Error :  Declared and not used variable will give u error

2. **Short Variable Declaration**

```
  Price:= 234
  fmt.Println(Price)
```
Please don't confuse in between := and = as := is a declaration and = is assignment


# Loops

-	There is no **while loop** in goLang 

- only **for loop exists**
	
- You can also use **break & continue** in goLang  and it works just like in any other language


# If Else

*There is no Ternary Operator in goLang till I am writing this go1.23.2*

Code Structure-
```
if condition {
	//Code
} else if condition{
   //Code
} else{
   //Code
}
```

It will give **syntax error** : **unexpected else** ,  if you write like this -
```
if condition {
	//Code
} 
else{
   //Code
}

OR

if condition {
	//Code
} 
else if condition{
   //Code
} 
```


# Switch

- Same use case just like in other languages

- default case is available

- multiple condition could check in single case

- For no Reason, Here Switch is used to do type check unlike any other language 

:smile: