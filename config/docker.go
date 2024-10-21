package config

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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
	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		log.Fatalf("Error listing containers: %v", err)
		return
	}

	for _, c := range containers {
		if c.Names[0] == "/"+containerName {
			fmt.Printf("Container %s exists\n", containerName)
			// If the container exists, delete it
			if err := cli.ContainerRemove(ctx, c.ID, container.RemoveOptions{Force: true}); err != nil {
				log.Fatalf("Error removing container: %v", err)
				return
			}
		}
	}

	fmt.Printf("Container %s does not exist, creating it\n", containerName)

	// If the container does not exist, create it
	networkNames, err := cli.NetworkList(ctx, network.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing networks: %v", err)
		return
	}

	networkExists := false
	var networkName string

	if len(networkNames) > 0 {
		for _, n := range networkNames {
			if strings.Contains(n.Name, "mynetwork") {
				networkExists = true
				networkName = n.Name
				break
			}
		}
	}

	if !networkExists {
		log.Fatalf("Error creating network: %v", err)
		return
	}

	fmt.Printf("Network %s exists\n", networkName)

	networkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			networkName: {},
		},
	}

	configDocker := &container.Config{
		Image: os.Getenv("MAILSERVER_IMAGE"),
		Env: []string{
			fmt.Sprintf("SMTP_SERVER=%s", config.SmtpHost),
			fmt.Sprintf("SMTP_USERNAME=%s", config.SmtpEmail),
			fmt.Sprintf("SMTP_PASSWORD=%s", config.SmtpPassword),
			fmt.Sprintf("SMTP_PORT=%d", config.SmtpPort),
			fmt.Sprintf("SERVER_HOSTNAME=%s", "localhost"),
			fmt.Sprintf("SUPPORT_EMAIL=%s", config.SupportEmail),
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

	io.Copy(os.Stdout, reader)

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
