# telegraphcl
telegra.ph from your friendly terminal.

## Introduction

[telegra.ph](https://telegra.ph) is a minimalist content publishing platform for users to quickly share richly formatted post. `telegraphcl` lets you use
telegra.ph from command-line, using Markdown file to draft your posts.

Here's what you can do with `telegraphcl`
[![asciicast](https://asciinema.org/a/371198.png)](https://asciinema.org/a/371198)

## Installation

Please make sure go binaries are accessible. i.e for UNIX,

    $ export PATH=$PATH:~/go/bin/


### Using golang installation

    $ go get github.com/plant99/telegraphcl/cmd/telegraphcl

### Using docker

Let's say `examples` directory has the markdown files to be used to create/edit pages.

    $ docker run -v ~/.telegraphcl:/root/.telegraphcl -v "$PWD"/examples:/root/telegraph-blogs/ -it plant99/telegraphcl --help

## Usage

    $ telegraphcl --help

```
Usage:
  telegraphcl [command]

Available Commands:
  help        Help about any command
  page        Manage your Telegra.ph pages
  user        Operations related to Telegraph account management
  version     Print the version number of telegraphcl

Flags:
  -h, --help   help for telegraphcl

Use "telegraphcl [command] --help" for more information about a command.
```

### user

With this command, one can handle an user account to manage your pages. A token is generated and stored at `~/.telegraphcl/telegraph.token`

```
Operations related to Telegraph account management

Usage:
  telegraphcl user [command]

Available Commands:
  create      Create an user
  edit        Edit current user information
  revoke      Revoke current access token, and regenerate access token.
  view        View current user information

Flags:
  -h, --help   help for user

Use "telegraphcl user [command] --help" for more information about a command.
```

Some things to note on user management.

1. If you run the following command after creating an user

        $ telegraphcl user view
   The `AuthURL` can be used in web-browsers to authenticate this user to perform create/edit operations from the user account referred to by the `telegraph.token`. 

2. `revoke` option revokes the current access token, and creates a new one for the `same` user account.

### page

With this option, one can manage pages. **Some things to note**

1. Markdown files are used to `create`, and `edit` a page. If one changes the pages on a browser these changes don't currently sync with the markdown files one used to create a page.
2. <path> means the identity of the 'page' with telegraph. For example, `https://telegra.ph/some-title-11-07`. The `<path>` for this URL is `some-title-11-07`.
3. `<markdown-path>` means path to the Markdown file used to create/edit a telegra.ph page.
4. Create a directory to keep all the Markdown files for telegra.ph pages. This would here on be referenced as `<markdown-dir>`

```
Manage your Telegra.ph pages

Usage:
  telegraphcl page [command]

Available Commands:
  create      Create Page from a Markdown file. Arguments: <markdown-path> <title>
  edit        Edit page with Telegra.ph path. Arguments: <path> <markdown-path>
  get         Get page with Telegra.ph path. Arguments: <path>
  list        List your Telegra.ph pages
  views       Count views on your Telegra.ph page. Arguments: <path>

Flags:
  -h, --help   help for page

Use "telegraphcl page [command] --help" for more information about a command.
```

#### create

Create a markdown file in `<markdown-dir>`, say `first_page.md`. Then run the following command

    $ telegraphcl page create <markdown-dir>/first_page.md "My first blog"

*Note for docker image users*: The `WORKDIR` is `/root/telegraph-blogs` so to reference files, you can assume you're running `telegraph` from `<markdown-dir>`.
So assuming `examples` is my `<markdown-dir>`, the following would create a page.

    $ docker run -v ~/.telegraphcl:/root/.telegraphcl -v "$PWD"/examples:/root/telegraph-blogs/ -it plant99/telegraphcl page create tiny_blogpost.md some-title


#### list

List the pages you created by running the following command

    $ telegraphcl page list

#### get

To get details of a page, record the `<path>` of the page. Then run
    
    $ telegraphcl page get <path>

#### edit

To edit a page, record the `<path>` of the page. Make modifications to `first_page.md`, then run

    $ telegraphcl page edit <path> <markdown-dir>/first_page.md

#### views

To get the count of views on a particular page, run the following command.

    $ telegraphcl page views <path>

### version

Prints the version of `telegraphcl`.


## Note


telegra.ph doesn't support all Markdown formatting

1. h1 elements, `#` in Markdown format.
2. h2 elements, `##` in Markdown format.

There could be others, please raise an issue in that case, the documentation would be updated.

## Contributing

Please file an issue, or make a patch via GitHub PRs if you have time.

## Credits

This is an educational project to practice some Golang, some code has been used unabridged from [https://github.com/toby3d/telegraph](https://github.com/toby3d/telegraph) and  
[https://github.com/meinside/telegraph-go](https://github.com/meinside/telegraph-go)

Thank you, for checking out `telegraphcl`.
