package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	dockerClient "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	Models "github.com/SinergiaManager/sinergiamanager-backend/models"
)

func DockerConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	cli, err := dockerClient.NewClientWithOpts(dockerClient.FromEnv, dockerClient.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
		return
	}

	ctx := context.Background()

	config := &Models.ConfigDb{}
	err = DB.Collection("configs").FindOne(ctx, bson.M{}).Decode(config)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Fatalf("Error finding config: %v", err)
		} else {
			log.Fatalf("Error finding config: %v", err)
		}
		return
	}

	fmt.Printf("Config: %+v\n", config)

	// Check if the container already exists
	containerName := "mailserver"
	containers, err := cli.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing containers: %v", err)
		return
	}

	for _, c := range containers {
		if c.Names[0] == "/"+containerName {
			// Container exists, start it if it's not already running
			if c.State != "running" {
				log.Printf("Starting existing container: %s\n", containerName)
				if err := cli.ContainerStart(ctx, c.ID, container.StartOptions{}); err != nil {
					log.Fatalf("Error starting existing container: %v", err)
				}
				fmt.Printf("Container %s started successfully\n", containerName)
			} else {
				log.Printf("Container %s is already running\n", containerName)

				err = cli.ContainerRestart(ctx, c.ID, container.StopOptions{})
				if err != nil {
					log.Fatalf("Error restarting container: %v", err)
				}
			}
			return
		}
	}

	// If the container does not exist, create it
	networkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"sinergiamanager-backend_mynetwork": {},
		},
	}

	configDocker := &container.Config{
		Image: os.Getenv("MAILSERVER_IMAGE"), // Image specified in .env
		Env: []string{
			fmt.Sprintf("SMTP_SERVER=%s", config.SmtpHost),
			fmt.Sprintf("SMTP_USERNAME=%s", config.SmtpEmail),
			fmt.Sprintf("SMTP_PASSWORD=%s", config.SmtpPassword),
			fmt.Sprintf("SMTP_PORT=%d", config.SmtpPort),
			fmt.Sprintf("SERVER_HOSTNAME=%s", "localhost"),
		},
		ExposedPorts: nat.PortSet{
			"2500/tcp": struct{}{},
		},
	}

	// Set container host configurations (ports, volumes, etc.)
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"2500/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "2500",
				},
			},
		},
	}

	// Pull the image if it doesn't exist
	reader, err := cli.ImagePull(ctx, configDocker.Image, image.PullOptions{})
	if err != nil {
		log.Fatalf("Error pulling image: %v", err)
		return
	}
	defer reader.Close()

	// Create the container
	resp, err := cli.ContainerCreate(ctx, configDocker, hostConfig, networkingConfig, nil, containerName)
	if err != nil {
		log.Fatalf("Error creating container: %v", err)
		return
	}

	// Start the container
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Fatalf("Error starting container: %v", err)
		return
	}

	fmt.Println("Container started successfully:", resp.ID)
}
