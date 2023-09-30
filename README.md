# goFiberHtmx
goFiber with HTMX and tailwind 
docker-compose up --build
# Build the Docker image
docker build -t myfiberapp .

# Run the Docker container
docker run -p 3000:3000 myfiberapp


go mod init app
go mod tidy


# 1. Build the services:
docker-compose build

# 2. Start the services in detached mode:
docker-compose up -d

# Alternatively, combine the build and start steps:
# docker-compose up --build -d

# 3. Stop the services:
docker-compose stop

# 4. Stop and remove the containers:
docker-compose down

# If you also want to remove the volumes:
# docker-compose down -v

