package directory

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// exportDirectory provide directory xml config content for FreeSWTICH
//
// @Summary provide directory xml config content for FreeSWTICH
// @Tag Directory
// @Param domain body string true "domain name of the FreeSWTICH"
// @Param user body string true "the user for asking"
// @Router /directory/export [post]
//
func (d *Directory) exportDirectory(ctx *gin.Context) {
	domainName := ctx.PostForm("domain")
	userID := ctx.PostForm("user")

	// TODO 查询密码
	password := "2345"

	include := include{
		Domain: domain{
			Name: domainName,
			Params: params{
				Param: []param{
					{
						Name:  "dial-string",
						Value: "{^^:sip_invite_domain=${dialed_domain}:presence_id=${dialed_user}@${dialed_domain}}${sofia_contact(*/${dialed_user}@${dialed_domain})},${verto_contact(${dialed_user}@${dialed_domain})}",
					},
					{
						Name:  "jsonrpc-allowed-methods",
						Value: "verto",
					},
				},
			},
			Variables: variables{
				Variable: []variable{
					{
						Name:  "record_stereo",
						Value: "true",
					},
					{
						Name:  "default_gateway",
						Value: "$${default_provider}",
					},
					{
						Name:  "default_areacode",
						Value: "$${default_areacode}",
					},
					{
						Name:  "transfer_fallback_extension",
						Value: "operator",
					},
				},
			},
			Groups: groups{
				Group: []group{
					{
						Name: "default",
						Users: users{
							User: []user{
								{
									ID: userID,
									Params: params{
										Param: []param{
											{
												Name:  "password",
												Value: password,
											},
											{
												Name:  "vm-password",
												Value: userID,
											},
										},
									},
									Variables: variables{
										Variable: []variable{
											{
												Name:  "toll_allow",
												Value: "domestic,international,local",
											},
											{
												Name:  "accountcode",
												Value: userID,
											},
											{
												Name:  "user_context",
												Value: "default",
											},
											{
												Name:  "effective_caller_id_name",
												Value: fmt.Sprintf("Extension %s", userID),
											},
											{
												Name:  "effective_caller_id_number",
												Value: userID,
											},
											{
												Name:  "outbound_caller_id_name",
												Value: "$${outbound_caller_name}",
											},
											{
												Name:  "outbound_caller_id_number",
												Value: "$${outbound_caller_id}",
											},
											{
												Name:  "callgroup",
												Value: "default",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	ctx.XML(http.StatusOK, include)
}
