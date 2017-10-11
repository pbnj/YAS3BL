# CONTRIBUTING

First, thank you for your interest in contributing to this project. I try to stay abreast of news that breaks regarding bucket leaks, but I could easily miss one or two.

Contributing to this project is relatively simple.

## Overview

1. Fork and clone the repository.
1. Add the information in [`yas3bl.json`](yas3bl.json).
1. Generate the [README.md](README.md).

And you're done!

## Details

### Prerequisites

You must have **ONE** of the following two options installed locally on your machine:
- Docker
- Go

Alternatively, if you have a [c9.io](https://c9.io) account, you can spin up a generic workspace (it is provisioned with Go out-of-the-box 😍 )

### Steps

1. Fork then clone the repo: `git clone https://github.com/<YOUR_GITHUB_USERNAME>/YAS3BL && cd YAS3BL`
1. Add new entry in `yas3bl.json`. Required JSON properties:
    ```json
    {
      "count": "<NUMBER OF RECORDS EXPOSED BY THIS LEAK>",
      "data": "<TYPE OF DATA EXPOSED BY THIS LEAK>",
      "organization": "<OWNER OF THIS LEAKED DATA>",
      "url": "<URL TO ORIGINAL BLOG ARTICLE BY SECURITY RESEARCHERS>"
    }
    ```
    - Note: The `url` should be that of the original security researchers that disclosed this leak. I'd rather not rely on secondary news sources that report on this leak.
1. Generate the markup using **ONE** of the two options below:
    - If you have Go installed locally or if you're on c9.io: `make` or `make readme`
    - If you have Docker installed locally: `make docker`
1. Push & open PR 👍
