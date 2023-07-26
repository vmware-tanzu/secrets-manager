# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

help:
	@echo "--------------------------------------------------------------------"
	@echo "          🛡️ VMware Secrets Manager: Keep your secrets… secret."
	@echo "          🛡️ https://vsecm.com/"
	@echo "--------------------------------------------------------------------"
	@echo "        ℹ️ This Makefile assumes you use Minikube and Docker"
	@echo "        ℹ️ for most operations."
	@echo "--------------------------------------------------------------------"

	@if [ "`uname`" = "Darwin" ]; then \
		if type docker > /dev/null 2>&1; then \
			echo "  Using Docker for Mac?"; \
			echo "        ➡ 'make mac-tunnel' to proxy to the internal registry."; \
		else \
			echo "  Docker is not installed on this Mac."; \
		fi; \
	fi

	@echo ""

	@if [ -z "$(DOCKER_HOST)" -o -z "$(MINIKUBE_ACTIVE_DOCKERD)" ]; then \
		echo "  Using Minikube? If DOCKER_HOST and MINIKUBE_ACTIVE_DOCKERD are"; \
		echo '  not set, then run: eval $$(minikube -p minikube docker-env)'; \
		echo "        ➡ \$$DOCKER_HOST            : ${DOCKER_HOST}"; \
		echo "        ➡ \$$MINIKUBE_ACTIVE_DOCKERD: ${MINIKUBE_ACTIVE_DOCKERD}"; \
	else \
	    echo "  Make sure DOCKER_HOST and MINIKUBE_ACTIVE_DOCKERD are current:"; \
		echo '          eval $$(minikube -p minikube docker-env)'; \
	    echo "          (they may change if you reinstall Minikube)"; \
		echo "        ➡ \$$DOCKER_HOST            : ${DOCKER_HOST}"; \
		echo "        ➡ \$$MINIKUBE_ACTIVE_DOCKERD: ${MINIKUBE_ACTIVE_DOCKERD}"; \
	fi

	@echo "--------------------------------------------------------------------"
	@echo "  Prep/Cleanup:"
	@echo "          ˃ make k8s-delete;make k8s-start;"
	@echo "          ˃ make clean;"
	@echo "--------------------------------------------------------------------"
	@echo "  Testing:"
	@echo "    ⦿ Istanbul images:"
	@echo "          ˃ make build-local;make deploy-local;make test-local;"
	@echo "    ⦿ Istanbul FIPS images:"
	@echo "          ˃ make build-local;make deploy-fips-local;make test-local;"
	@echo "    ⦿ Photon images:"
	@echo "          ˃ make build-local;make deploy-photon-local;make test-local;"
	@echo "    ⦿ Photon FIPS images:"
	@echo "          ˃ make build-local;make deploy-photon-fips-local;make test-local;"
	@echo "--------------------------------------------------------------------"
	@echo "  Example Use Cases:"
	@echo "    Using local images:"
	@echo "          ˃ make example-sidecar-deploy-local;"
	@echo "          ˃ make example-sdk-deploy-local;"
	@echo "          ˃ make example-multiple-secrets-deploy-local;"
	@echo "    Using remote images:"
	@echo "          ˃ make example-sidecar-deploy;"
	@echo "          ˃ make example-sdk-deploy;"
	@echo "          ˃ make example-multiple-secrets-deploy;"
	@echo "--------------------------------------------------------------------"
	@echo " Do these at the build server only!"
	@echo "   Building and Remote Testing:"
	@echo "     ⦿ Istanbul (remote) images:"
	@echo "         ˃ make build;make deploy;make test-remote;"
	@echo "     ⦿ Istanbul FIPS (remote) images:"
	@echo "         ˃ make build;make deploy-fips;make test-remote;"
	@echo "     ⦿ Photon (remote) images:"
	@echo "         ˃ make build;make deploy-photon;make test-remote;"
	@echo "     ⦿ Photon FIPS (remote) images:"
	@echo "         ˃ make build;make deploy-photon-fips;make test-remote;"
	@echo "   Tagging:"
	@echo "         ˃ make tag;"
	@echo "--------------------------------------------------------------------"

h:
	@echo "➡ 'make mac-tunnel'";
	@echo "˃ make k8s-delete;make k8s-start;"
	@echo '⦿ eval $$(minikube -p minikube docker-env)';
	@echo "˃ make clean;"
	@echo "˃ make build-local;make deploy-local;make test-local;"
	@echo "˃ make build;make deploy;make test-remote;"
	@echo "˃ make tag;"
