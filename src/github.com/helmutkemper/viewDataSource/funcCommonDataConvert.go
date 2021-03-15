package viewDataSource

import (
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/dataFormat"
)

/*
dataSource: [
                    {
                        text: "Kemper.com.br",
                        spriteCssClass: "menuWith fas fa-code-branch",
                        items: [
                            {
                                text: "Github",
                                dataSpriteCssClassField: "k-test-menu",
                                spriteCssClass: "menuWith fab fa-github",
                                url: "https://github.com/helmutkemper"
                            }
                        ]
                    },
                    {
                        text: "Códigos",
                        spriteCssClass: "menuWith fas fa-code",
                        items: [
                            {
                                text: "Conversando com o dev",
                                spriteCssClass: "menuWith fas fa-fire-extinguisher",
                                items: [
                                    {
                                        text: "Conversando com o dev",
                                        imageUrl: "../content/shared/icons/16/star.png"
                                    },
                                    {
                                        text: "Sérgio Desenvolvedor",
                                        imageUrl: "../content/shared/icons/16/photo.png"
                                    }
                                ]
                            },
                            {
                                text: "Sérgio Desenvolvedor",
                                spriteCssClass: "menuWith fas fa-fire",
                                items: [
                                    {
                                        text: "Hello World"
                                    },
                                    {
                                        text: "Sérgio Desenvolvedor"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        text: "Admin",
                        spriteCssClass: "menuWith fas fa-cogs",
                        items: [
                            {
                                text: "Login"
                            },
                            {
                                text: "New content"
                            }
                        ]
                    },
                    {
                        text: "Doação",
                        spriteCssClass: "menuWith fas fa-donate"
                    }
                ]}
*/

type populateList struct {
	menuFromDataSource *[]dataFormat.Menu
	menuToPopulate     *[]Item
}

func (e *Menu) getRefs(menuFromDataSource *[]dataFormat.Menu, menuToPopulate *[]Item) (pointer []populateList) {
	pointer = make([]populateList, 0)
	*menuToPopulate = make([]Item, len(*menuFromDataSource))

	pointer = append(
		pointer,
		populateList{
			menuFromDataSource: menuFromDataSource,
			menuToPopulate:     menuToPopulate,
		},
	)
	for k := range *menuFromDataSource {
		if (*menuFromDataSource)[k].Menu != nil {
			pointer = append(
				pointer,
				populateList{
					menuFromDataSource: &(*menuFromDataSource)[k].Menu,
					menuToPopulate:     &(*menuToPopulate)[k].Items,
				},
			)
			pointer = append(pointer, e.getRefs(&(*menuFromDataSource)[k].Menu, &(*menuToPopulate)[k].Items)...)
		}
	}

	return
}

func (e *Menu) CommonDataConvert(menuFromDataSource *[]dataFormat.Menu) {
	menuToPopulate := make([]Item, 0)

	refs := e.getRefs(menuFromDataSource, &menuToPopulate)

	for refKey := range refs {
		items := *(refs[refKey]).menuToPopulate

		for k, v := range *(refs[refKey]).menuFromDataSource {
			items[k].Text = v.Text
			items[k].DataSpriteCssClassField = "k-test-menu"
			items[k].Url = v.Url

			if v.Icon != "" {
				items[k].SpriteCssClass = "menuWith " + v.Icon
			}
		}
	}

	*e = menuToPopulate
}
