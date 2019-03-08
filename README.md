# keyextract

[![Go Report Card](https://goreportcard.com/badge/github.com/Keyology/key-extract)](https://goreportcard.com/report/github.com/Keyology/key-extract)

How to install:

```
run git clone https://github.com/Keyology/key-extract.git inside your go workspace
```

How to use - run the code by entering:

`./main`

once the project is running enter the command `extract` followed by the link.

**_Example_**

```
$keyextract extract <your website link here>

```

the output will look similar to this:

```

$Keyextract extract https://keyology.herokuapp.com/
***THIS IS ARGS:*** [https://keyology.herokuapp.com/]
extracting site....
data extracted
{"status":"success","data":{"author":"Keoni Murray","lang":"en","title":"Keoni Murray::Portfolio","description":"Keoni Murray portfolio website, softeware engineer","url":"https://keyology.herokuapp.com/","image":{"width":688,"height":330,"size":107995,"size_pretty":"108 kB","type":"jpg","url":"https://keyology.herokuapp.com/images/coding.jpg"},"logo":{"width":128,"height":128,"size":8281,"size_pretty":"8.28 kB","type":"png","url":"https://logo.clearbit.com/keyology.herokuapp.com"}}}
exec: "extract": executable file not found in $PATH
$Keyextract


```

How to stop the application enter `exit`

social:

[Twitter](https://twitter.com/Keyology1)
[myWebsite](https://tinyurl.com/yc39f7f4)
