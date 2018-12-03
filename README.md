# blackfriday-slack
A [Blackfriday](https://github.com/russross/blackfriday) v2 Renderer which translates markdown to slack styling

[![godoc](https://img.shields.io/badge/godoc-reference-orange.svg?style=flat-square)](https://godoc.org/github.com/karriereat/blackfriday-slack)
[![license](https://img.shields.io/github/license/karriereat/blackfriday-slack.svg?style=flat-square)](https://github.com/karriereat/blackfriday-slack/blob/master/LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/karriereat/blackfriday-slack)](https://goreportcard.com/report/github.com/karriereat/blackfriday-slack)

## Installation
```
$ go get -u github.com/karriereat/blackfriday-slack
```

## Examples

### Input
i had to mask the code block for this example
```
# head1

## head2
- list1
- list2
- list3

### head3
* list4
* list5
* list6

### head3
* list1
  * list 2
  * list 3
    * list 4
    * list 5
      * list 6
      * list 7
  * list 8
* list 9

### head4
1. list1
2. list2
3. list3
  1. list4
  2. list5
4. list6

---
`code`
---

\``` go
    code block 
    such code block - much wow
```\

```

### Output
![output image](https://github.com/karriereat/blackfriday-slack/blob/master/output.png)

## Documentation
[GoDoc](https://godoc.org/github.com/karriereat/blackfriday-slack)


## Thanks
Blackfriday-Slack is heavily inspired by [Blackfriday-Confluence](https://github.com/kentaro-m/blackfriday-confluence)


## License
[Apache 2.0](https://github.com/karriereat/blackfriday-slack/blob/master/LICENSE)