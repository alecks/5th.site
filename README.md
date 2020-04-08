# 5G

![Go Build](https://github.com/elitisgroup/5G/workflows/Go%20Build/badge.svg)

A website in favour of 5G; we love fake news.
Written in Go with its template lib.

## Contributors

You're able to contribute by making PRs (pull requests).

To add a question/answer to the homepage, create a new file in the `points` directory.
Make sure the contents follow that of `points/example.md`.

## Hosting

While hosting the website isn't exactly required in the slightest,
you're able to easily with [Docker](https://hub.docker.com/r/alexeek/5g).

Run the following command:

```shell script
$ docker run -it --rm -p 80:80 --name 5g alexeek/5g
```