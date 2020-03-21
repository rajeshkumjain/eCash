
https://github.com/golang-standards/project-layout

Standard Go Project Layout

This is a basic layout for Go application projects. It's not an official standard defined by the core Go dev team; however, it is a set of common historical and emerging project layout patterns in the Go ecosystem. Some of these patterns are more popular than others. It also has a number of small enhancements along with several supporting directories common to any large enough real world application.

If you are trying to learn Go or if you are building a PoC or a toy project for yourself this project layout is an overkill. Start with something really simple (a single main.go file is more than enough). As your project grows keep in mind that it'll be important to make sure your code is well structured otherwise you'll end up with a messy code with lots of hidden dependencies and global state. When you have more people working on the project you'll need even more structure. That's when it's important to introduce a common way to manage packages/libraries. When you have an open source project or when you know other projects import the code from your project repository that's when it's important to have private (aka internal) packages and code. Clone the repository, keep what you need and delete everything else! Just because it's there it doesn't mean you have to use it all. None of these patterns are used in every single project. Even the vendor pattern is not universal.

Note that Go modules and the related capabilities will have an impact on your project layout. The repo will be updated to include Go modules once it's fully enabled by default. In the meantime, feel free to add your thoughts and ideas in this Github issue.

This project layout is intentionally generic and it doesn't try to impose a specific Go package structure.

This is a community effort. Open an issue if you see a new pattern or if you think one of the existing patterns needs to be updated.


/cmd

Main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g., /cmd/myapp).

Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the /pkg directory. If the code is not reusable or if you don't want others to reuse it, put that code in the /internal directory. You'll be surprised what others will do, so be explicit about your intentions!

It's common to have a small main function that imports and invokes the code from the /internal and /pkg directories and nothing else.
See the /cmd directory for examples.

/internal

Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself. See the Go 1.4 release notes for more details. Note that you are not limited to the top level internal directory. You can have more than one internal directory at any level of your project tree.

You can optionally add a bit of extra structure to your internal packages to separate your shared and non-shared internal code. It's not required (especially for smaller projects), but it's nice to have visual clues showing the intended package use. Your actual application code can go in the /internal/app directory (e.g., /internal/app/myapp) and the code shared by those apps in the /internal/pkg directory (e.g., /internal/pkg/myprivlib).

/pkg

Library code that's ok to use by external applications (e.g., /pkg/mypubliclib). Other projects will import these libraries expecting them to work, so think twice before you put something here :-) Note that the internal directory is a better way to ensure your private packages are not importable because it's enforced by Go. The /pkg directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. The I'll take pkg over internal blog post by Travis Jeffery provides a good overview of the pkg and internal directories and when it might make sense to use them.

It's also a way to group Go code in one place when your root directory contains lots of non-Go components and directories making it easier to run various Go tools (as mentioned in these talks: Best Practices for Industrial Programming from GopherCon EU 2018, GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps and GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go).

See the /pkg directory if you want to see which popular Go repos use this project layout pattern. This is a common layout pattern, but it's not universally accepted and some in the Go community don't recommend it.

Ok not to use it if your app project is really small and where an extra level of nesting doesn't add much value (unless you really want to :-)). Think about it when it's getting big enough and your root directory gets pretty busy (especially if you have a lot of non-Go app components).

/vendor

Application dependencies (managed manually or by your favorite dependency management tool like the new built-in, but still experimental, modules feature).
Don't commit your application dependencies if you are building a library.

Note that since 1.13 Go also enabled the module proxy feature (using https://proxy.golang.org as their module proxy server by default). Read more about it here to see if it fits all of your requirements and constraints. If it does, then you won't need the vendor directory at all.

=================================
Service Application Directories
=================================

/api

OpenAPI/Swagger specs, JSON schema files, protocol definition files.
See the /api directory for examples.

===========================
Web Application Directories
===========================
/web

Web application specific components: static web assets, server side templates and SPAs.

================================
Common Application Directories
================================

/configs

Configuration file templates or default configs.
Put your confd or consul-template template files here.

/init

System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.

/scripts

Scripts to perform various build, install, analysis, etc operations.
These scripts keep the root level Makefile small and simple (e.g., https://github.com/hashicorp/terraform/blob/master/Makefile).
See the /scripts directory for examples.

/build

Packaging and Continuous Integration.

Put your cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the /build/package directory.

Put your CI (travis, circle, drone) configurations and scripts in the /build/ci directory. Note that some of the CI tools (e.g., Travis CI) are very picky about the location of their config files. Try putting the config files in the /build/ci directory linking them to the location where the CI tools expect them (when possible).

/deployments

IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).

/test

Additional external test apps and test data. Feel free to structure the /test directory anyway you want. For bigger projects it makes sense to have a data subdirectory. For example, you can have /test/data or /test/testdata if you need Go to ignore what's in that directory. Note that Go will also ignore directories or files that begin with "." or "_", so you have more flexibility in terms of how you name your test data directory.

See the /test directory for examples.

=======================
Other Directories
=======================
/docs

Design and user documents (in addition to your godoc generated documentation).
See the /docs directory for examples.

/tools

Supporting tools for this project. Note that these tools can import code from the /pkg and /internal directories.
See the /tools directory for examples.

/examples

Examples for your applications and/or public libraries.
See the /examples directory for examples.

/third_party

External helper tools, forked code and other 3rd party utilities (e.g., Swagger UI).

/githooks

Git hooks.

/assets

Other assets to go along with your repository (images, logos, etc).

/website

This is the place to put your project's website data if you are not using Github pages.
See the /website directory for examples.

