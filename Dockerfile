FROM golang:1.17

WORKDIR /app

# Copy source
COPY go.mod ./src/
COPY go.sum ./src/
COPY main.go ./src/
COPY muon ./src/muon

# Build service
RUN cd /app/src && go mod download
RUN cd /app/src && go build -o /app/muon
RUN rm -rf /app/src

# Prepare data
COPY data/i18n /app/data/i18n
COPY data/templates /app/data/templates
COPY data/static /app/data/static
RUN rm -rf /app/data/static/files
RUN mkdir /app/data/static/files /app/data/articles

# Set up permissions
RUN useradd -s /bin/true muon
RUN passwd -l muon
RUN chown -R root:muon /app/data
RUN find /app/data -type d -exec chmod 750 {} \;
RUN find /app/data -type f -exec chmod 640 {} \;

# Start service
USER muon
ENV MUON_LISTEN=:8080
ENV MUON_DATA=/app/data
ENTRYPOINT ["/app/muon"]
