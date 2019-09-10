# taggo
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Frelastle%2Ftaggo.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Frelastle%2Ftaggo?ref=badge_shield)


[![CircleCI](https://circleci.com/gh/relastle/taggo/tree/master.svg?style=shield)](https://circleci.com/gh/relastle/taggo/tree/master)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Frelastle%2Ftaggo.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Frelastle%2Ftaggo?ref=badge_shield)
[![Go Report Card](https://goreportcard.com/badge/github.com/relastle/taggo)](https://goreportcard.com/report/github.com/relastle/taggo)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/a3eaf1d737d54d86b9727477519439c0)](https://www.codacy.com/app/relastle/taggo?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=relastle/taggo&amp;utm_campaign=Badge_Grade)

taggo provides powerful decoration I/O stream. It is mainly designed for giving
powerful features to [fzf](https://github.com/junegunn/fzf) (with `--ansi` option).
It also provides reversion of decorated lines into original lines.

## Installation

```sh
go get -u github.com/relastle/taggo
```
## Usage

taggo can simply decorate table-like inputs like this.

<img width="872" alt="Screen Shot 2019-06-30 at 20 43 22" src="https://user-images.githubusercontent.com/6816040/60396070-bd3fcf80-9b77-11e9-9332-32f48d336999.png">

Furthermore, taggo can revert decorated lines into original lines with a `-r (--revert)` option (other options must be the same as used in decorating phase).

<img width="765" alt="Screen Shot 2019-06-30 at 20 49 44" src="https://user-images.githubusercontent.com/6816040/60396149-a64dad00-9b78-11e9-91ee-09a9a9111421.png">

This reversion feature is very useful when taggo is used as an intermidiate stream of fzf-related tools, because you usually
need non-decorated output lines after they are selected by fzf interface.

## [License](LICENSE)

The MIT License (MIT)


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Frelastle%2Ftaggo.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Frelastle%2Ftaggo?ref=badge_large)