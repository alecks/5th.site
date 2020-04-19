# 5G

![Rust](https://github.com/elitisgroup/5G/workflows/Rust/badge.svg)

A website in favour of 5G; we love fake news.
Written in Rust and Handlebars.

## Contributors

You're able to contribute by making PRs (pull requests). Make sure to follow the testing section below.

To add a question/answer to the homepage, create a new file in the `sections` directory.
Make sure the content follows that of `sections/example.md`.

## Testing

Making a contribution requires testing three things:
- Frontend - checking if the website looks okay;
- Build - checking if the Rust code builds;
- Security - making sure no LGTM errors are outputted.

The last two - build and security checking - are done automatically when you create a PR (pull request).

Frontend testing, however, needs to be done by you. The instructions to do this are as follows:
1. Install Rust - you can get it [here](https://www.rust-lang.org/tools/install);
2. Run `FIFTHSITE_PORT=8080 cargo run` in the repo directory. If `{ code: 13, kind: PermissionDenied, message: "Permission denied" }` is thrown, run it as sudo;
3. Visit http://localhost:8080, check if the page looks alright, then screenshot it.

## Production

While hosting for production isn't exactly required in the slightest,
you're able to easily with [Docker](https://hub.docker.com/r/alexeek/5g).

Run the following command:

```shell script
$ docker run -d --rm -p 80:80 --name 5g alexeek/5g
```

Please note that the Docker image can be delayed by up to 15 minutes.
