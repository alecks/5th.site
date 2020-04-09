# 5G

![Go Build](https://github.com/elitisgroup/5G/workflows/Go%20Build/badge.svg)

A website in favour of 5G; we love fake news.
Written in Go with its template lib.

## Contributors

You're able to contribute by making PRs (pull requests). Make sure to follow the testing section below.

To add a question/answer to the homepage, create a new file in the `points` directory.
Make sure the contents follow that of `points/example.md`.

## Testing

Making a contribution requires testing three things:
- Frontend - checking if the website looks okay;
- Build - checking if the Go code builds;
- Security - making sure no LGTM errors are outputted.

The last two - build and security checking - are done automatically when you create a PR (pull request).

Frontend testing, however, needs to be done by you. The instructions to do this are as follows:
1. Install Go - you can download it [here](https://golang.org/dl);
2. Run `go get .` in the repository's directory (assuming that you cloned it);
3. Run `go run .` in the repo directory. You may need to run this as the superuser (`sudo`);
4. Visit http://localhost, check if the page looks alright, then screenshot it.

## Production

While hosting for production isn't exactly required in the slightest,
you're able to easily with [Docker](https://hub.docker.com/r/alexeek/5g).

Run the following command:

```shell script
$ docker run -d --rm -p 80:80 --name 5g alexeek/5g
```

Please note that the Docker image can be delayed by up to 15 minutes.
