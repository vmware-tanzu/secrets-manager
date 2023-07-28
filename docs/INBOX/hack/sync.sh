#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

./hack/update-timestamp.sh

JEKYLL_ENV=production jekyll build

rm -rf _site/versions
aws s3 sync _site/ s3://aegis.ist/

aws cloudfront create-invalidation --distribution-id EZFGMY32S3BBS --paths "/*"
