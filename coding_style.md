# Coding style

## History

Grammalecte is a fork of Lightproof which used a lot of different ways to name variables.
An attempt to unify them has been made. Though not perfect, we try to follow the same coding style in all files of this project.

As a software written in Python and JavaScript and mainly provided as extensions for other softwares, it’s not always possible to follow the same principles.


## Variable names

### Style

* Local variable: lowerCamelCase (even in Python).
* function name: lowerCamelCase (even in Python) or lowercase.
* attributes: lowerCamelCase (even in Python) or lowercase.
* class name: UpperCamelCase
* Module name: lowercase_separated_by_underscores

One-letter names should be avoided, though tolerated for specific usages, as long as the variable is used only for few lines.


### Hungarian notation

As much as possible, we use the Hungarian notation for variable names, meaning we add one (or several) lowercase letters as a prefix to each variables. These prefixes are mainly a way to remember the type of the variables.

```
b:      boolean

i:      integer (mainly for index)
j:      integer (mainly for index)
k:      integer (mainly for index)
n:      integer (count)

f:      float

c:      string (one char)
s:      string
sf:     string (filename, no path)
sp:     string (path, folder only)
spf:    string (path, folder and filename)

by:     binary string or array

z:      compiled regular expression (with re.compile() or new RegExp() or /.../)
m:      match object

d:      dictionary (or Map in JavaScript)

l:      list or array
a:      list or array or set
t:      tuple

o:      object defined by our project
x:      object defined by 3rd part software (library, DOM element, etc.)

e:      event or error object
h:      file handler
```
