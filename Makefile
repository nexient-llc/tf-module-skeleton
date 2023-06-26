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

# this will always point here
REPO_MANIFESTS_URL ?= https://github.com/nexient-llc/common-automation-framework.git
# this eventually will point to a git tag
REPO_BRANCH ?= main
# this should point to a seed manifest
REPO_MANIFEST ?= manifests/terraform_modules/seed/manifest.xml

# Settings to pull in nexient version of (google) repo utility that supports environment substitution:
REPO_URL ?= https://github.com/nexient-llc/git-repo.git
REPO_REV ?= main
export REPO_REV REPO_URL

# Example variable to substituted after init, but before sync in repo manifests.
GITBASE ?= https://github.com/nexient-llc/
GITREV ?= main
export GITBASE GITREV

# Set to true in a pipeline context
IS_PIPELINE ?= false

IS_AUTHENTICATED ?= false

JOB_NAME ?= job
JOB_EMAIL ?= job@job.job

COMPONENTS_DIR = components
-include $(COMPONENTS_DIR)/Makefile

.PHONY: configure-git-hooks
configure-git-hooks:
	pre-commit install

ifeq ($(IS_PIPELINE),true)
.PHONY: git-config
git-config:
	@set -ex; \
	git config --global user.name "$(JOB_NAME)"; \
	git config --global user.email "$(JOB_EMAIL)"; \
	git config --global color.ui false

configure: git-config
endif

ifeq ($(IS_AUTHENTICATED),true)
.PHONY: git-auth
git-auth:
	$(call config,Bearer $(GIT_TOKEN))

define config
	@set -ex; \
	git config --global http.extraheader "AUTHORIZATION: $(1)"; \
	git config --global http.https://gerrit.googlesource.com/git-repo/.extraheader ''; \
	git config --global http.version HTTP/1.1;
endef

configure: git-auth
endif

.PHONY: configure
configure: configure-git-hooks
	repo --color=never init --no-repo-verify \
		-u "$(REPO_MANIFESTS_URL)" \
		-b "$(REPO_BRANCH)" \
		-m "$(REPO_MANIFEST)"
	repo envsubst
	repo sync

# The first line finds and removes all the directories pulled in by repo
# The second line finds and removes all the broken symlinks from removing things
# https://stackoverflow.com/questions/42828021/removing-files-with-rm-using-find-and-xargs
.PHONY: clean
clean:
	-repo list | awk '{ print $1; }' | cut -d '/' -f1 | uniq | xargs rm -rf
	find . -type l ! -exec test -e {} \; -print | xargs rm -rf

.PHONY: init-clean
init-clean:
	rm -rf .git
	git init --initial-branch=main
ifneq (,$(wildcard ./TEMPLATED_README.md))
	mv TEMPLATED_README.md README.MD
endif
