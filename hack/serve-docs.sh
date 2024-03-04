#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

cd ./docs || exit

# This script is for local development only.
#
# You will need ruby, bundler, and jekyll installed.
# You will also need to run `cd ./docs && bundler install`
# to install the dependencies first.
#
# Here is what you typically need to do just once:
#   cd ./docs
#   gem install bundler jekyll
#   bundle install
#   jekyll serve

# bundle install
bundle exec jekyll serve
