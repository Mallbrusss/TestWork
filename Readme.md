# Website Availability Monitor

This application monitors the availability of websites and provides information on access time.

## Prerequisites

Make sure you have Docker and Docker Compose installed on your system.

- Docker: https://docs.docker.com/get-docker/
- Docker Compose: https://docs.docker.com/compose/install/

## Usage

1. Clone the repository:

git clone <https://github.com/Mallbrusss/TestWork.git>


2. Navigate to the project directory:

cd /root/storage/of/Project


3. Customize the website list (optional):

Open the `data/sites.go` file and modify the `Sites` variable to include the websites you want to monitor. Each website should have a unique name and URL.

4. Build the Docker image:

docker-compose build

5. Start the application:

docker-compose up -d

This command will start the application in detached mode.

6. Access the application:

- To get the access time for a specific site, make a GET request to `http://localhost:8080/access-time?site=<site_name>`, where `<site_name>` is the name of the site you want to check.

- To get the site with the minimum access time, make a GET request to `http://localhost:8080/min-access-time`.

- To get the site with the maximum access time, make a GET request to `http://localhost:8080/max-access-time`.

7. Stop the application:

docker-compose down


This command will stop and remove the containers.

## Monitoring Multiple Instances

If you want to monitor multiple instances of the application, you can modify the `docker-compose.yml` file and specify different ports for each instance. By default, the application runs on port 8080.
