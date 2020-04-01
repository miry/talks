# Talks
My slides and content for presentations

* [2020.04.02 Intigrate Sentry in GO](https://miry-talks.herokuapp.com/go-gin-sentry/go-gin-sentry.slide)

## Development


### Heroku

Install Heroku client before use

```shell
$ brew bundle
```

Create a new application

```shell
$ heroku create miry-talks
Creating ⬢ miry-talks... done
https://miry-talks.herokuapp.com/ | https://git.heroku.com/miry-talks.git
$ heroku stack:set container
```

Deploy a new release

```shell
$ git push heroku master
```
