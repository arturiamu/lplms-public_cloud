package compute

import (
	"errors"
	"fmt"
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type CreateSecurityGroupArgs struct {
	SecurityGroupName *string `json:"security_group_name"`
	Description       *string `json:"description"`
}

func (args *CreateSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupCreateArg {
	var arg = entity.SecurityGroupCreateArg{
		ProjectID:         u.ProjectID,
		SecurityGroupName: args.SecurityGroupName,
		Description:       args.Description,
	}

	return &arg
}

func (rc *RouterCompute) CreateSecurityGroup(c *gin.Context) {
	var (
		args   CreateSecurityGroupArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	createArgs := args.toEntityArgs(u)
	_, err := rc.ci.CreateSecurityGroup(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type DeleteSecurityGroupArgs struct {
	SecurityGroupID *string `json:"security_group_id"`
}

func (args *DeleteSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupDeleteArg {
	var arg = entity.SecurityGroupDeleteArg{
		SecurityGroupID: *args.SecurityGroupID,
	}

	return &arg
}

func (rc *RouterCompute) DeleteSecurityGroup(c *gin.Context) {
	var (
		args   DeleteSecurityGroupArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	createArgs := args.toEntityArgs(u)
	_, err := rc.ci.DeleteSecurityGroup(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type UpdateSecurityGroupArgs struct {
	SecurityGroupID   *string `json:"security_group_id"`
	SecurityGroupName *string `json:"security_group_name"`
	Description       *string `json:"description"`
}

func (args *UpdateSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupUpdateArg {
	var arg = entity.SecurityGroupUpdateArg{
		SecurityGroupID:   *args.SecurityGroupID,
		SecurityGroupName: args.SecurityGroupName,
		Description:       args.Description,
	}

	return &arg
}

func (rc *RouterCompute) UpdateSecurityGroup(c *gin.Context) {
	var (
		args   UpdateSecurityGroupArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	createArgs := args.toEntityArgs(u)
	_, err := rc.ci.UpdateSecurityGroup(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type GetSecurityGroupArgs struct {
	SecurityGroupID string `pos:"path:id"`

	Direction *common.Direction `pos:"query:direction"`
}

func (args *GetSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupGetArg {
	var arg = entity.SecurityGroupGetArg{
		SecurityGroupID: args.SecurityGroupID,
		Direction:       args.Direction,
	}

	return &arg
}

func (rc *RouterCompute) GetSecurityGroup(c *gin.Context) {
	var (
		id     = c.Param("id")
		args   GetSecurityGroupArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)

	createArgs := args.toEntityArgs(u)
	createArgs.SecurityGroupID = id
	resp, err := rc.ci.GetSecurityGroup(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessWith("", resp.SecurityGroup))
}

type ListSecurityGroupArgs struct {
	SecurityGroupIDs []string `json:"security_group_i_ds"`
}

func (args *ListSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupListArg {
	var arg = entity.SecurityGroupListArg{
		ProjectID:        u.ProjectID,
		SecurityGroupIDs: args.SecurityGroupIDs,
	}

	return &arg
}

func (rc *RouterCompute) ListSecurityGroup(c *gin.Context) {
	var (
		args   ListSecurityGroupArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)

	createArgs := args.toEntityArgs(u)
	resp, err := rc.ci.ListSecurityGroup(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessWith("", resp.SecurityGroups))
}

type CreateSecurityGroupRuleArgs struct {
	SecurityGroupID string `pos:"path:id"`

	Direction             common.Direction             `json:"direction"`
	Type                  common.SecurityRuleGrantType `json:"type"`
	RemoteSecurityGroupID []*string                    `json:"remoteSecurityGroupID"`
	CIDR                  []*string                    `json:"CIDR"`
	Protocol              common.NetworkProtocol       `json:"protocol"`
	PortRange             []*string                    `json:"portRange"`
	Description           *string                      `json:"description"`
}

func (args *CreateSecurityGroupRuleArgs) SecurityGroupRuleCreateArgsBaseCheck() (err error) {
	if !args.Direction.IsValid() {
		return errors.New("err direction")
	}
	if !args.Type.IsValid() {
		return errors.New("err security rule grant type")
	}
	if args.Type == common.SecurityGroupType {
		if args.RemoteSecurityGroupID == nil {
			return errors.New("err remote security id")
		}
	} else {
		if args.CIDR == nil {
			return errors.New("err cidr")
		}
	}
	if !args.Protocol.IsValid() {
		return errors.New("err protocol")
	}
	if args.PortRange == nil || len(args.PortRange) == 0 {
		return errors.New("err port range")
	}
	return
}

func (args *CreateSecurityGroupRuleArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupRuleCreateArg {
	var arg = entity.SecurityGroupRuleCreateArg{
		SecurityGroupID: args.SecurityGroupID,
		Direction:       args.Direction,
		Protocol:        args.Protocol,
		Type:            args.Type,
		Description:     args.Description,
	}

	return &arg
}

func (rc *RouterCompute) CreateSecurityGroupRule(c *gin.Context) {
	var (
		args   CreateSecurityGroupRuleArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	if err := args.SecurityGroupRuleCreateArgsBaseCheck(); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	res, err := rc.ci.GetSecurityGroup(&entity.SecurityGroupGetArg{
		SecurityGroupID: args.SecurityGroupID,
		Direction:       &args.Direction,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	rmap := make(map[string]struct{})
	r := res.SecurityGroup
	for _, ri := range r.SecurityGroupRules {
		str := fmt.Sprintf("%s-%s-%d-%d-%d-%s", ri.Direction, ri.Protocol, *ri.FromPort, *ri.ToPort, ri.Type, *ri.CIDR)
		rmap[str] = struct{}{}
	}

	createArgs := args.toEntityArgs(u)
	//_, err := rc.ci.CreateSecurityGroupRule(createArgs)
	var createArgsList []entity.SecurityGroupRuleCreateArg
	for _, portRange := range args.PortRange {
		fromPort, toPort, err1 := extractPortRange(*portRange)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
			return
		}
		createArgs.FromPort = fromPort
		createArgs.ToPort = toPort
		var authList []*string
		switch args.Type {
		case common.SecurityGroupType:
			authList = args.RemoteSecurityGroupID
		case common.CIDRType:
			authList = args.CIDR
		default:
			c.JSON(http.StatusBadRequest, common.FailWith("err security rule grant type", nil))
			return
		}

		cas := buildCreateSecurityGroupRuleArgs(*createArgs, args.Type, authList)
		for i, cai := range cas {
			str := fmt.Sprintf("%s-%s-%d-%d-%d-%s", cai.Direction, cai.Protocol, cai.FromPort, cai.ToPort, cai.Type, *cai.CIDR)
			if _, ok := rmap[str]; ok {
				c.JSON(http.StatusBadRequest, common.FailWith("rule already exist", nil))
				return
			}
			rmap[str] = struct{}{}
			createArgsList = append(createArgsList, cas[i])
		}
	}

	c.JSON(http.StatusOK, common.Success())
}

type DeleteSecurityGroupRuleArgs struct {
	SecurityGroupID     string `json:"security_group_id"`
	SecurityGroupRuleID string `json:"security_group_rule_id"`
}

func (args *DeleteSecurityGroupRuleArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupRuleDeleteArg {
	var arg = entity.SecurityGroupRuleDeleteArg{
		SecurityGroupID:     args.SecurityGroupID,
		SecurityGroupRuleID: args.SecurityGroupRuleID,
	}

	return &arg
}

func (rc *RouterCompute) DeleteSecurityGroupRule(c *gin.Context) {
	var (
		args   DeleteSecurityGroupRuleArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	createArgs := args.toEntityArgs(u)
	_, err := rc.ci.DeleteSecurityGroupRule(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type GetSecurityGroupRuleArgs struct {
}

func (args *GetSecurityGroupRuleArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupRuleGetArg {
	return nil
}

func (rc *RouterCompute) GetSecurityGroupRule(c *gin.Context) {
}

type ListSecurityGroupRuleArgs struct {
	SecurityGroupID string `json:"security_group_id"`
}

func (args *ListSecurityGroupRuleArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupRuleListArg {
	var arg = entity.SecurityGroupRuleListArg{
		SecurityGroupID: args.SecurityGroupID,
	}

	return &arg
}

func (rc *RouterCompute) ListSecurityGroupRule(c *gin.Context) {
	var (
		id     = c.Param("id")
		args   ListSecurityGroupRuleArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)

	createArgs := args.toEntityArgs(u)
	createArgs.SecurityGroupID = id
	describeRes, err := rc.ci.ListSecurityGroupRule(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	data := make([]*entity.SecurityGroupRule, len(describeRes.SecurityGroupRoles))
	for i, sgr := range describeRes.SecurityGroupRoles {
		data[i] = &entity.SecurityGroupRule{
			SecurityGroupRuleID:   sgr.SecurityGroupRuleID,
			Direction:             sgr.Direction,
			Type:                  sgr.Type,
			RemoteSecurityGroupID: sgr.RemoteSecurityGroupID,
			CIDR:                  sgr.CIDR,
			Protocol:              sgr.Protocol,
			FromPort:              sgr.FromPort,
			ToPort:                sgr.ToPort,
			Description:           sgr.Description,
			CreatedAt:             sgr.CreatedAt,
		}
	}
	c.JSON(http.StatusOK, common.SuccessWith("", data))
}

func extractPortRange(portRange string) (fromPort, toPort int, err error) {
	splitPorts := strings.Split(portRange, "/")
	if len(splitPorts) != 2 {
		err = fmt.Errorf("invlaid portRange:%s", portRange)
		return
	}

	fromPort, err = strconv.Atoi(splitPorts[0])
	if err != nil {
		err = fmt.Errorf("strconv.Atoi(%s):%v", splitPorts[0], err)
		return
	}

	toPort, err = strconv.Atoi(splitPorts[1])
	if err != nil {
		err = fmt.Errorf("strconv.Atoi(%s):%v", splitPorts[1], err)
	}

	return
}

func buildCreateSecurityGroupRuleArgs(createArgs entity.SecurityGroupRuleCreateArg, grantType common.SecurityRuleGrantType, authList []*string) (createArgsList []entity.SecurityGroupRuleCreateArg) {
	for _, source := range authList {
		if grantType == common.SecurityGroupType {
			createArgs.RemoteSecurityGroupID = source
		} else {
			createArgs.CIDR = source
		}
		if createArgs.Direction == common.AllDirection {
			createArgs.Direction = common.IngressDirection
			createArgsList = append(createArgsList, createArgs)
			createArgs.Direction = common.EgressDirection
		}
		createArgsList = append(createArgsList, createArgs)
	}

	return createArgsList
}
