#!/usr/bin/env bash

# Copyright 2022 The OpenYurt Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e
set -u

YURT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"

make manifests

if [ ! -z "$(git diff ${YURT_ROOT}/charts)" ]; then
    echo "diff"
    cat a/charts/yurt-manager/crds/iot.openyurt.io_platformadmins.yaml
    echo "-------------"
    cat b/charts/yurt-manager/crds/iot.openyurt.io_platformadmins.yaml
    echo "-------------"
    git diff ${YURT_ROOT}/charts
    echo "Verify manifests error, please run make manifests"
    exit -1
fi 
