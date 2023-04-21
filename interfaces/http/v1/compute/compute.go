package compute

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/interfaces/http/v1/storage"
	"github.com/arturiamu/lplms-public_cloud/utils/ctx"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

type Compute struct {
	ci application.ComputeAppInterface
}

func NewCompute(ci application.ComputeAppInterface) *Compute {
	return &Compute{
		ci: ci,
	}
}

type CreateServerArg struct {
	ProjectID       string                 `validate:"required"`
	ServerName      string                 `json:"server_name" validate:"required"`
	ImageID         string                 `json:"image_id" validate:"required"`
	FlavorID        string                 `json:"flavor_id" validate:"required"`
	VSwitchID       string                 `json:"v_switch_id"`
	SecurityGroupID string                 `json:"security_group_id"`
	HostName        string                 `json:"host_name"`
	SystemDisk      *storage.CreateDiskArg `json:"system_disk" validate:"required"`
	DiskDisk        *storage.CreateDiskArg `json:"disk_disk" validate:"omitempty"`
	VPCID           string                 `json:"vpc_id"`
	Keypair         string                 `json:"keypair"`
	Password        string                 `json:"password" validate:"required"`
	UserData        string                 `json:"user_data"`
	EIP             *ipInfo                `json:"eip"`
	Description     string                 `json:"description"`
}

type ipInfo struct {
	EIPName   string `json:"eip_name"`
	Bandwidth uint   `json:"bandwidth"`
	MNO       string `json:"mno"`
}

func (co *Compute) CreateServer(c *gin.Context) {
	var (
		arg CreateServerArg
		pid = ctx.GetProject(c)
	)
	if err := c.BindJSON(&arg); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	arg.ProjectID = pid
	if err := validate.Struct(&arg); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	co.ci.SaveServer(&entity.Server{})
}

func (co *Compute) DeleteServer(c *gin.Context) {

}
