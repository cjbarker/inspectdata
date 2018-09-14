Create Docker Image

```bash
docker build -t registry.gitlab.com/cjbarker/go-inspect-data .
```

Connect to verify

```bash
docker run -i -t registry.gitlab.com/cjbarker/go-inspect-data /bin/sh
```

To push to GitLab Registry. Note: Assumes login to registry via GitLab access token is already completed.

```bash
# Get image id, tag it with private registry then push
DOCKER_IMG_ID=`docker images "registry.gitlab.com/cjbarker/go-inspect-data" | awk 'FNR==2{print $3}'`
docker tag ${DOCKER_IMG_ID} registry.gitlab.com/cjbarker/go-inspect-data:latest
docker push registry.gitlab.com/cjbarker/go-inspect-data:latest
```

Pull Down Image from remote repo

```bash
docker pull registry.gitlab.com/cjbarker/go-inspect-data:latest
```
