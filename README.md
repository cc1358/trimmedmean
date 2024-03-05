The complete steps to use this package are quite elementary. 

First, create a new Go module on the terminal with **go mod init projectname** 

Then, create a new .go file in that directory, and in the terminal enter **go get github.com/cc1358/trimmedmean** 

Now, in the project imports section, enter **import github.com/cc1358/trimmedmean** 

Finally, you can create instances of the **'TrimmedMean'** struct and use its **'Calculate'** method in the Go code. 

In the repository, we can see the use of the package by using samples of at least 100 integers and 100 floats to test the Go trimmed mean function against symmetric trimming results from the base mean function in R with 0.05 observations taken from both ends of the distribution: **mean(x, trim = 0.05)**. The comparasion can be seen in trimmedmeanimplementationpackage.go and trimmedmeanimplementation.R. 
