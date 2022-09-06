package tabs

import "github.com/jfcote87/esign/v2.1/model"

// Merge merges the map data into the tabs structure
func Merge(tb *model.Tabs, m map[string]interface{}) *model.Tabs {
	r := &model.Tabs{}

	for _, t := range tb.ApproveTabs {
		r.ApproveTabs = append(r.ApproveTabs, t)
	}
	for _, t := range tb.CheckboxTabs {
		if v, ok := m[t.TabLabel]; ok {
			bval, _ := v.(bool)
			t.Selected = model.DSBool(bval)
		}
		r.CheckboxTabs = append(r.CheckboxTabs, t)
	}

	for _, t := range tb.CompanyTabs {
		r.CompanyTabs = append(r.CompanyTabs, t)
	}

	for _, t := range tb.DateSignedTabs {
		r.DateSignedTabs = append(r.DateSignedTabs, t)
	}
	for _, t := range tb.DeclineTabs {
		r.DeclineTabs = append(r.DeclineTabs, t)
	}
	for _, t := range tb.EmailTabs {
		if v, ok := m[t.TabLabel]; ok {
			t.Value, _ = v.(string)
		}
		r.EmailTabs = append(r.EmailTabs, t)
	}
	for _, t := range tb.EnvelopeIDTabs {
		r.EnvelopeIDTabs = append(r.EnvelopeIDTabs, t)
	}
	for _, t := range tb.FullNameTabs {
		r.FullNameTabs = append(r.FullNameTabs, t)
	}
	for _, t := range tb.InitialHereTabs {
		r.InitialHereTabs = append(r.InitialHereTabs, t)
	}
	for _, t := range tb.ListTabs {
		if v, ok := m[t.TabLabel]; ok {
			if tval, ok := v.(string); ok {
				for i := range t.ListItems {
					if t.ListItems[i].Value == tval {
						t.ListItems[i].Selected = true
					} else {
						t.ListItems[i].Selected = false
					}
				}
			}
		}
		r.ListTabs = append(r.ListTabs, t)
	}
	for _, t := range tb.NoteTabs {
		if v, ok := m[t.TabLabel]; ok {
			t.Value, _ = v.(string)
		}
		r.NoteTabs = append(r.NoteTabs, t)
	}
	for _, t := range tb.NumberTabs {
		if v, ok := m[t.TabLabel]; ok {
			t.Value, _ = v.(string)
		}
		r.NumberTabs = append(r.NumberTabs, t)
	}

	for _, t := range tb.RadioGroupTabs {
		if v, ok := m[t.GroupName]; ok {
			if tval, ok := v.(string); ok {
				for i := range t.Radios {
					if t.Radios[i].Value == tval {
						t.Radios[i].Selected = true
					} else {
						t.Radios[i].Selected = false
					}
				}
			}
		}
		r.RadioGroupTabs = append(r.RadioGroupTabs, t)
	}
	for _, t := range tb.SignHereTabs {
		r.SignHereTabs = append(r.SignHereTabs, t)
	}
	for _, t := range tb.SignerAttachmentTabs {
		r.SignerAttachmentTabs = append(r.SignerAttachmentTabs, t)
	}
	for _, t := range tb.SSNTabs {
		if v, ok := m[t.TabLabel]; ok {
			t.Value, _ = v.(string)
		}
		r.SSNTabs = append(r.SSNTabs, t)
	}
	for _, t := range tb.TextTabs {
		if v, ok := m[t.TabLabel]; ok {
			t.Value, _ = v.(string)
		}
		r.TextTabs = append(r.TextTabs, t)
	}
	for _, t := range tb.TitleTabs {
		r.TitleTabs = append(r.TitleTabs, t)
	}
	for _, t := range tb.ZipTabs {
		if v, ok := m[t.TabLabel]; ok {
			t.Value, _ = v.(string)
		}
		r.ZipTabs = append(r.ZipTabs, t)
	}
	return r
}

var tabsForAsst = &model.Tabs{
	InitialHereTabs: []model.InitialHere{
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabPosition: model.TabPosition{
				PageNumber: "2",
				XPosition:  "535",
				YPosition:  "703",
				TabLabel:   "initConf",
			},
			ScaleValue: "1",
			Name:       "Initial",
		},
	},
}

var tabsForDir = &model.Tabs{
	SignHereTabs: []model.SignHere{
		{
			TabBase: model.TabBase{
				DocumentID: "1",
			},
			TabPosition: model.TabPosition{
				AnchorString:  "dir00_Signature",
				AnchorXOffset: "0",
				AnchorYOffset: "0",
				AnchorUnits:   "mms",
			},
		},
	},
}

var tabsForFellow = &model.Tabs{
	SignHereTabs: []model.SignHere{
		{
			TabBase: model.TabBase{
				DocumentID: "1",
			},
			TabPosition: model.TabPosition{
				AnchorString:  "lf000_Signature",
				AnchorXOffset: "0",
				AnchorYOffset: "0",
				AnchorUnits:   "mms",
			},
		},
	},
	TextTabs: []model.Text{
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Conference Date",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "129",
				YPosition:  "40",
				TabLabel:   "lockConfDate",
			},
			Locked:   true,
			Required: model.REQUIRED_TRUE,
			Shared:   true,
			Width:    "180",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Location",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "129",
				YPosition:  "67",
				TabLabel:   "lockLocation",
			},
			Locked:   true,
			Required: model.REQUIRED_TRUE,
			Shared:   true,
			Width:    "354",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Conference Name",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size16",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "129",
				YPosition:  "15",
				TabLabel:   "lockConfName",
			},
			Locked:   true,
			Required: model.REQUIRED_TRUE,
			Shared:   true,
			Width:    "432",
			Height:   "22",
		},
	},
}

var tabsForConferees = &model.Tabs{
	CheckboxTabs: []model.Checkbox{
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabPosition: model.TabPosition{
				PageNumber: "2",
				XPosition:  "24",
				YPosition:  "655",
				TabLabel:   "cbSSN",
			},
			TabStyle: model.TabStyle{
				Name: "I have a US Social Security Number",
			},
			Selected: false,
			Shared:   true,
		},
	},
	DateSignedTabs: []model.DateSigned{
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name: "Date Signed",
			},
			TabPosition: model.TabPosition{
				PageNumber: "2",
				XPosition:  "391",
				YPosition:  "738",
				TabLabel:   "Date Signed",
			},
		},
	},
	EmailAddressTabs: []model.EmailAddress{
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Email Address for conferee list (if different)",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "12",
				YPosition:  "60",
				TabLabel:   "txtEmail",
			},
			Width: "546",
		},
	},
	ListTabs: []model.List{
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "11",
				YPosition:  "643",
				TabLabel:   "txtCorrespondence",
			},
			ListItems: []model.ListItem{
				{
					Selected: false,
					Text:     "Home",
					Value:    "Home",
				},
				{
					Selected: false,
					Text:     "Work",
					Value:    "Work",
				},
				{
					Selected: false,
					Text:     "Other (describe below)",
					Value:    "Other (describe below)",
				},
			},
			Font:              "verdana",
			FontColor:         "black",
			FontSize:          "size12",
			Required:          model.REQUIRED_TRUE,
			Shared:            true,
			Width:             "156",
			ListSelectedValue: "",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "307",
				YPosition:  "643",
				TabLabel:   "txtConfereeList",
			},
			ListItems: []model.ListItem{
				{
					Selected: false,
					Text:     "Home",
					Value:    "Home",
				},
				{
					Selected: false,
					Text:     "Work",
					Value:    "Work",
				},
				{
					Selected: false,
					Text:     "Other (describe below)",
					Value:    "Other (describe below)",
				},
			},
			Font:              "verdana",
			FontColor:         "black",
			FontSize:          "size12",
			Required:          model.REQUIRED_TRUE,
			Shared:            true,
			Width:             "156",
			ListSelectedValue: "",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "11",
				YPosition:  "727",
				TabLabel:   "txtPayment",
			},
			ListItems: []model.ListItem{
				{
					Selected: false,
					Text:     "Home",
					Value:    "Home",
				},
				{
					Selected: false,
					Text:     "Work",
					Value:    "Work",
				},
				{
					Selected: false,
					Text:     "Other (describe below)",
					Value:    "Other (describe below)",
				},
			},
			Font:              "verdana",
			FontColor:         "black",
			FontSize:          "size12",
			Required:          model.REQUIRED_TRUE,
			Shared:            true,
			Width:             "156",
			ListSelectedValue: "",
		},
	},
	RadioGroupTabs: []model.RadioGroup{
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			GroupName: "rbGov",
			Radios: []model.Radio{
				{
					PageNumber: "2",
					XPosition:  "89",
					YPosition:  "545",
					Required:   model.REQUIRED_TRUE,
					Value:      "no",
				},
				{
					PageNumber: "2",
					XPosition:  "24",
					YPosition:  "545",
					Required:   model.REQUIRED_TRUE,
					Value:      "yes",
				},
			},
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			GroupName: "rbUSCit",
			Radios: []model.Radio{
				{
					PageNumber: "2",
					XPosition:  "22",
					YPosition:  "616",
					Required:   model.REQUIRED_TRUE,
					Value:      "cit",
				},
				{
					PageNumber: "2",
					XPosition:  "116",
					YPosition:  "616",
					Required:   model.REQUIRED_TRUE,
					Value:      "resident",
				},
				{
					PageNumber: "2",
					XPosition:  "247",
					YPosition:  "616",
					Required:   model.REQUIRED_TRUE,
					Value:      "other",
				},
			},
		},
	},
	SignHereTabs: []model.SignHere{
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			Name: "Date Signed",
			TabPosition: model.TabPosition{
				PageNumber: "2",
				XPosition:  "43",
				YPosition:  "705",
				TabLabel:   "sigConferee",
			},
		},
	},
	TextTabs: []model.Text{
		{
			TabBase: model.TabBase{
				DocumentID:             "2",
				ConditionalParentLabel: "rbUSCit",
				ConditionalParentValue: "other",
			},
			TabStyle: model.TabStyle{
				Name:      "Country of Residence",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "2",
				XPosition:  "330",
				YPosition:  "633",
				TabLabel:   "txtCountry",
			},
			Required:               model.REQUIRED_TRUE,
			Shared:                 true,
			Width:                  "204",
			Height:                 "22",
			ConcealValueOnDocument: false,
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Conferee First Name",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "188",
				YPosition:  "124",
				TabLabel:   "txtFName",
			},
			Required: model.REQUIRED_TRUE,
			Shared:   true,
			Width:    "264",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Conferee Affiliation",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "7",
				YPosition:  "196",
				TabLabel:   "txtAff",
			},
			Required: model.REQUIRED_TRUE,
			Shared:   true,
			Width:    "336",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "347",
				YPosition:  "196",
				TabLabel:   "txtAffTitle",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   true,
			Width:    "264",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Conferee Home Phone",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "12",
				YPosition:  "399",
				TabLabel:   "txtHomePhone",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   true,
			Width:    "234",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Conferee Home Fax",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "257",
				YPosition:  "399",
				TabLabel:   "txtHomeFax",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   true,
			Width:    "234",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Conferee BusFax",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "261",
				YPosition:  "574",
				TabLabel:   "txtBusFax",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   true,
			Width:    "234",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Business Mailing Address",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "12",
				YPosition:  "453",
				TabLabel:   "txtBusAddr",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   false,
			Width:    "468",
			Height:   "88",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Conferee MiddleName/Initial",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "468",
				YPosition:  "123",
				TabLabel:   "txtMiddle",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   true,
			Width:    "132",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Conferee Last Name",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "11",
				YPosition:  "160",
				TabLabel:   "txtLName",
			},
			Required: model.REQUIRED_TRUE,
			Shared:   true,
			Width:    "210",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Prefix (Jr., Sr., III, etc)",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "226",
				YPosition:  "161",
				TabLabel:   "txtSuffix",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   true,
			Width:    "96",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Conferee Business Phone",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "10",
				YPosition:  "573",
				TabLabel:   "txtBusPhone",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   true,
			Width:    "234",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Home Mailing Address",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "12",
				YPosition:  "280",
				TabLabel:   "txtHomeAddr",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   false,
			Width:    "444",
			Height:   "99",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Name Tags and Table Signs",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "2",
				XPosition:  "12",
				YPosition:  "111",
				TabLabel:   "txtTableTent",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   true,
			Width:    "552",
			Height:   "22",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Dietary Restrictions",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "2",
				XPosition:  "12",
				YPosition:  "156",
				TabLabel:   "txtDiet",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   true,
			Width:    "540",
			Height:   "66",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Other Information",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "2",
				XPosition:  "11",
				YPosition:  "342",
				TabLabel:   "txtAnythingElse",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   true,
			Width:    "540",
			Height:   "154",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Other accomodations",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "2",
				XPosition:  "11",
				YPosition:  "247",
				TabLabel:   "txtHotel",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   true,
			Width:    "552",
			Height:   "55",
		},
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabStyle: model.TabStyle{
				Name:      "Conferee Prefix",
				Font:      "verdana",
				FontColor: "black",
				FontSize:  "size12",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "11",
				YPosition:  "124",
				TabLabel:   "txtPrefix",
			},
			Required: model.REQUIRED_DEFAULT,
			Shared:   true,
			Width:    "162",
			Height:   "22",
		},
	},
	DeclineTabs: []model.Decline{
		{
			TabBase: model.TabBase{
				DocumentID: "2",
			},
			TabPosition: model.TabPosition{
				PageNumber: "1",
				XPosition:  "532",
				YPosition:  "64",
				TabLabel:   "btnDecline",
			},
			Font:      "verdana",
			FontColor: "black",
			FontSize:  "size12",
			Width:     "162",
			Height:    "22",
		},
	},
}

func getTabs() map[string]model.Tabs {
	return map[string]model.Tabs{
		"asstTabs": {
			InitialHereTabs: []model.InitialHere{
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabPosition: model.TabPosition{
						PageNumber: "2",
						XPosition:  "535",
						YPosition:  "703",
						TabLabel:   "initConf",
					},
					ScaleValue: "1",
					Name:       "Initial",
				},
			},
		},
		"dir": {
			SignHereTabs: []model.SignHere{
				{
					TabBase: model.TabBase{
						DocumentID: "1",
					},
					TabPosition: model.TabPosition{
						AnchorString:  "dir00_Signature",
						AnchorXOffset: "0",
						AnchorYOffset: "0",
						AnchorUnits:   "mms",
					},
				},
			},
		},
		"fellow": {
			SignHereTabs: []model.SignHere{
				{
					TabBase: model.TabBase{
						DocumentID: "1",
					},
					TabPosition: model.TabPosition{
						AnchorString:  "lf000_Signature",
						AnchorXOffset: "0",
						AnchorYOffset: "0",
						AnchorUnits:   "mms",
					},
				},
			},
			TextTabs: []model.Text{
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Conference Date",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "129",
						YPosition:  "40",
						TabLabel:   "lockConfDate",
					},
					Locked:   true,
					Required: model.REQUIRED_TRUE,
					Shared:   true,
					Width:    "180",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Location",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "129",
						YPosition:  "67",
						TabLabel:   "lockLocation",
					},
					Locked:   true,
					Required: model.REQUIRED_TRUE,
					Shared:   true,
					Width:    "354",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Conference Name",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size16",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "129",
						YPosition:  "15",
						TabLabel:   "lockConfName",
					},
					Locked:   true,
					Required: model.REQUIRED_TRUE,
					Shared:   true,
					Width:    "432",
					Height:   "22",
				},
			},
		},
		"conferee": {
			CheckboxTabs: []model.Checkbox{
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabPosition: model.TabPosition{
						PageNumber: "2",
						XPosition:  "24",
						YPosition:  "655",
						TabLabel:   "cbSSN",
					},
					TabStyle: model.TabStyle{
						Name: "I have a US Social Security Number",
					},
					Selected: false,
					Shared:   true,
				},
			},
			DateSignedTabs: []model.DateSigned{
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name: "Date Signed",
					},
					TabPosition: model.TabPosition{
						PageNumber: "2",
						XPosition:  "391",
						YPosition:  "738",
						TabLabel:   "Date Signed",
					},
				},
			},
			EmailAddressTabs: []model.EmailAddress{
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Email Address for conferee list (if different)",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "12",
						YPosition:  "60",
						TabLabel:   "txtEmail",
					},
					Width: "546",
				},
			},
			ListTabs: []model.List{
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "11",
						YPosition:  "643",
						TabLabel:   "txtCorrespondence",
					},
					ListItems: []model.ListItem{
						{
							Selected: false,
							Text:     "Home",
							Value:    "Home",
						},
						{
							Selected: false,
							Text:     "Work",
							Value:    "Work",
						},
						{
							Selected: false,
							Text:     "Other (describe below)",
							Value:    "Other (describe below)",
						},
					},
					Font:              "verdana",
					FontColor:         "black",
					FontSize:          "size12",
					Required:          model.REQUIRED_TRUE,
					Shared:            true,
					Width:             "156",
					ListSelectedValue: "",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "307",
						YPosition:  "643",
						TabLabel:   "txtConfereeList",
					},
					ListItems: []model.ListItem{
						{
							Selected: false,
							Text:     "Home",
							Value:    "Home",
						},
						{
							Selected: false,
							Text:     "Work",
							Value:    "Work",
						},
						{
							Selected: false,
							Text:     "Other (describe below)",
							Value:    "Other (describe below)",
						},
					},
					Font:              "verdana",
					FontColor:         "black",
					FontSize:          "size12",
					Required:          model.REQUIRED_TRUE,
					Shared:            true,
					Width:             "156",
					ListSelectedValue: "",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "11",
						YPosition:  "727",
						TabLabel:   "txtPayment",
					},
					ListItems: []model.ListItem{
						{
							Selected: false,
							Text:     "Home",
							Value:    "Home",
						},
						{
							Selected: false,
							Text:     "Work",
							Value:    "Work",
						},
						{
							Selected: false,
							Text:     "Other (describe below)",
							Value:    "Other (describe below)",
						},
					},
					Font:              "verdana",
					FontColor:         "black",
					FontSize:          "size12",
					Required:          model.REQUIRED_TRUE,
					Shared:            true,
					Width:             "156",
					ListSelectedValue: "",
				},
			},
			RadioGroupTabs: []model.RadioGroup{
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					GroupName: "rbGov",
					Radios: []model.Radio{
						{
							PageNumber: "2",
							XPosition:  "89",
							YPosition:  "545",
							Required:   model.REQUIRED_TRUE,
							Value:      "no",
						},
						{
							PageNumber: "2",
							XPosition:  "24",
							YPosition:  "545",
							Required:   model.REQUIRED_TRUE,
							Value:      "yes",
						},
					},
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					GroupName: "rbUSCit",
					Radios: []model.Radio{
						{
							PageNumber: "2",
							XPosition:  "22",
							YPosition:  "616",
							Required:   model.REQUIRED_TRUE,
							Value:      "cit",
						},
						{
							PageNumber: "2",
							XPosition:  "116",
							YPosition:  "616",
							Required:   model.REQUIRED_TRUE,
							Value:      "resident",
						},
						{
							PageNumber: "2",
							XPosition:  "247",
							YPosition:  "616",
							Required:   model.REQUIRED_TRUE,
							Value:      "other",
						},
					},
				},
			},
			SignHereTabs: []model.SignHere{
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					Name: "Date Signed",
					TabPosition: model.TabPosition{
						PageNumber: "2",
						XPosition:  "43",
						YPosition:  "705",
						TabLabel:   "sigConferee",
					},
				},
			},
			TextTabs: []model.Text{
				{
					TabBase: model.TabBase{
						DocumentID:             "2",
						ConditionalParentLabel: "rbUSCit",
						ConditionalParentValue: "other",
					},
					TabStyle: model.TabStyle{
						Name:      "Country of Residence",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "2",
						XPosition:  "330",
						YPosition:  "633",
						TabLabel:   "txtCountry",
					},
					Required:               model.REQUIRED_TRUE,
					Shared:                 true,
					Width:                  "204",
					Height:                 "22",
					ConcealValueOnDocument: false,
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Conferee First Name",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "188",
						YPosition:  "124",
						TabLabel:   "txtFName",
					},
					Required: model.REQUIRED_TRUE,
					Shared:   true,
					Width:    "264",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Conferee Affiliation",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "7",
						YPosition:  "196",
						TabLabel:   "txtAff",
					},
					Required: model.REQUIRED_TRUE,
					Shared:   true,
					Width:    "336",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "347",
						YPosition:  "196",
						TabLabel:   "txtAffTitle",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   true,
					Width:    "264",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Conferee Home Phone",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "12",
						YPosition:  "399",
						TabLabel:   "txtHomePhone",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   true,
					Width:    "234",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Conferee Home Fax",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "257",
						YPosition:  "399",
						TabLabel:   "txtHomeFax",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   true,
					Width:    "234",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Conferee BusFax",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "261",
						YPosition:  "574",
						TabLabel:   "txtBusFax",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   true,
					Width:    "234",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Business Mailing Address",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "12",
						YPosition:  "453",
						TabLabel:   "txtBusAddr",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   false,
					Width:    "468",
					Height:   "88",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Conferee MiddleName/Initial",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "468",
						YPosition:  "123",
						TabLabel:   "txtMiddle",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   true,
					Width:    "132",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Conferee Last Name",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "11",
						YPosition:  "160",
						TabLabel:   "txtLName",
					},
					Required: model.REQUIRED_TRUE,
					Shared:   true,
					Width:    "210",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Prefix (Jr., Sr., III, etc)",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "226",
						YPosition:  "161",
						TabLabel:   "txtSuffix",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   true,
					Width:    "96",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Conferee Business Phone",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "10",
						YPosition:  "573",
						TabLabel:   "txtBusPhone",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   true,
					Width:    "234",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Home Mailing Address",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "12",
						YPosition:  "280",
						TabLabel:   "txtHomeAddr",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   false,
					Width:    "444",
					Height:   "99",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Name Tags and Table Signs",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "2",
						XPosition:  "12",
						YPosition:  "111",
						TabLabel:   "txtTableTent",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   true,
					Width:    "552",
					Height:   "22",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Dietary Restrictions",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "2",
						XPosition:  "12",
						YPosition:  "156",
						TabLabel:   "txtDiet",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   true,
					Width:    "540",
					Height:   "66",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Other Information",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "2",
						XPosition:  "11",
						YPosition:  "342",
						TabLabel:   "txtAnythingElse",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   true,
					Width:    "540",
					Height:   "154",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Other accomodations",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "2",
						XPosition:  "11",
						YPosition:  "247",
						TabLabel:   "txtHotel",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   true,
					Width:    "552",
					Height:   "55",
				},
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabStyle: model.TabStyle{
						Name:      "Conferee Prefix",
						Font:      "verdana",
						FontColor: "black",
						FontSize:  "size12",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "11",
						YPosition:  "124",
						TabLabel:   "txtPrefix",
					},
					Required: model.REQUIRED_DEFAULT,
					Shared:   true,
					Width:    "162",
					Height:   "22",
				},
			},
			DeclineTabs: []model.Decline{
				{
					TabBase: model.TabBase{
						DocumentID: "2",
					},
					TabPosition: model.TabPosition{
						PageNumber: "1",
						XPosition:  "532",
						YPosition:  "64",
						TabLabel:   "btnDecline",
					},
					Font:      "verdana",
					FontColor: "black",
					FontSize:  "size12",
					Width:     "162",
					Height:    "22",
				},
			},
		},
	}
}
