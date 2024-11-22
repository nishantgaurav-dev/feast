<<<<<<< HEAD
FROM registry.access.redhat.com/ubi9/python-311:9.5

ARG VERSION
RUN pip install "feast[aws,gcp,snowflake,redis,go,mysql,postgres,opentelemetry,grpcio,k8s,duckdb,milvus]"==${VERSION}

=======
FROM python:3.11-slim-bullseye

RUN pip install --no-cache-dir pip --upgrade
RUN pip install --no-cache-dir "feast[aws,gcp,snowflake,redis,go,mysql,postgres,opentelemetry,grpcio]"


RUN apt update && apt install -y -V ca-certificates lsb-release wget && \
    wget https://apache.jfrog.io/artifactory/arrow/$(lsb_release --id --short | tr 'A-Z' 'a-z')/apache-arrow-apt-source-latest-$(lsb_release --codename --short).deb && \
    apt install -y -V ./apache-arrow-apt-source-latest-$(lsb_release --codename --short).deb && apt update && \
    apt install -y \
    jq \
    libarrow-dev && \
    apt remove -y lsb-release wget && \
    apt-get clean && rm -rf /var/cache/apt/lists

>>>>>>> 4c971f6a1 (fix: Reduce feast-server container image size & fix dev image build (#4781))
# modify permissions to support running with a random uid
RUN chmod g+w $(python -c "import feast.ui as ui; print(ui.__path__)" | tr -d "[']")/build/projects-list.json
