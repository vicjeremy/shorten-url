<div align="center" id="readme-top">

  <h1>URL Shortening Service</h1>

  <p>
    A simple RESTful API that allows users to shorten long URLs </p>
 <p>Built with Go (Gin Gonic, Gorm) and PostgreSQL, this tool is my implementation of the <a href="https://roadmap.sh/projects/url-shortening-service">URL Shortening Service</a> challenge from <a href="https://roadmap.sh">roadmap.sh</a>.
  </p>

<!-- Badges -->
<p>
  <a href="https://github.com/vicjeremy/shorten-url/graphs/contributors">
    <img src="https://img.shields.io/github/contributors/vicjeremy/shorten-url" alt="contributors" />
  </a>
  <a href="">
    <img src="https://img.shields.io/github/last-commit/vicjeremy/shorten-url" alt="last update" />
  </a>
  <a href="https://github.com/vicjeremy/shorten-url/network/members">
    <img src="https://img.shields.io/github/forks/vicjeremy/shorten-url" alt="forks" />
  </a>
  <a href="https://github.com/vicjeremy/shorten-url/stargazers">
    <img src="https://img.shields.io/github/stars/vicjeremy/shorten-url" alt="stars" />
  </a>
  <a href="https://github.com/vicjeremy/shorten-url/issues/">
    <img src="https://img.shields.io/github/issues/vicjeremy/shorten-url" alt="open issues" />
  </a>
  <a href="https://github.com/vicjeremy/shorten-url/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/vicjeremy/shorten-url.svg" alt="license" />
  </a>
</p>

<h4>
    <a href="https://github.com/vicjeremy/shorten-url/issues/">Report Bug</a>
  <span> · </span>
    <a href="https://github.com/vicjeremy/shorten-url/issues/">Request Feature</a>
  </h4>
</div>

<br />

<!-- Table of Contents -->

# :notebook_with_decorative_cover: Table of Contents

- [About the Project](#star2-about-the-project)
  <!-- - [Screenshots](#camera-screenshots) -->
  - [Tech Stack](#space_invader-tech-stack)
  - [Features](#dart-features)
- [Getting Started](#toolbox-getting-started)
  - [Prerequisites](#bangbang-prerequisites)
  <!-- [Installation](#gear-installation) -->
  - [Run Locally](#running-run-locally)
- [Usage](#eyes-usage)
- [Contributing](#wave-contributing)
- [License](#warning-license)
- [Contact](#handshake-contact)
- [Acknowledgements](#gem-acknowledgements)

<!-- About the Project -->

## :star2: About the Project

<!-- Screenshots -->

<!-- ### :camera: Screenshots

<div align="center">
  <img src="img/example-test.png" style="width:600px;height:400px" alt="screenshot" />
</div> -->

<!-- TechStack -->

### :space_invader: Tech Stack

- <a href="https://golang.org" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="30" height="30"/>[![Go][Go]][Go-url]</a>

<!-- Features -->

### :dart: Features

- Create a new short URL
- Retrieve an original URL from a short URL
- Update an existing short URL
- Delete an existing short URL
- Get statistics on the short URL (e.g., number of times accessed)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Getting Started -->

## :toolbox: Getting Started

<!-- Prerequisites -->

### :bangbang: Prerequisites

This project uses Go as its main language. Make sure you have it installed on your machine. If not, you can install it from the official website [here](https://golang.org/).

```bash
  go version
```

<!-- Run Locally -->

### :running: Run Locally

Clone the project

```bash
  git clone https://github.com/vicjeremy/shorten-url.git
```

Go to the project directory

```bash
  cd shorten-url
```

Run the project

```bash
  go run cmd/main.go
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Usage -->

## :eyes: Usage
Use Postman or Curl in terminal

```bash
# Create Short URL
curl -d '{"url": "https://www.example.com/some/long/url"}' -H 'Content-Type: application/json' http://localhost:3000/shorten
# Example output "shortCode" : "abc123"

# Retrieve Original URL
curl -v http://localhost:3000/shorten/abc123

# Update Short URL
curl -X PUT -d '{"url": "https://www.example.com/some/updated/url"}' -H 'Content-Type: application/json' http://localhost:3000/shorten/abc123

# Delete Short URL
curl -X DELETE http://localhost:3000/shorten/abc123

# Get URL Statistics
curl -v http://localhost:3000/shorten/abc123/stats
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Contributing -->

## :wave: Contributing

<a href="https://github.com/vicjeremy/shorten-url/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=vicjeremy/shorten-url" />
</a>

Contributions are always welcome!

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- License -->

## :warning: License

Distributed under the MIT License. See [LICENSE.txt](LICENSE.txt) for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Contact -->

## :handshake: Contact

Vic Jeremy - [@viccjeremy](https://instagram.com/viccjeremy) - [vicjeremyp@gmail.com](mailto:vicjeremyp@gmail.com)

Project Link: [https://github.com/vicjeremy/shorten-url](https://github.com/vicjeremy/shorten-url)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Acknowledgments -->

## :gem: Acknowledgements

- [Backend Project ideas from roadmap.sh](https://roadmap.sh/backend/projects)


<p align="right">(<a href="#readme-top">back to top</a>)</p>

[Go]: https://img.shields.io/badge/GOlang-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://golang.org/
