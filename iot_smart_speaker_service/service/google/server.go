package google

import (
	"github.com/gin-gonic/gin"
)

const AgentUserIdHeader = "agentUserId"

func SetAgentUserIdHeader(c *gin.Context, agentUserId string) {
	c.Request.Header.Set(AgentUserIdHeader, agentUserId)
}

func GetAgentUserIdFromHeader(c *gin.Context) string {
	return c.GetHeader(AgentUserIdHeader)
}
