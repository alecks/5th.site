# 5G

![Go Build](https://github.com/elitisgroup/5G/workflows/Go%20Build/badge.svg)

A website in favour of 5G; we love fake news.
Written in Go with its template lib.

## Contributors

You're able to contribute by making PRs (pull requests).

Sections are called "points".
Simply add a markdown file to the /points folder;
make sure it follows points/example.md's style.

## Hosting

While hosting the website isn't exactly required in the slightest,
you're able to easily with Docker.

```shell script
$ docker run -it --rm -p 80:80 --name 5g alexeek/5g
```