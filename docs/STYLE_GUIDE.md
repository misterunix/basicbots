# STYLE GUIDE

I am going to try and distille the multi guides in to one document. Starting with **Effective Go** and then **Ubers Style Guide**.

## Editors

I cannot tell you which editor to use, but I do recomend Microsoft's Visual Studio Code. With the Go extentions, many of the format guides are done automaticly.

## Code Formatting

- gofmt gofmt will reformat the source to the GoLang standards and find errors in the source code.

## Commenting
- blocks of comments should be sorounded by `/*` `*/`, single line comments should use `//`. 

## Code Documentation
- Every package should have a package comment, a block comment preceding the package clause. For multi-file packages, the package comment only needs to be present in one file, and any one will do. The package comment should introduce the package and provide information relevant to the package as a whole.

