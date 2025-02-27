package provider

import (
	"github.com/coding-ia/terraform-provider-cloudops/internal/conn"
)

type Meta struct {
	AWSClient conn.AWSClient
}
