# Git commit messages for lazy people

You feed it a verbose commit status and it will spit out commit message via
OpenAI API.

## Install

```console
go install github.com/mitjafelicijan/lazycommit
```

Then you need to put OpenAI API key in somewhere so it's accessible in your
shell. I put mine in `.bashrc` file with `export OPENAI_API_KEY=""`.

You can then test it with:

```console
echo "My diff" | lazygit
```

## Usage

If you try to commit with `git commit --verbose` this will open up your default
terminal editor and also provide diffs of your changes.

You can also provide a setting in your `.gitconfig` and this way you will not
be required to provide verbose flag each time.

```gitconfig
[commit]
    verbose = true
```

## Use in VIM

If your default editor is VIM and when Git asks you for commit message you can
then execute the command below at the top of the buffer you will get back the
commit message from LLM.

To try a different message just undo the change and try executing command again.

```vimrc
:r !< % lazycommit
```

You could keybind this like this (Leader gc).

```vimrc
nnoremap <Leader>gc :r !< % lazycommit<cr>
```

## Ideas

- Add local model suppor through Ollama.
- Add Anthropic Claude support.

## Shoutout to

- [Commit Message Writer's Block? Here's the Cure](https://www.youtube.com/watch?v=cxapgvGkOJY)
