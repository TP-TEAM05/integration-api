# ReCo integration API

## Using Golang private packages

- Create a personal access token on GitHub
- Add the following to your `~/.bashrc` or `~/.zshrc`:

```bash
export GOPRIVATE="github.com/ReCoFIIT"
```

If you have multiple private repositories, separate them with a comma:

```bash
export GOPRIVATE="github.com/ReCoFIIT,github.com/your-other-repo"
```

- Run `source ~/.bashrc` or `source ~/.zshrc`
- In `.gitconfig` add the following:

```bash
[url "ssh://git@github.com/"]
 insteadOf = https://github.com/
```

- create a file `~/.netrc` with the following content:

```bash
machine github.com
login <your-github-username>
password <your-github-token>
```