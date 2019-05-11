#!/bin/bash
#
# Copyright 2017 Istio Authors
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

set -o errexit

if [ "$#" -ne 1 ]; then
    echo Missing version parameter
    echo Usage: build-services.sh \<version\>
    exit 1
fi

VERSION=$1
SCRIPTDIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
REPO_SPEC="soloio"


pushd "$SCRIPTDIR/reviews"
  #java build the app.
  docker run --rm -u root -v "$(pwd)":/home/gradle/project -w /home/gradle/project gradle:4.8.1 gradle clean build
  pushd reviews-wlpcfg
    #plain build -- no ratings
  docker build -t "${REPO_SPEC}/examples-bookinfo-reviews-v1:${VERSION}" -t "${REPO_SPEC}/examples-bookinfo-reviews-v1:latest" --build-arg service_version=v1 .
    #with ratings black stars
    docker build -t "${REPO_SPEC}/examples-bookinfo-reviews-v2:${VERSION}" -t "${REPO_SPEC}/examples-bookinfo-reviews-v2:latest" --build-arg service_version=v2 \
	   --build-arg enable_ratings=true .
    #with ratings red stars
    docker build -t "${REPO_SPEC}/examples-bookinfo-reviews-v3:${VERSION}" -t "${REPO_SPEC}/examples-bookinfo-reviews-v3:latest" --build-arg service_version=v3 \
	   --build-arg enable_ratings=true --build-arg star_color=red .
    #with ratings red stars, with cascading failure
    docker build -t "${REPO_SPEC}/examples-bookinfo-reviews-v4:${VERSION}" -t "${REPO_SPEC}/examples-bookinfo-reviews-v4:latest" --build-arg service_version=v4 \
	   --build-arg enable_ratings=true --build-arg star_color=red --build-arg propagate_failure=true .
  popd
popd
