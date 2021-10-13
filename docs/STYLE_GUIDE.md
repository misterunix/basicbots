# STYLE GUIDE

I am going to try and distill the multiple guides into one document. Starting with **Effective Go** and then **Ubers Style Guide**.

## Editors

I cannot tell you which editor to use, but I do recommend Microsoft's Visual Studio Code. With the Go extensions, many of the format guides are done automatically.

## Code Formatting

- gofmt will reformat the source to the Go standards and can find many errors in the source code.

## Commenting
- blocks of comments should be surrounded by `/*` `*/`, single-line comments should use `//`. 

## Code Documentation
- Every package should have a package comment, a block comment preceding the package clause. For multi-file packages, the package comment only needs to be present in one file, and anyone will do. The package comment should introduce the package and provide information relevant to the package as a whole.

## MixedCaps {#mixed-caps}

The convention in Go is to use `MixedCaps` or `mixedCaps` rather than underscores to write multiword names.

## Interface names

By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun:`Reader`, `Writer`, `Formatter`, `CloseNotifier` etc.

There are several such names and it's productive to honor them and the function names they capture. `Read`, `Write`, `Close`, `Flush`, `String`, and so on have canonical signatures and meanings. To avoid confusion, don't give your method one of those names unless it has the same signature and meaning. Conversely, if your type implements a method with the same meaning as a method on a well-known type, give it the same name and signature; call your string-converter method `String` not `ToString`.

### Getters {#Getters}

Go doesn't provide automatic support for getters and setters. There's nothing wrong with providing getters and setters yourself, and it's often appropriate to do so, but it's neither idiomatic nor necessary to put `Get` into the getter's name. If you have a field called `owner` (lower case, unexported), the getter method should be called `Owner` (upper case, exported), not `GetOwner`. The use of upper-case names for export provides the hook to discriminate the field from the method. A setter function, if needed, will likely be called `SetOwner`. Both names read well in practice:

    owner := obj.Owner()
    if owner != user {
        obj.SetOwner(user)
    }

